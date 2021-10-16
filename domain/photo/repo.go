package photo

import (
	"github.com/m0thm4n/UploadService/mongo"
	"github.com/globalsign/mgo/bson"
)

type Repository struct {
	mongoClient *mongo.Client
}

func NewRepository(mongoClient *mongo.Client) *Repository {
	repo := Repository{mongoClient}
	return &repo
}

const (
	databaseName = "photos"
	collectionName = "photos"

)

func (r *Repository) GetById(id int) (*Photo, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()

	query := bson.M{"Id": id}

	var photo *Photo
	err := session.DB(databaseName).C(collectionName).Find(query).One(&photo)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (r *Repository) Insert(photo *Photo) error {
	var session = r.mongoClient.NewSession()
	defer session.Close()

	err := session.DB(databaseName).C(collectionName).Insert(photo)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(id int) error {
	var session = r.mongoClient.NewSession()
	defer session.Close()

	query := bson.M{"Id": id}

	err := session.DB(databaseName).C(collectionName).Remove(query)
	if err != nil {
		return err
	}

	return nil
}