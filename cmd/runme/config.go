package main

type Config struct {
	MysqlDSN string
	RedisDSN string
	// All configurations should be present here for this binary MAIN.go

	// All valuse should be passed to domain object
}

func LoadConfig() *Config {
	return &Config{}
}
