package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress  string      `json:"server_address"`
	DataDir        string      `json:"data_dir"`
	BackupDir      string      `json:"backup_dir"`
	JWTSecret      string      `json:"jwt_secret"`
	Email          EmailConfig `json:"email"`
	BaseURL        string      `json:"base_url"`
	TOTPEncryptKey string      `json:"totp_encrypt_key"`
	SkipSSLVerify  bool        `json:"skip_ssl_verify"`
}

type EmailConfig struct {
	Enabled     bool   `json:"enabled"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	FromEmail   string `json:"from_email"`
	FromName    string `json:"from_name"`
	ReplyTo     string `json:"reply_to"`
	EnableTLS   bool   `json:"enable_tls"`
	RequireAuth bool   `json:"require_auth"`
}

func Load() (*Config, error) {
	// Инициализация конфига по умолчанию
	cfg := &Config{
		ServerAddress: ":8080",
		DataDir:       "./data",
		BackupDir:     "./backup",
		JWTSecret:     "your-jwt-secret",
		Email: EmailConfig{
			Enabled:     false,
			Host:        "smtp.example.com",
			Port:        587,
			Username:    "your-email@example.com",
			Password:    "your-email-password",
			FromEmail:   "Denis@example.com",
			FromName:    "Denis",
			ReplyTo:     "your-email@example.com",
			EnableTLS:   true,
			RequireAuth: true,
		},
		BaseURL:        "http://localhost:8080",
		TOTPEncryptKey: "your-totp-encrypt-key",
		SkipSSLVerify:  false,
	}
	// Загрузка конфига из файла
	envPath := ".env"
	if _, err := os.Stat(envPath); err == nil {
		if err := godotenv.Load(envPath); err != nil {
			return nil, err
		}
		// Загрузка переменных окружения
		if serverAddress := os.Getenv("SERVER_ADDRESS"); serverAddress != "" {
			cfg.ServerAddress = serverAddress
		}
		if dataDir := os.Getenv("DATA_DIR"); dataDir != "" {
			cfg.DataDir = dataDir
		}
		if backupDir := os.Getenv("BACKUP_DIR"); backupDir != "" {
			cfg.BackupDir = backupDir
		}
		if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
			cfg.JWTSecret = jwtSecret
		}
		if emailHost := os.Getenv("EMAIL_HOST"); emailHost != "" {
			cfg.Email.Host = emailHost
		}
		if emailPort, err := strconv.Atoi(os.Getenv("EMAIL_PORT")); err == nil {
			cfg.Email.Port = emailPort
		}
		if emailUsername := os.Getenv("EMAIL_USERNAME"); emailUsername != "" {
			cfg.Email.Username = emailUsername
		}
		if emailPassword := os.Getenv("EMAIL_PASSWORD"); emailPassword != "" {
			cfg.Email.Password = emailPassword
		}
		if emailFromEmail := os.Getenv("EMAIL_FROM_EMAIL"); emailFromEmail != "" {
			cfg.Email.FromEmail = emailFromEmail
		}
		if emailFromName := os.Getenv("EMAIL_FROM_NAME"); emailFromName != "" {
			cfg.Email.FromName = emailFromName
		}
		if emailReplyTo := os.Getenv("EMAIL_REPLY_TO"); emailReplyTo != "" {
			cfg.Email.ReplyTo = emailReplyTo
		}
		if emailEnableTLS, err := strconv.ParseBool(os.Getenv("EMAIL_ENABLE_TLS")); err == nil {
			cfg.Email.EnableTLS = emailEnableTLS
		}
		if emailRequireAuth, err := strconv.ParseBool(os.Getenv("EMAIL_REQUIRE_AUTH")); err == nil {
			cfg.Email.RequireAuth = emailRequireAuth
		}
		if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
			cfg.BaseURL = baseURL
		}
		if totpEncryptKey := os.Getenv("TOTP_ENCRYPT_KEY"); totpEncryptKey != "" {
			cfg.TOTPEncryptKey = totpEncryptKey
		}
		if skipSSLVerify, err := strconv.ParseBool(os.Getenv("SKIP_SSL_VERIFY")); err == nil {
			cfg.SkipSSLVerify = skipSSLVerify
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	// Создание директории для хранения данных, если она не существует
	// Если директория уже существует, то ошибки не будет
	if err := os.MkdirAll(cfg.DataDir, 0755); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(cfg.BackupDir, 0755); err != nil {
		return nil, err
	}

	return cfg, nil
}
