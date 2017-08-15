package configuration

import (
	"fmt"
	"log"
	"os"
)

// Dialect of the db to use for saving
func Dialect() string {
	return environmentVariable("DB_DIALECT")
}

// Connection string to connect to preferred db
func Connection() string {
	return environmentVariable("DATABASE_URL")
}

// Port to use for the server
func Port() string {
	return environmentVariable("PORT")
}

//GithubKey to use for oAuth2
func GithubKey() string {
	return environmentVariable("GITHUB_KEY")
}

//GithubSecret to use for oAuth2
func GithubSecret() string {
	return environmentVariable("GITHUB_SECRET")
}

//GithubAuthCallback url for oAuth2
func GithubAuthCallback() string {
	return fmt.Sprintf("%s/auth/github/callback", environmentVariable("CALLBACK_HOST"))
}

func environmentVariable(variable string) string {
	value := os.Getenv(variable)
	if value == "" {
		log.Fatal(fmt.Sprintf("$%s must be set", variable))
	}
	return value
}
