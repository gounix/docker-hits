module docker-hits/main

go 1.22.2

replace docker-hits/data => ../data

require docker-hits/data v0.0.0-00010101000000-000000000000 // indirect

replace docker-hits/producer => ../producer

require (
	docker-hits/consumer v0.0.0-00010101000000-000000000000
	docker-hits/producer v0.0.0-00010101000000-000000000000
)

replace docker-hits/consumer => ../consumer
