package service

import (
	"context"
	"log/slog"
	pb "question/genproto/question"
	"question/storage"
)

type QuestionService struct {
	pb.UnimplementedQuestionServiceServer
	Question storage.Istorage
	Logger   *slog.Logger
}

func NewQuestionService(Logger *slog.Logger, istorage storage.Istorage) *QuestionService {
	return &QuestionService{
		Question: istorage,
		Logger:   Logger,
	}
}

func (s *QuestionService) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.QuestionId, error) {
	s.Logger.Info("CreateQuestion rpc method is started")
	res, err := s.Question.Question().CreateQuestion(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("CreateQuestion finished succesfully")
	return res, nil
}

func (s *QuestionService) GetQuestion(ctx context.Context, req *pb.QuestionId) (*pb.GetQuestionResponse, error) {
	s.Logger.Info("GetQuestion rpc method is started")
	res, err := s.Question.Question().GetQuestion(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("GetQuestion finished succesfully")
	return res, nil
}

func (s *QuestionService) GetAllQuestions(ctx context.Context, req *pb.GetAllQuestionsRequest) (*pb.GetAllQuestionsResponse, error) {
	s.Logger.Info("GetAllQuestions rpc method is started")
	res, err := s.Question.Question().GetAllQuestions(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("GetAllQuestions finished succesfully")
	return res, nil
}

func (s *QuestionService) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.Void, error) {
	s.Logger.Info("UpdateQuestion rpc method is started")
	res, err := s.Question.Question().UpdateQuestion(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("UpdateQuestion finished succesfully")
	return res, nil
}

func (s *QuestionService) DeleteQuestion(ctx context.Context, req *pb.DeleteQuestionRequest) (*pb.Void, error) {
	s.Logger.Info("DeleteQuestion rpc method is started")
	res, err := s.Question.Question().DeleteQuestion(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("DeleteQuestion finished succesfully")
	return res, nil
}

func (s *QuestionService) UploadImageQuestion(ctx context.Context, req *pb.UploadImageQuestionRequest) (*pb.Void, error) {
	s.Logger.Info("UploadImageQuestion rpc method is started")
	res, err := s.Question.Question().UploadImageQuestion(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("UploadImageQuestion finished succesfully")
	return res, nil
}

func (s *QuestionService) DeleteImageQuestion(ctx context.Context, req *pb.DeleteImageQuestionRequest) (*pb.Void, error) {
	s.Logger.Info("DeleteImageQuestion rpc method is started")
	res, err := s.Question.Question().DeleteImageQuestion(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("DeleteImageQuestion finished succesfully")
	return res, nil
}

func (s *QuestionService) IsQuestionExist(ctx context.Context, req *pb.QuestionId) (*pb.Void, error) {
	s.Logger.Info("IsQuestionExist rpc method is started")
	res, err := s.Question.Question().IsQuestionExist(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("IsQuestionExist finished succesfully")
	return res, nil
}
