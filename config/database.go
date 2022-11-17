package config

const (
	ENVIRONMMENT string = "development" // your environment (testing, development, production)

	PORT_PRODDUCTION   int    = 8080  // database port for production
	HOST_PRODDUCTION   string = "prod_host" // database host for production
	UNAME_PRODDUCTION  string = "prod_uname" // database username for production
	PASS_PRODDUCTION   string = "prod_password" // database password for production
	DBNAME_PRODDUCTION string = "prod_dbname" // database name for production

	PORT_DEVELOPMENT   int    = 8081  // database port for development
	HOST_DEVELOPMENT   string = "dev_host" // database host for development
	UNAME_DEVELOPMENT  string = "dev_uname" // database username for development
	PASS_DEVELOPMENT   string = "dev_password" // database password for development
	DBNAME_DEVELOPMENT string = "dev_dbname" // database name for development

	PORT_TESTING   int    = 8082  // database port for testing
	HOST_TESTING   string = "test_host" // database host for testing
	UNAME_TESTING  string = "test_uname" // database username for testing
	PASS_TESTING   string = "test_password" // database password for testing
	DBNAME_TESTING string = "test_dbname" // database name for testing
)
