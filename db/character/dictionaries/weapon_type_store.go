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

const weaponTypeCol = "weapon-types"

type WeaponTypeStore interface {
	GetWeaponTypes(ctx context.Context) ([]*dictionaries.WeaponType, error)
	CreateWeaponType(ctx context.Context, weaponType *dictionaries.WeaponTypeCreateParams) (*dictionaries.WeaponTypeCreateParams, error)
	DeleteWeaponType(ctx context.Context, id string) error
	UpdateWeaponType(ctx context.Context, id string, values *dictionaries.WeaponTypeUpdateParams) (*dictionaries.WeaponType, error)
}

type MongoWeaponTypeStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoWeaponTypeStore(client *mongo.Client, dbcreds *db.Creds) *MongoWeaponTypeStore {
	return &MongoWeaponTypeStore{
		client: client,
		coll:   client.Database(dbcreds.DBNAME).Collection(weaponTypeCol),
	}
}

func (s *MongoWeaponTypeStore) GetWeaponTypes(ctx context.Context) ([]*dictionaries.WeaponType, error) {
	var weaponTypes []*dictionaries.WeaponType

	cur, err := s.coll.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	if err = cur.All(ctx, &weaponTypes); err != nil {
		return nil, err
	}

	return weaponTypes, nil
}

func (s *MongoWeaponTypeStore) CreateWeaponType(ctx context.Context, weaponType *dictionaries.WeaponTypeCreateParams) (*dictionaries.WeaponTypeCreateParams, error) {
	res, err := s.coll.InsertOne(ctx, weaponType)

	if err != nil {
		return nil, err
	}

	weaponType.ID = res.InsertedID.(primitive.ObjectID)

	return weaponType, nil
}

func (s *MongoWeaponTypeStore) DeleteWeaponType(ctx context.Context, id string) error {
	oid, err := db.ToObjectID(id)

	if err != nil {
		return err
	}

	if _, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid}); err != nil {
		return err
	}

	return nil
}

func (s *MongoWeaponTypeStore) UpdateWeaponType(ctx context.Context, id string, values *dictionaries.WeaponTypeUpdateParams) (*dictionaries.WeaponType, error) {
	var weaponType *dictionaries.WeaponType
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

	err = res.Decode(&weaponType)
	if err != nil {
		return nil, err
	}

	return weaponType, nil
}
