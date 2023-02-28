package data_layer

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type ResourceModel struct {
	DB *mongo.Client
}

func (m ResourceModel) InsertResource() error {

	return errors.New("memes")
}
