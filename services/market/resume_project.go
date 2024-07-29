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

// ResumeProject invokes the market.ResumeProject API synchronously
func (client *Client) ResumeProject(request *ResumeProjectRequest) (response *ResumeProjectResponse, err error) {
	response = CreateResumeProjectResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeProjectWithChan invokes the market.ResumeProject API asynchronously
func (client *Client) ResumeProjectWithChan(request *ResumeProjectRequest) (<-chan *ResumeProjectResponse, <-chan error) {
	responseChan := make(chan *ResumeProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeProject(request)
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

// ResumeProjectWithCallback invokes the market.ResumeProject API asynchronously
func (client *Client) ResumeProjectWithCallback(request *ResumeProjectRequest, callback func(response *ResumeProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeProjectResponse
		var err error
		defer close(result)
		response, err = client.ResumeProject(request)
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

// ResumeProjectRequest is the request struct for api ResumeProject
type ResumeProjectRequest struct {
	*requests.RpcRequest
	Remark     string           `position:"Query" name:"Remark"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	NodeId     requests.Integer `position:"Query" name:"NodeId"`
}

// ResumeProjectResponse is the response struct for api ResumeProject
type ResumeProjectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateResumeProjectRequest creates a request to invoke ResumeProject API
func CreateResumeProjectRequest() (request *ResumeProjectRequest) {
	request = &ResumeProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Market", "2015-11-01", "ResumeProject", "yunmarket", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResumeProjectResponse creates a response to parse from ResumeProject response
func CreateResumeProjectResponse() (response *ResumeProjectResponse) {
	response = &ResumeProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
