package jobs

import "time"

type JobsContentDB struct {
	En string `json:"en" bson:"en"`
	Pt string `json:"pt" bson:"pt"`
}

type JobsDB struct {
	Title       string        `json:"title" bson:"title"`
	CompanyName string        `json:"companyName" bson:"companyName"`
	Content     JobsContentDB `json:"content" bson:"content"`
	Location    string        `json:"location" bson:"location"`
	Stacks      []string      `json:"stacks" bson:"stacks"`
	StartDate   time.Time     `json:"startDate" bson:"startDate"`
	ExitDate    *time.Time    `json:"exitDate,omitempty" bson:"exitDate,omitempty"`
}
