package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeaponTypeCreateParams struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	WeaponKey string             `json:"weapon_key" validate:"required"`
}

type WeaponTypeUpdateParams struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	WeaponKey string             `json:"weapon_key"`
}

type WeaponType struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	WeaponKey string             `bson:"weapon_key" json:"weapon_key"`
}
