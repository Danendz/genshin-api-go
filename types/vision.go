package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type VisionCreateParams struct {
	ID          primitive.ObjectID `json:"_id,omitempty"`
	VisionKey   string             `json:"vision_key" validate:"required"`
	VisionImage string             `json:"vision_image" validate:"required"`
}

type VisionUpdateParams struct {
	ID          primitive.ObjectID `json:"_id,omitempty"`
	VisionKey   string             `json:"vision_key"`
	VisionImage string             `json:"vision_image"`
}

type Vision struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	VisionKey   string             `bson:"vision_key" json:"vision_key"`
	VisionImage string             `bson:"vision_image" json:"vision_image"`
}
