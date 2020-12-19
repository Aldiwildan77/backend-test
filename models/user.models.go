package models

// ReadUserResponse model
type ReadUserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Photo    string `json:"photo"`
}
