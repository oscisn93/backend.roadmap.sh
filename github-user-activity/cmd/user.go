package cmd

type UserRequest struct {
	Username  string `json:"username"`
	Timestamp int64  `json:"timestamp"`
	Etag      string `json:"etag"`
	RateLimit int    `json:"rateLimit"`
}

func CreateUserRequest(username string) {
	return
}
