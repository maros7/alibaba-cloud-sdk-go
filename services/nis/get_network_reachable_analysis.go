package nis

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

// GetNetworkReachableAnalysis invokes the nis.GetNetworkReachableAnalysis API synchronously
func (client *Client) GetNetworkReachableAnalysis(request *GetNetworkReachableAnalysisRequest) (response *GetNetworkReachableAnalysisResponse, err error) {
	response = CreateGetNetworkReachableAnalysisResponse()
	err = client.DoAction(request, response)
	return
}

// GetNetworkReachableAnalysisWithChan invokes the nis.GetNetworkReachableAnalysis API asynchronously
func (client *Client) GetNetworkReachableAnalysisWithChan(request *GetNetworkReachableAnalysisRequest) (<-chan *GetNetworkReachableAnalysisResponse, <-chan error) {
	responseChan := make(chan *GetNetworkReachableAnalysisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNetworkReachableAnalysis(request)
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

// GetNetworkReachableAnalysisWithCallback invokes the nis.GetNetworkReachableAnalysis API asynchronously
func (client *Client) GetNetworkReachableAnalysisWithCallback(request *GetNetworkReachableAnalysisRequest, callback func(response *GetNetworkReachableAnalysisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNetworkReachableAnalysisResponse
		var err error
		defer close(result)
		response, err = client.GetNetworkReachableAnalysis(request)
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

// GetNetworkReachableAnalysisRequest is the request struct for api GetNetworkReachableAnalysis
type GetNetworkReachableAnalysisRequest struct {
	*requests.RpcRequest
	NetworkReachableAnalysisId string `position:"Query" name:"NetworkReachableAnalysisId"`
}

// GetNetworkReachableAnalysisResponse is the response struct for api GetNetworkReachableAnalysis
type GetNetworkReachableAnalysisResponse struct {
	*responses.BaseResponse
	NetworkPathId                  string `json:"NetworkPathId" xml:"NetworkPathId"`
	NetworkReachableAnalysisId     string `json:"NetworkReachableAnalysisId" xml:"NetworkReachableAnalysisId"`
	NetworkReachableAnalysisStatus string `json:"NetworkReachableAnalysisStatus" xml:"NetworkReachableAnalysisStatus"`
	NetworkReachableAnalysisResult string `json:"NetworkReachableAnalysisResult" xml:"NetworkReachableAnalysisResult"`
	Reachable                      bool   `json:"Reachable" xml:"Reachable"`
	CreateTime                     string `json:"CreateTime" xml:"CreateTime"`
	AliUid                         int64  `json:"AliUid" xml:"AliUid"`
	RequestId                      string `json:"RequestId" xml:"RequestId"`
	NetworkPathParameter           string `json:"NetworkPathParameter" xml:"NetworkPathParameter"`
}

// CreateGetNetworkReachableAnalysisRequest creates a request to invoke GetNetworkReachableAnalysis API
func CreateGetNetworkReachableAnalysisRequest() (request *GetNetworkReachableAnalysisRequest) {
	request = &GetNetworkReachableAnalysisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("nis", "2021-12-16", "GetNetworkReachableAnalysis", "networkana", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetNetworkReachableAnalysisResponse creates a response to parse from GetNetworkReachableAnalysis response
func CreateGetNetworkReachableAnalysisResponse() (response *GetNetworkReachableAnalysisResponse) {
	response = &GetNetworkReachableAnalysisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
