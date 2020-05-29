server:
	go run ./src/main.go

package:
	aws cloudformation package \
	--template-file ./infra/main.yaml \
	--s3-bucket aws-sam-cli-managed-default-samclisourcebucket-1k8g6cex8e00g \
	--output-template-file ./infra/packaged.template.yaml

deploy: package
	aws cloudformation deploy \
	--template-file ./infra/packaged.template.yaml \
	--stack-name fargate-demo2 \
	--capabilities CAPABILITY_AUTO_EXPAND CAPABILITY_IAM

