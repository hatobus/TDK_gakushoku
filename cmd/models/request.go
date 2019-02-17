package models

type PostReq struct {
	UserID   string `json:"userid"`
	Category int64  `json:"category"`
}
