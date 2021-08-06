package edas

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

// ListCluster invokes the edas.ListCluster API synchronously
func (client *Client) ListCluster(request *ListClusterRequest) (response *ListClusterResponse, err error) {
	response = CreateListClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterWithChan invokes the edas.ListCluster API asynchronously
func (client *Client) ListClusterWithChan(request *ListClusterRequest) (<-chan *ListClusterResponse, <-chan error) {
	responseChan := make(chan *ListClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCluster(request)
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

// ListClusterWithCallback invokes the edas.ListCluster API asynchronously
func (client *Client) ListClusterWithCallback(request *ListClusterRequest, callback func(response *ListClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterResponse
		var err error
		defer close(result)
		response, err = client.ListCluster(request)
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

// ListClusterRequest is the request struct for api ListCluster
type ListClusterRequest struct {
	*requests.RoaRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	LogicalRegionId string `position:"Query" name:"LogicalRegionId"`
}

// ListClusterResponse is the response struct for api ListCluster
type ListClusterResponse struct {
	*responses.BaseResponse
	Code        int                      `json:"Code" xml:"Code"`
	Message     string                   `json:"Message" xml:"Message"`
	RequestId   string                   `json:"RequestId" xml:"RequestId"`
	ClusterList ClusterListInListCluster `json:"ClusterList" xml:"ClusterList"`
}

// CreateListClusterRequest creates a request to invoke ListCluster API
func CreateListClusterRequest() (request *ListClusterRequest) {
	request = &ListClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListCluster", "/pop/v5/resource/cluster_list", "edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListClusterResponse creates a response to parse from ListCluster response
func CreateListClusterResponse() (response *ListClusterResponse) {
	response = &ListClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
