package user

import "time"

type UserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (u *UserRequest) User() User {
	return User{
		Name: u.Name,
	}
}

type UserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func response(u User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
