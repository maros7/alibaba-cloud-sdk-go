package ehpc

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

// ListClusterLogs invokes the ehpc.ListClusterLogs API synchronously
func (client *Client) ListClusterLogs(request *ListClusterLogsRequest) (response *ListClusterLogsResponse, err error) {
	response = CreateListClusterLogsResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterLogsWithChan invokes the ehpc.ListClusterLogs API asynchronously
func (client *Client) ListClusterLogsWithChan(request *ListClusterLogsRequest) (<-chan *ListClusterLogsResponse, <-chan error) {
	responseChan := make(chan *ListClusterLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterLogs(request)
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

// ListClusterLogsWithCallback invokes the ehpc.ListClusterLogs API asynchronously
func (client *Client) ListClusterLogsWithCallback(request *ListClusterLogsRequest, callback func(response *ListClusterLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterLogsResponse
		var err error
		defer close(result)
		response, err = client.ListClusterLogs(request)
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

// ListClusterLogsRequest is the request struct for api ListClusterLogs
type ListClusterLogsRequest struct {
	*requests.RpcRequest
	ClusterId  string           `position:"Query" name:"ClusterId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListClusterLogsResponse is the response struct for api ListClusterLogs
type ListClusterLogsResponse struct {
	*responses.BaseResponse
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	ClusterId  string `json:"ClusterId" xml:"ClusterId"`
	Logs       Logs   `json:"Logs" xml:"Logs"`
}

// CreateListClusterLogsRequest creates a request to invoke ListClusterLogs API
func CreateListClusterLogsRequest() (request *ListClusterLogsRequest) {
	request = &ListClusterLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2017-07-14", "ListClusterLogs", "", "")
	request.Method = requests.GET
	return
}

// CreateListClusterLogsResponse creates a response to parse from ListClusterLogs response
func CreateListClusterLogsResponse() (response *ListClusterLogsResponse) {
	response = &ListClusterLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
