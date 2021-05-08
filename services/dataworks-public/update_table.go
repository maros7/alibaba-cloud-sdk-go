package dataworks_public

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

// UpdateTable invokes the dataworks_public.UpdateTable API synchronously
func (client *Client) UpdateTable(request *UpdateTableRequest) (response *UpdateTableResponse, err error) {
	response = CreateUpdateTableResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTableWithChan invokes the dataworks_public.UpdateTable API asynchronously
func (client *Client) UpdateTableWithChan(request *UpdateTableRequest) (<-chan *UpdateTableResponse, <-chan error) {
	responseChan := make(chan *UpdateTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTable(request)
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

// UpdateTableWithCallback invokes the dataworks_public.UpdateTable API asynchronously
func (client *Client) UpdateTableWithCallback(request *UpdateTableRequest, callback func(response *UpdateTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTableResponse
		var err error
		defer close(result)
		response, err = client.UpdateTable(request)
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

// UpdateTableRequest is the request struct for api UpdateTable
type UpdateTableRequest struct {
	*requests.RpcRequest
	Columns           *[]UpdateTableColumns `position:"Body" name:"Columns"  type:"Repeated"`
	LifeCycle         requests.Integer      `position:"Query" name:"LifeCycle"`
	Themes            *[]UpdateTableThemes  `position:"Body" name:"Themes"  type:"Repeated"`
	LogicalLevelId    requests.Integer      `position:"Query" name:"LogicalLevelId"`
	Endpoint          string                `position:"Body" name:"Endpoint"`
	EnvType           requests.Integer      `position:"Body" name:"EnvType"`
	HasPart           requests.Integer      `position:"Query" name:"HasPart"`
	TableName         string                `position:"Query" name:"TableName"`
	AppGuid           string                `position:"Query" name:"AppGuid"`
	ProjectId         requests.Integer      `position:"Query" name:"ProjectId"`
	CategoryId        requests.Integer      `position:"Query" name:"CategoryId"`
	Visibility        requests.Integer      `position:"Query" name:"Visibility"`
	PhysicsLevelId    requests.Integer      `position:"Query" name:"PhysicsLevelId"`
	OwnerId           string                `position:"Query" name:"OwnerId"`
	IsView            requests.Integer      `position:"Query" name:"IsView"`
	ExternalTableType string                `position:"Query" name:"ExternalTableType"`
	Location          string                `position:"Query" name:"Location"`
	Comment           string                `position:"Query" name:"Comment"`
	CreateIfNotExists requests.Boolean      `position:"Query" name:"CreateIfNotExists"`
}

// UpdateTableColumns is a repeated param struct in UpdateTableRequest
type UpdateTableColumns struct {
	SeqNumber      string `name:"SeqNumber"`
	IsPartitionCol string `name:"IsPartitionCol"`
	ColumnNameCn   string `name:"ColumnNameCn"`
	Length         string `name:"Length"`
	Comment        string `name:"Comment"`
	ColumnName     string `name:"ColumnName"`
	ColumnType     string `name:"ColumnType"`
}

// UpdateTableThemes is a repeated param struct in UpdateTableRequest
type UpdateTableThemes struct {
	ThemeLevel string `name:"ThemeLevel"`
	ThemeId    string `name:"ThemeId"`
}

// UpdateTableResponse is the response struct for api UpdateTable
type UpdateTableResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	TaskInfo  TaskInfo `json:"TaskInfo" xml:"TaskInfo"`
}

// CreateUpdateTableRequest creates a request to invoke UpdateTable API
func CreateUpdateTableRequest() (request *UpdateTableRequest) {
	request = &UpdateTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateTable", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateTableResponse creates a response to parse from UpdateTable response
func CreateUpdateTableResponse() (response *UpdateTableResponse) {
	response = &UpdateTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
