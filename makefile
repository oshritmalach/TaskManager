build:
	docker build -t task-manager .
run:
	docker run -d -p 8081:8081 task-manager
test:
	 go test ./repository ./handler -v
lint:
	 golangci-lint run

