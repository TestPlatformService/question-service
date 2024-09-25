package postgres

import (
	"database/sql"
	pb "question/genproto/topic"
	"question/logs"
	"testing"
)

var logger = logs.NewLogger()

func DB()*sql.DB{
	db, err := ConnectDB()
	if err != nil{
		panic(err)
	}
	return db
}

func Test_CreateTopic(t *testing.T){
	db := DB()
	defer db.Close()

	topic := NewTopicRepo(db, logger)

	_, err := topic.CreateTopic(&pb.CreateTopicReq{
		Name: "For loop",
		Description: "Takroriy operator",
		SubjectId: "6c4f4072-db59-4fab-a584-6e6b7b50be61",
	})
	if err != nil{
		t.Fatal(err)
	}
}

func Test_UpdateTopic(t *testing.T){
	db := DB()
	defer db.Close()

	topic := NewTopicRepo(db, logger)

	_, err := topic.UpdateTopic(&pb.UpdateTopicReq{
		Id: "b7d16159-ea38-4df6-95f1-34e703a7fa3f",
		SubjectId: "6c4f4072-db59-4fab-a584-6e6b7b50be61",
		Name: "While",
		Description: "Takrorlash operatori",
	})
	if err != nil{
		t.Fatal(err)
	}
}

func Test_GetAllTopics(t *testing.T){
	db := DB()
	defer db.Close()

	topic := NewTopicRepo(db, logger)

	_, err := topic.GetAllTopics(&pb.GetAllTopicsReq{
		SubjectId: "6c4f4072-db59-4fab-a584-6e6b7b50be61",
		Limit: 10,
		Offset: 0,
	})
	if err != nil{
		t.Fatal(err)
	}
}

func Test_DeleteTopic(t *testing.T){
	db := DB()
	defer db.Close()

	topic := NewTopicRepo(db, logger)

	_, err := topic.DeleteTopic(&pb.DeleteTopicReq{
		TopicId: "b7d16159-ea38-4df6-95f1-34e703a7fa3f",
	})
	if err != nil{
		t.Fatal(err)
	}
}