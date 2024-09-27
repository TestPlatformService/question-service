package repo

import (
	"context"
	pb3 "question/genproto/question"
	pb "question/genproto/subject"
	pb2 "question/genproto/topic"
)

type IStorage interface {
	Subject() ISubjectStorage
	Topic() ITopicStorage
	Close()
}

type ISubjectStorage interface {
	CreateSubject(context.Context, *pb.CreateSubjectRequest) (*pb.Void, error)
	GetSubject(context.Context, *pb.GetSubjectRequest) (*pb.GetSubjectResponse, error)
	GetAllSubjects(context.Context, *pb.GetAllSubjectsRequest) (*pb.GetAllSubjectsResponse, error)
	UpdateSubject(context.Context, *pb.UpdateSubjectRequest) (*pb.Void, error)
	DeleteSubject(context.Context, *pb.DeleteSubjectRequest) (*pb.Void, error)
}

type ITopicStorage interface {
	CreateTopic(req *pb2.CreateTopicReq) (*pb2.CreateTopicResp, error)
	UpdateTopic(req *pb2.UpdateTopicReq) (*pb2.UpdateTopicResp, error)
	DeleteTopic(req *pb2.DeleteTopicReq) (*pb2.DeleteTopicResp, error)
	GetAllTopics(req *pb2.GetAllTopicsReq) (*pb2.GetAllTopicsResp, error)
}

type IQuestionStorage interface {
	CreateQuestion(context.Context, *pb3.CreateQuestionRequest) (*pb3.QuestionId, error)
	GetQuestion(context.Context, *pb3.QuestionId) (*pb3.GetQuestionResponse, error)
	GetAllQuestions(context.Context, *pb3.GetAllQuestionsRequest) (*pb3.GetAllQuestionsResponse, error)
	UpdateQuestion(context.Context, *pb3.UpdateQuestionRequest) (*pb3.Void, error)
	DeleteQuestion(context.Context, *pb3.DeleteQuestionRequest) (*pb3.Void, error)
	UploadImageQuestion(context.Context, *pb3.UploadImageQuestionRequest) (*pb3.Void, error)
	DeleteImageQuestion(context.Context, *pb3.DeleteImageQuestionRequest) (*pb3.Void, error)
	IsQuestionExist(context.Context, *pb3.QuestionId) (*pb3.Void, error)
}
