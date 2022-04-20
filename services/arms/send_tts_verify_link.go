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

// SendTTSVerifyLink invokes the arms.SendTTSVerifyLink API synchronously
func (client *Client) SendTTSVerifyLink(request *SendTTSVerifyLinkRequest) (response *SendTTSVerifyLinkResponse, err error) {
	response = CreateSendTTSVerifyLinkResponse()
	err = client.DoAction(request, response)
	return
}

// SendTTSVerifyLinkWithChan invokes the arms.SendTTSVerifyLink API asynchronously
func (client *Client) SendTTSVerifyLinkWithChan(request *SendTTSVerifyLinkRequest) (<-chan *SendTTSVerifyLinkResponse, <-chan error) {
	responseChan := make(chan *SendTTSVerifyLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendTTSVerifyLink(request)
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

// SendTTSVerifyLinkWithCallback invokes the arms.SendTTSVerifyLink API asynchronously
func (client *Client) SendTTSVerifyLinkWithCallback(request *SendTTSVerifyLinkRequest, callback func(response *SendTTSVerifyLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendTTSVerifyLinkResponse
		var err error
		defer close(result)
		response, err = client.SendTTSVerifyLink(request)
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

// SendTTSVerifyLinkRequest is the request struct for api SendTTSVerifyLink
type SendTTSVerifyLinkRequest struct {
	*requests.RpcRequest
	ContactId requests.Integer `position:"Body" name:"ContactId"`
	Phone     string           `position:"Body" name:"Phone"`
}

// SendTTSVerifyLinkResponse is the response struct for api SendTTSVerifyLink
type SendTTSVerifyLinkResponse struct {
	*responses.BaseResponse
	RequestId                  string `json:"RequestId" xml:"RequestId"`
	SendTTSVerifyLinkIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
}

// CreateSendTTSVerifyLinkRequest creates a request to invoke SendTTSVerifyLink API
func CreateSendTTSVerifyLinkRequest() (request *SendTTSVerifyLinkRequest) {
	request = &SendTTSVerifyLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "SendTTSVerifyLink", "", "")
	request.Method = requests.POST
	return
}

// CreateSendTTSVerifyLinkResponse creates a response to parse from SendTTSVerifyLink response
func CreateSendTTSVerifyLinkResponse() (response *SendTTSVerifyLinkResponse) {
	response = &SendTTSVerifyLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
