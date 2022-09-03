package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var DB *DataBase

type DataBase struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func (db *DataBase) Course() *mongo.Collection {
	return db.DB.Collection("course")
}

func (db *DataBase) Info() *mongo.Collection {
	return db.DB.Collection("info")
}

// Connect connects to the mongodb
func Connect() {
	url := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbName)
	DB = &DataBase{
		Client: client,
		DB:     db,
	}
}
