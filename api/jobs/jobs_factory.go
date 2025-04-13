package jobs

import "portfolio/api/infra/db"

func MakeJobsController() *JobsController {
	jobsService := NewJobsService(db.GetMongoDB())
	jobsController := NewJobsController(jobsService)

	return jobsController
}
