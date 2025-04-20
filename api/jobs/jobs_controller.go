package jobs

import (
	"fmt"
	"net/http"
	response "portfolio/api/utils"

	"github.com/gin-gonic/gin"
)

type JobsController struct {
	jobsService *JobsService
}

func NewJobsController(
	jobsService *JobsService,
) *JobsController {
	return &JobsController{
		jobsService: jobsService,
	}
}

func (ctx *JobsController) GetAll(c *gin.Context) {
	data, err := ctx.jobsService.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		response.Error(c, "Error to get jobs.")
		return
	}

	response.Data(c, data, "Jobs got successfully.", http.StatusOK)
}
