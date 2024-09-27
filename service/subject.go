package service

import (
	"context"
	"log/slog"
	pb "question/genproto/subject"
	"question/storage"
)

type SubjectService struct {
	pb.UnimplementedSubjectServiceServer
	Subject storage.Istorage
	Logger  *slog.Logger
}

func NewSubjectService(Logger *slog.Logger, istorage storage.Istorage) *SubjectService {
	return &SubjectService{
		Subject: istorage,
		Logger:  Logger,
	}
}

func (s *SubjectService) CreateSubject(ctx context.Context, req *pb.CreateSubjectRequest) (*pb.Void, error) {
	_, err := s.Subject.Subject().CreateSubject(ctx, req)
	if err != nil {
		s.Logger.Error("failed to create subject", "error", err)
		return nil, err
	}

	return &pb.Void{}, nil
}

func (s *SubjectService) GetSubject(ctx context.Context, req *pb.GetSubjectRequest) (*pb.GetSubjectResponse, error) {
	res, err := s.Subject.Subject().GetSubject(ctx, req)
	if err != nil {
		s.Logger.Error("failed to get subject", "error", err)
		return nil, err
	}

	return res, nil
}

func (s *SubjectService) UpdateSubject(ctx context.Context, req *pb.UpdateSubjectRequest) (*pb.Void, error) {
	_, err := s.Subject.Subject().UpdateSubject(ctx, req)
	if err != nil {
		s.Logger.Error("failed to update subject", "error", err)
		return nil, err
	}

	return &pb.Void{}, nil
}

func (s *SubjectService) DeleteSubject(ctx context.Context, req *pb.DeleteSubjectRequest) (*pb.Void, error) {
	_, err := s.Subject.Subject().DeleteSubject(ctx, req)
	if err != nil {
		s.Logger.Error("failed to delete subject", "error", err)
		return nil, err
	}

	return &pb.Void{}, nil
}
