package mongosh

import (
	"context"

	pb "question/genproto/question"
	"question/storage/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionOutput struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	QuestionID string             `bson:"question_id"`
	InputID    string             `bson:"input_id"`
	Answer     string             `bson:"answer"`
}

type OutputRepository struct {
	Coll *mongo.Collection
}

func NewOutputRepository(db *mongo.Database) repo.IOutputStorage {
	return &OutputRepository{Coll: db.Collection("output")}
}

func (repo *OutputRepository) CreateQuestionOutput(ctx context.Context, req *pb.CreateQuestionOutputRequest) (*pb.QuestionOutputId, error) {
	// Convert request data to the QuestionOutput struct
	output := QuestionOutput{
		QuestionID: req.QuestionId,
		InputID:    req.InputId,
		Answer:     req.Answer,
	}

	// Insert into MongoDB collection
	result, err := repo.Coll.InsertOne(ctx, output)
	if err != nil {
		return nil, err
	}

	// Return the newly created QuestionOutputId
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &pb.QuestionOutputId{Id: id}, nil
}

func (repo *OutputRepository) GetQuestionOutput(ctx context.Context, req *pb.QuestionOutputId) (*pb.GetQuestionOutputResponse, error) {
	var output QuestionOutput

	// Find a single question output by its ID
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	err = repo.Coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(&output)
	if err != nil {
		return nil, err
	}

	// Return the question output in response format
	return &pb.GetQuestionOutputResponse{
		Id:         output.ID.Hex(),
		QuestionId: output.QuestionID,
		InputId:    output.InputID,
		Answer:     output.Answer,
	}, nil
}

func (repo *OutputRepository) GetAllQuestionOutputsByQuestionId(ctx context.Context, req *pb.GetAllQuestionOutputsByQuestionIdRequest) (*pb.GetAllQuestionOutputsByQuestionIdResponse, error) {
	filter := bson.M{"question_id": req.QuestionId}

	cursor, err := repo.Coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var outputs []*pb.GetQuestionOutputResponse
	for cursor.Next(ctx) {
		var output QuestionOutput
		if err := cursor.Decode(&output); err != nil {
			return nil, err
		}

		outputs = append(outputs, &pb.GetQuestionOutputResponse{
			Id:         output.ID.Hex(),
			QuestionId: output.QuestionID,
			InputId:    output.InputID,
			Answer:     output.Answer,
		})
	}

	return &pb.GetAllQuestionOutputsByQuestionIdResponse{QuestionOutputs: outputs}, nil
}

func (repo *OutputRepository) UpdateQuestionOutput(ctx context.Context, req *pb.UpdateQuestionOutputRequest) (*pb.Void, error) {
	// Convert ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	// Update the question output fields
	update := bson.M{
		"$set": bson.M{
			"question_id": req.QuestionId,
			"input_id":    req.InputId,
			"answer":      req.Answer,
		},
	}

	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (repo *OutputRepository) DeleteQuestionOutput(ctx context.Context, req *pb.DeleteQuestionOutputRequest) (*pb.Void, error) {
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	_, err = repo.Coll.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (repo *OutputRepository) GetQuestionOutputByInputId(ctx context.Context, req *pb.GetQUestionOutPutByInputIdRequest) (*pb.GetQUestionOutPutByInputIdRes, error) {
	filter := bson.M{"input_id": req.InputId}

	cursor, err := repo.Coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var questionOutputs []*pb.GetQuestionOutputResponse
	for cursor.Next(ctx) {
		var output QuestionOutput
		if err := cursor.Decode(&output); err != nil {
			return nil, err
		}

		questionOutputs = append(questionOutputs, &pb.GetQuestionOutputResponse{
			Id:         output.ID.Hex(),
			QuestionId: output.QuestionID,
			InputId:    output.InputID,
			Answer:     output.Answer,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.GetQUestionOutPutByInputIdRes{
		QuestionOutputs: questionOutputs,
	}, nil
}
