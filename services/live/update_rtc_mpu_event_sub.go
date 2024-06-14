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

// UpdateRtcMPUEventSub invokes the live.UpdateRtcMPUEventSub API synchronously
func (client *Client) UpdateRtcMPUEventSub(request *UpdateRtcMPUEventSubRequest) (response *UpdateRtcMPUEventSubResponse, err error) {
	response = CreateUpdateRtcMPUEventSubResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRtcMPUEventSubWithChan invokes the live.UpdateRtcMPUEventSub API asynchronously
func (client *Client) UpdateRtcMPUEventSubWithChan(request *UpdateRtcMPUEventSubRequest) (<-chan *UpdateRtcMPUEventSubResponse, <-chan error) {
	responseChan := make(chan *UpdateRtcMPUEventSubResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRtcMPUEventSub(request)
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

// UpdateRtcMPUEventSubWithCallback invokes the live.UpdateRtcMPUEventSub API asynchronously
func (client *Client) UpdateRtcMPUEventSubWithCallback(request *UpdateRtcMPUEventSubRequest, callback func(response *UpdateRtcMPUEventSubResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRtcMPUEventSubResponse
		var err error
		defer close(result)
		response, err = client.UpdateRtcMPUEventSub(request)
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

// UpdateRtcMPUEventSubRequest is the request struct for api UpdateRtcMPUEventSub
type UpdateRtcMPUEventSubRequest struct {
	*requests.RpcRequest
	AppId       string `position:"Query" name:"AppId"`
	CallbackUrl string `position:"Query" name:"CallbackUrl"`
	ChannelIds  string `position:"Query" name:"ChannelIds"`
}

// UpdateRtcMPUEventSubResponse is the response struct for api UpdateRtcMPUEventSub
type UpdateRtcMPUEventSubResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateRtcMPUEventSubRequest creates a request to invoke UpdateRtcMPUEventSub API
func CreateUpdateRtcMPUEventSubRequest() (request *UpdateRtcMPUEventSubRequest) {
	request = &UpdateRtcMPUEventSubRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateRtcMPUEventSub", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateRtcMPUEventSubResponse creates a response to parse from UpdateRtcMPUEventSub response
func CreateUpdateRtcMPUEventSubResponse() (response *UpdateRtcMPUEventSubResponse) {
	response = &UpdateRtcMPUEventSubResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
