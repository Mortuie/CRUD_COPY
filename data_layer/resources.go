package data_layer

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mortuie/CRUD_COPY/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResourceModel struct {
	DB *mongo.Client
}

func (m ResourceModel) InsertResource() error {

	return errors.New("memes")
}

func (m ResourceModel) CreateCollection(name string) error {
	db := m.DB.Database(constants.RESOURCE_DB)

	err := db.CreateCollection(context.TODO(), name)

	if err != nil {
		if strings.Contains(err.Error(), "Collection already exists.") {
			fmt.Printf("Collection: %s already exists.\n", name)
		} else {
			fmt.Println("ERROR CREATING COLLECTION", err)
			return err
		}
	}
	return nil
}
