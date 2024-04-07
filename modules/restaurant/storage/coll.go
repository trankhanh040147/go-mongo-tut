package storage

import "go.mongodb.org/mongo-driver/mongo"

type CollStore struct {
	coll *mongo.Collection
}

func NewCollStore(coll *mongo.Collection) *CollStore {
	return &CollStore{coll}
}
