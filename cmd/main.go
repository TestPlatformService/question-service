package main

import (
	"context"
	"log"
	"net"
	"question/config"
	"question/genproto/question"
	"question/genproto/subject"
	"question/genproto/task"
	"question/genproto/topic"
	"question/logs"
	"question/service"
	"question/storage"
	"question/storage/mongosh"
	"question/storage/postgres"

	"google.golang.org/grpc"
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

	storage := storage.NewStoragePro(mdb, db, logger)

	listener, err := net.Listen("tcp", cfg.QUESTION_SERVICE)
	if err != nil{
		panic(err)
	}
	defer listener.Close()

	topicService := service.NewTopicService(storage, logger)
	subjectService := service.NewSubjectService(logger, storage)
	caseService := service.NewCaseService(logger, storage)
	inputService := service.NewInputService(logger, storage)
	outputService := service.NewOutputService(logger, storage)
	questionService := service.NewQuestionService(logger, storage)
	taskService := service.NewTaskService(storage, logger)

	s := grpc.NewServer()

	topic.RegisterTopicServiceServer(s, topicService)
	subject.RegisterSubjectServiceServer(s, subjectService)
	question.RegisterTestCaseServiceServer(s, caseService)
	question.RegisterInputServiceServer(s, inputService)
	question.RegisterOutputServiceServer(s, outputService)
	question.RegisterQuestionServiceServer(s, questionService)
	task.RegisterTaskServiceServer(s, taskService)

	log.Printf("Service is run: %v", cfg.QUESTION_SERVICE)
	if err = s.Serve(listener); err != nil{
		panic(err)
	}
}