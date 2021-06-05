package repository

import (
	"context"
	"fmt"
	"log"
	"server/models"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository interface {
	Create(models.Task) primitive.M
	GetAll() primitive.M
	Complete(string) primitive.M
	Undo(string) primitive.M
	Delete(string) primitive.M
}

type TaskRepositoryImpl struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) TaskRepository {
	return &TaskRepositoryImpl{collection}
}

var errorResult = primitive.M{
	"error": true,
}

func (r TaskRepositoryImpl) Create(task models.Task) primitive.M {
	insertResult, err := r.collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
		return errorResult
	}

	fmt.Println("Inserted a single Record ", insertResult.InsertedID.(primitive.ObjectID).Hex())
	return primitive.M{
		"error": false,
	}
}

func (r TaskRepositoryImpl) GetAll() primitive.M {
	cur, err := r.collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
		return errorResult
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
			return errorResult
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return errorResult
	}

	fmt.Println("Records retrieved: ", len(results))

	cur.Close(context.Background())
	return primitive.M{
		"error":  false,
		"result": results,
	}
}

func (r TaskRepositoryImpl) Complete(taskId string) primitive.M {
	fmt.Println(taskId)
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return errorResult
	}

	fmt.Println("modified count: ", result.ModifiedCount)
	return primitive.M{
		"error": false,
	}
}

func (r TaskRepositoryImpl) Undo(taskId string) primitive.M {
	fmt.Println(taskId)
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return errorResult
	}

	fmt.Println("modified count: ", result.ModifiedCount)
	return primitive.M{
		"error": false,
	}
}

func (r TaskRepositoryImpl) Delete(taskId string) primitive.M {
	fmt.Println(taskId)
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	d, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return errorResult
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return primitive.M{
		"error": false,
	}
}
