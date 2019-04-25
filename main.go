package main

import "os"

// APIPortEvnName - name of evn var with API port
const APIPortEvnName = "API_PORT"

// APIDBUserEvnName - name of evn var with db user
const APIDBUserEvnName = "API_DB_USER"

// APIDBPasswordEvnName - name of evn var with db password
const APIDBPasswordEvnName = "API_DB_PASSWORD"

// APIDBNameEvnName - name of evn var with db name
const APIDBNameEvnName = "API_DB_NAME"

// APIDBHostEvnName - name of evn var with db host
const APIDBHostEvnName = "API_DB_HOST"

// DefaultAPIPort - default port
const DefaultAPIPort = "8080"

// DefaultAPIDBUser - default db user
const DefaultAPIDBUser = "gonotes"

// DefaultAPIDBPassword - default db password
const DefaultAPIDBPassword = "1Q2w3e4r"

// DefaultAPIDBName - default db name
const DefaultAPIDBName = "gonotes"

// DefaultAPIDBHost - default db host
const DefaultAPIDBHost = "localhost"

func main() {
	a := NewApp(getDbUser(), getDbPassword(), getDbName(), getDbHost())
	a.Run(getPort())
}

func getEnv(envName, defaultVal string) string {
	value := os.Getenv(APIPortEvnName)
	if value == "" {
		return defaultVal
	}
	return value

}

func getDbUser() string {
	return getEnv(APIDBUserEvnName, DefaultAPIDBUser)
}

func getDbPassword() string {
	return getEnv(APIDBPasswordEvnName, DefaultAPIDBPassword)
}

func getDbName() string {
	return getEnv(APIDBNameEvnName, DefaultAPIDBName)
}

func getDbHost() string {
	return getEnv(APIDBHostEvnName, DefaultAPIDBHost)
}

func getPort() string {
	return ":" + getEnv(APIPortEvnName, DefaultAPIPort)
}
