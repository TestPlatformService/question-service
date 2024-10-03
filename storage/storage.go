package storage

import (
	"database/sql"
	"log/slog"
	"question/logs"
	"question/storage/mongosh"
	"question/storage/postgres"
	"question/storage/repo"

	"go.mongodb.org/mongo-driver/mongo"
)

type Istorage interface {
	Question() repo.IQuestionStorage
	Output() repo.IOutputStorage
	Input() repo.IInputStorage
	TestCase() repo.ITestCaseStorage
	Subject() repo.ISubjectStorage
	Topic() repo.ITopicStorage
}

type StoragePro struct {
	Mdb *mongo.Database
	PDB *sql.DB
	Logger *slog.Logger
}

func NewStoragePro(mdb *mongo.Database, pdb *sql.DB, logger *slog.Logger) Istorage {
	return &StoragePro{
		Mdb: mdb,
		PDB: pdb,
		Logger: logger,
	}
}

func (pro *StoragePro) Question() repo.IQuestionStorage {
	return mongosh.NewQuestionRepository(pro.Mdb)
}

func (pro *StoragePro) Subject() repo.ISubjectStorage {
	return postgres.NewSubjectRepo(pro.PDB)
}

func (pro *StoragePro) Topic() repo.ITopicStorage {
	return postgres.NewTopicRepo(pro.PDB, logs.NewLogger())
}

func (pro *StoragePro) Input() repo.IInputStorage {
	return mongosh.NewInputRepository(pro.Mdb)
}

func (pro *StoragePro) Output() repo.IOutputStorage {
	return mongosh.NewOutputRepository(pro.Mdb)
}

func (pro *StoragePro) TestCase() repo.ITestCaseStorage {
	return mongosh.NewTestCaseRepository(pro.Mdb)
}

func (pro *StoragePro) Task() repo.ITaskStorage {
	return postgres.NewTaskService(pro.PDB, pro.Logger)
}
