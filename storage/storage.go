package storage

import (
	"database/sql"
	"question/logs"
	"question/storage/mongosh"
	"question/storage/postgres"
	"question/storage/repo"

	"go.mongodb.org/mongo-driver/mongo"
)

type Istorage interface {
	Question() repo.IQuestionStorage
	Subject() repo.ISubjectStorage
	Topic() repo.ITopicStorage
}

type StoragePro struct {
	Mdb *mongo.Database
	PDB *sql.DB
}

func NewStoragePro(mdb *mongo.Database, pdb *sql.DB) Istorage {
	return &StoragePro{
		Mdb: mdb,
		PDB: pdb,
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
