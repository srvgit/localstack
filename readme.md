<!-- START LOCAL STACK: -->
localstack start 
<!-- if it's not already running. -->
check opensearch -
aws --endpoint-url=http://localhost:4566 --region us-east-1 opensearch describe-domain --domain-name my-local-domain
create if not exists:
aws --endpoint-url=http://localhost:4566 --region us-east-1 opensearch create-domain --domain-name my-local-domain

Elastic search endpoint:
my-local-domain.us-east-1.opensearch.localhost.localstack.cloud:4566


#Gateway :

aws --endpoint-url=http://localhost:4566 --profile localstack apigateway create-rest-api --name 'MyAPI'



#EKS
aws confgiure 
srujans-mbp localstack % aws configure
AWS Access Key ID [****************test]: test
AWS Secret Access Key [****************test]: test
Default region name [us-east-1]: us-east-1
Default output format [json]: json

export AWS_ENDPOINT_URL=http://localhost:4566

aws iam create-role --role-name eksServiceRole --assume-role-policy-document file://eks-trust-policy.json


aws --endpoint-url=http://localhost:4566 --profile localstack eks create-cluster --name my-cluster --role-arn arn:aws:iam::000000000000:role/eksServiceRole --resources-vpc-config subnetIds=subnet-12345,securityGroupIds=sg-12345

check status of cluster:
aws eks describe-cluster --name my-cluster --query cluster.status

configure :
aws eks update-kubeconfig --name my-cluster


create cluster :
eksctl create cluster --name my-cluster --region us-west-2 --nodegroup-name my-nodes --node-type t3.medium --nodes 3

