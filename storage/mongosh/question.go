package mongosh

import (
	"context"
	pb "question/genproto/question"
	"question/storage/repo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Question struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	TopicID     string             `bson:"topic_id"`
	Type        string             `bson:"type"`
	Name        string             `bson:"name"`
	Number      int64              `bson:"number"`
	Difficulty  string             `bson:"difficulty"`
	Description string             `bson:"description"`
	Image       string             `bson:"image"`
	Constraints string             `bson:"constrains"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	DeletedAt   *time.Time         `bson:"deleted_at,omitempty"`
}

type QuestionRepository struct {
	Coll *mongo.Collection
}

func NewQuestionRepository(db *mongo.Database) repo.IQuestionStorage {
	return &QuestionRepository{Coll: db.Collection("question")}
}

func (repo *QuestionRepository) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.QuestionId, error) {
	question := Question{
		TopicID:     req.TopicId,
		Type:        req.Type,
		Name:        req.Name,
		Number:      req.Number,
		Difficulty:  req.Difficulty,
		Description: req.Description,
		Image:       req.Image,
		Constraints: req.Constrains,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	res, err := repo.Coll.InsertOne(ctx, question)
	if err != nil {
		return nil, err
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, mongo.ErrInvalidIndexValue
	}

	return &pb.QuestionId{Id: oid.Hex()}, nil
}
func (repo *QuestionRepository) GetQuestion(ctx context.Context, id *pb.QuestionId) (*pb.GetQuestionResponse, error) {
	var question Question
	err := repo.Coll.FindOne(ctx, bson.M{"_id": id.Id, "deleted_at": bson.M{"$exists": false}}).Decode(&question)
	if err != nil {
		return nil, err
	}

	return &pb.GetQuestionResponse{
		TopicId:     question.TopicID,
		Type:        question.Type,
		Name:        question.Name,
		Number:      question.Number,
		Difficulty:  question.Difficulty,
		Description: question.Description,
		Image:       question.Image,
		Constrains:  question.Constraints,
	}, nil
}

func (repo *QuestionRepository) GetAllQuestions(ctx context.Context, req *pb.GetAllQuestionsRequest) (*pb.GetAllQuestionsResponse, error) {
	filter := bson.M{}

	if req.TopicId != "" {
		filter["topic_id"] = req.TopicId
	}
	if req.Type != "" {
		filter["type"] = req.Type
	}
	if req.Name != "" {
		filter["name"] = bson.M{"$regex": req.Name, "$options": "i"} // Case insensitive regex for name
	}
	if req.Number > 0 {
		filter["number"] = req.Number
	}
	if req.Difficulty != "" {
		filter["difficulty"] = req.Difficulty
	}

	filter["deleted_at"] = bson.M{"$exists": false}

	totalCount, err := repo.Coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Calculate pagination values
	skip := (req.Offset - 1) * req.Limit
	if skip < 0 {
		skip = 0 // Ensure skip is not negative
	}

	// Find the questions with limit and offset for pagination
	cursor, err := repo.Coll.Find(ctx, filter, options.Find().SetLimit(req.Limit).SetSkip(int64(skip)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var questions []*pb.GetQuestionResponse
	for cursor.Next(ctx) {
		var question Question
		if err := cursor.Decode(&question); err != nil {
			return nil, err
		}

		questions = append(questions, &pb.GetQuestionResponse{
			TopicId:     question.TopicID,
			Type:        question.Type,
			Name:        question.Name,
			Number:      question.Number,
			Difficulty:  question.Difficulty,
			Description: question.Description,
			Image:       question.Image,
			Constrains:  question.Constraints,
		})
	}

	return &pb.GetAllQuestionsResponse{
		Questions: questions,
		Total:     totalCount,
		Page:      req.Offset, // Return the requested page number
	}, nil
}

// UpdateQuestion updates an existing question
func (repo *QuestionRepository) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.Void, error) {
	update := bson.M{
		"$set": bson.M{
			"type":        req.Type,
			"name":        req.Name,
			"number":      req.Number,
			"difficulty":  req.Difficulty,
			"description": req.Description,
			"image":       req.Image,
			"constrains":  req.Constrains,
			"updated_at":  time.Now(),
		},
	}

	_, err := repo.Coll.UpdateOne(ctx, bson.M{"_id": req.Id, "deleted_at": bson.M{"$exists": false}}, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

// DeleteQuestion marks a question as deleted
func (repo *QuestionRepository) DeleteQuestion(ctx context.Context, req *pb.DeleteQuestionRequest) (*pb.Void, error) {
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	}

	_, err := repo.Coll.UpdateOne(ctx, bson.M{"_id": req.Id}, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

// UploadImageQuestion uploads an image for a question
func (repo *QuestionRepository) UploadImageQuestion(ctx context.Context, req *pb.UploadImageQuestionRequest) (*pb.Void, error) {
	update := bson.M{
		"$set": bson.M{
			"image": req.Image,
		},
	}

	_, err := repo.Coll.UpdateOne(ctx, bson.M{"_id": req.QuestionId, "deleted_at": bson.M{"$exists": false}}, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

// DeleteImageQuestion removes the image from a question
func (repo *QuestionRepository) DeleteImageQuestion(ctx context.Context, req *pb.DeleteImageQuestionRequest) (*pb.Void, error) {
	update := bson.M{
		"$unset": bson.M{
			"image": "",
		},
	}

	_, err := repo.Coll.UpdateOne(ctx, bson.M{"_id": req.QuestionId, "deleted_at": bson.M{"$exists": false}}, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (repo *QuestionRepository) IsQuestionExist(ctx context.Context, id *pb.QuestionId) (*pb.Void, error) {
	filter := bson.M{
		"_id":        id.Id,
		"deleted_at": bson.M{"$exists": false},
	}

	count, err := repo.Coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &pb.Void{}, nil
}