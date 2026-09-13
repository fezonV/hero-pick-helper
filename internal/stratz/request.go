package stratz

type graphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}
