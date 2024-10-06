package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	"question/genproto/group"
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
	query = `
				SELECT 
					question_id
				FROM 
					topic_questions
				ORDER BY 
					RANDOM()
				LIMIT $1`
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
		questions := []string{}
		rows, err := tr.Query(query, questionCount)
		if err != nil {
			T.Logger.Error(err.Error())
			tr.Rollback()
			return nil, err
		}
		for rows.Next() {
			var question string
			err = rows.Scan(&question)
			if err != nil {
				T.Logger.Error(err.Error())
				tr.Rollback()
				return nil, err
			}
			questions = append(questions, question)
		}
		for i := 0; i < questionCount; i++ {
			_, err = tr.Exec(query2, id, student.Id, req.TopicId, questions[i])
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
