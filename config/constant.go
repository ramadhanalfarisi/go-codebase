package config

const (
	LOCAL_FILEPATH        string = `C:\bahaya\Berkembang\go\src\github.com\ramadhanalfarisi\go-codebase-kocak` // local log filepath
	FILEPATH              string = `./`                                                                        // production / development log filepath
	DEBUG                 bool   = true                                                                        // is debug ?
	MIGRATIONS_LOCAL_PATH string = "file://../migrations/"                                                     // migrations local path
	MIGRATIONS_PATH       string = "file://./migrations/"                                                      // migrations production / development path
)
