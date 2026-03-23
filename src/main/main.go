package main

import (
	"fmt"
	"docker-hits/producer"
	"docker-hits/consumer"
	"os"
	"strconv"
)
var (
	port int
	interval int
	namespace string
	image string
)

func getReqNumFromEnv(variable string, deflt int) int {

	var (
		err error
		str string
		value int
	)

	str = os.Getenv(variable)
	if str == "" {
		return deflt
	}
	value, err = strconv.Atoi(str)
	if err != nil {
		fmt.Printf("FATAL main.getEnv(%s) %s", variable, err)
		os.Exit(1)
	}
	return value
}

func getEnv() {

	port = getReqNumFromEnv("PORT", 8080)
	interval = getReqNumFromEnv("INTERVAL", 300)

	namespace = os.Getenv("NAMESPACE")
	image = os.Getenv("IMAGE")
	if namespace == "" && image == ""{
		fmt.Println("FATAL: environment variable NAMESPACE or IMAGE is required")
		os.Exit(1)
	}

	fmt.Printf("Environment:\nPORT: %d\nINTERVAL: %d\nNAMESPACE: %s\nIMAGE: %s\n", port, interval, namespace, image)
}

func main() {
	getEnv()
	go producer.Put(interval, namespace, image)
	consumer.Get(port, interval)
}
