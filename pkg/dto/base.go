package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BaseDTO struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at"`
	CreatedBy string             `json:"created_by"`
	UpdatedBy string             `json:"updated_by"`
}
