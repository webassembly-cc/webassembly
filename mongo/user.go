package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//用户相关的
type User struct {
	UUID     string
	Name     string
	Password string
}

var user *mgo.Collection

func init() {
	initMongo()
	user = db.C("user")
}

//Add user
func (m *User) Insert() (err error) {
	err = user.Insert(m)
	return
}

//find by UUID
func (m *User) GetByUUID() error {
	return user.Find(bson.M{"UUID": m.UUID}).One(m)
}

//find by Name
func (m *User) GetByName() error {
	return user.Find(bson.M{"Name": m.Name}).One(m)
}
