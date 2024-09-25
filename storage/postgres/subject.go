package postgres

import (
	pb "question/genproto/subject"
	"context"
	"database/sql"
	"log/slog"
	"question/logs"
	"question/storage"
	"time"
)

type subjectRepo struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewSubjectRepo(DB *sql.DB) storage.ISubjectStorage {
	return &subjectRepo{DB: DB, Log: logs.NewLogger()}
}

func (s *subjectRepo) CreateSubject(ctx context.Context, req *pb.CreateSubjectRequest) (*pb.Void, error) {
	query := `INSERT INTO subjects (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.ExecContext(ctx, query, req.Name, req.Description, time.Now(), time.Now())
	if err != nil {
		s.Log.Error("failed to create subject", "error", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *subjectRepo) GetSubject(ctx context.Context, req *pb.GetSubjectRequest) (*pb.GetSubjectResponse, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM subjects WHERE id = $1 and deleted_at IS NULL`
	row := s.DB.QueryRowContext(ctx, query, req.Id)
	var subject pb.GetSubjectResponse
	err := row.Scan(&subject.Id, &subject.Name, &subject.Description, &subject.CreatedAt, &subject.UpdatedAt)
	if err != nil {
		s.Log.Error("failed to get subject", "error", err)
		return nil, err
	}
	return &subject, nil
}


func (s *subjectRepo) GetAllSubjects(ctx context.Context, req *pb.GetAllSubjectsRequest) (*pb.GetAllSubjectsResponse, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM subjects WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := s.DB.QueryContext(ctx, query, req.Limit, req.Offset)
	if err != nil {
		s.Log.Error("failed to get all subjects", "error", err)
		return nil, err
	}
	defer rows.Close()

	var subjects []*pb.GetAll
	for rows.Next() {
		var subject pb.GetAll
		if err := rows.Scan(&subject.Id, &subject.Name, &subject.Description); err != nil {
			s.Log.Error("failed to scan subject row", "error", err)
			return nil, err
		}
		subjects = append(subjects, &subject)
	}

	if err = rows.Err(); err != nil {
		s.Log.Error("error during rows iteration", "error", err)
		return nil, err
	}

	return &pb.GetAllSubjectsResponse{
		Subjects: subjects,
	}, nil
}


func (s *subjectRepo) UpdateSubject(ctx context.Context, req *pb.UpdateSubjectRequest) (*pb.Void, error) {
	query := `UPDATE subjects SET name = $1, description = $2, updated_at = $3 WHERE id = $4 and deleted_at IS NULL`
	_, err := s.DB.ExecContext(ctx, query, req.Name, req.Description, time.Now(), req.Id)
	if err != nil {
		s.Log.Error("failed to update subject", "error", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *subjectRepo) DeleteSubject(ctx context.Context, req *pb.DeleteSubjectRequest) (*pb.Void, error) {
	query := `UPDATE subjects SET deleted_at = $1 WHERE id = $2 and deleted_at IS NULL`
	_, err := s.DB.ExecContext(ctx, query, time.Now(), req.Id)
	if err != nil {
		s.Log.Error("failed to delete subject", "error", err)
		return nil, err
	}
	return &pb.Void{}, nil
}
	