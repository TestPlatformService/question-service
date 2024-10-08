package service

import (
	"context"
	"log/slog"
	pb "question/genproto/question"
	"question/storage"
)

type OutputService struct {
	pb.UnimplementedOutputServiceServer
	Output storage.Istorage
	Logger *slog.Logger
}

func NewOutputService(Logger *slog.Logger, istorage storage.Istorage) *OutputService {
	return &OutputService{
		Output: istorage,
		Logger: Logger,
	}
}

func (s *OutputService) CreateQuestionOutput(ctx context.Context, req *pb.CreateQuestionOutputRequest) (*pb.QuestionOutputId, error) {
	s.Logger.Info("CreateQuestionOutput rpc method is started")
	res, err := s.Output.Output().CreateQuestionOutput(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("CreateQuestionOutput finished succesfully")
	return res, nil
}

func (s *OutputService) GetQuestionOutput(ctx context.Context, req *pb.QuestionOutputId) (*pb.GetQuestionOutputResponse, error) {
	s.Logger.Info("GetQuestionOutput rpc method is started")
	res, err := s.Output.Output().GetQuestionOutput(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("GetQuestionOutput finished succesfully")
	return res, nil
}

func (s *OutputService) GetAllQuestionOutputsByQuestionId(ctx context.Context, req *pb.GetAllQuestionOutputsByQuestionIdRequest) (*pb.GetAllQuestionOutputsByQuestionIdResponse, error) {
	s.Logger.Info("GetAllQuestionOutputsByQuestionId rpc method is started")
	res, err := s.Output.Output().GetAllQuestionOutputsByQuestionId(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("GetAllQuestionOutputsByQuestionId finished succesfully")
	return res, nil
}

func (s *OutputService) UpdateQuestionOutput(ctx context.Context, req *pb.UpdateQuestionOutputRequest) (*pb.Void, error) {
	s.Logger.Info("UpdateQuestionOutput rpc method is started")
	res, err := s.Output.Output().UpdateQuestionOutput(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("UpdateQuestionOutput finished succesfully")

	return res, nil
}

func (s *OutputService) DeleteQuestionOutput(ctx context.Context, req *pb.DeleteQuestionOutputRequest) (*pb.Void, error) {
	s.Logger.Info("DeleteQuestionOutput rpc method is started")
	res, err := s.Output.Output().DeleteQuestionOutput(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("DeleteQuestionOutput finished succesfully")
	return res, nil
}

func (s *OutputService) GetQUestionOutPutByInputId(ctx context.Context, req *pb.GetQUestionOutPutByInputIdRequest) (*pb.GetQUestionOutPutByInputIdRes, error) {
	s.Logger.Info("GetQuestionOutputByInputId rpc method is started")
	res, err := s.Output.Output().GetQuestionOutputByInputId(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("GetQuestionOutputByInputId finished succesfully")
	return res, nil
}
