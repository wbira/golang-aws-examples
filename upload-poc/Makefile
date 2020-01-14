.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./upload-url/upload-url

build:
	GOOS=linux GOARCH=amd64 go build -o upload-url/upload-url ./upload-url

deploy: build
	sam deploy \
	--template-file ./template.yaml \
	--stack-name upload-poc \
	--capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND