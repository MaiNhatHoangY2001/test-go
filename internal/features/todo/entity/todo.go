package entity

import (
"time"

"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
ID          primitive.ObjectID `bson:"_id,omitempty"`
Title       string             `bson:"title"`
Description string             `bson:"description"`
Completed   bool               `bson:"completed"`
CreatedAt   time.Time          `bson:"created_at"`
UpdatedAt   time.Time          `bson:"updated_at"`
}
