rungo:
	go run cmd/prweb/main.go

dockerbuild:
	docker build -t getting-started .

dockerrun:
	docker run -dp 8080:8080 getting-started