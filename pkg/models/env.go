package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Env struct {
	context          context.Context
	client           *mongo.Client
	connectionString string
}

func NewEnv() *Env {
	return &Env{client: nil}
}

func (e *Env) Open(ctx context.Context, connectionString string) error {
	var err error

	e.context = ctx
	e.connectionString = connectionString

	clientOptions := options.Client().ApplyURI(e.connectionString)

	log.Printf("Connecting to MongoDB '%s' ...", e.connectionString)

	// create a new MongoDB client and connect to your running MongoDB server.
	// clientOptions is used to set the connection string and other driver settings.
	// Ref https://godoc.org/go.mongodb.org/mongo-driver/mongo/options
	//
	// Context is like a timeout or deadline that indicates when an operation should stop
	// running and return.
	//
	// Ref https://golang.org/pkg/context/
	if e.client, err = mongo.Connect(e.context, clientOptions); err != nil {
		return err
	}

	return e.Ping()
}

// Ping is used to test the connection to the MongoDB server.
func (e *Env) Ping() error {
	// Ensure that our MongoDB server was found and connected to successfully using the Ping() method.
	log.Printf("Pinging MongoDB instance '%s' ...", e.connectionString)
	return e.client.Ping(e.context, nil)
}

func (e *Env) Context() context.Context {
	return e.context
}

func (e *Env) Close(ctx context.Context) error {
	return e.client.Disconnect(ctx)
}

func (e *Env) Client() *mongo.Client {
	return e.client
}

func (e *Env) Database(name string) *mongo.Database {
	return e.client.Database(name)
}

func (e *Env) Collection(database, collection string) *mongo.Collection {
	return e.Database(database).Collection(collection)
}
