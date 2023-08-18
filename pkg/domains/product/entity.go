package product

import "time"

type Entity struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
	DeletedAt time.Time `json:"deletedAt" bson:"deleted_at"`
}
