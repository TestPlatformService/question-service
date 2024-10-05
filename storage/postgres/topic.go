package postgres

import (
	"database/sql"
	"log/slog"
	pb "question/genproto/topic"
	"question/storage/repo"
	"strconv"
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
	count, err := strconv.Atoi(req.QuestionCount)
	if err != nil {
		return nil, err
	}
	query := `
				INSERT INTO subject_topics(
					id, name, subject_id, description, question_count)
				VALUES
					($1, $2, $3, $4)`
	_, err = T.DB.Exec(query, id, req.Name, req.SubjectId, req.Description, count)
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
	resp := &pb.GetAllTopicsResp{
		Limit: req.Limit,
		Page:  req.Page,
	}

	baseQuery := `
        SELECT id, subject_id, name, description, created_at
        FROM subject_topics
        WHERE deleted_at IS NULL
    `

	countQuery := `
        SELECT COUNT(*)
        FROM subject_topics
        WHERE deleted_at IS NULL
    `

	filters := []interface{}{}
	conditions := ""

	if req.SubjectId != "" {
		conditions += " AND subject_id = $1"
		filters = append(filters, req.SubjectId)
	}

	offset := (req.Page - 1) * req.Limit
	filters = append(filters, req.Limit, offset)

	if req.SubjectId != "" {
		baseQuery += conditions + " LIMIT $2 OFFSET $3"
	} else {
		baseQuery += " LIMIT $1 OFFSET $2"
	}

	rows, err := T.DB.Query(baseQuery, filters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var topic pb.Topic
		err := rows.Scan(
			&topic.Id,
			&topic.SubjectId,
			&topic.Name,
			&topic.Description,
			&topic.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp.Topics = append(resp.Topics, &topic)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if req.SubjectId != "" {
		countQuery += conditions
	}
	var totalCount int
	err = T.DB.QueryRow(countQuery, filters[0]).Scan(&totalCount)
	if err != nil && req.SubjectId == "" {
		err = T.DB.QueryRow(countQuery).Scan(&totalCount)
	}

	if err != nil {
		return nil, err
	}

	resp.Count = int32(totalCount)

	return resp, nil
}
