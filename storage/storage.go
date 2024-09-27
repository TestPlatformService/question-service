package storage

import (
	"database/sql"
	"question/storage/mongosh"
	"question/storage/repo"

	"go.mongodb.org/mongo-driver/mongo"
)

type Istorage interface {
	Question() repo.IQuestionStorage
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
