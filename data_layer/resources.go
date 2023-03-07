package data_layer

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mortuie/CRUD_COPY/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResourceModel struct {
	DB *mongo.Client
}

func (m ResourceModel) GetResource(collectionName string, id string) (map[string]interface{}, error) {
	db := m.DB.Database(constants.RESOURCE_DB)
	collection := db.Collection(collectionName)

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("Invalid resource id", id)
		return nil, errors.New("Invalid resource id: " + id)
	}

	res := collection.FindOne(context.TODO(), bson.M{"_id": objectId})
	anyJson := map[string]interface{}{}
	res.Decode(anyJson)
	return anyJson, nil
}

func (m ResourceModel) InsertResource(collectionName string, resource map[string]interface{}) (map[string]interface{}, error) {
	db := m.DB.Database(constants.RESOURCE_DB)
	collection := db.Collection(collectionName)

	b, err := collection.InsertOne(context.TODO(), resource)

	if err != nil {
		return nil, err
	}

	fmt.Println("ID: ", b.InsertedID.(primitive.ObjectID).Hex())

	idString := b.InsertedID.(primitive.ObjectID).Hex()

	resource["_id"] = idString

	return resource, nil
}

func (m ResourceModel) CreateCollection(collectionName string) error {
	db := m.DB.Database(constants.RESOURCE_DB)

	err := db.CreateCollection(context.TODO(), collectionName)

	if err != nil {
		if strings.Contains(err.Error(), "Collection already exists.") {
			fmt.Printf("Collection: %s already exists.\n", collectionName)
		} else {
			fmt.Println("ERROR CREATING COLLECTION", err)
			return err
		}
	}
	return nil
}
