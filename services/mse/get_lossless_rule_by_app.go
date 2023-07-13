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

// GetLosslessRuleByApp invokes the mse.GetLosslessRuleByApp API synchronously
func (client *Client) GetLosslessRuleByApp(request *GetLosslessRuleByAppRequest) (response *GetLosslessRuleByAppResponse, err error) {
	response = CreateGetLosslessRuleByAppResponse()
	err = client.DoAction(request, response)
	return
}

// GetLosslessRuleByAppWithChan invokes the mse.GetLosslessRuleByApp API asynchronously
func (client *Client) GetLosslessRuleByAppWithChan(request *GetLosslessRuleByAppRequest) (<-chan *GetLosslessRuleByAppResponse, <-chan error) {
	responseChan := make(chan *GetLosslessRuleByAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLosslessRuleByApp(request)
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

// GetLosslessRuleByAppWithCallback invokes the mse.GetLosslessRuleByApp API asynchronously
func (client *Client) GetLosslessRuleByAppWithCallback(request *GetLosslessRuleByAppRequest, callback func(response *GetLosslessRuleByAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLosslessRuleByAppResponse
		var err error
		defer close(result)
		response, err = client.GetLosslessRuleByApp(request)
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

// GetLosslessRuleByAppRequest is the request struct for api GetLosslessRuleByApp
type GetLosslessRuleByAppRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	AppName        string `position:"Query" name:"AppName"`
	AppId          string `position:"Query" name:"AppId"`
	Namespace      string `position:"Query" name:"Namespace"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// GetLosslessRuleByAppResponse is the response struct for api GetLosslessRuleByApp
type GetLosslessRuleByAppResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           int    `json:"Code" xml:"Code"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetLosslessRuleByAppRequest creates a request to invoke GetLosslessRuleByApp API
func CreateGetLosslessRuleByAppRequest() (request *GetLosslessRuleByAppRequest) {
	request = &GetLosslessRuleByAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetLosslessRuleByApp", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetLosslessRuleByAppResponse creates a response to parse from GetLosslessRuleByApp response
func CreateGetLosslessRuleByAppResponse() (response *GetLosslessRuleByAppResponse) {
	response = &GetLosslessRuleByAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
