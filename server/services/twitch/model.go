package twitch

type User struct {
	ID          string `json:"_id"`
	Bio         string `json:"bio"`
	CreatedAt   string `json:"created_at"`
	DisplayName string `json:"display_name"`
	Logo        string `json:"logo"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	UpdatedAt   string `json:"updated_at"`
}

type Users []User

type TwitchUserResponse struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Total   int    `json:"_total"`
	Users   Users  `json:"users"`
}

type Stream struct {
	Id      int     `json:"_id"`
	Game    string  `json:"game"`
	Viewers int     `json:"viewers"`
	Channel Channel `json:"channel"`
}

type Channel struct {
	Status      string `json:"status"`
	DisplayName string `json:"display_name"`
	Logo        string `json:"logo"`
	Url         string `json:"url"`
}

type TwitchStreamResponse struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Stream  Stream `json:"stream"`
}
