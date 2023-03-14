package repository

import (
	"context"
	"time"

	"github.com/saweima12/secretary-bot/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMsgRepository interface {
	InsertOne(msg entity.UserMsg) (*mongo.InsertOneResult, error)
	FetchOne(userId string) *mongo.SingleResult
}

type userMsgRepository struct {
	collection *mongo.Collection
}

func NewBotMessageRepository(collection *mongo.Collection) UserMsgRepository {
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	// create collection index.
	collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.M{"userid": 1},
			Options: options.Index().SetUnique(true),
		},
	}, opts)
	return &userMsgRepository{collection: collection}
}

func (r *userMsgRepository) InsertOne(msg entity.UserMsg) (*mongo.InsertOneResult, error) {
	result, err := r.collection.InsertOne(context.Background(), msg)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userMsgRepository) FetchOne(userId string) *mongo.SingleResult {
	filter := bson.M{
		"userId": userId,
	}
	return r.collection.FindOne(context.Background(), filter)
}
