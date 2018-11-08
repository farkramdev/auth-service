package api

import (
	"auth-service/users/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserRepository dddd UserRepository
type UserRepository struct {
	C *mgo.Collection
}

// FindUser from store
func (r *UserRepository) FindUser(username, password string) (*model.User, error) {
	// var user model.User
	// err = r.C.Find({username: username, password: password})
	// q := datastore.
	// 	NewQuery(kindUser).
	// 	Filter("Username =", username).
	// 	Limit(1)
	// key, err := client.Run(ctx, q).Next(&user)
	// if err == iterator.Done {
	// 	// Not found
	// 	return nil, nil
	// }
	// if err != nil {
	// 	return nil, err
	// }
	// user.SetKey(key)
	// if !user.ComparePassword(password) {
	// 	// wrong password return like user not found
	// 	return nil, nil
	// }
	// return &user, nil
	return nil, nil
}

// SaveUser to datastore
func (r *UserRepository) SaveUser(user *model.User) error {

	user.ID = bson.NewObjectId()
	user.Stamp()
	// user.SetKey(key)
	err := r.C.Insert(&user)
	return err
}
