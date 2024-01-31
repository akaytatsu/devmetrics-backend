package infra_jira

import (
	"app/config"

	"github.com/andygrunwald/go-jira"
)

func Connect() *jira.Client {

	JiraBaseUrl := config.EnvironmentVariables.JIRA_BASE_URL
	JiraUser := config.EnvironmentVariables.JIRA_USER
	JiraToken := config.EnvironmentVariables.JIRA_TOKEN

	tp := jira.BasicAuthTransport{
		Username: JiraUser,
		Password: JiraToken,
	}

	jiraClient, err := jira.NewClient(tp.Client(), JiraBaseUrl)

	if err != nil {
		panic(err)
	}

	return jiraClient
}
