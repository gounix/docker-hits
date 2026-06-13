/*
MIT License

Copyright (c) 2026 gounix

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
