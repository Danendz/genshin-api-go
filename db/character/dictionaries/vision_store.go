package dictionaries

import (
	"context"
	"github.com/Danendz/genshin-api-go/db"
	"github.com/Danendz/genshin-api-go/types/character/dictionaries"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const visionCol = "visions"

type VisionStore interface {
	GetVisions(ctx context.Context) ([]*dictionaries.Vision, error)
	CreateVision(ctx context.Context, vision *dictionaries.VisionCreateParams) (*dictionaries.VisionCreateParams, error)
	DeleteVision(ctx context.Context, id string) error
	UpdateVision(ctx context.Context, id string, values *dictionaries.VisionUpdateParams) (*dictionaries.Vision, error)
}

type MongoVisionStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoVisionStore(client *mongo.Client, dbcreds *db.Creds) *MongoVisionStore {
	return &MongoVisionStore{
		client: client,
		coll:   client.Database(dbcreds.DBNAME).Collection(visionCol),
	}
}

func (s *MongoVisionStore) GetVisions(ctx context.Context) ([]*dictionaries.Vision, error) {
	var visions []*dictionaries.Vision

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

func (s *MongoVisionStore) CreateVision(ctx context.Context, vision *dictionaries.VisionCreateParams) (*dictionaries.VisionCreateParams, error) {
	res, err := s.coll.InsertOne(ctx, vision)

	if err != nil {
		return nil, err
	}

	vision.ID = res.InsertedID.(primitive.ObjectID)

	return vision, nil
}

func (s *MongoVisionStore) DeleteVision(ctx context.Context, id string) error {
	oid, err := db.ToObjectID(id)

	if err != nil {
		return err
	}

	if _, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid}); err != nil {
		return err
	}

	return nil
}

func (s *MongoVisionStore) UpdateVision(ctx context.Context, id string, values *dictionaries.VisionUpdateParams) (*dictionaries.Vision, error) {
	var vision *dictionaries.Vision
	oid, err := db.ToObjectID(id)

	if err != nil {
		return nil, err
	}

	update := db.MakeUpdateFormat(values)

	res := s.coll.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "_id", Value: oid}},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if res.Err() != nil {
		return nil, res.Err()
	}

	err = res.Decode(&vision)

	if err != nil {
		return nil, err
	}

	return vision, nil
}
