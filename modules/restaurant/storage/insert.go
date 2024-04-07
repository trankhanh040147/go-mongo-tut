package storage

import (
	"context"

	"github.com/trankhanh040147/go-mongo-tut/modules/restaurant/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *CollStore) Insert(ctx context.Context, data *model.Restaurant) (*mongo.InsertOneResult, error) {
	result, err := s.coll.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}

	return result, err
}
