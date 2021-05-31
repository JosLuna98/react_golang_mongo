package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResultData struct {
	Error  bool          `json:"error,omitempty"`
	Result []primitive.M `json:"result,omitempty"`
}
