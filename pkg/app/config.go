package app

// Config ...
type Config struct {
	Port   int64  `mapstructure:"PORT"`
	JWTKey string `mapstructure:"JWT_KEY"`

	DbUri  string `mapstructure:"DB_URI"`
	DbName string `mapstructure:"DB_NAME"`
}
