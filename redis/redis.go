package redis

import (
  "fmt"
  "github.com/gomodule/redigo/redis"
  "log"
)


var conn redis.Conn

func GetConn() redis.Conn{

  if (conn == nil) {
      conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", 6379))

      if err != nil {
        log.Println("connect to redis error ", err)
        return nil
      }

      return conn
  }

  return conn
}

func Get(key string) (cache interface{}) {

    c := GetConn()
    cache, err := c.Do("GET", key);

    if err != nil {
      log.Println(" redis get failure", err)
      return nil
    }

    return cache
}

func Set(key string, val string, expire int) bool {
    c := GetConn()
    _, err := c.Do("SET", key, val, "EX", expire);

    if err != nil {
        log.Println(" redis set failure", err)
        return false
    }

    return true
}

func Close()  {
    c := GetConn()
    c.Close();
}
