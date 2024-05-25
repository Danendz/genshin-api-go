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

	if len(dbname) == 0 {
		dbname = "genshin-api"
	}

	dburi := os.Getenv("DBURI")

	if len(dburi) == 0 {
		dburi = "mongodb://localhost:27017"
	}

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

func MakeUpdateFormat(values *bson.M) *bson.D {
	return &bson.D{
		{
			Key: "$set", Value: values,
		},
	}
}