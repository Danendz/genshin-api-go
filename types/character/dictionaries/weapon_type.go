package dictionaries

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeaponTypeCreateParams struct {
	ID         primitive.ObjectID `json:"_id,omitempty"`
	WeaponKey  string             `json:"weapon_key" validate:"required"`
	WeaponIcon string             `json:"weapon_icon" validate:"required"`
}

type WeaponTypeUpdateParams struct {
	ID         primitive.ObjectID `json:"_id,omitempty"`
	WeaponKey  string             `json:"weapon_key"`
	WeaponIcon string             `json:"weapon_icon"`
}

type WeaponType struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	WeaponKey  string             `bson:"weapon_key" json:"weapon_key"`
	WeaponIcon string             `bson:"weapon_icon" json:"weapon_icon"`
}
