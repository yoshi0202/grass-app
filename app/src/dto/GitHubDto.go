package dto

import (
	"encoding/json"

	"gorm.io/datatypes"
)

type GitHub struct {
	Count string `json:"count"`
	Date  string `json:"date"`
}

type GitHubs []GitHub

func NewGitHub() *GitHub {
	return &GitHub{}
}

func NewGitHubs() *GitHubs {
	return &GitHubs{}
}

func (g *GitHub) ToJSON() string {
	json, _ := json.Marshal(g)
	return string(json)
}

func (g *GitHubs) ToJSON() datatypes.JSON {
	json, _ := json.Marshal(g)
	return json
}
