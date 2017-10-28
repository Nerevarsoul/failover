package failover

import (
    "github.com/go-redis/redis"
)


var client *redis.Client

func GetKey()(string) {
    client = redis.NewClient(&redis.Options{
        Addr:         "172.17.0.2:6379",
        Password: "", // no password set
        DB: 0, // use default DB
    })

    // result, err := client.BLPop(1*time.Second, "queue").Result()
    result, err := client.Get("foo").Result()
    if err != nil {
	    panic(err)
    }

    return result
}

func SetKey() {
    client = redis.NewClient(&redis.Options{
        Addr:         "172.17.0.2:6379",
        Password: "", // no password set
        DB: 0, // use default DB
    })

    // result, err := client.BLPop(1*time.Second, "queue").Result()
    err := client.Set("foo", "bar", 0).Err()
    if err != nil {
	    panic(err)
    }
}
