package db

import (
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBCreds struct {
	DBNAME  string
	DBURI   string
	DBCREDS options.Credential
}

func NewDBCreds() *DBCreds {
	dbname := os.Getenv("DBNAME")
	dburi := os.Getenv("DBURI")
	username := os.Getenv("DBUSERNAME")
	password := os.Getenv("DBPASSWORD")

	if len(username) == 0 || len(password) == 0 {
		username = "dev"
		password = "root"
	}

	return &DBCreds{
		DBNAME: dbname,
		DBURI: dburi,
		DBCREDS: options.Credential{
			Username: username,
			Password: password,
		},
	}
}

func ToObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

func MakeUpdateFormat(values interface{}) *bson.D {
	return &bson.D{
		{
			Key: "$set", Value: values,
		},
	}
}