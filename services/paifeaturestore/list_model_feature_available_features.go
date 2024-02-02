package paifeaturestore

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

// ListModelFeatureAvailableFeatures invokes the paifeaturestore.ListModelFeatureAvailableFeatures API synchronously
func (client *Client) ListModelFeatureAvailableFeatures(request *ListModelFeatureAvailableFeaturesRequest) (response *ListModelFeatureAvailableFeaturesResponse, err error) {
	response = CreateListModelFeatureAvailableFeaturesResponse()
	err = client.DoAction(request, response)
	return
}

// ListModelFeatureAvailableFeaturesWithChan invokes the paifeaturestore.ListModelFeatureAvailableFeatures API asynchronously
func (client *Client) ListModelFeatureAvailableFeaturesWithChan(request *ListModelFeatureAvailableFeaturesRequest) (<-chan *ListModelFeatureAvailableFeaturesResponse, <-chan error) {
	responseChan := make(chan *ListModelFeatureAvailableFeaturesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListModelFeatureAvailableFeatures(request)
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

// ListModelFeatureAvailableFeaturesWithCallback invokes the paifeaturestore.ListModelFeatureAvailableFeatures API asynchronously
func (client *Client) ListModelFeatureAvailableFeaturesWithCallback(request *ListModelFeatureAvailableFeaturesRequest, callback func(response *ListModelFeatureAvailableFeaturesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListModelFeatureAvailableFeaturesResponse
		var err error
		defer close(result)
		response, err = client.ListModelFeatureAvailableFeatures(request)
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

// ListModelFeatureAvailableFeaturesRequest is the request struct for api ListModelFeatureAvailableFeatures
type ListModelFeatureAvailableFeaturesRequest struct {
	*requests.RoaRequest
	ModelFeatureId string `position:"Path" name:"ModelFeatureId"`
	FeatureName    string `position:"Query" name:"FeatureName"`
	InstanceId     string `position:"Path" name:"InstanceId"`
}

// ListModelFeatureAvailableFeaturesResponse is the response struct for api ListModelFeatureAvailableFeatures
type ListModelFeatureAvailableFeaturesResponse struct {
	*responses.BaseResponse
	RequestId         string              `json:"requestId" xml:"requestId"`
	TotalCount        int64               `json:"TotalCount" xml:"TotalCount"`
	AvaliableFeatures []AvailableFeatures `json:"AvaliableFeatures" xml:"AvaliableFeatures"`
}

// CreateListModelFeatureAvailableFeaturesRequest creates a request to invoke ListModelFeatureAvailableFeatures API
func CreateListModelFeatureAvailableFeaturesRequest() (request *ListModelFeatureAvailableFeaturesRequest) {
	request = &ListModelFeatureAvailableFeaturesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "ListModelFeatureAvailableFeatures", "/api/v1/instances/[InstanceId]/modelfeatures/[ModelFeatureId]/availablefeatures", "", "")
	request.Method = requests.GET
	return
}

// CreateListModelFeatureAvailableFeaturesResponse creates a response to parse from ListModelFeatureAvailableFeatures response
func CreateListModelFeatureAvailableFeaturesResponse() (response *ListModelFeatureAvailableFeaturesResponse) {
	response = &ListModelFeatureAvailableFeaturesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
