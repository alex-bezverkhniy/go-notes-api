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

// APIDBPortEvnName - name of evn var with db port
const APIDBPortEvnName = "API_DB_PORT"

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

// DefaultAPIDBPort - default db port
const DefaultAPIDBPort = "3306"

func main() {
	a := NewApp(getDbUser(), getDbPassword(), getDbName(), getDbHost(), getDbPort())
	a.Run(getPort())
}

func getEnv(envName, defaultVal string) string {
	value := os.Getenv(envName)
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

func getDbPort() string {
	return getEnv(APIDBPortEvnName, DefaultAPIDBPort)
}

func getPort() string {
	return ":" + getEnv(APIPortEvnName, DefaultAPIPort)
}
