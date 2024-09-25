package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"question/config"
	"question/storage"

	_ "github.com/lib/pq"
)

type postgresStorage struct {
	db *sql.DB
	logger *slog.Logger
}

func ConnectDB() (*sql.DB, error) {

	conf := config.LoadConfig()
	conDb := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_NAME, conf.DB_PASSWORD)
	log.Printf("connecting to postgres: %s\n", conDb)
	db, err := sql.Open("postgres", conDb)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}

func (p *postgresStorage) Close() {
	p.db.Close()
}

func (p *postgresStorage) Subject() storage.ISubjectStorage {
	return NewSubjectRepo(p.db)
}

func (p *postgresStorage) Topic() storage.ITopicStorage{
	return NewTopicRepo(p.db, p.logger)
}

func NewIstorage(db *sql.DB, logger *slog.Logger) storage.IStorage {
	return &postgresStorage{
		db: db,
		logger: logger,
	}
}
