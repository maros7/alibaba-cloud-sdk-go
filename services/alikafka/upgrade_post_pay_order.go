package alikafka

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

// UpgradePostPayOrder invokes the alikafka.UpgradePostPayOrder API synchronously
func (client *Client) UpgradePostPayOrder(request *UpgradePostPayOrderRequest) (response *UpgradePostPayOrderResponse, err error) {
	response = CreateUpgradePostPayOrderResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradePostPayOrderWithChan invokes the alikafka.UpgradePostPayOrder API asynchronously
func (client *Client) UpgradePostPayOrderWithChan(request *UpgradePostPayOrderRequest) (<-chan *UpgradePostPayOrderResponse, <-chan error) {
	responseChan := make(chan *UpgradePostPayOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradePostPayOrder(request)
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

// UpgradePostPayOrderWithCallback invokes the alikafka.UpgradePostPayOrder API asynchronously
func (client *Client) UpgradePostPayOrderWithCallback(request *UpgradePostPayOrderRequest, callback func(response *UpgradePostPayOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradePostPayOrderResponse
		var err error
		defer close(result)
		response, err = client.UpgradePostPayOrder(request)
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

// UpgradePostPayOrderRequest is the request struct for api UpgradePostPayOrder
type UpgradePostPayOrderRequest struct {
	*requests.RpcRequest
	ServerlessConfig UpgradePostPayOrderServerlessConfig `position:"Query" name:"ServerlessConfig"  type:"Struct"`
	DiskSize         requests.Integer                    `position:"Query" name:"DiskSize"`
	IoMax            requests.Integer                    `position:"Query" name:"IoMax"`
	EipModel         requests.Boolean                    `position:"Query" name:"EipModel"`
	IoMaxSpec        string                              `position:"Query" name:"IoMaxSpec"`
	TopicQuota       requests.Integer                    `position:"Query" name:"TopicQuota"`
	EipMax           requests.Integer                    `position:"Query" name:"EipMax"`
	SpecType         string                              `position:"Query" name:"SpecType"`
	InstanceId       string                              `position:"Query" name:"InstanceId"`
	PartitionNum     requests.Integer                    `position:"Query" name:"PartitionNum"`
}

// UpgradePostPayOrderServerlessConfig is a repeated param struct in UpgradePostPayOrderRequest
type UpgradePostPayOrderServerlessConfig struct {
	ReservedPublishCapacity   string `name:"ReservedPublishCapacity"`
	ReservedSubscribeCapacity string `name:"ReservedSubscribeCapacity"`
}

// UpgradePostPayOrderResponse is the response struct for api UpgradePostPayOrder
type UpgradePostPayOrderResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpgradePostPayOrderRequest creates a request to invoke UpgradePostPayOrder API
func CreateUpgradePostPayOrderRequest() (request *UpgradePostPayOrderRequest) {
	request = &UpgradePostPayOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "UpgradePostPayOrder", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpgradePostPayOrderResponse creates a response to parse from UpgradePostPayOrder response
func CreateUpgradePostPayOrderResponse() (response *UpgradePostPayOrderResponse) {
	response = &UpgradePostPayOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
