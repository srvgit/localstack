import requests
from requests_aws4auth import AWS4Auth
import boto3

# Your OpenSearch domain endpoint
host = 'http://my-local-domain.us-east-1.opensearch.localhost.localstack.cloud:4566'  # For example: 'https://search-mydomain.us-west-1.es.amazonaws.com'

# AWS Region
region = 'us-east-1'  # For example: 'us-west-1'

# AWS Credentials
# AWS Credentials
service = 'es'
session = boto3.Session()
credentials = session.get_credentials()
awsauth = AWS4Auth(credentials.access_key, credentials.secret_key, region, service, session_token=credentials.token)

# Function to create an index
def create_index(index_name):
    url = host + '/' + index_name
    response = requests.put(url, auth=awsauth)

    if response.status_code == 200 or response.status_code == 201:
        print(f"Index '{index_name}' created successfully.")
    else:
        print(f"Error creating index: {response.text}")

# Function to add a document to an index
def add_document(index_name, document, doc_id):
    url = f"{host}/{index_name}/_doc/{doc_id}"
    headers = { "Content-Type": "application/json" }
    response = requests.post(url, auth=awsauth, json=document, headers=headers)

    if response.status_code == 200 or response.status_code == 201:
        print(f"Document added to '{index_name}' successfully.")
    else:
        print(f"Error adding document: {response.text}")

# Example Usage
index_name = "sample-index"
sample_document = {
    "title": "Hello OpenSearch",
    "content": "This is a test document."
}

create_index(index_name)
add_document(index_name, sample_document, doc_id=1)