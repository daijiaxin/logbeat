package config

type Output struct {
	Type      string   `json:"type"`
	Codec     string   `json:"codec"`
	Path      string   `json:"path"`
	Broker    []string `json:"broker"`
	Topic     string   `json:"topic"`
	Key       string   `json:"key"`
	Host      []string `json:"host"`
	Index     string   `json:"index"`
	IndexType string   `json:"index_type"`
}
