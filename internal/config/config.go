package config


import(
	"log"
	"os"

	"github.com/joho/godotenv"
)


// Config структура для хранения параметров базы данных
type Config struct {
    DBHost string
    DBUser string
    DBPass string
    DBPort string
    DBName string
}

// LoadEnv загружает переменные окружения из .env файла
func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Ошибка при загрузке .env файла: %v", err)
    }
}

// getEnv получает переменную окружения по ключу
func getEnv(key string) string {
    value := os.Getenv(key)
    if value == "" {
        log.Fatalf("Переменная окружения %s не найдена", key)
    }
    return value
}

func NewConfig() *Config {
    return &Config{
        DBHost: getEnv("DB_HOST"),
        DBUser: getEnv("DB_USER"),
        DBPass: getEnv("DB_PASS"),
        DBPort: getEnv("DB_PORT"),
        DBName: getEnv("DB_BASE"),
    }
}