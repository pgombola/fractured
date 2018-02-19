.PHONY: db
db:
	docker run -d \
	--name=roach1 \
	-p 26257:26257 \
	-p 8080:8080 \
	-v "${PWD/../}/cockroach-data/roach1:/cockroach/cockroach-data" \
	cockroachdb/cockroach:v1.1.5 \
	start \
	--insecure

.PHONY: build
build:
	GOOS=linux go build -o app .
	docker build -t app:v1 .
	rm app

.PHONY: run
run: build
	docker run --rm app:v1