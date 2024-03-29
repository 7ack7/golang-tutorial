package github

//ErrorResponse struct
type ErrorResponse struct {
	StatusCode       int     `json:"status_code"`
	Message          string  `json:"message"`
	DocumentationURL string  `json:"documentation_url"`
	Errors           []Error `json:"errors"`
}

//Error struct in ErrorResponse
type Error struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
