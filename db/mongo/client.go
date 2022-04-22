package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/go-swagger/go-swagger/examples/accountmgr/db"
	"github.com/go-swagger/go-swagger/examples/accountmgr/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	db.RegisterDataStore("mongo", NewClient)
}
func NewClient() (db.DataStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client1, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("The error is", err)
		return nil, err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client1.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("The error is", err)
		return &client{dbc: nil}, err
	}
	return &client{dbc: client1.Database("account_app")}, nil
}

type client struct {
	dbc *mongo.Database
}

func (c *client) AddAccount(account *domain.Account) (string, error) {
	collection := c.dbc.Collection("accounts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := collection.ReplaceOne(ctx,
		bson.D{{Key: "_id", Value: account.ID}},
		account,
		options.Replace().SetUpsert(true)); err != nil {
		fmt.Println("The error is", err)
		return "", err
	}

	return account.ID, nil
}
