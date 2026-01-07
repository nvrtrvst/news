package config

import "github.com/spf13/viper"

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`

	JwtSecretKey string `json:"jwt_secret_key"`
	JwtIssuer    string `json:"jwt_issuer"`
}
type PsqlDB struct {
	Host           string `json:"host"`
	Port           string `json:"port"`
	User           string `json:"user"`
	Password       string `json:"password"`
	DBName         string `json:"db_name"`
	DBMaxOpenConns int    `json:"db_max_open_conns"`
	DBMaxIdleConns int    `json:"db_max_idle_conns"`
}

type CloudflareR2 struct {
	Name      string `json:"endpoint"`
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	Token     string `json:"token"`
	AccountID string `json:"account_id"`
	PublicURL string `json:"public_url"`
}

type Config struct {
	App  App          `json:"app"`
	Psql PsqlDB       `json:"psql_db"`
	R2   CloudflareR2 `json:"cloudflare_r2"`
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort:      viper.GetString("APP_PORT"),
			AppEnv:       viper.GetString("APP_ENV"),
			JwtSecretKey: viper.GetString("JWT_SECRET_KEY"),
			JwtIssuer:    viper.GetString("JWT_ISSUER"),
		},
		Psql: PsqlDB{
			Host:           viper.GetString("DB_HOST"),
			Port:           viper.GetString("DB_PORT"),
			User:           viper.GetString("DB_USER"),
			Password:       viper.GetString("DB_PASSWORD"),
			DBName:         viper.GetString("DB_NAME"),
			DBMaxOpenConns: viper.GetInt("DB_MAX_OPEN_CONNS"),
			DBMaxIdleConns: viper.GetInt("DB_MAX_IDLE_CONNS"),
		},
		R2: CloudflareR2{
			Name:      viper.GetString("CLOUDFLARE_R2_NAME"),
			ApiKey:    viper.GetString("CLOUDFLARE_R2_API_KEY"),
			ApiSecret: viper.GetString("CLOUDFLARE_R2_API_SECRET"),
			Token:     viper.GetString("CLOUDFLARE_R2_TOKEN"),
			AccountID: viper.GetString("CLOUDFLARE_R2_ACCOUNT_ID"),
			PublicURL: viper.GetString("CLOUDFLARE_R2_PUBLIC_URL"),
		},
	}
}
