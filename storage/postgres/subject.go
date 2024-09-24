package postgres

import (
	"database/sql"
	"log/slog"
	"question/logs"
	"question/storage"
)

type subjectRepo struct {
	DB *sql.DB
	Log *slog.Logger
}

func NewSubjectRepo(DB *sql.DB) storage.ISubjectStorage {
	return &subjectRepo{DB: DB, Log: logs.NewLogger()}
}