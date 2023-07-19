package rds

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

// DescribeWhitelistTemplate invokes the rds.DescribeWhitelistTemplate API synchronously
func (client *Client) DescribeWhitelistTemplate(request *DescribeWhitelistTemplateRequest) (response *DescribeWhitelistTemplateResponse, err error) {
	response = CreateDescribeWhitelistTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWhitelistTemplateWithChan invokes the rds.DescribeWhitelistTemplate API asynchronously
func (client *Client) DescribeWhitelistTemplateWithChan(request *DescribeWhitelistTemplateRequest) (<-chan *DescribeWhitelistTemplateResponse, <-chan error) {
	responseChan := make(chan *DescribeWhitelistTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWhitelistTemplate(request)
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

// DescribeWhitelistTemplateWithCallback invokes the rds.DescribeWhitelistTemplate API asynchronously
func (client *Client) DescribeWhitelistTemplateWithCallback(request *DescribeWhitelistTemplateRequest, callback func(response *DescribeWhitelistTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWhitelistTemplateResponse
		var err error
		defer close(result)
		response, err = client.DescribeWhitelistTemplate(request)
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

// DescribeWhitelistTemplateRequest is the request struct for api DescribeWhitelistTemplate
type DescribeWhitelistTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	TemplateId           requests.Integer `position:"Query" name:"TemplateId"`
}

// DescribeWhitelistTemplateResponse is the response struct for api DescribeWhitelistTemplate
type DescribeWhitelistTemplateResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateDescribeWhitelistTemplateRequest creates a request to invoke DescribeWhitelistTemplate API
func CreateDescribeWhitelistTemplateRequest() (request *DescribeWhitelistTemplateRequest) {
	request = &DescribeWhitelistTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeWhitelistTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeWhitelistTemplateResponse creates a response to parse from DescribeWhitelistTemplate response
func CreateDescribeWhitelistTemplateResponse() (response *DescribeWhitelistTemplateResponse) {
	response = &DescribeWhitelistTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
