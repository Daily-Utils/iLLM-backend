package utils

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/daily-utils/iLLM-backend/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func MongoInsertResponse(ctx context.Context, mongoClient *mongo.Client, llmResponse models.Response, database string, collectionName string) error {
	dbResponse := models.DBResponse{
		Model:                llmResponse.Model,
		Created_at:           llmResponse.Created_at,
		Response:             llmResponse.Response,
		Done:                 llmResponse.Done,
		Total_Duration:       llmResponse.Total_Duration,
		Load_Duration:        llmResponse.Load_Duration,
		Prompt_Eval_Count:    llmResponse.Prompt_Eval_Count,
		Prompt_Eval_Duration: llmResponse.Prompt_Eval_Duration,
		Eval_Count:           llmResponse.Eval_Count,
		Eval_Duration:        llmResponse.Eval_Duration,
		Context:              []byte(base64.StdEncoding.EncodeToString(ConvertInt64ToBytesArr(llmResponse.Context))),
	}

	collection := mongoClient.Database(database).Collection(collectionName)

	_, err := collection.InsertOne(ctx, dbResponse)
	if err != nil {
		log.Println("Error inserting response into MongoDB:", err)
		return err
	}

	return nil
}
