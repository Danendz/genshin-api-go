package dictionaries

import "go.mongodb.org/mongo-driver/bson/primitive"

type VisionCreateParams struct {
	ID         primitive.ObjectID `json:"_id,omitempty"`
	VisionKey  string             `json:"vision_key" validate:"required"`
	VisionIcon string             `json:"vision_icon" validate:"required"`
}

type VisionUpdateParams struct {
	ID         primitive.ObjectID `json:"_id,omitempty"`
	VisionKey  string             `json:"vision_key"`
	VisionIcon string             `json:"vision_icon"`
}

type Vision struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	VisionKey  string             `bson:"vision_key" json:"vision_key"`
	VisionIcon string             `bson:"vision_icon" json:"vision_icon"`
}
