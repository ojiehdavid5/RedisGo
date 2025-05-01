package main

import (
	"fmt"
	"net/http"
	"github.com/redis/go-redis/v9"
	"context"
	"log"
	"time"

)
func main() {
	var ctx = context.Background()

	opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
 panic(err)
}

client := redis.NewClient(opt)

// Set a key-value pair
err := client.Set(ctx, "greeting", "Hello, Redis!", 0).Err()
if err != nil {
 log.Fatal(err)
}


// Get the value of the key
val, err := client.Get(ctx, "greeting").Result()
if err != nil {	
	 log.Fatal(err)
}
fmt.Println("Value of 'greeting':", val)










	
// fmt.Println("Building Rest API using Go and REDIS")
// mux:= http.NewServeMux();
// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// fmt.Fprint(w, "Hello, World!")
// })

// if err:= http.ListenAndServe(":8080", mux); err != nil {
// 	fmt.Print(err.Error())
// }
}