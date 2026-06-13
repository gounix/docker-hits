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
