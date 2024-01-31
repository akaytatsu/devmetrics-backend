package config

type EnvironmentVars struct {
	POSTGRES_DB       string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_PORT     int

	KAFKA_BOOTSTRAP_SERVER string
	KAFKA_CLIENT_ID        string
	KAFKA_GROUP_ID         string

	MONGO_URL           string
	MONGO_DATABASE_NAME string

	JIRA_USER     string
	JIRA_TOKEN    string
	JIRA_BASE_URL string

	DEFAULT_ADMIN_MAIL     string
	DEFAULT_ADMIN_PASSWORD string
}
