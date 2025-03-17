package configs

type DatabaseType string

const (
	DatabaseTypePostgres DatabaseType = "postgres"
)

type Database struct {
	Type     DatabaseType `yaml:"type"`
	Host     string       `yaml:"host"`
	Port     int          `yaml:"port"`
	Username string       `yaml:"username"`
	Password string       `yaml:"password"`
	Database string       `yaml:"database"`
}
