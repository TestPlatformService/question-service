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
	InputInfo   string             `bson:"input_info"`
	OutputInfo  string             `bson:"output_info"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	DeletedAt   *time.Time         `bson:"deleted_at,omitempty"`
	Language    string             `bson:"language"`
	TimeLimit   int64              `bson:"time_limit"`
	MemoryLimit int64              `bson:"memory_limit"`
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
		InputInfo:   req.InputInfo,
		OutputInfo:  req.OutputInfo,
		Language:    req.Language,    // New field
		TimeLimit:   req.TimeLimit,   // New field
		MemoryLimit: req.MemoryLimit, // New field
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
		Id:          question.ID.Hex(),
		TopicId:     question.TopicID,
		Type:        question.Type,
		Name:        question.Name,
		Number:      question.Number,
		Difficulty:  question.Difficulty,
		Description: question.Description,
		Image:       question.Image,
		Constrains:  question.Constraints,
		InputInfo:   question.InputInfo,
		OutputInfo:  question.OutputInfo,
		Language:    question.Language,    // New field
		TimeLimit:   question.TimeLimit,   // New field
		MemoryLimit: question.MemoryLimit, // New field
		CreatedAt:   question.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   question.UpdatedAt.Format(time.RFC3339),
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
		filter["name"] = bson.M{"$regex": req.Name, "$options": "i"}
	}
	if req.Number > 0 {
		filter["number"] = req.Number
	}
	if req.Difficulty != "" {
		filter["difficulty"] = req.Difficulty
	}
	if req.Language != "" {
		filter["language"] = req.Language
	}
	filter["deleted_at"] = bson.M{"$exists": false}

	totalCount, err := repo.Coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	skip := (req.Offset - 1) * req.Limit
	if skip < 0 {
		skip = 0
	}

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
			Id:          question.ID.Hex(),
			TopicId:     question.TopicID,
			Type:        question.Type,
			Name:        question.Name,
			Number:      question.Number,
			Difficulty:  question.Difficulty,
			Description: question.Description,
			Image:       question.Image,
			Constrains:  question.Constraints,
			InputInfo:   question.InputInfo,
			OutputInfo:  question.OutputInfo,
			Language:    question.Language,    // New field
			TimeLimit:   question.TimeLimit,   // New field
			MemoryLimit: question.MemoryLimit, // New field
			CreatedAt:   question.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   question.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &pb.GetAllQuestionsResponse{
		Questions: questions,
		Total:     totalCount,
		Page:      skip,
	}, nil
}

func (repo *QuestionRepository) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.Void, error) {
	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err // return an error if the ID is not valid
	}

	update := bson.M{
		"$set": bson.M{
			"type":         req.Type,
			"name":         req.Name,
			"number":       req.Number,
			"difficulty":   req.Difficulty,
			"description":  req.Description,
			"image":        req.Image,
			"constrains":   req.Constrains,
			"input_info":   req.InputInfo,
			"output_info":  req.OutputInfo,
			"language":     req.Language,
			"time_limit":   req.TimeLimit,
			"memory_limit": req.MemoryLimit,
			"updated_at":   time.Now(),
		},
	}

	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": objectID, "deleted_at": bson.M{"$exists": false}}, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

// DeleteQuestion marks a question as deleted
func (repo *QuestionRepository) DeleteQuestion(ctx context.Context, req *pb.DeleteQuestionRequest) (*pb.Void, error) {
	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err // return an error if the ID is not valid
	}

	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	}

	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (repo *QuestionRepository) UploadImageQuestion(ctx context.Context, req *pb.UploadImageQuestionRequest) (*pb.Void, error) {
	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.QuestionId)
	if err != nil {
		return nil, err // return an error if the ID is not valid
	}

	update := bson.M{
		"$set": bson.M{
			"image": req.Image,
		},
	}

	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": objectID, "deleted_at": bson.M{"$exists": false}}, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (repo *QuestionRepository) DeleteImageQuestion(ctx context.Context, req *pb.DeleteImageQuestionRequest) (*pb.Void, error) {
	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.QuestionId)
	if err != nil {
		return nil, err // return an error if the ID is not valid
	}

	update := bson.M{
		"$unset": bson.M{
			"image": "", // This will remove the image field from the document
		},
	}

	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": objectID, "deleted_at": bson.M{"$exists": false}}, update)
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
