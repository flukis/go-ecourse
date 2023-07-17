package domain

import "e-course/pkg/resp"

type ProfileRequestLogoutBody struct {
	Authorization string `header:"authorization"`
}

type ProfileResponseBody struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}

func CreateProfileResponse(user User) ProfileResponseBody {
	isVerified := false

	if user.EmailVerifiedAt != nil {
		isVerified = true
	}

	return ProfileResponseBody{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		IsVerified: isVerified,
	}
}

type ProfileUsecase interface {
	FindProfile(id int) (*ProfileResponseBody, *resp.ErrorResp)
	Update(id int, dto UserUpdateRequestBody) (*User, *resp.ErrorResp)
	Deactive(id int) *resp.ErrorResp
	Logout(accessToken string) *resp.ErrorResp
}
