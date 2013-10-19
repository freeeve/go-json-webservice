package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type User struct {
	Email    string `bson:"email"`
	Passhash []byte `bson:"passhash"`
	Confirm  string `bson:"confirm"`
}

func CreateUser(ucr UserCreateRequest) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Println("error connecting to mongo:", err)
		return err
	}

	users := session.DB("webapp").C("users")
	passhash, _ := bcrypt.GenerateFromPassword([]byte(ucr.Password), 10)
	user := User{Email: ucr.Email,
		Passhash: passhash,
		Confirm:  bson.NewObjectId().Hex(),
	}
	users.Insert(user)
	SendEmail(user.Email)
	return nil
}

func SendEmail(email string) {
	// TODO
}
