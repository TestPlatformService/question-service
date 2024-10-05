package postgres

import (
	"context"
	"database/sql"
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
					id, name, subject_id, description, question_count)
				VALUES
					($1, $2, $3, $4)`
	_, err := T.DB.Exec(query, id, req.Name, req.SubjectId, req.Description, req.QuestionCount)
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
					name = $1, subject_id = $2, description = $3, updated_at = $5, question_count = $6
				WHERE 
					id = $4 AND deleted_at IS NULL`
	_, err := T.DB.Exec(query, req.Name, req.SubjectId, req.Description, req.Id, time.Now(), req.QuestionCount)
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
			id, name, subject_id, description, question_count
		FROM 
			subject_topics
		WHERE 
			deleted_at IS NULL`

	if len(req.SubjectId) > 0 {
		query += " AND subject_id = ?"
	}
	query += " LIMIT ? OFFSET ?"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := T.DB.QueryContext(ctx, query, req.SubjectId, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		T.Logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close() // Resurslarni tozalash

	for rows.Next() {
		var topic pb.Topic
		if err = rows.Scan(&topic.Id, &topic.Name, &topic.SubjectId, &topic.Description, &topic.QuestionCount); err != nil {
			T.Logger.Error(err.Error())
			return nil, err
		}
		topics = append(topics, &topic)
	}

	var count int32
	countQuery := `
		SELECT COUNT(id)
		FROM subject_topics
		WHERE deleted_at IS NULL`
	if err = T.DB.QueryRowContext(ctx, countQuery).Scan(&count); err != nil {
		T.Logger.Error(err.Error())
		return nil, err
	}

	return &pb.GetAllTopicsResp{
		Topics: topics,
		Limit:  req.Limit,
		Page:   req.Page,
		Count:  count,
	}, nil
}
