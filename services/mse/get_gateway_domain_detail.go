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

// GetGatewayDomainDetail invokes the mse.GetGatewayDomainDetail API synchronously
func (client *Client) GetGatewayDomainDetail(request *GetGatewayDomainDetailRequest) (response *GetGatewayDomainDetailResponse, err error) {
	response = CreateGetGatewayDomainDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetGatewayDomainDetailWithChan invokes the mse.GetGatewayDomainDetail API asynchronously
func (client *Client) GetGatewayDomainDetailWithChan(request *GetGatewayDomainDetailRequest) (<-chan *GetGatewayDomainDetailResponse, <-chan error) {
	responseChan := make(chan *GetGatewayDomainDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGatewayDomainDetail(request)
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

// GetGatewayDomainDetailWithCallback invokes the mse.GetGatewayDomainDetail API asynchronously
func (client *Client) GetGatewayDomainDetailWithCallback(request *GetGatewayDomainDetailRequest, callback func(response *GetGatewayDomainDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGatewayDomainDetailResponse
		var err error
		defer close(result)
		response, err = client.GetGatewayDomainDetail(request)
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

// GetGatewayDomainDetailRequest is the request struct for api GetGatewayDomainDetail
type GetGatewayDomainDetailRequest struct {
	*requests.RpcRequest
	GatewayUniqueId string `position:"Query" name:"GatewayUniqueId"`
	AcceptLanguage  string `position:"Query" name:"AcceptLanguage"`
	Id              string `position:"Query" name:"Id"`
}

// GetGatewayDomainDetailResponse is the response struct for api GetGatewayDomainDetail
type GetGatewayDomainDetailResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetGatewayDomainDetailRequest creates a request to invoke GetGatewayDomainDetail API
func CreateGetGatewayDomainDetailRequest() (request *GetGatewayDomainDetailRequest) {
	request = &GetGatewayDomainDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetGatewayDomainDetail", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetGatewayDomainDetailResponse creates a response to parse from GetGatewayDomainDetail response
func CreateGetGatewayDomainDetailResponse() (response *GetGatewayDomainDetailResponse) {
	response = &GetGatewayDomainDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
