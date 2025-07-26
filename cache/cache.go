package cache

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

// InitRedis initializes the Redis client
func InitRedis() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Failed to load .env file: ", err)
    }

    // Retrieve environment variables
    redisHost := os.Getenv("REDIS_HOST")
    redisPort := os.Getenv("REDIS_PORT")
    redisPassword := os.Getenv("REDIS_PASSWORD")
    redisDB := 0 // Default DB, can be made configurable if needed
    redisPoolSize := 500 // Default pool size, can be made configurable

    // Construct Redis address
    redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

    // Initialize Redis client
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: redisPassword, // Empty if no password
        DB:       redisDB,
        PoolSize: redisPoolSize,
    })

    // Verify Redis connection
    if _, err := RedisClient.Ping(ctx).Result(); err != nil {
        log.Fatal("Failed to connect to Redis: ", err)
    }
    log.Println("Redis connected!")
}

// GetURL retrieves a URL from Redis by short_code
func GetURL(shortCode string) (string, error) {
    cacheKey := "url:" + shortCode
    return RedisClient.Get(ctx, cacheKey).Result()
}

// SetURL caches a URL in Redis with a TTL
func SetURL(shortCode, originalURL string, ttl time.Duration) error {
    cacheKey := "url:" + shortCode
    return RedisClient.Set(ctx, cacheKey, originalURL, ttl).Err()
}