package dts

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

// SummaryJobDetail invokes the dts.SummaryJobDetail API synchronously
func (client *Client) SummaryJobDetail(request *SummaryJobDetailRequest) (response *SummaryJobDetailResponse, err error) {
	response = CreateSummaryJobDetailResponse()
	err = client.DoAction(request, response)
	return
}

// SummaryJobDetailWithChan invokes the dts.SummaryJobDetail API asynchronously
func (client *Client) SummaryJobDetailWithChan(request *SummaryJobDetailRequest) (<-chan *SummaryJobDetailResponse, <-chan error) {
	responseChan := make(chan *SummaryJobDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SummaryJobDetail(request)
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

// SummaryJobDetailWithCallback invokes the dts.SummaryJobDetail API asynchronously
func (client *Client) SummaryJobDetailWithCallback(request *SummaryJobDetailRequest, callback func(response *SummaryJobDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SummaryJobDetailResponse
		var err error
		defer close(result)
		response, err = client.SummaryJobDetail(request)
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

// SummaryJobDetailRequest is the request struct for api SummaryJobDetail
type SummaryJobDetailRequest struct {
	*requests.RpcRequest
	JobCode                  string `position:"Query" name:"JobCode"`
	DtsJobId                 string `position:"Query" name:"DtsJobId"`
	DtsInstanceId            string `position:"Query" name:"DtsInstanceId"`
	SynchronizationDirection string `position:"Query" name:"SynchronizationDirection"`
}

// SummaryJobDetailResponse is the response struct for api SummaryJobDetail
type SummaryJobDetailResponse struct {
	*responses.BaseResponse
	HttpStatusCode         int                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code                   string                  `json:"Code" xml:"Code"`
	RequestId              string                  `json:"RequestId" xml:"RequestId"`
	Success                bool                    `json:"Success" xml:"Success"`
	JobId                  string                  `json:"JobId" xml:"JobId"`
	ProgressSummaryDetails []ProgressSummaryDetail `json:"ProgressSummaryDetails" xml:"ProgressSummaryDetails"`
}

// CreateSummaryJobDetailRequest creates a request to invoke SummaryJobDetail API
func CreateSummaryJobDetailRequest() (request *SummaryJobDetailRequest) {
	request = &SummaryJobDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "SummaryJobDetail", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSummaryJobDetailResponse creates a response to parse from SummaryJobDetail response
func CreateSummaryJobDetailResponse() (response *SummaryJobDetailResponse) {
	response = &SummaryJobDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
