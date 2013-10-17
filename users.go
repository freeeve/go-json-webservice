package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/rand"
	"io"
	"labix.org/v2/mgo"
	"log"
)

type User struct {
	Email       string `bson:"email"`
	Passhash    []byte `bson:"passhash"`
	Confirmed   bool   `bson:"confirmed"`
	ConfirmHash string `bson:"confirmHash"`
}

func CreateUser(ucr UserCreateRequest) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Println("error connecting to mongo:", err)
		return err
	}

	users := session.DB("webapp").C("users")
	passhash, _ := bcrypt.GenerateFromPassword([]byte(ucr.Password), 10)
	len := 32
	hash := make([]byte, len)
	io.ReadFull(rand.Reader, hash)
	user := User{Email: ucr.Email,
		Passhash:    passhash,
		Confirmed:   false,
		ConfirmHash: string(hash),
	}
	users.Insert(user)
	return nil
}
