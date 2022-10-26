package polardb

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

// ModifySQLExplorerPolicy invokes the polardb.ModifySQLExplorerPolicy API synchronously
func (client *Client) ModifySQLExplorerPolicy(request *ModifySQLExplorerPolicyRequest) (response *ModifySQLExplorerPolicyResponse, err error) {
	response = CreateModifySQLExplorerPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySQLExplorerPolicyWithChan invokes the polardb.ModifySQLExplorerPolicy API asynchronously
func (client *Client) ModifySQLExplorerPolicyWithChan(request *ModifySQLExplorerPolicyRequest) (<-chan *ModifySQLExplorerPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifySQLExplorerPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySQLExplorerPolicy(request)
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

// ModifySQLExplorerPolicyWithCallback invokes the polardb.ModifySQLExplorerPolicy API asynchronously
func (client *Client) ModifySQLExplorerPolicyWithCallback(request *ModifySQLExplorerPolicyRequest, callback func(response *ModifySQLExplorerPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySQLExplorerPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifySQLExplorerPolicy(request)
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

// ModifySQLExplorerPolicyRequest is the request struct for api ModifySQLExplorerPolicy
type ModifySQLExplorerPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SQLCollectorStatus   string           `position:"Query" name:"SQLCollectorStatus"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifySQLExplorerPolicyResponse is the response struct for api ModifySQLExplorerPolicy
type ModifySQLExplorerPolicyResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
}

// CreateModifySQLExplorerPolicyRequest creates a request to invoke ModifySQLExplorerPolicy API
func CreateModifySQLExplorerPolicyRequest() (request *ModifySQLExplorerPolicyRequest) {
	request = &ModifySQLExplorerPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifySQLExplorerPolicy", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySQLExplorerPolicyResponse creates a response to parse from ModifySQLExplorerPolicy response
func CreateModifySQLExplorerPolicyResponse() (response *ModifySQLExplorerPolicyResponse) {
	response = &ModifySQLExplorerPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
