// Copyright 2019 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongo

import (
	"context"
	"errors"
	"math"
	"sync"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
	"github.com/micro/go-micro/metadata"
	"github.com/opensds/multi-cloud/api/pkg/common"
	"github.com/opensds/multi-cloud/backend/pkg/model"
)

type mongoRepository struct {
	session *mgo.Session
}

var defaultDBName = "multi-cloud"
var defaultCollection = "backends"
var mutex sync.Mutex
var mongoRepo = &mongoRepository{}

func Init(host string) *mongoRepository {
	mutex.Lock()
	defer mutex.Unlock()

	if mongoRepo.session != nil {
		return mongoRepo
	}

	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	mongoRepo.session = session
	return mongoRepo
}

// The implementation of Repository
func UpdateFilter(m bson.M, filter map[string]string) error {
	for k, v := range filter {
		m[k] = interface{}(v)
	}
	return nil
}

func UpdateContextFilter(ctx context.Context, m bson.M) error {
	// if context is admin, no need filter by tenantId.
	md, ok := metadata.FromContext(ctx)
	if !ok {
		log.Error("get context failed")
		return errors.New("get context failed")
	}

	isAdmin, _ := md[common.CTX_KEY_IS_ADMIN]
	if isAdmin != common.CTX_VAL_TRUE {
		tenantId, ok := md[common.CTX_KEY_TENANT_ID]
		if !ok {
			log.Error("get tenantid failed")
			return errors.New("get tenantid failed")
		}
		m["tenantId"] = tenantId
	}

	return nil
}

func (repo *mongoRepository) CreateBackend(ctx context.Context, backend *model.Backend) (*model.Backend, error) {
	session := repo.session.Copy()
	defer session.Close()

	if backend.Id == "" {
		backend.Id = bson.NewObjectId()
	}

	err := session.DB(defaultDBName).C(defaultCollection).Insert(backend)
	if err != nil {
		return nil, err
	}
	return backend, nil
}

func (repo *mongoRepository) DeleteBackend(ctx context.Context, id string) error {
	session := repo.session.Copy()
	defer session.Close()

	m := bson.M{"_id": bson.ObjectIdHex(id)}
	err := UpdateContextFilter(ctx, m)
	if err != nil {
		return err
	}

	return session.DB(defaultDBName).C(defaultCollection).Remove(m)
}

func (repo *mongoRepository) UpdateBackend(ctx context.Context,
	backend *model.Backend) (*model.Backend, error) {
	session := repo.session.Copy()
	defer session.Close()

	m := bson.M{"_id": backend.Id}
	err := UpdateContextFilter(ctx, m)
	if err != nil {
		return nil, err
	}

	err = session.DB(defaultDBName).C(defaultCollection).Update(m, backend)
	if err != nil {
		return nil, err
	}

	return backend, nil
}

func (repo *mongoRepository) GetBackend(ctx context.Context, id string) (*model.Backend,
	error) {
	session := repo.session.Copy()
	defer session.Close()

	m := bson.M{"_id": bson.ObjectIdHex(id)}
	err := UpdateContextFilter(ctx, m)
	if err != nil {
		return nil, err
	}

	var backend = &model.Backend{}
	collection := session.DB(defaultDBName).C(defaultCollection)
	err = collection.Find(m).One(backend)
	if err != nil {
		return nil, err
	}
	return backend, nil
}

func (repo *mongoRepository) ListBackend(ctx context.Context, limit, offset int,
	query interface{}) ([]*model.Backend, error) {

	session := repo.session.Copy()
	defer session.Close()

	if limit == 0 {
		limit = math.MinInt32
	}
	var backends []*model.Backend

	m := bson.M{}
	UpdateFilter(m, query.(map[string]string))
	err := UpdateContextFilter(ctx, m)
	if err != nil {
		return nil, err
	}
	log.Infof("ListBackend, limit=%d, offset=%d, m=%+v\n", limit, offset, m)

	err = session.DB(defaultDBName).C(defaultCollection).Find(m).Skip(offset).Limit(limit).All(&backends)
	if err != nil {
		return nil, err
	}

	return backends, nil
}

func (repo *mongoRepository) Close() {
	repo.session.Close()
}
