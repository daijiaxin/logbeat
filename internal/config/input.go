package config

type Input struct {
	Type           string   `json:"type"`
	Codec          string   `json:"codec"`
	Path           []string `json:"path"`
	Broker         []string `json:"broker"`
	Topic          []string `json:"topic"`
	GroupId        string   `json:"group_id"`
	CommitInterval int      `json:"commit_interval"`
}
