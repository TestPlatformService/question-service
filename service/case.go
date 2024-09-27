package service

import (
	"context"
	"log/slog"
	pb "question/genproto/question"
	"question/storage"
)

type CaseService struct {
	pb.UnimplementedTestCaseServiceServer
	Case   storage.Istorage
	Logger *slog.Logger
}

func NewCaseService(Logger *slog.Logger, istorage storage.Istorage) *CaseService {
	return &CaseService{
		Case:   istorage,
		Logger: Logger,
	}
}

func (s *CaseService) CreateTestCase(ctx context.Context, req *pb.CreateTestCaseRequest) (*pb.TestCaseId, error) {
	s.Logger.Info("CreateTestCase rpc method is started")
	res, err := s.Case.TestCase().CreateTestCase(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("CreateTestCase finished succesfully")
	return res, nil
}

func (s *CaseService) GetTestCase(ctx context.Context, req *pb.TestCaseId) (*pb.GetTestCaseResponse, error) {
	s.Logger.Info("GetTestCase rpc method is started")
	res, err := s.Case.TestCase().GetTestCase(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("GetTestCase finished succesfully")
	return res, nil
}

func (s *CaseService) GetAllTestCasesByQuestionId(ctx context.Context, req *pb.GetAllTestCasesByQuestionIdRequest) (*pb.GetAllTestCasesByQuestionIdResponse, error) {
	s.Logger.Info("GetAllTestCasesByQuestionId rpc method is started")
	res, err := s.Case.TestCase().GetAllTestCasesByQuestionId(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("GetAllTestCasesByQuestionId finished succesfully")
	return res, nil
}

func (s *CaseService) UpdateTestCase(ctx context.Context, req *pb.UpdateTestCaseRequest) (*pb.Void, error) {
	s.Logger.Info("UpdateTestCase rpc method is started")
	res, err := s.Case.TestCase().UpdateTestCase(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("UpdateTestCase finished succesfully")
	return res, nil
}

func (s *CaseService) DeleteTestCase(ctx context.Context, req *pb.DeleteTestCaseRequest) (*pb.Void, error) {
	s.Logger.Info("DeleteTestCase rpc method is started")
	res, err := s.Case.TestCase().DeleteTestCase(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	s.Logger.Info("DeleteTestCase finished succesfully")
	return res, nil
}
