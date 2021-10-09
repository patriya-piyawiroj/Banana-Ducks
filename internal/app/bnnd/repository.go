package bnnd

import (
	log "github.com/sirupsen/logrus"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BNNDRepo interface {
	FindAll() ([]BananaDuck, error)
	Insert(d BananaDuck) error
}

type BNNDMongoRepo struct {
	c *mongo.Collection
}

// NewBNNDRepo
func NewBNNDRepo() BNNDRepo {
	return &BNNDMongoRepo{}
}

// Connect to mongo db repo at given address
func (r *BNNDMongoRepo) Connect(ctx context.Context, addr, db string) error {
	// Attempt connection
	log.Println("Attemptign to connect to %s %s", addr, db)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(addr))
	if err != nil {
		log.Error(err)
		return err
	}

	// Check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	r.c = client.Database(addr).Collection(db)

	// TODO: Work on db indexing and performance
	return nil
}

// FindAll BananaDucks from db
func (r *BNNDMongoRepo) FindAll() ([]BananaDuck, error) {
	return []BananaDuck{}, nil
}

// Insert BananaDuck to db
func (r *BNNDMongoRepo) Insert(d BananaDuck) error {
	return nil
}
