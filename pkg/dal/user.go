package dal

import (
	"fmt"
	"github.com/ocrosby/go-mongo-tutorial/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type UserModel struct {
	env *models.Env
}

func NewUserModel(env *models.Env) *UserModel {
	return &UserModel{env: env}
}

// Create a new user
func (u *UserModel) Create(user *models.User) (*models.User, error) {

	collection := u.env.Client().Database("tasker").Collection("users")

	insertResult, err := collection.InsertOne(u.env.Context(), user)
	if err != nil {
		return nil, err
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)

	return user, nil
}

// FetchById returns a single user by id
func (u *UserModel) FetchById(userId int) (*models.User, error) {
	return nil, fmt.Errorf("not implemented")
}

// All returns a slice of all users
func (u *UserModel) All() ([]*models.User, error) {
	var (
		err   error
		users []*models.User
	)

	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{{}}
	ctx := u.env.Context()
	client := u.env.Client()

	collection := client.Database("tasker").Collection("users")

	findOptions := options.Find()
	// findOptions.SetLimit(2)

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user models.User

		if err = cur.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	if err = cur.Close(u.env.Context()); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("no documents found")
	}

	return users, nil
}

func (u *UserModel) Update() (*models.User, error) {
	return nil, fmt.Errorf("not implemented")
}

func (u *UserModel) Delete() error {
	return fmt.Errorf("not implemented")
}
