package repo

import (
	"context"
	pb3 "question/genproto/question"
	pb "question/genproto/subject"
	pb4 "question/genproto/task"
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
	GetQuestionsByIds(ctx context.Context, ids []string) ([]*pb4.Question, error)
	GetQuestionRandomly(ctx context.Context, req *pb3.GetQuestionRandomlyRequest) (pb3.GetQuestionRandomlyResponse, error)
}

type IInputStorage interface {
	CreateQuestionInput(context.Context, *pb3.CreateQuestionInputRequest) (*pb3.QuestionInputId, error)
	GetQuestionInput(context.Context, *pb3.QuestionInputId) (*pb3.GetQuestionInputResponse, error)
	GetAllQuestionInputsByQuestionId(context.Context, *pb3.GetAllQuestionInputsByQuestionIdRequest) (*pb3.GetAllQuestionInputsByQuestionIdResponse, error)
	UpdateQuestionInput(context.Context, *pb3.UpdateQuestionInputRequest) (*pb3.Void, error)
	DeleteQuestionInput(context.Context, *pb3.DeleteQuestionInputRequest) (*pb3.Void, error)
}

type IOutputStorage interface {
	CreateQuestionOutput(context.Context, *pb3.CreateQuestionOutputRequest) (*pb3.QuestionOutputId, error)
	GetQuestionOutput(context.Context, *pb3.QuestionOutputId) (*pb3.GetQuestionOutputResponse, error)
	GetAllQuestionOutputsByQuestionId(context.Context, *pb3.GetAllQuestionOutputsByQuestionIdRequest) (*pb3.GetAllQuestionOutputsByQuestionIdResponse, error)
	UpdateQuestionOutput(context.Context, *pb3.UpdateQuestionOutputRequest) (*pb3.Void, error)
	DeleteQuestionOutput(context.Context, *pb3.DeleteQuestionOutputRequest) (*pb3.Void, error)
}

type ITestCaseStorage interface {
	CreateTestCase(context.Context, *pb3.CreateTestCaseRequest) (*pb3.TestCaseId, error)
	GetTestCase(context.Context, *pb3.TestCaseId) (*pb3.GetTestCaseResponse, error)
	GetAllTestCasesByQuestionId(context.Context, *pb3.GetAllTestCasesByQuestionIdRequest) (*pb3.GetAllTestCasesByQuestionIdResponse, error)
	UpdateTestCase(context.Context, *pb3.UpdateTestCaseRequest) (*pb3.Void, error)
	DeleteTestCase(context.Context, *pb3.DeleteTestCaseRequest) (*pb3.Void, error)
}

type ITaskStorage interface {
	CreateTask(req *pb4.CreateTaskReq) (*pb4.CreateTaskResp, error)
	DeleteTask(req *pb4.DeleteTaskReq) (*pb4.DeleteTaskResp, error)
	GetTask(req *pb4.GetTaskReq) ([]string, error)
}
