package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

type Config struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Bucket    string `yaml:"bucket"`
	Region    string `yaml:"region"`
	Endpoint  string `yaml:"endpoint"`
	S3Client  *s3.S3
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось загрузить .env файл")
	}

	config := &Config{
		AccessKey: os.Getenv("YANDEX_ACCESS_KEY"),
		SecretKey: os.Getenv("YANDEX_SECRET_KEY"),
		Bucket:    os.Getenv("YANDEX_BUCKET"),
		Region:    os.Getenv("region"),
		Endpoint:  os.Getenv("endpoint"),
	}
	if config.AccessKey == "" || config.SecretKey == "" || config.Bucket == "" {
		log.Fatal("Переменные окружения YANDEX_ACCESS_KEY, YANDEX_SECRET_KEY и YANDEX_BUCKET обязательны")
	}

	var sess *session.Session
	sess, err = session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(
			config.AccessKey,
			config.SecretKey,
			"",
		),
		Endpoint:         aws.String(config.Endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("Ошибка при создании сессии AWS: %v", err)
	}

	config.S3Client = s3.New(sess)
	return config
}
