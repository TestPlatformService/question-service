package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	pb "question/genproto/topic"
	"question/storage/repo"
	"time"

	"github.com/google/uuid"
)

type topicRepo struct {
	DB     *sql.DB
	Logger *slog.Logger
}

func NewTopicRepo(db *sql.DB, logger *slog.Logger) repo.ITopicStorage {
	return &topicRepo{
		DB:     db,
		Logger: logger,
	}
}

func (T *topicRepo) CreateTopic(req *pb.CreateTopicReq) (*pb.CreateTopicResp, error) {
	id := uuid.NewString()
	query := `
				INSERT INTO subject_topics(
					id, name, subject_id, description)
				VALUES
					($1, $2, $3, $4)`
	_, err := T.DB.Exec(query, id, req.Name, req.SubjectId, req.Description)
	if err != nil {
		T.Logger.Error(err.Error())
		return nil, err
	}
	return &pb.CreateTopicResp{
		Id:        id,
		CreatedAt: time.Now().String(),
	}, nil
}

func (T *topicRepo) UpdateTopic(req *pb.UpdateTopicReq) (*pb.UpdateTopicResp, error) {
	query := `
				UPDATE subject_topics SET
					name = $1, subject_id = $2, description = $3, updated_at = $5
				WHERE 
					id = $4 AND deleted_at IS NULL`
	_, err := T.DB.Exec(query, req.Name, req.SubjectId, req.Description, req.Id, time.Now())
	if err != nil {
		T.Logger.Error(err.Error())
		return nil, err
	}
	return &pb.UpdateTopicResp{
		Id:        req.Id,
		UpdatedAt: time.Now().String(),
	}, nil
}

func (T *topicRepo) DeleteTopic(req *pb.DeleteTopicReq) (*pb.DeleteTopicResp, error) {
	query := `
				UPDATE subject_topics SET
					deleted_at = $1
				WHERE 
					id = $2 AND deleted_at IS NULL`
	_, err := T.DB.Exec(query, time.Now(), req.TopicId)
	if err != nil {
		T.Logger.Error(err.Error())
		return &pb.DeleteTopicResp{
			Status: "Topic o'chirilmadi",
		}, err
	}
	return &pb.DeleteTopicResp{
		Status: "Topic muvaffaqiyatli o'chirildi",
	}, nil
}

func (T *topicRepo) GetAllTopics(req *pb.GetAllTopicsReq) (*pb.GetAllTopicsResp, error) {
	var topics = []*pb.Topic{}
	query := `
				SELECT 
					id, name, subject_id, description
				FROM 
					subject_topics
				WHERE 
					deleted_at IS NULL`
	if len(req.SubjectId) > 0 {
		query += fmt.Sprintf(" AND subject_id = '%s'", req.SubjectId)
	}
	query += fmt.Sprintf(" limit %d offset %d", req.Limit, (req.Page-1)*req.Limit)
	rows, err := T.DB.Query(query)
	if err != nil {
		T.Logger.Error(err.Error())
		return nil, err
	}
	for rows.Next() {
		var topic = pb.Topic{}
		err = rows.Scan(&topic.Id, &topic.Name, &topic.SubjectId, &topic.Description)
		if err != nil {
			T.Logger.Error(err.Error())
			return nil, err
		}
		topics = append(topics, &topic)
	}
	return &pb.GetAllTopicsResp{
		Topics: topics,
		Limit:  req.Limit,
		Page: req.Page,
	}, nil
}

