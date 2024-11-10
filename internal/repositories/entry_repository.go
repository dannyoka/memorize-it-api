package repositories

import (
	"context"

	"github.com/dannyoka/memorize-it-api/internal/data"
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

func (repo *EntryRepository) CreateEntry(entry data.Entry) {
	// create entry in db
}

func (repo *EntryRepository) GetEntries() []data.Entry {
	var results []data.Entry
	cursor, err := repo.coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result data.Entry
		cursor.Decode(&result)
		results = append(results, result)
	}
	return results
}
