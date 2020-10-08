package nsq

// {"topics":["test"]}
type Topics struct {
	Topics []string `json:"topics"`
}

// {"channels":["test1"]}
type Channels struct {
	Channels []string `json:"channels"`
}
