package service

import (
	"database/sql"
	"log/slog"
	pb "question/genproto/question"
	"question/storage"
)

type OutputService struct {
	pb.UnimplementedOutputServiceServer
	Output storage.Istorage
	Logger *slog.Logger
}

func NewOutputService(db *sql.DB, Logger *slog.Logger, istorage storage.Istorage) *OutputService {
	return &OutputService{
		Output: istorage,
		Logger: Logger,
	}
}
