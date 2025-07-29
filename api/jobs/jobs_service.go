package jobs

import (
	"context"

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

func (ctx *JobsService) GetAll() ([]JobsDB, error) {
	jobColletion := ctx.mongoDB.Collection("jobs")

	ctxBg := context.Background()

	jobsDb, err := jobColletion.Find(ctxBg, bson.M{})
	if err != nil {
		return []JobsDB{}, err
	}

	defer jobsDb.Close(ctxBg)

	var jobs []JobsDB
	for jobsDb.Next(ctxBg) {
		var job JobsDB
		if err := jobsDb.Decode(&job); err != nil {
			return []JobsDB{}, err
		}

		jobs = append(jobs, job)
	}

	if err := jobsDb.Err(); err != nil {
		return []JobsDB{}, err
	}

	if jobs == nil {
		jobs = []JobsDB{}
	}

	return jobs, nil
}
