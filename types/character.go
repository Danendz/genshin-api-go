package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Character struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	VisionID       primitive.ObjectID `bson:"vision_id" json:"vision_id"`
	WeaponTypeID   primitive.ObjectID `bson:"weapon_type_id" json:"weapon_type_id"`
	Name           string             `bson:"name" json:"name"`
	NameKey        string             `bson:"name_key" json:"name_key"`
	NationKey      string             `bson:"nation_key" json:"nation_key"`
	PassiveTalents []PassiveTalent    `bson:"passive_talents" json:"passive_talents"`
	Rarity         uint8              `bson:"rarity" json:"rarity"`
	SkillTalents   SkillTalent        `bson:"skill_talent" json:"skill_talent"`
	Affiliation    string             `bson:"affiliation" json:"affiliation"`
	Birthday       string             `bson:"birthday" json:"birthday"`
	Constellation  string             `bson:"constellation" json:"constellation"`
	Constellations []Constellation    `bson:"constellations" json:"constellations"`
	Description    string             `bson:"description" json:"description"`
	Title          string             `bson:"title" json:"title"`
}

type Constellation struct {
	Name        string `bson:"name" json:"name"`
	Level       uint8  `bson:"level" json:"level"`
	Description string `bson:"description" json:"description"`
}

type PassiveTalent struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}

type SkillTalent struct {
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Type        SkillType `bson:"type" json:"type"`
}

type SkillType uint8

const (
	NORMAL_ATTACK SkillType = iota
	ELEMENTAL_SKILL
	ELEMENTAL_BURST
)
