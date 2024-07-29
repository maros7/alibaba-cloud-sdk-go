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

// DescribeDistributionProductsLink invokes the market.DescribeDistributionProductsLink API synchronously
func (client *Client) DescribeDistributionProductsLink(request *DescribeDistributionProductsLinkRequest) (response *DescribeDistributionProductsLinkResponse, err error) {
	response = CreateDescribeDistributionProductsLinkResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDistributionProductsLinkWithChan invokes the market.DescribeDistributionProductsLink API asynchronously
func (client *Client) DescribeDistributionProductsLinkWithChan(request *DescribeDistributionProductsLinkRequest) (<-chan *DescribeDistributionProductsLinkResponse, <-chan error) {
	responseChan := make(chan *DescribeDistributionProductsLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDistributionProductsLink(request)
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

// DescribeDistributionProductsLinkWithCallback invokes the market.DescribeDistributionProductsLink API asynchronously
func (client *Client) DescribeDistributionProductsLinkWithCallback(request *DescribeDistributionProductsLinkRequest, callback func(response *DescribeDistributionProductsLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDistributionProductsLinkResponse
		var err error
		defer close(result)
		response, err = client.DescribeDistributionProductsLink(request)
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

// DescribeDistributionProductsLinkRequest is the request struct for api DescribeDistributionProductsLink
type DescribeDistributionProductsLinkRequest struct {
	*requests.RpcRequest
	Codes *[]string `position:"Query" name:"Codes"  type:"Json"`
}

// DescribeDistributionProductsLinkResponse is the response struct for api DescribeDistributionProductsLink
type DescribeDistributionProductsLinkResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int64  `json:"TotalCount" xml:"TotalCount"`
	Success    bool   `json:"Success" xml:"Success"`
	Result     []Item `json:"Result" xml:"Result"`
}

// CreateDescribeDistributionProductsLinkRequest creates a request to invoke DescribeDistributionProductsLink API
func CreateDescribeDistributionProductsLinkRequest() (request *DescribeDistributionProductsLinkRequest) {
	request = &DescribeDistributionProductsLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Market", "2015-11-01", "DescribeDistributionProductsLink", "yunmarket", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDistributionProductsLinkResponse creates a response to parse from DescribeDistributionProductsLink response
func CreateDescribeDistributionProductsLinkResponse() (response *DescribeDistributionProductsLinkResponse) {
	response = &DescribeDistributionProductsLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
