package dto

import (
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Grass struct {
	gorm.Model
	GitHubID  string         `json:"githubId"`
	CountDate datatypes.JSON `json:"countDate"`
}

type Grasses []Grass

func NewGrass() *Grass {
	return &Grass{}
}

func NewGrasses() *Grasses {
	return &Grasses{}
}

func (g *Grass) ToJSON() string {
	json, _ := json.Marshal(g)
	return string(json)
}

func (g *Grasses) ToJSON() string {
	json, _ := json.Marshal(g)
	return string(json)
}
