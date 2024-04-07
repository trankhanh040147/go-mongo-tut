package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *CollStore) Update(ctx context.Context, filter, data *bson.D) (*mongo.UpdateResult, error) {
	result, err := s.coll.UpdateOne(ctx, filter, data)
	if err != nil {
		panic(err)
	}

	return result, err
}
