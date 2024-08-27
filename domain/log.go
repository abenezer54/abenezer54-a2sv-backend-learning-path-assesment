package domain

import (
	"context"
	"time"
)

type Log struct {
	ID        string    `json:"id"`
	Event     string    `json:"event"`
	Details   string    `json:"details"`
	Timestamp time.Time `json:"timestamp"`
	UserID    string    `json:"user_id"`
}

type LogRepository interface {
	StoreLog(ctx context.Context, log *Log) error
	GetLogs(ctx context.Context, eventType string, order string) ([]Log, error)
}

type LogUsecase interface {
	GetLogs(ctx context.Context, eventType string, order string) ([]Log, error)
}
