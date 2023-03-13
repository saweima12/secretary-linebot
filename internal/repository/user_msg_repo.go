package repository

import "go.mongodb.org/mongo-driver/mongo"

type UserMsgRepository interface {
	InsertOne()
	FetchOne()
}

type userMsgRepository struct {
	collection *mongo.Collection
}

func NewBotMessageRepository(collection *mongo.Collection) UserMsgRepository {
	return &userMsgRepository{collection: collection}
}

func (r *userMsgRepository) InsertOne() {
}

func (r *userMsgRepository) FetchOne() {

}
