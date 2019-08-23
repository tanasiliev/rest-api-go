package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Admin     bool               `bson:"admin"`
	Approved  bool               `bson:"approved"`
	CreatedAt primitive.DateTime `bson:"created_at"`
}

type UserCollection struct {
	database   string
	collection string
	col        *mongo.Collection
}

func (uc *UserCollection) Connect(client *mongo.Client) {
	uc.col = client.Database(uc.database).Collection(uc.collection)
}

func (uc *UserCollection) FindAll() ([]User, error) {
	var users []User

	ctx := context.Background()
	cur, err := uc.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var user User
		err = cur.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (uc *UserCollection) FindOne(id primitive.ObjectID) (User, error) {
	var user User
	filter := bson.M{"_id": id}
	err := uc.col.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

func (uc *UserCollection) InsertOne(user User) (*mongo.InsertOneResult, error) {

	user.Id = primitive.NewObjectID()
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	insertResult, err := uc.col.InsertOne(context.Background(), user)
	return insertResult, err
}

func (uc *UserCollection) UpdateOne(id primitive.ObjectID, user User) (*mongo.UpdateResult, error) {

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"username": user.Username, "admin": user.Admin, "approved": user.Approved}}

	result, err := uc.col.UpdateOne(context.Background(), filter, update)
	return result, err
}

func (uc *UserCollection) DeleteOne(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	res, err := uc.col.DeleteOne(context.Background(), filter)
	return res, err
}
