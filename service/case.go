package service

import (
	"database/sql"
	"log/slog"
	pb "question/genproto/question"
	"question/storage"
)

type CaseService struct {
	pb.UnimplementedTestCaseServiceServer
	Case   storage.Istorage
	Logger *slog.Logger
}

func NewCaseService(db *sql.DB, Logger *slog.Logger, istorage storage.Istorage) *CaseService {
	return &CaseService{
		Case:   istorage,
		Logger: Logger,
	}
}
