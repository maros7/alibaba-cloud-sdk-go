package arms

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

// CreateDispatchRule invokes the arms.CreateDispatchRule API synchronously
func (client *Client) CreateDispatchRule(request *CreateDispatchRuleRequest) (response *CreateDispatchRuleResponse, err error) {
	response = CreateCreateDispatchRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDispatchRuleWithChan invokes the arms.CreateDispatchRule API asynchronously
func (client *Client) CreateDispatchRuleWithChan(request *CreateDispatchRuleRequest) (<-chan *CreateDispatchRuleResponse, <-chan error) {
	responseChan := make(chan *CreateDispatchRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDispatchRule(request)
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

// CreateDispatchRuleWithCallback invokes the arms.CreateDispatchRule API asynchronously
func (client *Client) CreateDispatchRuleWithCallback(request *CreateDispatchRuleRequest, callback func(response *CreateDispatchRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDispatchRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateDispatchRule(request)
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

// CreateDispatchRuleRequest is the request struct for api CreateDispatchRule
type CreateDispatchRuleRequest struct {
	*requests.RpcRequest
	DispatchRule string `position:"Query" name:"DispatchRule"`
	ProxyUserId  string `position:"Query" name:"ProxyUserId"`
}

// CreateDispatchRuleResponse is the response struct for api CreateDispatchRule
type CreateDispatchRuleResponse struct {
	*responses.BaseResponse
	DispatchRuleId int64  `json:"DispatchRuleId" xml:"DispatchRuleId"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDispatchRuleRequest creates a request to invoke CreateDispatchRule API
func CreateCreateDispatchRuleRequest() (request *CreateDispatchRuleRequest) {
	request = &CreateDispatchRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CreateDispatchRule", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDispatchRuleResponse creates a response to parse from CreateDispatchRule response
func CreateCreateDispatchRuleResponse() (response *CreateDispatchRuleResponse) {
	response = &CreateDispatchRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
