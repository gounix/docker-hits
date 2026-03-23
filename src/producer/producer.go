package producer

import (
	"docker-hits/data"
	"net/http"
	"io"
	"time"
	"fmt"
)

func GetData(namespace string, image string) {

	var url string

	if image == "" {
		url = fmt.Sprintf("https://hub.docker.com/v2/repositories/%s", namespace)
	} else {
		url = fmt.Sprintf("https://hub.docker.com/v2/repositories/%s", image)
	}
	fmt.Printf("producer.GetData: url=%s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("producer.GetData(%s,%s) error %s\n", namespace, image, err)
	} else {
		body, err2 := io.ReadAll(resp.Body)
		if err2 == nil {
			data.Put(body)
		} else {
			fmt.Printf("body err %d\n", err2)
		}
		resp.Body.Close()
	}
}

func Put(interval int, namespace string, image string) {
	for {
		GetData(namespace, image)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
