package mse

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

// GetEngineNamepace invokes the mse.GetEngineNamepace API synchronously
func (client *Client) GetEngineNamepace(request *GetEngineNamepaceRequest) (response *GetEngineNamepaceResponse, err error) {
	response = CreateGetEngineNamepaceResponse()
	err = client.DoAction(request, response)
	return
}

// GetEngineNamepaceWithChan invokes the mse.GetEngineNamepace API asynchronously
func (client *Client) GetEngineNamepaceWithChan(request *GetEngineNamepaceRequest) (<-chan *GetEngineNamepaceResponse, <-chan error) {
	responseChan := make(chan *GetEngineNamepaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEngineNamepace(request)
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

// GetEngineNamepaceWithCallback invokes the mse.GetEngineNamepace API asynchronously
func (client *Client) GetEngineNamepaceWithCallback(request *GetEngineNamepaceRequest, callback func(response *GetEngineNamepaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEngineNamepaceResponse
		var err error
		defer close(result)
		response, err = client.GetEngineNamepace(request)
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

// GetEngineNamepaceRequest is the request struct for api GetEngineNamepace
type GetEngineNamepaceRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	Id             string `position:"Query" name:"Id"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// GetEngineNamepaceResponse is the response struct for api GetEngineNamepace
type GetEngineNamepaceResponse struct {
	*responses.BaseResponse
	HttpCode          string `json:"HttpCode" xml:"HttpCode"`
	Type              string `json:"Type" xml:"Type"`
	Quota             string `json:"Quota" xml:"Quota"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
	Message           string `json:"Message" xml:"Message"`
	ConfigCount       string `json:"ConfigCount" xml:"ConfigCount"`
	NamespaceShowName string `json:"NamespaceShowName" xml:"NamespaceShowName"`
	ErrorCode         string `json:"ErrorCode" xml:"ErrorCode"`
	Success           bool   `json:"Success" xml:"Success"`
	NamespaceDesc     string `json:"NamespaceDesc" xml:"NamespaceDesc"`
	Namespace         string `json:"Namespace" xml:"Namespace"`
}

// CreateGetEngineNamepaceRequest creates a request to invoke GetEngineNamepace API
func CreateGetEngineNamepaceRequest() (request *GetEngineNamepaceRequest) {
	request = &GetEngineNamepaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetEngineNamepace", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetEngineNamepaceResponse creates a response to parse from GetEngineNamepace response
func CreateGetEngineNamepaceResponse() (response *GetEngineNamepaceResponse) {
	response = &GetEngineNamepaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
