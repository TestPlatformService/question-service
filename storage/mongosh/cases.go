package mongosh

import (
	"context"

	pb "question/genproto/question"
	"question/storage/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestCase struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	QuestionID string             `bson:"question_id"`
	Case       string             `bson:"case"`
	IsCorrect  bool               `bson:"is_correct"`
}

type TestCaseRepository struct {
	Coll *mongo.Collection
}

func NewTestCaseRepository(db *mongo.Database) repo.ITestCaseStorage {
	return &TestCaseRepository{Coll: db.Collection("testcase")}
}

func (repo *TestCaseRepository) CreateTestCase(ctx context.Context, req *pb.CreateTestCaseRequest) (*pb.TestCaseId, error) {
	// Convert request data to the TestCase struct
	testCase := TestCase{
		QuestionID: req.QuestionId,
		Case:       req.Case,
		IsCorrect:  req.IsCorrect,
	}

	// Insert into MongoDB collection
	result, err := repo.Coll.InsertOne(ctx, testCase)
	if err != nil {
		return nil, err
	}

	// Return the newly created TestCaseId
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &pb.TestCaseId{Id: id}, nil
}

func (repo *TestCaseRepository) GetTestCase(ctx context.Context, req *pb.TestCaseId) (*pb.GetTestCaseResponse, error) {
	var testCase TestCase

	// Convert the ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	// Find the test case in the MongoDB collection
	err = repo.Coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(&testCase)
	if err != nil {
		return nil, err
	}

	// Return the test case as a response
	return &pb.GetTestCaseResponse{
		Id:         testCase.ID.Hex(),
		QuestionId: testCase.QuestionID,
		Case:       testCase.Case,
		IsCorrect:  testCase.IsCorrect,
	}, nil
}

func (repo *TestCaseRepository) GetAllTestCasesByQuestionId(ctx context.Context, req *pb.GetAllTestCasesByQuestionIdRequest) (*pb.GetAllTestCasesByQuestionIdResponse, error) {
	filter := bson.M{"question_id": req.QuestionId}

	cursor, err := repo.Coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var testCases []*pb.GetTestCaseResponse
	for cursor.Next(ctx) {
		var testCase TestCase
		if err := cursor.Decode(&testCase); err != nil {
			return nil, err
		}

		testCases = append(testCases, &pb.GetTestCaseResponse{
			Id:         testCase.ID.Hex(),
			QuestionId: testCase.QuestionID,
			Case:       testCase.Case,
			IsCorrect:  testCase.IsCorrect,
		})
	}

	return &pb.GetAllTestCasesByQuestionIdResponse{TestCases: testCases}, nil
}

func (repo *TestCaseRepository) UpdateTestCase(ctx context.Context, req *pb.UpdateTestCaseRequest) (*pb.Void, error) {
	// Convert ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	// Update the test case fields
	update := bson.M{
		"$set": bson.M{
			"question_id": req.QuestionId,
			"case":        req.Case,
			"is_correct":  req.IsCorrect,
		},
	}

	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (repo *TestCaseRepository) DeleteTestCase(ctx context.Context, req *pb.DeleteTestCaseRequest) (*pb.Void, error) {
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
