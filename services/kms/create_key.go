package kms

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

// CreateKey invokes the kms.CreateKey API synchronously
func (client *Client) CreateKey(request *CreateKeyRequest) (response *CreateKeyResponse, err error) {
	response = CreateCreateKeyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateKeyWithChan invokes the kms.CreateKey API asynchronously
func (client *Client) CreateKeyWithChan(request *CreateKeyRequest) (<-chan *CreateKeyResponse, <-chan error) {
	responseChan := make(chan *CreateKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateKey(request)
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

// CreateKeyWithCallback invokes the kms.CreateKey API asynchronously
func (client *Client) CreateKeyWithCallback(request *CreateKeyRequest, callback func(response *CreateKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateKeyResponse
		var err error
		defer close(result)
		response, err = client.CreateKey(request)
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

// CreateKeyRequest is the request struct for api CreateKey
type CreateKeyRequest struct {
	*requests.RpcRequest
	ProtectionLevel         string           `position:"Query" name:"ProtectionLevel"`
	KeyUsage                string           `position:"Query" name:"KeyUsage"`
	Origin                  string           `position:"Query" name:"Origin"`
	Description             string           `position:"Query" name:"Description"`
	KeySpec                 string           `position:"Query" name:"KeySpec"`
	RotationInterval        string           `position:"Query" name:"RotationInterval"`
	EnableAutomaticRotation requests.Boolean `position:"Query" name:"EnableAutomaticRotation"`
}

// CreateKeyResponse is the response struct for api CreateKey
type CreateKeyResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	KeyMetadata KeyMetadata `json:"KeyMetadata" xml:"KeyMetadata"`
}

// CreateCreateKeyRequest creates a request to invoke CreateKey API
func CreateCreateKeyRequest() (request *CreateKeyRequest) {
	request = &CreateKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "CreateKey", "kms-service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateKeyResponse creates a response to parse from CreateKey response
func CreateCreateKeyResponse() (response *CreateKeyResponse) {
	response = &CreateKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
