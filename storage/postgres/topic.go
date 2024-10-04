package postgres

import (
	"database/sql"
	"log/slog"
	pb "question/genproto/topic"
	"question/storage/repo"
	"time"
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
	query := `
				INSERT INTO subject_topics(
					name, subject_id, description)
				VALUES
					($1, $2, $3)
				RETURNING id`

	var id string

	err := T.DB.QueryRow(query, req.Name, req.SubjectId, req.Description).Scan(&id)
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
	// Initialize response
	resp := &pb.GetAllTopicsResp{
		Limit: req.Limit,
		Page:  req.Page,
	}

	// Base query
	baseQuery := `
        SELECT id, subject_id, name, description, created_at
        FROM subject_topics
        WHERE deleted_at IS NULL
    `

	// Base count query
	countQuery := `
        SELECT COUNT(*)
        FROM subject_topics
        WHERE deleted_at IS NULL
    `

	// Filters
	filters := []interface{}{}
	conditions := ""

	if req.SubjectId != "" {
		conditions += " AND subject_id = $1"
		filters = append(filters, req.SubjectId)
	}

	// Pagination: Calculate offset based on page and limit
	offset := (req.Page - 1) * req.Limit
	filters = append(filters, req.Limit, offset)

	// Final query with filters and pagination
	finalQuery := baseQuery + conditions + " LIMIT $2 OFFSET $3"

	// Execute the final query
	rows, err := T.DB.Query(finalQuery, filters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Collect topics
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

	// Error check for rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Count total topics (without limit and offset)
	countQueryFinal := countQuery + conditions
	var totalCount int
	err = T.DB.QueryRow(countQueryFinal, filters[0]).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	// Set the total count in response
	resp.Count = int32(totalCount)

	return resp, nil
}
