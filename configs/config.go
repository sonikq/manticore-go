package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP      HTTPConf
	DB        PostgresConfig
	Manticore ManticoreConfig

	// context timeout in seconds
	CtxTimeout int

	LogLevel string
}

type HTTPConf struct {
	Host string
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type ManticoreConfig struct {
	Host        string
	Port        uint16
	IndexerPort uint16
	Limit       int
	MaxMatches  int
	IndexerUrl  string
	MergeUrl    string
	SearchUrl   string
	CLIUrl      string
}

// Load loads configs
func Load() (cfg Config) {

	if err := godotenv.Load("debug.env"); err != nil {
		log.Fatal(err)
		return
	}

	cfg.HTTP.Host = cast.ToString(os.Getenv("HTTP_HOST"))
	cfg.HTTP.Port = cast.ToString(os.Getenv("HTTP_PORT"))

	cfg.DB.Host = cast.ToString(os.Getenv("POSTGRES_HOST"))
	cfg.DB.Port = cast.ToString(os.Getenv("POSTGRES_PORT"))
	cfg.DB.Username = cast.ToString(os.Getenv("POSTGRES_USERNAME"))
	cfg.DB.Password = cast.ToString(os.Getenv("POSTGRES_PASSWORD"))
	cfg.DB.DBName = cast.ToString(os.Getenv("POSTGRES_DB"))
	cfg.DB.SSLMode = cast.ToString(os.Getenv("POSTGRES_SSLMODE"))

	cfg.Manticore.Host = cast.ToString(os.Getenv("MANTICORE_HOST"))
	cfg.Manticore.Port = cast.ToUint16(os.Getenv("MANTICORE_PORT"))
	cfg.Manticore.IndexerPort = cast.ToUint16(os.Getenv("MANTICORE_INDEXER_PORT"))
	cfg.Manticore.Limit = cast.ToInt(os.Getenv("MANTICORE_LIMIT"))
	cfg.Manticore.MaxMatches = cast.ToInt(os.Getenv("MAX_MATCHES"))
	cfg.Manticore.IndexerUrl = fmt.Sprintf("http://%s:%d/indexer/rotateindex", cfg.Manticore.Host, cfg.Manticore.IndexerPort)
	cfg.Manticore.MergeUrl = fmt.Sprintf("http://%s:%d/indexer/mergeindexes", cfg.Manticore.Host, cfg.Manticore.IndexerPort)
	cfg.Manticore.SearchUrl = fmt.Sprintf("http://%s:%d/json/search", cfg.Manticore.Host, cfg.Manticore.Port)
	cfg.Manticore.CLIUrl = fmt.Sprintf("http://%s:%d/cli", cfg.HTTP.Host, cfg.Manticore.Port)

	cfg.CtxTimeout = cast.ToInt(os.Getenv("CTX_TIMEOUT"))

	cfg.LogLevel = cast.ToString(os.Getenv("LOG_LEVEL"))

	return cfg
}
