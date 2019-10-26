package main

import (
	"LoadBalancingBackend/cfg"
	"LoadBalancingBackend/metric"
	"LoadBalancingBackend/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 0 {
		log.Fatal("Config not provided")
	}

	cfg.Init(args[0])
	go metric.ExposeMetrics()

	port := cfg.GetInt("port")
	mux := router.GetRouter()
	fmt.Println("Starting listening on port", port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), mux))
}
