package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	c   *mongo.Client
	db  *mongo.Database
	ctx context.Context
}

func New(uri string, dbName string) (*DB, error) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	return &DB{
		c:   client,
		db:  db,
		ctx: ctx,
	}, nil
}

func (m *DB) GetCollection(n string) *mongo.Collection {
	return m.db.Collection(n)
}

func TransformId(id string) (primitive.ObjectID, error) {
	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return i, nil
}

func TransformIds(ids []string) ([]primitive.ObjectID, error) {
	oids := []primitive.ObjectID{}
	for _, id := range ids {
		i, err := TransformId(id)
		if err != nil {
			return nil, err
		}
		oids = append(oids, i)
	}
	return oids, nil
}
