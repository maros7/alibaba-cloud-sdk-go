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

// ModifyEventStatus invokes the sddp.ModifyEventStatus API synchronously
func (client *Client) ModifyEventStatus(request *ModifyEventStatusRequest) (response *ModifyEventStatusResponse, err error) {
	response = CreateModifyEventStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyEventStatusWithChan invokes the sddp.ModifyEventStatus API asynchronously
func (client *Client) ModifyEventStatusWithChan(request *ModifyEventStatusRequest) (<-chan *ModifyEventStatusResponse, <-chan error) {
	responseChan := make(chan *ModifyEventStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyEventStatus(request)
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

// ModifyEventStatusWithCallback invokes the sddp.ModifyEventStatus API asynchronously
func (client *Client) ModifyEventStatusWithCallback(request *ModifyEventStatusRequest, callback func(response *ModifyEventStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyEventStatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyEventStatus(request)
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

// ModifyEventStatusRequest is the request struct for api ModifyEventStatus
type ModifyEventStatusRequest struct {
	*requests.RpcRequest
	Backed      requests.Boolean `position:"Query" name:"Backed"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	HandleInfo  string           `position:"Query" name:"HandleInfo"`
	DealReason  string           `position:"Query" name:"DealReason"`
	Id          requests.Integer `position:"Query" name:"Id"`
	Lang        string           `position:"Query" name:"Lang"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	Status      requests.Integer `position:"Query" name:"Status"`
}

// ModifyEventStatusResponse is the response struct for api ModifyEventStatus
type ModifyEventStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyEventStatusRequest creates a request to invoke ModifyEventStatus API
func CreateModifyEventStatusRequest() (request *ModifyEventStatusRequest) {
	request = &ModifyEventStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "ModifyEventStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyEventStatusResponse creates a response to parse from ModifyEventStatus response
func CreateModifyEventStatusResponse() (response *ModifyEventStatusResponse) {
	response = &ModifyEventStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
