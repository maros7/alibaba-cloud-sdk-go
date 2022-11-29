package sae

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateSecret invokes the sae.CreateSecret API synchronously
func (client *Client) CreateSecret(request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
	response = CreateCreateSecretResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSecretWithChan invokes the sae.CreateSecret API asynchronously
func (client *Client) CreateSecretWithChan(request *CreateSecretRequest) (<-chan *CreateSecretResponse, <-chan error) {
	responseChan := make(chan *CreateSecretResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSecret(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateSecretWithCallback invokes the sae.CreateSecret API asynchronously
func (client *Client) CreateSecretWithCallback(request *CreateSecretRequest, callback func(response *CreateSecretResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSecretResponse
		var err error
		defer close(result)
		response, err = client.CreateSecret(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateSecretRequest is the request struct for api CreateSecret
type CreateSecretRequest struct {
	*requests.RoaRequest
	SecretType  string `position:"Query" name:"SecretType"`
	NamespaceId string `position:"Query" name:"NamespaceId"`
	SecretData  string `position:"Body" name:"SecretData"`
	SecretName  string `position:"Query" name:"SecretName"`
}

// CreateSecretResponse is the response struct for api CreateSecret
type CreateSecretResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateSecretRequest creates a request to invoke CreateSecret API
func CreateCreateSecretRequest() (request *CreateSecretRequest) {
	request = &CreateSecretRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "CreateSecret", "/pop/v1/sam/secret/secret", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSecretResponse creates a response to parse from CreateSecret response
func CreateCreateSecretResponse() (response *CreateSecretResponse) {
	response = &CreateSecretResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
