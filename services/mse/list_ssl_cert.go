package mse

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

// ListSSLCert invokes the mse.ListSSLCert API synchronously
func (client *Client) ListSSLCert(request *ListSSLCertRequest) (response *ListSSLCertResponse, err error) {
	response = CreateListSSLCertResponse()
	err = client.DoAction(request, response)
	return
}

// ListSSLCertWithChan invokes the mse.ListSSLCert API asynchronously
func (client *Client) ListSSLCertWithChan(request *ListSSLCertRequest) (<-chan *ListSSLCertResponse, <-chan error) {
	responseChan := make(chan *ListSSLCertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSSLCert(request)
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

// ListSSLCertWithCallback invokes the mse.ListSSLCert API asynchronously
func (client *Client) ListSSLCertWithCallback(request *ListSSLCertRequest, callback func(response *ListSSLCertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSSLCertResponse
		var err error
		defer close(result)
		response, err = client.ListSSLCert(request)
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

// ListSSLCertRequest is the request struct for api ListSSLCert
type ListSSLCertRequest struct {
	*requests.RpcRequest
	GatewayUniqueId string `position:"Query" name:"GatewayUniqueId"`
	AcceptLanguage  string `position:"Query" name:"AcceptLanguage"`
}

// ListSSLCertResponse is the response struct for api ListSSLCert
type ListSSLCertResponse struct {
	*responses.BaseResponse
	RequestId      string    `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int       `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string    `json:"Message" xml:"Message"`
	Code           int       `json:"Code" xml:"Code"`
	Success        bool      `json:"Success" xml:"Success"`
	Data           []Domains `json:"Data" xml:"Data"`
}

// CreateListSSLCertRequest creates a request to invoke ListSSLCert API
func CreateListSSLCertRequest() (request *ListSSLCertRequest) {
	request = &ListSSLCertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListSSLCert", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListSSLCertResponse creates a response to parse from ListSSLCert response
func CreateListSSLCertResponse() (response *ListSSLCertResponse) {
	response = &ListSSLCertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
