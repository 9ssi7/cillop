package product

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/ssibrahimbas/cillop/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, name string) (string, *i18np.Error)
	Get(ctx context.Context, id string) (*Entity, *i18np.Error)
}

type repo struct {
	db mongodb.Collection
}

func NewRepository(db mongodb.Collection) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, name string) (string, *i18np.Error) {
	res, err := r.db.InsertOne(ctx, &Entity{
		Name:      name,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return "", i18np.NewError(messages.Failed)
	}
	return res.InsertedID.(string), nil
}

func (r *repo) Get(ctx context.Context, id string) (*Entity, *i18np.Error) {
	var entity Entity
	oid, err := mongodb.TransformId(id)
	if err != nil {
		return nil, i18np.NewError(messages.InvalidID)
	}
	res := r.db.FindOne(ctx, bson.M{"_id": oid})
	if error := res.Err(); error != nil {
		if error == mongo.ErrNoDocuments {
			return nil, i18np.NewError(messages.NotFound)
		}
		return nil, i18np.NewError(messages.Failed)
	}
	if error := res.Decode(&entity); error != nil {
		return nil, i18np.NewError(messages.Failed)
	}
	return &entity, nil
}
