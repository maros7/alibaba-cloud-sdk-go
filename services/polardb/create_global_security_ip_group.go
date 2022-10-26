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

// CreateGlobalSecurityIPGroup invokes the polardb.CreateGlobalSecurityIPGroup API synchronously
func (client *Client) CreateGlobalSecurityIPGroup(request *CreateGlobalSecurityIPGroupRequest) (response *CreateGlobalSecurityIPGroupResponse, err error) {
	response = CreateCreateGlobalSecurityIPGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGlobalSecurityIPGroupWithChan invokes the polardb.CreateGlobalSecurityIPGroup API asynchronously
func (client *Client) CreateGlobalSecurityIPGroupWithChan(request *CreateGlobalSecurityIPGroupRequest) (<-chan *CreateGlobalSecurityIPGroupResponse, <-chan error) {
	responseChan := make(chan *CreateGlobalSecurityIPGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGlobalSecurityIPGroup(request)
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

// CreateGlobalSecurityIPGroupWithCallback invokes the polardb.CreateGlobalSecurityIPGroup API asynchronously
func (client *Client) CreateGlobalSecurityIPGroupWithCallback(request *CreateGlobalSecurityIPGroupRequest, callback func(response *CreateGlobalSecurityIPGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGlobalSecurityIPGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateGlobalSecurityIPGroup(request)
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

// CreateGlobalSecurityIPGroupRequest is the request struct for api CreateGlobalSecurityIPGroup
type CreateGlobalSecurityIPGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	GIpList              string           `position:"Query" name:"GIpList"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	SecurityIPType       string           `position:"Query" name:"SecurityIPType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	WhitelistNetType     string           `position:"Query" name:"WhitelistNetType"`
	GlobalIgName         string           `position:"Query" name:"GlobalIgName"`
}

// CreateGlobalSecurityIPGroupResponse is the response struct for api CreateGlobalSecurityIPGroup
type CreateGlobalSecurityIPGroupResponse struct {
	*responses.BaseResponse
	RequestId             string                      `json:"RequestId" xml:"RequestId"`
	GlobalSecurityIPGroup []GlobalSecurityIPGroupItem `json:"GlobalSecurityIPGroup" xml:"GlobalSecurityIPGroup"`
}

// CreateCreateGlobalSecurityIPGroupRequest creates a request to invoke CreateGlobalSecurityIPGroup API
func CreateCreateGlobalSecurityIPGroupRequest() (request *CreateGlobalSecurityIPGroupRequest) {
	request = &CreateGlobalSecurityIPGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "CreateGlobalSecurityIPGroup", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateGlobalSecurityIPGroupResponse creates a response to parse from CreateGlobalSecurityIPGroup response
func CreateCreateGlobalSecurityIPGroupResponse() (response *CreateGlobalSecurityIPGroupResponse) {
	response = &CreateGlobalSecurityIPGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
