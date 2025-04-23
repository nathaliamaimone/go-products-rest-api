package config

import (
    "os"
    "path/filepath"
    "strconv"
    "github.com/joho/godotenv"
)

var (
    SecretKey string
    TokenExpiration int
)

func LoadConfig() error {

    currentDir, err := os.Getwd()
    if err != nil {
        return err
    }

    if filepath.Base(currentDir) == "cmd" {
        currentDir = filepath.Dir(currentDir)
    }

    if loadErr := godotenv.Load(filepath.Join(currentDir, ".env")); loadErr != nil {
        return err
    }

    SecretKey = os.Getenv("JWT_SECRET_KEY")
    expiration, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
    if err != nil {
        TokenExpiration = 24 // default value
    } else {
        TokenExpiration = expiration
    }

    return nil
}
