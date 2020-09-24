package dto

type UserTimeline struct {
	GitHubID string `json:"gitHubID"`
}

func NewUserTimeline() *UserTimeline {
	return &UserTimeline{}
}
