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

// GetOverview invokes the mse.GetOverview API synchronously
func (client *Client) GetOverview(request *GetOverviewRequest) (response *GetOverviewResponse, err error) {
	response = CreateGetOverviewResponse()
	err = client.DoAction(request, response)
	return
}

// GetOverviewWithChan invokes the mse.GetOverview API asynchronously
func (client *Client) GetOverviewWithChan(request *GetOverviewRequest) (<-chan *GetOverviewResponse, <-chan error) {
	responseChan := make(chan *GetOverviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOverview(request)
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

// GetOverviewWithCallback invokes the mse.GetOverview API asynchronously
func (client *Client) GetOverviewWithCallback(request *GetOverviewRequest, callback func(response *GetOverviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOverviewResponse
		var err error
		defer close(result)
		response, err = client.GetOverview(request)
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

// GetOverviewRequest is the request struct for api GetOverview
type GetOverviewRequest struct {
	*requests.RpcRequest
	Period         requests.Integer `position:"Query" name:"Period"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
	Region         string           `position:"Query" name:"Region"`
}

// GetOverviewResponse is the response struct for api GetOverview
type GetOverviewResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           string `json:"Data" xml:"Data"`
	Code           int    `json:"Code" xml:"Code"`
	Success        string `json:"Success" xml:"Success"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateGetOverviewRequest creates a request to invoke GetOverview API
func CreateGetOverviewRequest() (request *GetOverviewRequest) {
	request = &GetOverviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetOverview", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetOverviewResponse creates a response to parse from GetOverview response
func CreateGetOverviewResponse() (response *GetOverviewResponse) {
	response = &GetOverviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
