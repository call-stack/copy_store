package copyrepo

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/call-stack/copy_store.git/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct {
	mongoClient *mongo.Client
}

func NewRepo() *repo {
	mongoURI := fmt.Sprintf("mongodb+srv://dev-user:%s@copystore.hcale.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", os.Getenv("MONGODB_PASSWORD"))

	mongoClient, _ := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	err := mongoClient.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return &repo{mongoClient: mongoClient}
}

func (r *repo) Get(id string) (domain.Store, error) {
	store := domain.Store{}
	collection := r.mongoClient.Database("content").Collection("data_store")

	err := collection.FindOne(context.TODO(), bson.M{"url": id}).Decode(&store)
	if err != nil {
		return domain.Store{Content: ""}, nil
	}

	return store, nil
}

func (r *repo) Create(item domain.Store) error {
	bson_content, err := bson.Marshal(item)
	if err != nil {
		log.Println("Error in marshalling.")
		log.Fatal()
	}

	collection := r.mongoClient.Database("content").Collection("data_store")
	_, insertErr := collection.InsertOne(context.TODO(), bson_content)
	if insertErr != nil {
		log.Println("Error in storing the content.")
		log.Fatal()
	}
	return nil
}
