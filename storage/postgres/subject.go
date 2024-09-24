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
	query := `SELECT * FROM subjects WHERE id = $1`
	row := s.DB.QueryRowContext(ctx, query, req.Id)
	var subject pb.GetSubjectResponse
	err := row.Scan(&subject.Id, &subject.Name, &subject.Description, &subject.CreatedAt, &subject.UpdatedAt)
	if err != nil {
		s.Log.Error("failed to get subject", "error", err)
		return nil, err
	}
	return &subject, nil
}

func (s *subjectRepo) UpdateSubject(ctx context.Context, req *pb.UpdateSubjectRequest) (*pb.Void, error) {
	query := `UPDATE subjects SET name = $1, description = $2, updated_at = $3 WHERE id = $4`
	_, err := s.DB.ExecContext(ctx, query, req.Name, req.Description, time.Now(), req.Id)
	if err != nil {
		s.Log.Error("failed to update subject", "error", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *subjectRepo) DeleteSubject(ctx context.Context, req *pb.DeleteSubjectRequest) (*pb.Void, error) {
	query := `DELETE FROM subjects WHERE id = $1`
	_, err := s.DB.ExecContext(ctx, query, req.Id)
	if err != nil {
		s.Log.Error("failed to delete subject", "error", err)
		return nil, err
	}
	return &pb.Void{}, nil
}
	