package config

// import (
// 	"github.com/ilyakaznacheev/cleanenv"
// 	"github.com/joho/godotenv"
// )

// type AppSettings struct {
// 	ServerRPCPort     int64 `env:"SERVER_RPC_PORT" env-default:"7001"`
// 	ServerIngressPort int64 `env:"SERVER_INGRESS_PORT" env-default:"7002"`
// 	ClientForwardPort int64 `env:"CLIENT_FORWARD_PORT" env-default:"7003"`
// }

// var (
// 	Settings *AppSettings
// )

// func init() {
// 	Settings = &AppSettings{}
// 	godotenv.Load()
// 	cleanenv.ReadEnv(Settings)
// }

// func Get() *AppSettings {
// 	return Settings
// }
