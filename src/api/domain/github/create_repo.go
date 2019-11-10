package github

//CreateRepoRequest struct
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

//CreateRepoResponse struct
type CreateRepoResponse struct {
	ID           int64           `json:"id"`
	Name         string          `json:"name"`
	FullName     string          `json:"full_name"`
	Owner        RepoOwner       `json:"owner"`
	Perimissions RepoPermissions `json:"permissions"`
}

//RepoOwner struct in CreateRepoResponse
type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

//RepoPermissions struct in CreateRepoResponse
type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPull bool `json:"has_pull"`
	HasPush bool `json:"has_push"`
}
