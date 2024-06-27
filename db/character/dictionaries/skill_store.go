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

const skillTypeCol = "skill-types"

type SkillTypeStore interface {
	GetSkillTypes(ctx context.Context) ([]*dictionaries.SkillType, error)
	CreateSkillType(ctx context.Context, skillType *dictionaries.SkillTypeCreateParams) (*dictionaries.SkillTypeCreateParams, error)
	DeleteSkillType(ctx context.Context, id string) error
	UpdateSkillType(ctx context.Context, id string, values *dictionaries.SkillTypeUpdateParams) (*dictionaries.SkillType, error)
}

type MongoSkillTypeStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoSkillTypeStore(client *mongo.Client, dbcreds *db.Creds) *MongoSkillTypeStore {
	return &MongoSkillTypeStore{
		client: client,
		coll:   client.Database(dbcreds.DBNAME).Collection(skillTypeCol),
	}
}

func (s *MongoSkillTypeStore) GetSkillTypes(ctx context.Context) ([]*dictionaries.SkillType, error) {
	var skillTypes []*dictionaries.SkillType

	cur, err := s.coll.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	if err = cur.All(ctx, &skillTypes); err != nil {
		return nil, err
	}

	return skillTypes, nil
}

func (s *MongoSkillTypeStore) CreateSkillType(ctx context.Context, skillType *dictionaries.SkillTypeCreateParams) (*dictionaries.SkillTypeCreateParams, error) {
	res, err := s.coll.InsertOne(ctx, skillType)

	if err != nil {
		return nil, err
	}

	skillType.ID = res.InsertedID.(primitive.ObjectID)

	return skillType, nil
}

func (s *MongoSkillTypeStore) DeleteSkillType(ctx context.Context, id string) error {
	oid, err := db.ToObjectID(id)

	if err != nil {
		return err
	}

	if _, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid}); err != nil {
		return err
	}

	return nil
}

func (s *MongoSkillTypeStore) UpdateSkillType(ctx context.Context, id string, values *dictionaries.SkillTypeUpdateParams) (*dictionaries.SkillType, error) {
	var skillType *dictionaries.SkillType
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

	err = res.Decode(&skillType)
	if err != nil {
		return nil, err
	}

	return skillType, nil
}
