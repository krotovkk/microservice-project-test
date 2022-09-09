package config

const (
	DataGrpcServerPort       = 8001
	ValidationGrpcServerPort = 8002

	RedisHost     = "localhost"
	RedisPort     = 6379
	RedisDB       = 0
	RedisPassword = ""
)

var Brokers = []string{"localhost:9095", "localhost:9096", "localhost:9097"}
