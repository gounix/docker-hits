package consumer

import (
	"fmt"
	"net/http"
	"docker-hits/data"
	"encoding/json"
)

const promHeader = `# HELP docker_pull_count number of pulls of a docker repo
# TYPE docker_pull_count gauge
# HELP docker_star_count number of stars of a docker repo
# TYPE docker_star_count gauge`

var interval int

func logRequest(r *http.Request) {
	fmt.Printf("%s: %s %s %s\n", r.Host, r.Method, r.URL.Path, r.UserAgent())
}

func sendPromLines(w http.ResponseWriter, entry map[string]any) {
	var str string

	str = fmt.Sprintf("docker_pull_count{docker_namespace=\"%s\",repo=\"%s\"} %.1f\n", entry["namespace"], entry["name"], entry["pull_count"])
	fmt.Fprintf(w, str)
	fmt.Printf(str)
	str = fmt.Sprintf("docker_star_count{docker_namespace=\"%s\",repo=\"%s\"} %.1f\n", entry["namespace"], entry["name"], entry["star_count"])
	fmt.Fprintf(w, str)
	fmt.Printf(str)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
type repo map[string]interface{}
var dat repo

	logRequest(r)
	bytestr := data.Get()
	json_err := json.Unmarshal(bytestr, &dat)
	if json_err != nil {
		fmt.Println("unmarshal error")
	} else {
		fmt.Fprintln(w, promHeader)
		_, err := dat["results"]
		if err {
			// this is a namespace
			for _, value := range dat["results"].([]any) {
				entry := value.(map[string]any)
				sendPromLines(w, entry)
			}
		} else {
			// this is an image
			sendPromLines(w, dat)
		}
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	if data.Alive(interval) {
		fmt.Fprintf(w, "OK")
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func Get(port int, intrvl int) {
	interval = intrvl
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/health", healthHandler)
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}

