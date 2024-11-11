package repositories

import (
	"context"
	"fmt"

	"github.com/dannyoka/memorize-it-api/internal/data"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EntryRepository struct {
	coll *mongo.Collection
}

type IEntryRepository interface {
	GetEntry(id string) data.Entry
	CreateEntry()
	GetEntries() []data.Entry
}

func NewEntryRepository(db *mongo.Database) *EntryRepository {
	return &EntryRepository{
		coll: db.Collection("entries"),
	}
}

func (repo *EntryRepository) GetEntry(id string) data.Entry {
	var result data.Entry
	repo.coll.FindOne(context.Background(), bson.D{{"id", id}}).Decode(&result)
	return result
}

func (repo *EntryRepository) CreateEntry(entry data.EntryPayload) data.Entry {
	// create entry in db
	newEntry := data.Entry{
		Id:      uuid.New().String(),
		Name:    entry.Name,
		Content: entry.Content,
	}
	_, err := repo.coll.InsertOne(context.Background(), newEntry)
	if err != nil {
		fmt.Println("error", err)
	}
	return newEntry
}

func (repo *EntryRepository) GetEntries() []data.Entry {
	cursor, err := repo.coll.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println("error", err)
		return nil
	}
	defer cursor.Close(context.Background())
	var results []data.Entry
	cursor.All(context.Background(), &results)
	return results
}
