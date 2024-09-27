package mongosh

import (
	"context"

	pb "question/genproto/question"
	"question/storage/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InputRepository struct {
	Coll *mongo.Collection
}

func NewInputRepository(db *mongo.Database) repo.IInputStorage {
	return &InputRepository{Coll: db.Collection("input")}
}

type QuestionInput struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"` // MongoDB ObjectID
	QuestionID string             `bson:"question_id"`   // Associated question ID
	Input      string             `bson:"input"`         // The input field
}

func (repo *InputRepository) CreateQuestionInput(ctx context.Context, req *pb.CreateQuestionInputRequest) (*pb.QuestionInputId, error) {
	input := QuestionInput{
		QuestionID: req.QuestionId,
		Input:      req.Input,
	}

	// Insert the document into MongoDB
	result, err := repo.Coll.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}

	// Return the inserted document ID as a response
	return &pb.QuestionInputId{Id: result.InsertedID.(primitive.ObjectID).Hex()}, nil
}

func (repo *InputRepository) GetQuestionInput(ctx context.Context, id *pb.QuestionInputId) (*pb.GetQuestionInputResponse, error) {
	var input QuestionInput

	// Convert string ID to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, err
	}

	// Find the document by its ObjectID
	err = repo.Coll.FindOne(ctx, bson.M{"_id": objID}).Decode(&input)
	if err != nil {
		return nil, err
	}

	// Return the result as a proto response
	return &pb.GetQuestionInputResponse{
		Id:         input.ID.Hex(),
		QuestionId: input.QuestionID,
		Input:      input.Input,
	}, nil
}

func (repo *InputRepository) GetAllQuestionInputsByQuestionId(ctx context.Context, req *pb.GetAllQuestionInputsByQuestionIdRequest) (*pb.GetAllQuestionInputsByQuestionIdResponse, error) {
	filter := bson.M{"question_id": req.QuestionId}

	cursor, err := repo.Coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var inputs []*pb.GetQuestionInputResponse
	for cursor.Next(ctx) {
		var input QuestionInput

		if err := cursor.Decode(&input); err != nil {
			return nil, err
		}
		inputs = append(inputs, &pb.GetQuestionInputResponse{
			Id:         input.ID.Hex(),
			QuestionId: input.QuestionID,
			Input:      input.Input,
		})
	}

	return &pb.GetAllQuestionInputsByQuestionIdResponse{
		QuestionInputs: inputs,
	}, nil
}

func (repo *InputRepository) UpdateQuestionInput(ctx context.Context, req *pb.UpdateQuestionInputRequest) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	// Update the question input fields
	update := bson.M{
		"$set": bson.M{
			"question_id": req.QuestionId,
			"input":       req.Input,
		},
	}

	// Update the document in the collection
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (repo *InputRepository) DeleteQuestionInput(ctx context.Context, req *pb.DeleteQuestionInputRequest) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	_, err = repo.Coll.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
