package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogEntry struct {
	ID        string    `json:"id.omitempty" bson:"_id.omitempty"`
	Name      string    `json:"name" bson:"name"`
	Data      string    `json:"data" bson:"data"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Models struct {
	LogEntry LogEntry
}

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		LogEntry: LogEntry{},
	}
}

func (l *LogEntry) Insert(entry LogEntry) error {
	connection := client.Database("mydb").Collection("logs")

	_, err := connection.InsertOne(context.TODO(), LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting log entry:", err)

		return err
	}

	return nil
}

func (l *LogEntry) All() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	connection := client.Database("mydb").Collection("logs")
	opts := options.Find()
	opts.SetSort(bson.D{bson.E{Key: "created_at", Value: -1}})

	cursor, err := connection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Error finding logs:", err)

		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []*LogEntry

	for cursor.Next(ctx) {
		var item LogEntry

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("Error decoding log into slice:", err)

			return nil, err
		}

		logs = append(logs, &item)
	}

	return logs, nil
}

func (l *LogEntry) GetOne(id string) (*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("mydb").Collection("logs")

	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var entry LogEntry
	err = collection.FindOne(ctx, bson.M{"_id": docId}).Decode(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (l *LogEntry) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("mydb").Collection("logs")

	if err := collection.Drop(ctx); err != nil {
		return err
	}

	return nil
}

func (l *LogEntry) Update() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("mydb").Collection("logs")

	docId, err := primitive.ObjectIDFromHex(l.ID)
	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docId},
		bson.D{
			bson.E{Key: "$set", Value: bson.D{
				bson.E{Key: "name", Value: l.Name},
				bson.E{Key: "data", Value: l.Data},
				bson.E{Key: "updated_at", Value: time.Now()},
			}}},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
