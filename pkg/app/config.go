package app

// Config ...
type Config struct {
	DbUri  string `mapstructure:"DB_URI"`
	DbName string `mapstructure:"DB_NAME"`
}
