package dto

import (
	"encoding/json"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login  string `json:"login"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Users []User

type OAuthUser struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           interface{} `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               string      `json:"bio"`
	TwitterUsername   interface{} `json:"twitter_username"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

type OAuthUsers []OAuthUser

func NewUser() *User {
	return &User{}
}

func NewUsers() *Users {
	return &Users{}
}

func NewOAuthUser() *OAuthUser {
	return &OAuthUser{}
}

func NewOAuthUsers() *OAuthUsers {
	return &OAuthUsers{}
}

func (u *User) ToJSON() string {
	json, _ := json.Marshal(u)
	return string(json)
}

func (u *Users) ToJSON() datatypes.JSON {
	json, _ := json.Marshal(u)
	return json
}

func (u *OAuthUser) ToJSON() string {
	json, _ := json.Marshal(u)
	return string(json)
}
