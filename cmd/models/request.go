package models

type PostReq struct {
	UserID   string `json:"userid"`
	Category string `json:"category"`
}
