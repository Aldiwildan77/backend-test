package models

// CreateUserRequest model
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"-"`
	FullName string `json:"fullname"`
}

// UpdateUserRequest model
type UpdateUserRequest struct {
	Password string `json:"password"`
	FullName string `json:"fullname"`
}

// ReadUserResponse model
type ReadUserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Photo    string `json:"photo"`
}

// CreateUserResponse model
type CreateUserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
}

// UpdateUserResponse model
type UpdateUserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Photo    string `json:"photo"`
}
