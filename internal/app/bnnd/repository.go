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

func NewBNNDRepo() BNNDRepo {
	return &BNNDMongoRepo{}
}

// Connect to mongo db repo at given address
func (r *BNNDMongoRepo) Connect(ctx context.Context, addr, db string) error {
	// Attempt connection
	log.Println("Attemptign to connect to %s %s", addr, db)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.addr))
	if err != nil {
		log.Error(err)
		return err
	}

	// Check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	r.c = client.Database(addr).Collect(db)

	// TODO: Add indexes
	// idxm := mongo.IndexModel{
	// 	Keys: bson.M{
	// 		"" : 1,
	// 	}, Options: nil,
	// }
	// idxm.Options.SetBackground(true) // must run in background
	// _, err := r.c.Indexes().CreateOne(ctx, idxm)
	// if err != nil {
	// 	return err
	// }
}

// FindAll BananaDucks
func (r *BNNDMongoRepo) FindAll() ([]BananaDuck, error) {
	return []BananaDuck{}, nil
}

// Insert BananaDuck
func (r *BNNDMongoRepo) Insert(d BananaDuck) error {
	return nil
}
