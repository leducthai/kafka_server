package config

type Options struct {
	// Redis configs
	RedisConnect  string `long:"redis-connect" description:"the address of redis server" default:"198.1.1.222:6379" env:"REDIS_CONNECT"`
	RedisPassword string `long:"redis-password" description:"the password of redis server" default:"edit_your_own_password_for_redis" env:"REDIS_PASSWORD"`
}
