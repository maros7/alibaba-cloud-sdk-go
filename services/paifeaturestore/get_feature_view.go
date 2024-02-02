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

// GetFeatureView invokes the paifeaturestore.GetFeatureView API synchronously
func (client *Client) GetFeatureView(request *GetFeatureViewRequest) (response *GetFeatureViewResponse, err error) {
	response = CreateGetFeatureViewResponse()
	err = client.DoAction(request, response)
	return
}

// GetFeatureViewWithChan invokes the paifeaturestore.GetFeatureView API asynchronously
func (client *Client) GetFeatureViewWithChan(request *GetFeatureViewRequest) (<-chan *GetFeatureViewResponse, <-chan error) {
	responseChan := make(chan *GetFeatureViewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFeatureView(request)
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

// GetFeatureViewWithCallback invokes the paifeaturestore.GetFeatureView API asynchronously
func (client *Client) GetFeatureViewWithCallback(request *GetFeatureViewRequest, callback func(response *GetFeatureViewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFeatureViewResponse
		var err error
		defer close(result)
		response, err = client.GetFeatureView(request)
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

// GetFeatureViewRequest is the request struct for api GetFeatureView
type GetFeatureViewRequest struct {
	*requests.RoaRequest
	InstanceId    string `position:"Path" name:"InstanceId"`
	FeatureViewId string `position:"Path" name:"FeatureViewId"`
}

// GetFeatureViewResponse is the response struct for api GetFeatureView
type GetFeatureViewResponse struct {
	*responses.BaseResponse
	RequestId              string       `json:"RequestId" xml:"RequestId"`
	ProjectId              string       `json:"ProjectId" xml:"ProjectId"`
	ProjectName            string       `json:"ProjectName" xml:"ProjectName"`
	FeatureEntityName      string       `json:"FeatureEntityName" xml:"FeatureEntityName"`
	Name                   string       `json:"Name" xml:"Name"`
	Owner                  string       `json:"Owner" xml:"Owner"`
	Type                   string       `json:"Type" xml:"Type"`
	GmtCreateTime          string       `json:"GmtCreateTime" xml:"GmtCreateTime"`
	GmtModifiedTime        string       `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
	FeatureEntityId        string       `json:"FeatureEntityId" xml:"FeatureEntityId"`
	JoinId                 string       `json:"JoinId" xml:"JoinId"`
	WriteMethod            string       `json:"WriteMethod" xml:"WriteMethod"`
	RegisterTable          string       `json:"RegisterTable" xml:"RegisterTable"`
	RegisterDatasourceId   string       `json:"RegisterDatasourceId" xml:"RegisterDatasourceId"`
	RegisterDatasourceName string       `json:"RegisterDatasourceName" xml:"RegisterDatasourceName"`
	SyncOnlineTable        bool         `json:"SyncOnlineTable" xml:"SyncOnlineTable"`
	TTL                    int          `json:"TTL" xml:"TTL"`
	Config                 string       `json:"Config" xml:"Config"`
	GmtSyncTime            string       `json:"GmtSyncTime" xml:"GmtSyncTime"`
	LastSyncConfig         string       `json:"LastSyncConfig" xml:"LastSyncConfig"`
	PublishTableScript     string       `json:"PublishTableScript" xml:"PublishTableScript"`
	Tags                   []string     `json:"Tags" xml:"Tags"`
	Fields                 []FieldsItem `json:"Fields" xml:"Fields"`
}

// CreateGetFeatureViewRequest creates a request to invoke GetFeatureView API
func CreateGetFeatureViewRequest() (request *GetFeatureViewRequest) {
	request = &GetFeatureViewRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "GetFeatureView", "/api/v1/instances/[InstanceId]/featureviews/[FeatureViewId]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetFeatureViewResponse creates a response to parse from GetFeatureView response
func CreateGetFeatureViewResponse() (response *GetFeatureViewResponse) {
	response = &GetFeatureViewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
