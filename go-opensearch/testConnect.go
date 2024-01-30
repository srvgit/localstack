package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

const (
	host    = "http://my-local-domain.us-east-1.opensearch.localhost.localstack.cloud:4566"
	region  = "us-east-1"
	service = "es"
)

func main() {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	// Create an STS client
	stsSvc := sts.NewFromConfig(cfg)

	// Retrieve credentials
	// creds := stscreds.NewAssumeRoleProvider(stsSvc, "arn:aws:iam::123456789012:role/my-role")
	//

	credentials := stscreds.NewAssumeRoleProvider(stsSvc, "arn:aws:iam::123456789012:role/my-role")
	cfg.Credentials = credentials

	// Create a static credentials provider
	fmt.Println(credentials.Retrieve(context.Background()))

	// Create a v4 signer

	v4Signer := v4.NewSigner(func(signer *v4.SignerOptions) {})

	// Create a v4 signer

	// Example Usage
	indexName := "sample-index"
	sampleDocument := `{
	        "title": "Hello OpenSearch",
	        "content": "This is a test document."
	    }`

	createIndex(indexName, v4Signer, &cfg)
	addDocument(indexName, sampleDocument, "1", v4Signer, &cfg)
}

func createIndex(indexName string, signer v4.HTTPSigner, cfg *aws.Config) {
	url := host + "/" + indexName
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		panic("request creation failed, " + err.Error())
	}

	signHTTPRequest(req, signer, cfg)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("request failed, " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		fmt.Printf("Index '%s' created successfully.\n", indexName)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Error creating index: %s\n", string(body))
	}
}

func addDocument(indexName, document, docID string, signer v4.HTTPSigner, cfg *aws.Config) {
	url := fmt.Sprintf("%s/%s/_doc/%s", host, indexName, docID)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(document))
	if err != nil {
		panic("request creation failed, " + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	signHTTPRequest(req, signer, cfg)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("request failed, " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		fmt.Printf("Document added to '%s' successfully.\n", indexName)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Error adding document: %s\n", string(body))
	}
}

func signHTTPRequest(req *http.Request, signer v4.HTTPSigner, cfg *aws.Config) {
	// Sign the request
	creds, err := cfg.Credentials.Retrieve(context.TODO())
	if err != nil {
		panic("request signing failed, " + err.Error())
	}
	err = signer.SignHTTP(context.TODO(), creds, req, "", service, region, time.Now())
	if err != nil {
		panic("request signing failed, " + err.Error())
	}
}
