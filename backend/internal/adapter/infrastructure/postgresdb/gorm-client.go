// internal/adapter/infrastructure/postgresdb/gorm_client.go
package postgresdb

import (
	"fmt"
	"guessing-game/internal/domain/models"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

// Config holds database configuration
type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

// DefaultConfig returns default database configuration
func DefaultConfig() *Config {
	return &Config{
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		User:     getEnvOrDefault("DB_USER", "postgres"),
		Password: getEnvOrDefault("DB_PASSWORD", "password"),
		DBName:   getEnvOrDefault("DB_NAME", "postgres"),
		Port:     getEnvOrDefault("DB_PORT", "5432"),
		SSLMode:  getEnvOrDefault("DB_SSLMODE", "disable"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// BuildDSN builds the database connection string
func (c *Config) BuildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode)
}

// InitGormClient initializes the database connection with retry mechanism
func InitGormClient() {
	config := DefaultConfig()
	once.Do(func() {
		var err error
		maxRetries := 5
		retryDelay := time.Second * 5

		// Configure GORM logger
		gormLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		)

		// Try to connect with retries
		for i := 0; i < maxRetries; i++ {
			dbInstance, err = gorm.Open(postgres.Open(config.BuildDSN()), &gorm.Config{
				Logger: gormLogger,
			})

			if err == nil {
				log.Println("Successfully connected to database")

				// Auto migrate the models (include the City model)
				if err := autoMigrate(dbInstance); err != nil {
					log.Printf("Auto migration failed: %v", err)
					panic(err)
				}
				return
			}

			log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
			if i < maxRetries-1 {
				log.Printf("Retrying in %v...", retryDelay)
				time.Sleep(retryDelay)
			}
		}

		panic(fmt.Sprintf("Failed to connect to database after %d attempts: %v", maxRetries, err))
	})
}

func autoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &models.User{},
        &models.Destination{},
        &models.Challenge{},
    )
}


// GetGormClient returns the already initialized Gorm DB instance.
func GetGormClient() *gorm.DB {
	if dbInstance == nil {
		panic("Database not initialized. Call InitGormClient() first.")
	}
	return dbInstance
}
