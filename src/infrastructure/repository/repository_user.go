package repository

import (
	"app/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryUser struct {
	DB *mongo.Database
}

func NewUserPostgres(DB *mongo.Database) *RepositoryUser {
	return &RepositoryUser{DB: DB}
}

func (u *RepositoryUser) GetByID(id string) (user *entity.EntityUser, err error) {

	filter := bson.D{{Key: "id", Value: id}}

	u.DB.Collection("users").FindOne(context.Background(), filter).Decode(&user)

	return user, err
}

func (u *RepositoryUser) GetByMail(email string) (user *entity.EntityUser, err error) {

	filter := bson.D{{Key: "email", Value: email}}

	err = u.DB.Collection("users").FindOne(context.Background(), filter).Decode(&user)

	return user, err
}

func (u *RepositoryUser) CreateUser(user *entity.EntityUser) error {

	newID, _ := primitive.NewObjectID().MarshalText()

	user.ID = string(newID)

	_, err := u.DB.Collection("users").InsertOne(context.Background(), user)

	return err
}

func (u *RepositoryUser) UpdateUser(user *entity.EntityUser) error {

	_, err := u.GetByMail(user.Email)

	if err != nil {
		return err
	}

	filter := bson.D{{Key: "email", Value: user.Email}}
	update := bson.D{{Key: "$set", Value: user}}

	_, err = u.DB.Collection("users").UpdateOne(context.Background(), filter, update)

	return err
}

func (u *RepositoryUser) DeleteUser(user *entity.EntityUser) error {

	_, err := u.GetByMail(user.Email)

	if err != nil {
		return err
	}

	filter := bson.D{{Key: "email", Value: user.Email}}

	_, err = u.DB.Collection("users").DeleteOne(context.Background(), filter)

	return err
}

func (u *RepositoryUser) GetUsersFromIDs(ids []string) (users []entity.EntityUser, err error) {
	users = make([]entity.EntityUser, 0)

	filter := bson.D{{Key: "id", Value: bson.D{{Key: "$in", Value: ids}}}}

	cursor, err := u.DB.Collection("users").Find(context.Background(), filter)

	if err != nil {
		return users, err
	}

	if err = cursor.All(context.Background(), &users); err != nil {
		return users, err
	}

	return users, err
}

func (u *RepositoryUser) GetUsers(filters entity.EntityUserFilters) (users []entity.EntityUser, err error) {

	users = make([]entity.EntityUser, 0)

	filter := bson.D{}

	if filters.Search != "" {
		// DBFind = DBFind.Where("name LIKE ? or email LIKE ?", "%"+filters.Search+"%", "%"+filters.Search+"%")
		filter = append(filter, bson.E{Key: "$or", Value: bson.A{
			bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: filters.Search}}}},
			bson.D{{Key: "email", Value: bson.D{{Key: "$regex", Value: filters.Search}}}},
		}})
	}

	if filters.Active != "" {
		// DBFind = DBFind.Where("active = ?", filters.Active)
		filter = append(filter, bson.E{Key: "active", Value: filters.Active})
	}

	cursor, err := u.DB.Collection("users").Find(context.Background(), filter)

	if err != nil {
		return users, err
	}

	if err = cursor.All(context.Background(), &users); err != nil {
		return users, err
	}

	return users, err
}

func (u *RepositoryUser) GetUser(id string) (user *entity.EntityUser, err error) {
	filter := bson.D{{Key: "id", Value: id}}

	u.DB.Collection("users").FindOne(context.Background(), filter).Decode(&user)

	return user, err
}
