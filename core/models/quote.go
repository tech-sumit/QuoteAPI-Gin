package models

import "time"

//Quote Model
type Quote struct {
	Author    string    `json:"author,omitempty" bson:"author"`
	Quote     string    `json:"quote,omitempty" bson:"quote"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
}
