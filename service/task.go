package service

import (
	"context"
	"fmt"
	"log/slog"
	pb "question/genproto/task"
	"question/storage"
)

type TaskService struct {
	storage storage.Istorage
	pb.UnimplementedTaskServiceServer
	logger *slog.Logger
}

func NewTaskService(storage storage.Istorage, logger *slog.Logger) *TaskService {
	return &TaskService{
		storage: storage,
		logger:  logger,
	}
}

func (T *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskReq) (*pb.CreateTaskResp, error) {
	resp, err := T.storage.Task().CreateTask(req)
	if err != nil {
		T.logger.Error(fmt.Sprintf("CreateTask request error: %v", err))
		return nil, err
	}
	return resp, nil
}

func (T *TaskService) DeleteTask(ctx context.Context, req *pb.DeleteTaskReq) (*pb.DeleteTaskResp, error) {
	resp, err := T.storage.Task().DeleteTask(req)
	if err != nil {
		T.logger.Error(fmt.Sprintf("DeleteTask request error: %v", err))
		return nil, err
	}
	return resp, nil
}

func (T *TaskService) GetTask(ctx context.Context, req *pb.GetTaskReq) (*pb.GetTaskResp, error) {
	questionsId, err := T.storage.Task().GetTask(req)
	if err != nil {
		T.logger.Error(fmt.Sprintf("Savollarning idisini olishda xatolik: %v", err))
		return nil, err
	}
	questions, err := T.storage.Question().GetQuestionsByIds(ctx, questionsId)
	if err != nil {
		T.logger.Error(fmt.Sprintf("Savollarni olishda xatolik: %v", err))
		return nil, err
	}
	return &pb.GetTaskResp{
		TaskId:    req.TaskId,
		Questions: questions,
	}, nil
}
