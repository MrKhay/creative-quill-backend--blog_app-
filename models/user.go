package models

import "time"

type UpdateAccDetailsRequest struct {
	UserID string `json:"user_id"`
	Value  string `json:"value"`
}

type UpdateAccountReq struct {
	UserID        string        `json:"user_id"`
	Firstname     string        `json:"first_name"`
	Lastname      string        `json:"last_name"`
	Displayname   string        `json:"display_name"`
	Username      string        `json:"username"`
	Email         string        `json:"email"`
	WebsiteUrl    WebsiteUrl    `json:"website_url"`
	Description   string        `json:"description"`
	ProfilePicUrl ProfilePicUrl `json:"profilepic_url"`
	HeaderPicUrl  HeaderPicUrl  `json:"headerpic_url"`
}

type FollowRequest struct {
	UserID       string    `json:"user_id"`
	AccountID    string    `json:"acc_tofollow_id"`
	DateFollowed time.Time `json:"date_followed"`
}

func NewFollowRequestFunc(req *FollowRequest) *FollowRequest {

	return &FollowRequest{
		UserID:       req.UserID,
		AccountID:    req.AccountID,
		DateFollowed: time.Now(),
	}
}

type UnFollowRequest struct {
	UserID    string `json:"user_id"`
	AccountID string `json:"acc_tounfollow_id"`
}

type GetFollowersDetail struct {
	UserID string `json:"user_id"`
}
type GetAccFollowingDetails struct {
	UserID string `json:"user_id"`
}
