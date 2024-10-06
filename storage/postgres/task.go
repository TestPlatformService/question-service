package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"question/genproto/group"
	"question/genproto/question"
	pb "question/genproto/task"
	"question/pkg"
	"question/storage/repo"
	"time"

	"github.com/google/uuid"
)

type TaskRepo struct {
	DB     *sql.DB
	Logger *slog.Logger
	mng    repo.IQuestionStorage
}

func NewTaskService(db *sql.DB, log *slog.Logger, mongo repo.IQuestionStorage) *TaskRepo {
	return &TaskRepo{
		DB:     db,
		Logger: log,
		mng:    mongo,
	}
}

func (T *TaskRepo) CreateTask(req *pb.CreateTaskReq) (*pb.CreateTaskResp, error) {
	id := uuid.NewString()
	grp, err := pkg.GroupServiceClient()
	if err != nil {
		T.Logger.Error(err.Error())
		return nil, err
	}

	tr, err := T.DB.Begin()
	if err != nil {
		T.Logger.Error(err.Error())
		tr.Rollback()
		return nil, err
	}
	query := `
				SELECT 
					question_count
				FROM
					subject_topics
				WHERE
					id = $1 AND deleted_at IS NULL`
	var questionCount int
	err = tr.QueryRow(query, req.TopicId).Scan(&questionCount)
	if err != nil {
		T.Logger.Error(err.Error())
		tr.Rollback()
		return nil, err
	}
	query2 := `
				INSERT INTO user_tasks(
					id, user_id, topic_id, question_id)
				VALUES
					($1, $2, $3, $4)`
	students, err := grp.GetGroupStudents(context.Background(), &group.GroupId{Id: req.GroupId})
	if err != nil {
		T.Logger.Error(err.Error())
		tr.Rollback()
		return nil, err
	}
	for _, student := range students.Students {
		resp, err := T.mng.GetQuestionRandomly(context.Background(), &question.GetQuestionRandomlyRequest{
			TopicId: req.TopicId,
			Count:   int64(questionCount),
		})
		if err != nil {
			T.Logger.Error(err.Error())
			return nil, err
		}
		questions := resp.QuestionsId
		for i := 0; i < len(questions); i++ {
			_, err = tr.Exec(query2, id, student.Id, req.TopicId, *questions[i])
			T.Logger.Info(fmt.Sprintf("Id::::::::::::", questions[i], *questions[i]))
			if err != nil {
				T.Logger.Error(err.Error())
				tr.Rollback()
				return nil, err
			}
		}
	}
	return &pb.CreateTaskResp{
		TaskId:    id,
		CreatedAt: time.Now().String(),
	}, nil
}

func (T *TaskRepo) DeleteTask(req *pb.DeleteTaskReq) (*pb.DeleteTaskResp, error) {
	query := `
				DELETE
					FROM
				user_tasks
					WHERE
				task_id = $1 AND deleted_at IS NULL`
	_, err := T.DB.Exec(query, req.TaskId)
	if err != nil {
		T.Logger.Error(err.Error())
		return &pb.DeleteTaskResp{
			Status: "Task o'chirilmadi",
		}, err
	}
	return &pb.DeleteTaskResp{
		Status: "Task muvaffaqiyatli o'chirildi",
	}, nil
}

func (T *TaskRepo) GetTask(req *pb.GetTaskReq) ([]string, error) {
	var questionIds = []string{}
	query := `
				SELECT 
					question_id
				FROM
					user_tasks
				WHERE
					id = $1 AND user_id = $2 AND topic_id = $3 AND deleted_at IS NULL`
	tr, err := T.DB.Begin()
	if err != nil {
		T.Logger.Error(err.Error())
		tr.Rollback()
		return nil, err
	}
	rows, err := tr.Query(query, req.TaskId, req.UserId, req.TopicId)
	if err != nil {
		T.Logger.Error(err.Error())
		tr.Rollback()
		return nil, err
	}
	for rows.Next() {
		var questionId string
		err = rows.Scan(&questionId)
		if err != nil {
			T.Logger.Error(err.Error())
			tr.Rollback()
			return nil, err
		}
		questionIds = append(questionIds, questionId)
	}
	return questionIds, nil
}
