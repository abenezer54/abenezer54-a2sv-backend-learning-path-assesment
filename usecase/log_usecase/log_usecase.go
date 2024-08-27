package log_usecase

import (
	"context"
	"loan-api/domain"
)

type logUsecase struct {
	logRepo domain.LogRepository
}

func NewLogUsecase(logRepo domain.LogRepository) *logUsecase {
	return &logUsecase{
		logRepo: logRepo,
	}
}

func (uc *logUsecase) GetLogs(ctx context.Context, eventType string, order string) ([]domain.Log, error) {
	return uc.logRepo.GetLogs(ctx, eventType, order)
}
