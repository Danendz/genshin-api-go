package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vision struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	VisionKey string             `bson:"vision_key" json:"vision_key"`
}
