package dictionaries

import "go.mongodb.org/mongo-driver/bson/primitive"

type SkillTypeCreateParams struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	SkillKey  string             `json:"skill_key" validate:"required"`
	SkillIcon string             `json:"skill_icon" validate:"required"`
}

type SkillTypeUpdateParams struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	SkillKey  string             `json:"skill_key"`
	SkillIcon string             `json:"skill_icon"`
}

type SkillType struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SkillKey  string             `bson:"skill_key" json:"skill_key"`
	SkillIcon string             `bson:"skill_icon" json:"skill_icon"`
}
