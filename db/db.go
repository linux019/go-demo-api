package db

import (
	log "api-demo/apilogger"
	"api-demo/constants"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var mongoClient *mongo.Client
var mongoDB *mongo.Database

type User struct {
	Email  string
	Passwd string
}

func ConnectDB(dbUser, dbPassw *string) {
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.Client()

	opts.SetAuth(options.Credential{
		Username: *dbUser,
		Password: *dbPassw,
	})

	if mongoClient, err = mongo.Connect(ctx, opts.ApplyURI(constants.DatabaseAddress)); err != nil {
		log.Logger.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		_ = mongoClient.Disconnect(nil)
		log.Logger.Fatal("Connection failed", err)
	}

	mongoDB = mongoClient.Database(constants.DatabaseName)
	log.Logger.Info("Connected to MongoDB")
}

func DisconnectDB() {
	if mongoClient != nil {
		_ = mongoClient.Disconnect(nil)
		mongoClient = nil
		mongoDB = nil
	}
}

// ok - user created, false - user with such email already exists
func CreateUser(email string, password string) bool {
	if _, ok := GetUser(email); ok {
		return false
	}

	if _, err := mongoDB.
		Collection(constants.DbAccountsSchema).
		InsertOne(nil, User{
			Email:  email,
			Passwd: password,
		}); err != nil {
		log.Logger.Fatal("Insert user", err)
	}
	return true
}

func GetUser(email string) (*User, bool) {
	var result User

	userFilter := bson.D{{"email", email}}

	err := mongoDB.
		Collection(constants.DbAccountsSchema).
		FindOne(nil, userFilter).
		Decode(&result)

	switch err {
	case mongo.ErrNoDocuments:
		return nil, false
	case nil:
		return &result, true
	default:
		log.Logger.Fatal("Get user", err)
	}

	return nil, false
}
