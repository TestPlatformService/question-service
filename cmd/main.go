package main

import (
	"context"
	"net"
	"question/config"
	"question/genproto/subject"
	"question/genproto/topic"
	"question/logs"
	"question/service"
	"question/storage"
	"question/storage/mongosh"
	"question/storage/postgres"
)

func main(){
	logger := logs.NewLogger()
	cfg := config.LoadConfig()

	db, err := postgres.ConnectDB()
	if err != nil{
		panic(err)
	}
	defer db.Close()

	mdb, err := mongosh.Connect(context.Background())
	if err != nil{
		panic(err)
	}

	storage := storage.NewStoragePro(mdb, db)

	listener, err := net.Listen("tcp", cfg.QUESTION_SERVICE)
	if err != nil{
		panic(err)
	}
	defer listener.Close()

	topicService := service.NewTopicService(storage, logger)
	subjectService := service.NewSubjectService(logger, storage)
	
}