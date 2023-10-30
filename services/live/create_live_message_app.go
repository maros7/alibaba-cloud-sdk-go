package live

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

// CreateLiveMessageApp invokes the live.CreateLiveMessageApp API synchronously
func (client *Client) CreateLiveMessageApp(request *CreateLiveMessageAppRequest) (response *CreateLiveMessageAppResponse, err error) {
	response = CreateCreateLiveMessageAppResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLiveMessageAppWithChan invokes the live.CreateLiveMessageApp API asynchronously
func (client *Client) CreateLiveMessageAppWithChan(request *CreateLiveMessageAppRequest) (<-chan *CreateLiveMessageAppResponse, <-chan error) {
	responseChan := make(chan *CreateLiveMessageAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLiveMessageApp(request)
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

// CreateLiveMessageAppWithCallback invokes the live.CreateLiveMessageApp API asynchronously
func (client *Client) CreateLiveMessageAppWithCallback(request *CreateLiveMessageAppRequest, callback func(response *CreateLiveMessageAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLiveMessageAppResponse
		var err error
		defer close(result)
		response, err = client.CreateLiveMessageApp(request)
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

// CreateLiveMessageAppRequest is the request struct for api CreateLiveMessageApp
type CreateLiveMessageAppRequest struct {
	*requests.RpcRequest
	DataCenter       string           `position:"Query" name:"DataCenter"`
	AppName          string           `position:"Query" name:"AppName"`
	AuditType        requests.Integer `position:"Query" name:"AuditType"`
	AuditUrl         string           `position:"Query" name:"AuditUrl"`
	EventCallbackUrl string           `position:"Query" name:"EventCallbackUrl"`
}

// CreateLiveMessageAppResponse is the response struct for api CreateLiveMessageApp
type CreateLiveMessageAppResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	AppId     string `json:"AppId" xml:"AppId"`
	AppKey    string `json:"AppKey" xml:"AppKey"`
	AppSign   string `json:"AppSign" xml:"AppSign"`
}

// CreateCreateLiveMessageAppRequest creates a request to invoke CreateLiveMessageApp API
func CreateCreateLiveMessageAppRequest() (request *CreateLiveMessageAppRequest) {
	request = &CreateLiveMessageAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CreateLiveMessageApp", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLiveMessageAppResponse creates a response to parse from CreateLiveMessageApp response
func CreateCreateLiveMessageAppResponse() (response *CreateLiveMessageAppResponse) {
	response = &CreateLiveMessageAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
