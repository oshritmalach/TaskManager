build:
	docker build -t oshrit-test .
run:
	docker run -d -p 8081:8081 oshrit-test
test:
	 go test ./repository ./handler -v
lint:
	 golangci-lint run

