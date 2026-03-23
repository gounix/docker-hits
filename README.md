# Overview
The docker-hits container queries the pull statistics at docker.io. If a namespace is specified all imagis in that namespace will be reported. If an image is specified only that image will be reported.
The values are made available for scraping by prometheus. The scrape url is http://<service_address>:<PORT>/metrics. If you use the prometheus-operator deployment in combination with our helm chart the scrape config is not needed. The helm chart contains a serviceMonitor definition.

# Screenshot
![Grafana](https://raw.githubusercontent.com/gounix/docker-hits/main/grafana.png)

# Environment variables
The following enviroment variables are supported:
| Variable | Description |
| -------- | -------- |
| INTERVAL | The amount of seconds between successive polls of docker.io |
| PORT | The port that is used for publishing the metrics |
| NAMESPACE | The docker namespace to query, all images will be reported|
| IMAGE | The docker image to query|
only ine of NAMESPACE or IMAGE can be specified.

# Helm chart
[github](https://github.com/gounix/docker-hits/tree/main/helm-charts)

# Sources
[github](https://github.com/gounix/docker-hits/tree/main/src)

# Container
[docker hub](https://hub.docker.com/r/gounix/docker-hits)
