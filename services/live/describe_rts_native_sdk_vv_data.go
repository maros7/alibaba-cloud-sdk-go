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

// DescribeRTSNativeSDKVvData invokes the live.DescribeRTSNativeSDKVvData API synchronously
func (client *Client) DescribeRTSNativeSDKVvData(request *DescribeRTSNativeSDKVvDataRequest) (response *DescribeRTSNativeSDKVvDataResponse, err error) {
	response = CreateDescribeRTSNativeSDKVvDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRTSNativeSDKVvDataWithChan invokes the live.DescribeRTSNativeSDKVvData API asynchronously
func (client *Client) DescribeRTSNativeSDKVvDataWithChan(request *DescribeRTSNativeSDKVvDataRequest) (<-chan *DescribeRTSNativeSDKVvDataResponse, <-chan error) {
	responseChan := make(chan *DescribeRTSNativeSDKVvDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRTSNativeSDKVvData(request)
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

// DescribeRTSNativeSDKVvDataWithCallback invokes the live.DescribeRTSNativeSDKVvData API asynchronously
func (client *Client) DescribeRTSNativeSDKVvDataWithCallback(request *DescribeRTSNativeSDKVvDataRequest, callback func(response *DescribeRTSNativeSDKVvDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRTSNativeSDKVvDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeRTSNativeSDKVvData(request)
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

// DescribeRTSNativeSDKVvDataRequest is the request struct for api DescribeRTSNativeSDKVvData
type DescribeRTSNativeSDKVvDataRequest struct {
	*requests.RpcRequest
	EndTime        string    `position:"Query" name:"EndTime"`
	DomainNameList *[]string `position:"Query" name:"DomainNameList"  type:"Json"`
	StartTime      string    `position:"Query" name:"StartTime"`
	DataInterval   string    `position:"Query" name:"DataInterval"`
}

// DescribeRTSNativeSDKVvDataResponse is the response struct for api DescribeRTSNativeSDKVvData
type DescribeRTSNativeSDKVvDataResponse struct {
	*responses.BaseResponse
	DataInterval string `json:"DataInterval" xml:"DataInterval"`
	EndTime      string `json:"EndTime" xml:"EndTime"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	StartTime    string `json:"StartTime" xml:"StartTime"`
	VvData       []Data `json:"VvData" xml:"VvData"`
}

// CreateDescribeRTSNativeSDKVvDataRequest creates a request to invoke DescribeRTSNativeSDKVvData API
func CreateDescribeRTSNativeSDKVvDataRequest() (request *DescribeRTSNativeSDKVvDataRequest) {
	request = &DescribeRTSNativeSDKVvDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeRTSNativeSDKVvData", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRTSNativeSDKVvDataResponse creates a response to parse from DescribeRTSNativeSDKVvData response
func CreateDescribeRTSNativeSDKVvDataResponse() (response *DescribeRTSNativeSDKVvDataResponse) {
	response = &DescribeRTSNativeSDKVvDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
