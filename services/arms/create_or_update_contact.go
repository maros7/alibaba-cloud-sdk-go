package arms

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

// CreateOrUpdateContact invokes the arms.CreateOrUpdateContact API synchronously
func (client *Client) CreateOrUpdateContact(request *CreateOrUpdateContactRequest) (response *CreateOrUpdateContactResponse, err error) {
	response = CreateCreateOrUpdateContactResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrUpdateContactWithChan invokes the arms.CreateOrUpdateContact API asynchronously
func (client *Client) CreateOrUpdateContactWithChan(request *CreateOrUpdateContactRequest) (<-chan *CreateOrUpdateContactResponse, <-chan error) {
	responseChan := make(chan *CreateOrUpdateContactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrUpdateContact(request)
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

// CreateOrUpdateContactWithCallback invokes the arms.CreateOrUpdateContact API asynchronously
func (client *Client) CreateOrUpdateContactWithCallback(request *CreateOrUpdateContactRequest, callback func(response *CreateOrUpdateContactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrUpdateContactResponse
		var err error
		defer close(result)
		response, err = client.CreateOrUpdateContact(request)
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

// CreateOrUpdateContactRequest is the request struct for api CreateOrUpdateContact
type CreateOrUpdateContactRequest struct {
	*requests.RpcRequest
	ContactName string           `position:"Body" name:"ContactName"`
	ContactId   requests.Integer `position:"Body" name:"ContactId"`
	Phone       string           `position:"Body" name:"Phone"`
	IsVerify    requests.Boolean `position:"Body" name:"IsVerify"`
	ProxyUserId string           `position:"Body" name:"ProxyUserId"`
	Email       string           `position:"Body" name:"Email"`
}

// CreateOrUpdateContactResponse is the response struct for api CreateOrUpdateContact
type CreateOrUpdateContactResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	AlertContact AlertContact `json:"AlertContact" xml:"AlertContact"`
}

// CreateCreateOrUpdateContactRequest creates a request to invoke CreateOrUpdateContact API
func CreateCreateOrUpdateContactRequest() (request *CreateOrUpdateContactRequest) {
	request = &CreateOrUpdateContactRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CreateOrUpdateContact", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateOrUpdateContactResponse creates a response to parse from CreateOrUpdateContact response
func CreateCreateOrUpdateContactResponse() (response *CreateOrUpdateContactResponse) {
	response = &CreateOrUpdateContactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
