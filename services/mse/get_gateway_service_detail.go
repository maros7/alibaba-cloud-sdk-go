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

// GetGatewayServiceDetail invokes the mse.GetGatewayServiceDetail API synchronously
func (client *Client) GetGatewayServiceDetail(request *GetGatewayServiceDetailRequest) (response *GetGatewayServiceDetailResponse, err error) {
	response = CreateGetGatewayServiceDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetGatewayServiceDetailWithChan invokes the mse.GetGatewayServiceDetail API asynchronously
func (client *Client) GetGatewayServiceDetailWithChan(request *GetGatewayServiceDetailRequest) (<-chan *GetGatewayServiceDetailResponse, <-chan error) {
	responseChan := make(chan *GetGatewayServiceDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGatewayServiceDetail(request)
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

// GetGatewayServiceDetailWithCallback invokes the mse.GetGatewayServiceDetail API asynchronously
func (client *Client) GetGatewayServiceDetailWithCallback(request *GetGatewayServiceDetailRequest, callback func(response *GetGatewayServiceDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGatewayServiceDetailResponse
		var err error
		defer close(result)
		response, err = client.GetGatewayServiceDetail(request)
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

// GetGatewayServiceDetailRequest is the request struct for api GetGatewayServiceDetail
type GetGatewayServiceDetailRequest struct {
	*requests.RpcRequest
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
	ServiceId       requests.Integer `position:"Query" name:"ServiceId"`
}

// GetGatewayServiceDetailResponse is the response struct for api GetGatewayServiceDetail
type GetGatewayServiceDetailResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetGatewayServiceDetailRequest creates a request to invoke GetGatewayServiceDetail API
func CreateGetGatewayServiceDetailRequest() (request *GetGatewayServiceDetailRequest) {
	request = &GetGatewayServiceDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetGatewayServiceDetail", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetGatewayServiceDetailResponse creates a response to parse from GetGatewayServiceDetail response
func CreateGetGatewayServiceDetailResponse() (response *GetGatewayServiceDetailResponse) {
	response = &GetGatewayServiceDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
