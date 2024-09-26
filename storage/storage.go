package storage

import (
	"context"
	pb "question/genproto/subject"
)

type IStorage interface {
	Subject() ISubjectStorage
	Close()
}

type ISubjectStorage interface {
	CreateSubject(context.Context, *pb.CreateSubjectRequest) (*pb.Void, error)
	GetSubject(context.Context, *pb.GetSubjectRequest) (*pb.GetSubjectResponse, error)
	GetAllSubjects(context.Context, *pb.GetAllSubjectsRequest) (*pb.GetAllSubjectsResponse, error)
	UpdateSubject(context.Context, *pb.UpdateSubjectRequest) (*pb.Void, error)
	DeleteSubject(context.Context, *pb.DeleteSubjectRequest) (*pb.Void, error)
}
