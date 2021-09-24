package dytnsapi

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

// DescribePhoneNumberResale invokes the dytnsapi.DescribePhoneNumberResale API synchronously
func (client *Client) DescribePhoneNumberResale(request *DescribePhoneNumberResaleRequest) (response *DescribePhoneNumberResaleResponse, err error) {
	response = CreateDescribePhoneNumberResaleResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePhoneNumberResaleWithChan invokes the dytnsapi.DescribePhoneNumberResale API asynchronously
func (client *Client) DescribePhoneNumberResaleWithChan(request *DescribePhoneNumberResaleRequest) (<-chan *DescribePhoneNumberResaleResponse, <-chan error) {
	responseChan := make(chan *DescribePhoneNumberResaleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePhoneNumberResale(request)
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

// DescribePhoneNumberResaleWithCallback invokes the dytnsapi.DescribePhoneNumberResale API asynchronously
func (client *Client) DescribePhoneNumberResaleWithCallback(request *DescribePhoneNumberResaleRequest, callback func(response *DescribePhoneNumberResaleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePhoneNumberResaleResponse
		var err error
		defer close(result)
		response, err = client.DescribePhoneNumberResale(request)
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

// DescribePhoneNumberResaleRequest is the request struct for api DescribePhoneNumberResale
type DescribePhoneNumberResaleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PhoneNumber          string           `position:"Query" name:"PhoneNumber"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Since                string           `position:"Query" name:"Since"`
}

// DescribePhoneNumberResaleResponse is the response struct for api DescribePhoneNumberResale
type DescribePhoneNumberResaleResponse struct {
	*responses.BaseResponse
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	TwiceTelVerify TwiceTelVerify `json:"TwiceTelVerify" xml:"TwiceTelVerify"`
}

// CreateDescribePhoneNumberResaleRequest creates a request to invoke DescribePhoneNumberResale API
func CreateDescribePhoneNumberResaleRequest() (request *DescribePhoneNumberResaleRequest) {
	request = &DescribePhoneNumberResaleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dytnsapi", "2020-02-17", "DescribePhoneNumberResale", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePhoneNumberResaleResponse creates a response to parse from DescribePhoneNumberResale response
func CreateDescribePhoneNumberResaleResponse() (response *DescribePhoneNumberResaleResponse) {
	response = &DescribePhoneNumberResaleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
