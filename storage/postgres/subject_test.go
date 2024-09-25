package postgres

import (
	"context"
	pb "question/genproto/subject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSubject(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewSubjectRepo(db)
	req := &pb.CreateSubjectRequest{
		Name:        "Foundation",
		Description: "Is beginning of programming",
	}

	_, err = repo.CreateSubject(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSubject(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewSubjectRepo(db)

	req := &pb.GetSubjectRequest{
		Id: "b6c59395-b07d-4796-ac92-fc8c9cfe4cb5",
	}

	res, err := repo.GetSubject(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, res)
}

func TestUpdateSubject(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewSubjectRepo(db)

	req := &pb.UpdateSubjectRequest{
		Id: "b6c59395-b07d-4796-ac92-fc8c9cfe4cb5",
		Name: "Go",
		Description: "For Gophers",
	}

	_, err = repo.UpdateSubject(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDeleteSubject(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewSubjectRepo(db)

	req := &pb.DeleteSubjectRequest{
		Id: "b6c59395-b07d-4796-ac92-fc8c9cfe4cb5",
	}

	_, err = repo.DeleteSubject(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}