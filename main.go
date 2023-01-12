package main

import (
	"flag"
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

func createGroup() *gache.Group {
	return gache.NewGroup("scores", 2<<10, gache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB search key]", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, g *gache.Group) {
	peers := gache.NewHTTPPool(addr)
	peers.Set(addrs...)
	g.RegisterPeers(peers)
	log.Println("gache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, g *gache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := g.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())
		}))
	log.Println("frontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "Gache server port")
	flag.BoolVar(&api, "api", false, "Start an api server?")
	flag.Parse()

	apiAddr := "http://localhost:8080"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	g := createGroup()
	if api {
		go startAPIServer(apiAddr, g)
	}
	startCacheServer(addrMap[port], []string(addrs), g)
}
