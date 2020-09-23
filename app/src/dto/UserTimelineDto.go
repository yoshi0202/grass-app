package dto

type UserTimeline struct {
	RegistUserID string `json:"registUserID"`
}

func NewUserTimeline() *UserTimeline {
	return &UserTimeline{}
}
