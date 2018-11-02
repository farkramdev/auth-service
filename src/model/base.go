package model

import (
	"gopkg.in/mgo.v2/bson"
)

// Base type provides datastore-based model
type Base struct {
	key string        `bson:"-"`
	ID  bson.ObjectId `bson:"_id" json:"id"`
}

// Key return datastore key or nil
func (x *Base) Key() *Base {
	return x
}

// SetKey sets key and id to new given key
func (x *Base) SetKey(key *Base) {
	x.key = key.key
	x.ID = key.ID
}
