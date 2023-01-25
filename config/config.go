package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment              string // develop, staging, production
	CtxTimeout               int    // context timeout in seconds
	LogLevel                 string
	HTTPPort                 string
	PostgresHost             string
	PostgresPort             string
	PostgresUser             string
	PostgresPassword         string
	PostgresDatabase         string
	IpInfoToken              string
	YouTubeApiKey            string
	SignInKey                string
	AuthConfigPath           string
	CSVFilePath              string
	OpenAiApiKey             string
	RatesCountToBeCalculated int
	CronTimeLapse            int // in seconds
}

// Load loads environment vars and inflates Config
func Load() Config {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	c := Config{}
	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(GetOrReturnDefault("HTTP_PORT", ":8000"))
	c.CtxTimeout = cast.ToInt(GetOrReturnDefault("CTX_TIMEOUT", 10))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABASE", "universities"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "azizbek"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "Azizbek"))
	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", "5432"))
	c.IpInfoToken = cast.ToString(GetOrReturnDefault("IP_INFO_TOKEN", "YOUIPINFOTOKEN"))
	c.YouTubeApiKey = cast.ToString(GetOrReturnDefault("YOUTUBE_API_KEY", "youryoutubeapikey"))
	c.AuthConfigPath = cast.ToString(GetOrReturnDefault("AUTH_CONFIG_PATH", "./config/auth.conf"))
	c.CSVFilePath = cast.ToString(GetOrReturnDefault("CSV_FILE_PATH", "./config/auth.csv"))
	c.SignInKey = cast.ToString(GetOrReturnDefault("SIGN_IN_KEY", "aksdjf;asdjkfas;dklAsdfaWEAFadfae"))
	c.OpenAiApiKey = cast.ToString(GetOrReturnDefault("OPENAI_API_KEY", "yourOpenAiAPIKey"))
	c.RatesCountToBeCalculated = cast.ToInt(GetOrReturnDefault("RATES_COUNT_TO_BE_CALCULATED", 10))
	c.CronTimeLapse = cast.ToInt(GetOrReturnDefault("CRON_TIME_LAPS", 10))
	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
