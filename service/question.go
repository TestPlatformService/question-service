package service

import (
	"database/sql"
	"log/slog"
	pb "question/genproto/question"
	"question/storage"
)

type QuestionService struct {
	pb.UnimplementedQuestionServiceServer
	Question storage.Istorage
	Logger   *slog.Logger
}

func NewQuestionService(db *sql.DB, Logger *slog.Logger, istorage storage.Istorage) *QuestionService {
	return &QuestionService{
		Question: istorage,
		Logger:   Logger,
	}
}
