package dto

type GetMeRequest struct{}

type GetMeResponse struct {
	User User `json:"user"`
}
