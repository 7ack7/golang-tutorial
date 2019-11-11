package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken) // to set env in os: export <ENV_NAME>=<TOKEN>
)

//GetGithubAccessToken func
func GetGithubAccessToken() string {
	return githubAccessToken
}
