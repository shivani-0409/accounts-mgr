package domain

import (
	"time"
)

type Account struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	UserID    int32     `bson:"user_id"`
	UserName  string    `bson:"user_name"`
	Source    string    `bson:"source"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
