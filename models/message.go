package models

type Message struct {
	ClientID string `json:"clientID"`
	Text     string `json:"text"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Time     string `json:"time"`
}
