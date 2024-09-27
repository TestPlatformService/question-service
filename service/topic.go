package service

import (
	"context"
	"fmt"
	"log/slog"
	pb "question/genproto/topic"
	"question/storage"
)

type TopicService struct {
	storage storage.Istorage
	pb.UnimplementedTopicServiceServer
	logger *slog.Logger
}

func NewTopicService(storage storage.Istorage, logger *slog.Logger) *TopicService {
	return &TopicService{
		storage: storage,
		logger:  logger,
	}
}

func (T *TopicService) CreateTopic(ctx context.Context, req *pb.CreateTopicReq) (*pb.CreateTopicResp, error) {
	resp, err := T.storage.Topic().CreateTopic(req)
	if err != nil {
		T.logger.Error(fmt.Sprintf("CreateTopic request error: %v", err))
		return nil, err
	}
	return resp, nil
}

func (T *TopicService) UpdateTopic(ctx context.Context, req *pb.UpdateTopicReq) (*pb.UpdateTopicResp, error) {
	resp, err := T.storage.Topic().UpdateTopic(req)
	if err != nil {
		T.logger.Error(fmt.Sprintf("UpdateTopic request error: %v", err))
		return nil, err
	}
	return resp, nil
}

func (T *TopicService) DeleteTopic(ctx context.Context, req *pb.DeleteTopicReq) (*pb.DeleteTopicResp, error) {
	resp, err := T.storage.Topic().DeleteTopic(req)
	if err != nil {
		T.logger.Error(fmt.Sprintf("DeleteTopic request error: %v", err))
		return nil, err
	}
	return resp, nil
}

func (T *TopicService) GetAllTopics(ctx context.Context, req *pb.GetAllTopicsReq) (*pb.GetAllTopicsResp, error) {
	resp, err := T.storage.Topic().GetAllTopics(req)
	if err != nil {
		T.logger.Error(fmt.Sprintf("GetAllTopics request error: %v", err))
		return nil, err
	}
	return resp, nil
}
