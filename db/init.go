package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func NewMongoDB(creds *Creds) (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(creds.DBURI).SetAuth(creds.DBCREDS))
}

func NewPGDB() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}
