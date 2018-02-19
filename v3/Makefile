.PHONY: db
db: db-stop
	docker run -d \
	--name=roach1 \
	-p 26257:26257 \
	-p 8080:8080 \
	-v "${PWD}/../cockroach-data/roach1:/cockroach/cockroach-data" \
	cockroachdb/cockroach:v1.1.5 \
	start \
	--insecure

.PHONY: db-stop
db-stop:
	-docker stop roach1
	-docker rm roach1

.PHONY: build
build:
	GOOS=linux go build -o app .
	docker build -t app:v3 .
	rm app

.PHONY: run
run: build
	docker run --rm \
	-e "APP_DIR=/var/lib/data" \
	-e "APP_HOST=172.16.4.179" \
	-e "APP_PORT=26257" \
	-e "APP_USERNAME=maxroach" \
	-e "APP_DATABASE=bank" \
	app:v3