package main

import (
	"fmt"
	"log"
	"net/http"

	"gache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	gache.NewGroup("scores", 2<<10, gache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB search key]", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:8080"
	peers := gache.NewHTTPPool(addr)
	log.Println("gache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
