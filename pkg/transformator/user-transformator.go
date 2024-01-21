package transformator

import "caching-user-app/models"

func ToUserResponse(u models.User) models.UserResponse {
	userResponse := models.UserResponse{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		Surname:   u.Surname,
	}

	if !u.DeletedAt.Time.IsZero() {
		userResponse.DeletedAt = &u.DeletedAt.Time
	}

	return userResponse
}
