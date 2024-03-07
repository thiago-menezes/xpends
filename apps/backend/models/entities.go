package models

type Spend struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int64  `json:"price"`
	Done     bool   `json:"done"`
	DueDate  string `json:"dueDate"`
}
