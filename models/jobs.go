package models

type JobsDB struct {
	Title       string   `bson:"title"`
	CompanyName string   `bson:"companyName"`
	Content     string   `bson:"content"`
	Location    string   `bson:"location"`
	Stacks      []string `bson:"stacks"`
	StartDate   string   `bson:"startDate"`
	ExitDate    string   `bson:"exitDate,omitempty"`
}
