package config

import "os"

type Config struct {
	Port      string
	StaticDir string
	TemplDir  string
	DevMode   bool
}

func New() *Config {
	return &Config{
		Port:      getEnv("PORT", "8080"),
		StaticDir: getEnv("STATIC_DIR", "web/static"),
		TemplDir:  getEnv("TEMPL_DIR", "web/templates"),
		DevMode:   getEnv("DEV_MODE", "false") == "true",
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
