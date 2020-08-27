package main

import (
   "fmt"
   "github.com/go-redis/redis/v7"
   "log"
   "math/rand"
   "time"
)

func init() {
   rand.Seed(time.Now().Unix())
}

const key = "players"

func main() {
   c := redis.NewClient(&redis.Options{
      Addr: "localhost:6379",
   })
   c.Del(key) // remove the key from redis for clean start

   for i := 0; i < 10000; i++ {
      player := getWinnerPlayer()

      err := c.ZIncr(key, &redis.Z{
         Score:  1,
         Member: player,
      }).Err()

      if err != nil {
         log.Fatal(err)
      }
   }

   result, _ := c.ZRevRangeWithScores(key, 0, -1).Result()
   fmt.Println(result)

}

func getWinnerPlayer() string {
   players := []string{"Mohammad", "Ali", "John", "Abdullah", "Farida"}
   return players[rand.Intn(len(players))]
}
