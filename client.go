import (
    "fmt"
    "time"

    "github.com/go-redis/redis"
)


var client *redis.Client

func init() {
    client = redis.NewClient(&redis.Options{
        Addr:         "localhost:6379",
        Password: "", // no password set
        DB: 1, // use default DB
    })

    result, err := client.BLPop(1*time.Second, "queue").Result()
    if err != nil {
	panic(err)
    }
    
    fmt.Println(result[0], result[1])
}

