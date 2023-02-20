package app

// Config ...
type Config struct {
	Port int64 `mapstructure:"PORT"`

	DbUri  string `mapstructure:"DB_URI"`
	DbName string `mapstructure:"DB_NAME"`
}
