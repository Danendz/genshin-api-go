package db

import (
	"context"

	"github.com/Danendz/genshin-api-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const characterCol = "characters"

type CharacterStore interface {
	GetCharacters(ctx context.Context) ([]*types.Character, error)
	GetCharacter(ctx context.Context, id string) (*types.Character, error)
	CreateCharacter(ctx context.Context, character *types.Character) (*types.Character, error)
	DeleteCharacter(ctx context.Context, id string) error
	UpdateCharacter(ctx context.Context, id string, values *bson.M) (*types.Character, error)
}

type MongoCharacterStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoCharacterStore(client *mongo.Client, dbcreds *DBCreds) *MongoCharacterStore {
	return &MongoCharacterStore{
		client: client,
		coll:   client.Database(dbcreds.DBNAME).Collection(characterCol),
	}
}

func (s *MongoCharacterStore) GetCharacters(ctx context.Context) ([]*types.Character, error) {
	var characters []*types.Character

	cur, err := s.coll.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	if err = cur.All(ctx, &characters); err != nil {
		return nil, err
	}

	return characters, nil
}

func (s *MongoCharacterStore) GetCharacter(ctx context.Context, id string) (*types.Character, error) {
	var character *types.Character

	oid, err := ToObjectID(id)

	if err != nil {
		return nil, err
	}

	if err = s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&character); err != nil {
		return nil, err
	}

	return character, nil
}

func (s *MongoCharacterStore) CreateCharacter(ctx context.Context, character *types.Character) (*types.Character, error) {
	res, err := s.coll.InsertOne(ctx, character)

	if err != nil {
		return nil, err
	}

	character.ID = res.InsertedID.(primitive.ObjectID)

	return character, nil
}

func (s *MongoCharacterStore) DeleteCharacter(ctx context.Context, id string) error {
	oid, err := ToObjectID(id)

	if err != nil {
		return err
	}

	if _, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid}); err != nil {
		return err
	}

	return nil
}

func (s *MongoCharacterStore) UpdateCharacter(ctx context.Context, id string, values *bson.M) (*types.Character, error) {
	var character *types.Character
	oid, err := ToObjectID(id)

	if err != nil {
		return nil, err
	}

	update := MakeUpdateFormat(values)

	res := s.coll.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "_id", Value: oid}},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if res.Err() != nil {
		return nil, res.Err()
	}

	res.Decode(&character)

	return character, nil
}
