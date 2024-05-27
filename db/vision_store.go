package db

import (
	"context"

	"github.com/Danendz/genshin-api-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const visionCol = "vision"

type VisionStore interface {
	GetVisions(ctx context.Context) ([]*types.Vision, error)
}

type MongoVisionStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoVisionStore(client *mongo.Client, dbcreds *DBCreds) *MongoVisionStore {
	return &MongoVisionStore{
		client: client,
		coll:   client.Database(dbcreds.DBNAME).Collection(visionCol),
	}
}

func (s *MongoVisionStore) GetVisions(ctx context.Context) ([]*types.Vision, error) {
	var visions []*types.Vision

	cur, err := s.coll.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	if err = cur.All(ctx, &visions); err != nil {
		return nil, err
	}

	return visions, nil
}
