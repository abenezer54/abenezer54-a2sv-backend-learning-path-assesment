package log_repository

import (
	"context"
	"loan-api/domain"
	"loan-api/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type logRepository struct {
	collection mongo.Collection
}

func NewLogRepository(collection mongo.Collection) *logRepository {
	return &logRepository{
		collection: collection,
	}
}

func (r *logRepository) StoreLog(ctx context.Context, log *domain.Log) error {
	_, err := r.collection.InsertOne(ctx, log)
	return err
}

func (r *logRepository) GetLogs(ctx context.Context, eventType string, order string) ([]domain.Log, error) {
	filter := bson.M{}
	if eventType != "" {
		filter["event"] = eventType
	}

	sortOrder := 1 // Ascending by default
	if order == "desc" {
		sortOrder = -1
	}

	findOptions := options.Find().SetSort(bson.D{{"timestamp", sortOrder}})
	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []domain.Log
	if err := cursor.All(ctx, &logs); err != nil {
		return nil, err
	}

	return logs, nil
}
