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
		return nil, err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()
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
					id, hh_id, topic_id, question_id)
				VALUES
					($1, $2, $3, $4)`
	students, err := grp.GetGroupStudents(context.Background(), &group.GroupId{Id: req.GroupId})
	if err != nil {
		T.Logger.Error(err.Error())
		tr.Rollback()
		return nil, err
	}
	for _, student := range students.Students {
		questions, err := T.mng.GetQuestionRandomly(context.Background(), &question.GetQuestionRandomlyRequest{
			TopicId: req.TopicId,
			Count:   int64(questionCount),
		})
		if err != nil {
			T.Logger.Error(err.Error())
			return nil, err // Bu yerda xato qaytarilganda tranzaksiya bekor qilinmaydi
		}
		for _, questionId := range questions {
			id := uuid.NewString() // Har bir yozuv uchun yangi ID
			_, err = tr.Exec(query2, id, student.HhId, req.TopicId, questionId)
			if err != nil {
				T.Logger.Error(err.Error())
				return nil, err // Bu yerda xato qaytarilganda tranzaksiya bekor qilinmaydi
			}
			T.Logger.Info(fmt.Sprintf("QuestionId: %s", questionId))
		}
	}
	if err = tr.Commit(); err != nil {
		T.Logger.Error(err.Error())
		return nil, err
	}
	return &pb.CreateTaskResp{
		TaskId:    id,
		CreatedAt: time.Now().String(),
	}, nil
}

func (T *TaskRepo) DeleteTask(req *pb.DeleteTaskReq) (*pb.DeleteTaskResp, error) {
	query := `
        UPDATE user_tasks
        SET deleted_at = $1
        WHERE id = $2 AND deleted_at IS NULL
    `

	result, err := T.DB.Exec(query, time.Now(), req.TaskId)
	if err != nil {
		T.Logger.Error(fmt.Sprintf("Error deleting task: %s", err))
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		T.Logger.Error(fmt.Sprintf("Error getting rows affected: %s", err))
		return nil, err
	}

	var status string
	if rowsAffected == 0 {
		status = "Task not found or already deleted"
	} else {
		status = "Task successfully deleted"
	}

	return &pb.DeleteTaskResp{
		Status: status,
	}, nil
}

func (T *TaskRepo) GetTask(req *pb.GetTaskReq) ([]string, *string, error) {
	var questionIds []string
	var TaskId string
	query := `
				SELECT 
					question_id, id
				FROM
					user_tasks
				WHERE
					hh_id = $1 AND topic_id = $2 AND deleted_at IS NULL`
	
	tr, err := T.DB.Begin()
	if err != nil {
		T.Logger.Error("Error starting transaction: " + err.Error())
		return nil, nil, err
	}
	defer func() {
		if err != nil {
			T.Logger.Error("Rolling back transaction: " + err.Error())
			tr.Rollback()
		}
	}()

	rows, err := tr.Query(query, req.UserId, req.TopicId)
	if err != nil {
		T.Logger.Error("Error querying tasks: " + err.Error())
		return nil, nil, err
	}
	defer rows.Close() // Kursorni yopish

	for rows.Next() {
		var questionId string
		err = rows.Scan(&questionId, &TaskId)
		if err != nil {
			T.Logger.Error("Error scanning row: " + err.Error())
			return nil, nil, err
		}
		questionIds = append(questionIds, questionId)
	}

	// Check if TaskId is empty
	if TaskId == "" {
		T.Logger.Warn("TaskId is empty")
		return questionIds, nil, nil
	}

	// Transactionni tasdiqlash
	if err := tr.Commit(); err != nil {
		T.Logger.Error("Error committing transaction: " + err.Error())
		return nil, nil, err
	}

	return questionIds, &TaskId, nil
}

