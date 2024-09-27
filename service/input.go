package service

import (
	"context"
	"log/slog"
	pb "question/genproto/question"
	"question/storage"
)

type InputService struct {
	pb.UnimplementedInputServiceServer
	Input  storage.Istorage
	Logger *slog.Logger
}

func NewInputService(Logger *slog.Logger, istorage storage.Istorage) *InputService {
	return &InputService{
		Input:  istorage,
		Logger: Logger,
	}
}

func (s *InputService) CreateQuestionInput(ctx context.Context, req *pb.CreateQuestionInputRequest) (*pb.QuestionInputId, error) {
	s.Logger.Info("CreateQuestionInput rpc method is started")
	res, err := s.Input.Input().CreateQuestionInput(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("CreateQuestionInput finished succesfully")
	return res, nil
}

func (s *InputService) GetQuestionInput(ctx context.Context, req *pb.QuestionInputId) (*pb.GetQuestionInputResponse, error) {
	s.Logger.Info("GetQuestionInput rpc method is started")
	res, err := s.Input.Input().GetQuestionInput(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("GetQuestionInput finished succesfully")
	return res, nil
}

func (s *InputService) GetAllQuestionInputsByQuestionId(ctx context.Context, req *pb.GetAllQuestionInputsByQuestionIdRequest) (*pb.GetAllQuestionInputsByQuestionIdResponse, error) {
	s.Logger.Info("GetAllQuestionInputsByQuestionId rpc method is started")
	res, err := s.Input.Input().GetAllQuestionInputsByQuestionId(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("GetAllQuestionInputsByQuestionId finished succesfully")
	return res, nil
}

func (s *InputService) UpdateQuestionInput(ctx context.Context, req *pb.UpdateQuestionInputRequest) (*pb.Void, error) {
	s.Logger.Info("UpdateQuestionInput rpc method is started")
	res, err := s.Input.Input().UpdateQuestionInput(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("UpdateQuestionInput finished succesfully")
	return res, nil
}

func (s *InputService) DeleteQuestionInput(ctx context.Context, req *pb.DeleteQuestionInputRequest) (*pb.Void, error) {
	s.Logger.Info("DeleteQuestionInput rpc method is started")
	res, err := s.Input.Input().DeleteQuestionInput(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("DeleteQuestionInput finished succesfully")
	return res, nil
}
