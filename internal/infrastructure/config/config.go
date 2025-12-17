package config

import "os"

type Config struct {
	ServiceName string
	HTTPPort    string
	GRPCPort    string
	MongoURI    string
	NatsURL     string
}

func Load() Config {
	return Config{
		ServiceName: getEnv("SERVICE_NAME", "service-template"),
		HTTPPort:    getEnv("HTTP_PORT", "8080"),
		GRPCPort:    getEnv("GRPC_PORT", "9090"),
		MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
		NatsURL:     getEnv("NATS_URL", "nats://localhost:4222"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
