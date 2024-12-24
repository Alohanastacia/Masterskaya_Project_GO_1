install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml

install-uber-mock:	
	go	install	go.uber.org/mock/mockgen@latest	
mock-processors:	
	mockgen	-source=internal/processors/complaints.go	-destination=internal/processors/mocks/mocks.go
