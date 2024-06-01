package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeaponType struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	WeaponKey string             `bson:"weapon_key" json:"weapon_key"`
}
