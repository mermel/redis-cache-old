package cache

import (
  "fmt"
  "os"

  "github.com/fzzy/radix/redis"
)

func Get(redisKey string) (usr string, err error) {
  client := OpenConnection()
  defer client.Close()

  usr, errGet := client.Cmd("GET", redisKey).Str()
  return usr, errGet
}

func Set(redisKey string, redisValue string) (err error) {
  client := OpenConnection()
  defer client.Close()

  status, errSet := client.Cmd("SET", redisKey, redisValue).Str()
  fmt.Println(status) // status
  return errSet
}

func Del(redisKey string) (err error) {
  client := OpenConnection()
  defer client.Close()

  status, errSet := client.Cmd("DEL", redisKey).Str()
  fmt.Println(status) // status
  return errSet
}

func OpenConnection()(client *redis.Client){
  client, errConnection := redis.Dial("tcp", os.Getenv("REDIS_URL"))
  if errConnection != nil {
    panic (errConnection)
  } else {
    fmt.Println("open connection")
  }

  return client
}
