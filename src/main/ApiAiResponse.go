package main


type ApiAiResponse struct {
	Lang string `json:"lang"`
	DisplayText string `json:"displayText"`
	Speech string `json:"speech"`
	/** json object */
	Data string `json:"data"`
	ContextOut []string `json:"contextOut"`
	Source string `json:"source"`
	FollowupEvent []string `json:"followupEvent"`
}

