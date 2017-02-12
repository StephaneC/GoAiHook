package main


type ApiAiResponse struct {
	Lang string
	DisplayText string
	/** json object */
	Data string
	ContextOut []string
	Source string
	FollowupEvent []string
}

