package jobs

import (
	"context"
	"portfolio/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobsService struct {
	mongoDB *mongo.Database
}

func NewJobsService(mongoDB *mongo.Database) *JobsService {
	return &JobsService{
		mongoDB: mongoDB,
	}
}

func (ctx *JobsService) GetAll() ([]models.JobsDB, error) {
	jobColletion := ctx.mongoDB.Collection("job")

	ctxBg := context.Background()

	jobsDb, err := jobColletion.Find(ctxBg, bson.M{})
	if err != nil {
		return []models.JobsDB{}, err
	}
	defer jobsDb.Close(ctxBg)

	var jobs []models.JobsDB
	for jobsDb.Next(ctxBg) {
		var job models.JobsDB
		if err := jobsDb.Decode(&job); err != nil {
			return []models.JobsDB{}, err
		}

		jobs = append(jobs, job)
	}

	if err := jobsDb.Err(); err != nil {
		return []models.JobsDB{}, err
	}

	if jobs == nil {
		jobs = []models.JobsDB{}
	}

	return jobs, nil
}
