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

// ListCircuitBreakerRules invokes the mse.ListCircuitBreakerRules API synchronously
func (client *Client) ListCircuitBreakerRules(request *ListCircuitBreakerRulesRequest) (response *ListCircuitBreakerRulesResponse, err error) {
	response = CreateListCircuitBreakerRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListCircuitBreakerRulesWithChan invokes the mse.ListCircuitBreakerRules API asynchronously
func (client *Client) ListCircuitBreakerRulesWithChan(request *ListCircuitBreakerRulesRequest) (<-chan *ListCircuitBreakerRulesResponse, <-chan error) {
	responseChan := make(chan *ListCircuitBreakerRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCircuitBreakerRules(request)
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

// ListCircuitBreakerRulesWithCallback invokes the mse.ListCircuitBreakerRules API asynchronously
func (client *Client) ListCircuitBreakerRulesWithCallback(request *ListCircuitBreakerRulesRequest, callback func(response *ListCircuitBreakerRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCircuitBreakerRulesResponse
		var err error
		defer close(result)
		response, err = client.ListCircuitBreakerRules(request)
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

// ListCircuitBreakerRulesRequest is the request struct for api ListCircuitBreakerRules
type ListCircuitBreakerRulesRequest struct {
	*requests.RpcRequest
	MseSessionId      string           `position:"Query" name:"MseSessionId"`
	AppName           string           `position:"Query" name:"AppName"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	PageIndex         requests.Integer `position:"Query" name:"PageIndex"`
	Resource          string           `position:"Query" name:"Resource"`
	AppId             string           `position:"Query" name:"AppId"`
	Namespace         string           `position:"Query" name:"Namespace"`
	AcceptLanguage    string           `position:"Query" name:"AcceptLanguage"`
	ResourceSearchKey string           `position:"Query" name:"ResourceSearchKey"`
}

// ListCircuitBreakerRulesResponse is the response struct for api ListCircuitBreakerRules
type ListCircuitBreakerRulesResponse struct {
	*responses.BaseResponse
	Code           int                           `json:"Code" xml:"Code"`
	Message        string                        `json:"Message" xml:"Message"`
	RequestId      string                        `json:"RequestId" xml:"RequestId"`
	Success        bool                          `json:"Success" xml:"Success"`
	HttpStatusCode int                           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInListCircuitBreakerRules `json:"Data" xml:"Data"`
}

// CreateListCircuitBreakerRulesRequest creates a request to invoke ListCircuitBreakerRules API
func CreateListCircuitBreakerRulesRequest() (request *ListCircuitBreakerRulesRequest) {
	request = &ListCircuitBreakerRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListCircuitBreakerRules", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListCircuitBreakerRulesResponse creates a response to parse from ListCircuitBreakerRules response
func CreateListCircuitBreakerRulesResponse() (response *ListCircuitBreakerRulesResponse) {
	response = &ListCircuitBreakerRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
