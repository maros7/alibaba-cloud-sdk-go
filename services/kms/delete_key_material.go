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

// DeleteKeyMaterial invokes the kms.DeleteKeyMaterial API synchronously
func (client *Client) DeleteKeyMaterial(request *DeleteKeyMaterialRequest) (response *DeleteKeyMaterialResponse, err error) {
	response = CreateDeleteKeyMaterialResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteKeyMaterialWithChan invokes the kms.DeleteKeyMaterial API asynchronously
func (client *Client) DeleteKeyMaterialWithChan(request *DeleteKeyMaterialRequest) (<-chan *DeleteKeyMaterialResponse, <-chan error) {
	responseChan := make(chan *DeleteKeyMaterialResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteKeyMaterial(request)
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

// DeleteKeyMaterialWithCallback invokes the kms.DeleteKeyMaterial API asynchronously
func (client *Client) DeleteKeyMaterialWithCallback(request *DeleteKeyMaterialRequest, callback func(response *DeleteKeyMaterialResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteKeyMaterialResponse
		var err error
		defer close(result)
		response, err = client.DeleteKeyMaterial(request)
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

// DeleteKeyMaterialRequest is the request struct for api DeleteKeyMaterial
type DeleteKeyMaterialRequest struct {
	*requests.RpcRequest
	KeyId string `position:"Query" name:"KeyId"`
}

// DeleteKeyMaterialResponse is the response struct for api DeleteKeyMaterial
type DeleteKeyMaterialResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteKeyMaterialRequest creates a request to invoke DeleteKeyMaterial API
func CreateDeleteKeyMaterialRequest() (request *DeleteKeyMaterialRequest) {
	request = &DeleteKeyMaterialRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "DeleteKeyMaterial", "kms-service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteKeyMaterialResponse creates a response to parse from DeleteKeyMaterial response
func CreateDeleteKeyMaterialResponse() (response *DeleteKeyMaterialResponse) {
	response = &DeleteKeyMaterialResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
