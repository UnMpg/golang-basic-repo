package models

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

type EmailStruck struct {
	From           string `json:"from"`
	To             string `json:"to"`
	Subject        string `json:"subject"`
	Body           string `json:"body"`
	AddAlternative string `json:"AddAlternative"`
}
