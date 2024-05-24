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

// DescribeMeterLiveInteractionDau invokes the live.DescribeMeterLiveInteractionDau API synchronously
func (client *Client) DescribeMeterLiveInteractionDau(request *DescribeMeterLiveInteractionDauRequest) (response *DescribeMeterLiveInteractionDauResponse, err error) {
	response = CreateDescribeMeterLiveInteractionDauResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMeterLiveInteractionDauWithChan invokes the live.DescribeMeterLiveInteractionDau API asynchronously
func (client *Client) DescribeMeterLiveInteractionDauWithChan(request *DescribeMeterLiveInteractionDauRequest) (<-chan *DescribeMeterLiveInteractionDauResponse, <-chan error) {
	responseChan := make(chan *DescribeMeterLiveInteractionDauResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMeterLiveInteractionDau(request)
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

// DescribeMeterLiveInteractionDauWithCallback invokes the live.DescribeMeterLiveInteractionDau API asynchronously
func (client *Client) DescribeMeterLiveInteractionDauWithCallback(request *DescribeMeterLiveInteractionDauRequest, callback func(response *DescribeMeterLiveInteractionDauResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMeterLiveInteractionDauResponse
		var err error
		defer close(result)
		response, err = client.DescribeMeterLiveInteractionDau(request)
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

// DescribeMeterLiveInteractionDauRequest is the request struct for api DescribeMeterLiveInteractionDau
type DescribeMeterLiveInteractionDauRequest struct {
	*requests.RpcRequest
	StartTs     requests.Integer `position:"Query" name:"StartTs"`
	ServiceArea string           `position:"Query" name:"ServiceArea"`
	AppId       string           `position:"Query" name:"AppId"`
	EndTs       requests.Integer `position:"Query" name:"EndTs"`
	Interval    requests.Integer `position:"Query" name:"Interval"`
}

// DescribeMeterLiveInteractionDauResponse is the response struct for api DescribeMeterLiveInteractionDau
type DescribeMeterLiveInteractionDauResponse struct {
	*responses.BaseResponse
	PeakDau   string                                      `json:"PeakDau" xml:"PeakDau"`
	RequestId string                                      `json:"RequestId" xml:"RequestId"`
	Data      []DataItemInDescribeMeterLiveInteractionDau `json:"Data" xml:"Data"`
}

// CreateDescribeMeterLiveInteractionDauRequest creates a request to invoke DescribeMeterLiveInteractionDau API
func CreateDescribeMeterLiveInteractionDauRequest() (request *DescribeMeterLiveInteractionDauRequest) {
	request = &DescribeMeterLiveInteractionDauRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeMeterLiveInteractionDau", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMeterLiveInteractionDauResponse creates a response to parse from DescribeMeterLiveInteractionDau response
func CreateDescribeMeterLiveInteractionDauResponse() (response *DescribeMeterLiveInteractionDauResponse) {
	response = &DescribeMeterLiveInteractionDauResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
