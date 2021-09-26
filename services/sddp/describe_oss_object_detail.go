package sddp

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

// DescribeOssObjectDetail invokes the sddp.DescribeOssObjectDetail API synchronously
func (client *Client) DescribeOssObjectDetail(request *DescribeOssObjectDetailRequest) (response *DescribeOssObjectDetailResponse, err error) {
	response = CreateDescribeOssObjectDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOssObjectDetailWithChan invokes the sddp.DescribeOssObjectDetail API asynchronously
func (client *Client) DescribeOssObjectDetailWithChan(request *DescribeOssObjectDetailRequest) (<-chan *DescribeOssObjectDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeOssObjectDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOssObjectDetail(request)
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

// DescribeOssObjectDetailWithCallback invokes the sddp.DescribeOssObjectDetail API asynchronously
func (client *Client) DescribeOssObjectDetailWithCallback(request *DescribeOssObjectDetailRequest, callback func(response *DescribeOssObjectDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOssObjectDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeOssObjectDetail(request)
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

// DescribeOssObjectDetailRequest is the request struct for api DescribeOssObjectDetail
type DescribeOssObjectDetailRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Id          requests.Integer `position:"Query" name:"Id"`
	Lang        string           `position:"Query" name:"Lang"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
}

// DescribeOssObjectDetailResponse is the response struct for api DescribeOssObjectDetail
type DescribeOssObjectDetailResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	OssObjectDetail OssObjectDetail `json:"OssObjectDetail" xml:"OssObjectDetail"`
}

// CreateDescribeOssObjectDetailRequest creates a request to invoke DescribeOssObjectDetail API
func CreateDescribeOssObjectDetailRequest() (request *DescribeOssObjectDetailRequest) {
	request = &DescribeOssObjectDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeOssObjectDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeOssObjectDetailResponse creates a response to parse from DescribeOssObjectDetail response
func CreateDescribeOssObjectDetailResponse() (response *DescribeOssObjectDetailResponse) {
	response = &DescribeOssObjectDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
