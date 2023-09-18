package user

import "time"

type Tabler interface {
  TableName() string
}

type UserRequest struct {
	Name 				string      `json:"name" validate:"required"`
	Email 			string     `json:"email" validate:"required"`
	Password 		string      `json:"password" validate:"required"`
	CreatedAt		time.Time   `json:"created_at"`
	UpdatedAt		time.Time   `json:"updated_at"`
}

func (UserRequest) TableName() string {
  return "users"
}