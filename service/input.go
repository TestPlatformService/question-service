package service

import (
	"database/sql"
	"log/slog"
	pb "question/genproto/question"
	"question/storage"
)

type InputService struct {
	pb.UnimplementedInputServiceServer
	Input  storage.Istorage
	Logger *slog.Logger
}

func NewInputService(db *sql.DB, Logger *slog.Logger, istorage storage.Istorage) *InputService {
	return &InputService{
		Input:  istorage,
		Logger: Logger,
	}
}
