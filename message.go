package main

type DiscordJson struct {
	Username  string `json:"username,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Content   string `json:"content,omitempty"`
}
