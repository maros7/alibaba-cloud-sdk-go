package imageprocess

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

// PredictCVD invokes the imageprocess.PredictCVD API synchronously
func (client *Client) PredictCVD(request *PredictCVDRequest) (response *PredictCVDResponse, err error) {
	response = CreatePredictCVDResponse()
	err = client.DoAction(request, response)
	return
}

// PredictCVDWithChan invokes the imageprocess.PredictCVD API asynchronously
func (client *Client) PredictCVDWithChan(request *PredictCVDRequest) (<-chan *PredictCVDResponse, <-chan error) {
	responseChan := make(chan *PredictCVDResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PredictCVD(request)
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

// PredictCVDWithCallback invokes the imageprocess.PredictCVD API asynchronously
func (client *Client) PredictCVDWithCallback(request *PredictCVDRequest, callback func(response *PredictCVDResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PredictCVDResponse
		var err error
		defer close(result)
		response, err = client.PredictCVD(request)
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

// PredictCVDRequest is the request struct for api PredictCVD
type PredictCVDRequest struct {
	*requests.RpcRequest
	DataSourceType string               `position:"Body" name:"DataSourceType"`
	OrgName        string               `position:"Body" name:"OrgName"`
	DataFormat     string               `position:"Body" name:"DataFormat"`
	URLList        *[]PredictCVDURLList `position:"Body" name:"URLList"  type:"Repeated"`
	OrgId          string               `position:"Body" name:"OrgId"`
	Async          requests.Boolean     `position:"Body" name:"Async"`
}

// PredictCVDURLList is a repeated param struct in PredictCVDRequest
type PredictCVDURLList struct {
	URL string `name:"URL"`
}

// PredictCVDResponse is the response struct for api PredictCVD
type PredictCVDResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreatePredictCVDRequest creates a request to invoke PredictCVD API
func CreatePredictCVDRequest() (request *PredictCVDRequest) {
	request = &PredictCVDRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageprocess", "2020-03-20", "PredictCVD", "imageprocess", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePredictCVDResponse creates a response to parse from PredictCVD response
func CreatePredictCVDResponse() (response *PredictCVDResponse) {
	response = &PredictCVDResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
