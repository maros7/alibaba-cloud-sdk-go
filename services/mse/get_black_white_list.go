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

// GetBlackWhiteList invokes the mse.GetBlackWhiteList API synchronously
func (client *Client) GetBlackWhiteList(request *GetBlackWhiteListRequest) (response *GetBlackWhiteListResponse, err error) {
	response = CreateGetBlackWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// GetBlackWhiteListWithChan invokes the mse.GetBlackWhiteList API asynchronously
func (client *Client) GetBlackWhiteListWithChan(request *GetBlackWhiteListRequest) (<-chan *GetBlackWhiteListResponse, <-chan error) {
	responseChan := make(chan *GetBlackWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBlackWhiteList(request)
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

// GetBlackWhiteListWithCallback invokes the mse.GetBlackWhiteList API asynchronously
func (client *Client) GetBlackWhiteListWithCallback(request *GetBlackWhiteListRequest, callback func(response *GetBlackWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBlackWhiteListResponse
		var err error
		defer close(result)
		response, err = client.GetBlackWhiteList(request)
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

// GetBlackWhiteListRequest is the request struct for api GetBlackWhiteList
type GetBlackWhiteListRequest struct {
	*requests.RpcRequest
	GatewayUniqueId string `position:"Query" name:"GatewayUniqueId"`
	Type            string `position:"Query" name:"Type"`
	ResourceType    string `position:"Query" name:"ResourceType"`
	AcceptLanguage  string `position:"Query" name:"AcceptLanguage"`
}

// GetBlackWhiteListResponse is the response struct for api GetBlackWhiteList
type GetBlackWhiteListResponse struct {
	*responses.BaseResponse
	RequestId      string                  `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                  `json:"Message" xml:"Message"`
	Code           int                     `json:"Code" xml:"Code"`
	Success        bool                    `json:"Success" xml:"Success"`
	Data           DataInGetBlackWhiteList `json:"Data" xml:"Data"`
}

// CreateGetBlackWhiteListRequest creates a request to invoke GetBlackWhiteList API
func CreateGetBlackWhiteListRequest() (request *GetBlackWhiteListRequest) {
	request = &GetBlackWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetBlackWhiteList", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetBlackWhiteListResponse creates a response to parse from GetBlackWhiteList response
func CreateGetBlackWhiteListResponse() (response *GetBlackWhiteListResponse) {
	response = &GetBlackWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
