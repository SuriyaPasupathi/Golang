package config

import (
    "os"
)

type Config struct {
    GRPCServerAddress string
    HTTPServerAddress string
}

func LoadConfig() *Config {
    return &Config{
        GRPCServerAddress: getEnv("GRPC_SERVER_ADDRESS", ":50051"),
        HTTPServerAddress: getEnv("HTTP_SERVER_ADDRESS", ":8080"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
