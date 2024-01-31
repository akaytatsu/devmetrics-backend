package config

import (
	"os"
	"strconv"
)

var EnvironmentVariables EnvironmentVars

func ReadEnvironmentVars() {
	// Set environment variables
	EnvironmentVariables.POSTGRES_DB = os.Getenv("POSTGRES_DB")
	EnvironmentVariables.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	EnvironmentVariables.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	EnvironmentVariables.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	EnvironmentVariables.POSTGRES_PORT, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	EnvironmentVariables.KAFKA_BOOTSTRAP_SERVER = os.Getenv("KAFKA_BOOTSTRAP_SERVER")
	EnvironmentVariables.KAFKA_CLIENT_ID = os.Getenv("KAFKA_CLIENT_ID")
	EnvironmentVariables.KAFKA_GROUP_ID = os.Getenv("KAFKA_GROUP_ID")

	EnvironmentVariables.MONGO_URL = os.Getenv("MONGO_URL")
	EnvironmentVariables.MONGO_DATABASE_NAME = os.Getenv("MONGO_DATABASE_NAME")

	EnvironmentVariables.JIRA_USER = os.Getenv("JIRA_USER")
	EnvironmentVariables.JIRA_TOKEN = os.Getenv("JIRA_TOKEN")
	EnvironmentVariables.JIRA_BASE_URL = os.Getenv("JIRA_BASE_URL")

	EnvironmentVariables.DEFAULT_ADMIN_MAIL = os.Getenv("DEFAULT_ADMIN_MAIL")
	EnvironmentVariables.DEFAULT_ADMIN_PASSWORD = os.Getenv("DEFAULT_ADMIN_PASSWORD")
}
