ARTIFACT_PATH=./artifacts/hacienda

.PHONY: build clean

build:
	go build -o $(ARTIFACT_PATH) cmd/hacienda/main.go

clean:
	rm -Rf ./artifacts
