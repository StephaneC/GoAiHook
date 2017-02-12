package main

type Result struct {
	Parameters map[string][]string
	Context []string `json:"context"`
	ResolvedQuery string `json:"resolvedQuery"`
	Source string `json:"source"`
	Score float32 `json:"score"`
	ActionIncomplete bool `json:"actionIncomplete"`
	Action string `json:"action"`

}

type ApiAiRequest struct {
	Speech string `json:"speech"`
	Lang string `json:"lang"`
	Result Result `json:"result"`
}
