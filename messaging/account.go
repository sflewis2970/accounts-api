package messaging

const (
	RESULTS_DEFAULT int64 = -1
)

type UserRequest struct {
	UserID    string `json:"userid"`
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
}

type UserResponse struct {
	UserID    string `json:"userid"`
	Username  string `json:"username"`
	Command   string `json:"command"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"warning,omitempty"`
}
