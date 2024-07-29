package market

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

// DescribeProjectOperateLogs invokes the market.DescribeProjectOperateLogs API synchronously
func (client *Client) DescribeProjectOperateLogs(request *DescribeProjectOperateLogsRequest) (response *DescribeProjectOperateLogsResponse, err error) {
	response = CreateDescribeProjectOperateLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProjectOperateLogsWithChan invokes the market.DescribeProjectOperateLogs API asynchronously
func (client *Client) DescribeProjectOperateLogsWithChan(request *DescribeProjectOperateLogsRequest) (<-chan *DescribeProjectOperateLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeProjectOperateLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProjectOperateLogs(request)
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

// DescribeProjectOperateLogsWithCallback invokes the market.DescribeProjectOperateLogs API asynchronously
func (client *Client) DescribeProjectOperateLogsWithCallback(request *DescribeProjectOperateLogsRequest, callback func(response *DescribeProjectOperateLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProjectOperateLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeProjectOperateLogs(request)
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

// DescribeProjectOperateLogsRequest is the request struct for api DescribeProjectOperateLogs
type DescribeProjectOperateLogsRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeProjectOperateLogsResponse is the response struct for api DescribeProjectOperateLogs
type DescribeProjectOperateLogsResponse struct {
	*responses.BaseResponse
	Success   bool             `json:"Success" xml:"Success"`
	RequestId string           `json:"RequestId" xml:"RequestId"`
	Result    []ProjectMessage `json:"Result" xml:"Result"`
}

// CreateDescribeProjectOperateLogsRequest creates a request to invoke DescribeProjectOperateLogs API
func CreateDescribeProjectOperateLogsRequest() (request *DescribeProjectOperateLogsRequest) {
	request = &DescribeProjectOperateLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Market", "2015-11-01", "DescribeProjectOperateLogs", "yunmarket", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeProjectOperateLogsResponse creates a response to parse from DescribeProjectOperateLogs response
func CreateDescribeProjectOperateLogsResponse() (response *DescribeProjectOperateLogsResponse) {
	response = &DescribeProjectOperateLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
