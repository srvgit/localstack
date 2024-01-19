<!-- START LOCAL STACK: -->
localstack start 
<!-- if it's not already running. -->
check opensearch -
aws --endpoint-url=http://localhost:4566 --region us-east-1 opensearch describe-domain --domain-name my-local-domain
create if not exists:
aws --endpoint-url=http://localhost:4566 --region us-east-1 opensearch create-domain --domain-name my-local-domain


my-local-domain.us-east-1.opensearch.localhost.localstack.cloud:4566



