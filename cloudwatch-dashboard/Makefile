.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world

deploy: build
	sam deploy \
	--template-file ./template.yaml \
	--stack-name cloudwatchDashboard \
	--capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND