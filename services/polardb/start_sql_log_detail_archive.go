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

// StartSqlLogDetailArchive invokes the polardb.StartSqlLogDetailArchive API synchronously
func (client *Client) StartSqlLogDetailArchive(request *StartSqlLogDetailArchiveRequest) (response *StartSqlLogDetailArchiveResponse, err error) {
	response = CreateStartSqlLogDetailArchiveResponse()
	err = client.DoAction(request, response)
	return
}

// StartSqlLogDetailArchiveWithChan invokes the polardb.StartSqlLogDetailArchive API asynchronously
func (client *Client) StartSqlLogDetailArchiveWithChan(request *StartSqlLogDetailArchiveRequest) (<-chan *StartSqlLogDetailArchiveResponse, <-chan error) {
	responseChan := make(chan *StartSqlLogDetailArchiveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartSqlLogDetailArchive(request)
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

// StartSqlLogDetailArchiveWithCallback invokes the polardb.StartSqlLogDetailArchive API asynchronously
func (client *Client) StartSqlLogDetailArchiveWithCallback(request *StartSqlLogDetailArchiveRequest, callback func(response *StartSqlLogDetailArchiveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartSqlLogDetailArchiveResponse
		var err error
		defer close(result)
		response, err = client.StartSqlLogDetailArchive(request)
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

// StartSqlLogDetailArchiveRequest is the request struct for api StartSqlLogDetailArchive
type StartSqlLogDetailArchiveRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Columns              string           `position:"Query" name:"Columns"`
	MinScanRows          requests.Integer `position:"Query" name:"MinScanRows"`
	StartTime            string           `position:"Query" name:"StartTime"`
	HostAddress          string           `position:"Query" name:"HostAddress"`
	AccountName          string           `position:"Query" name:"AccountName"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	State                string           `position:"Query" name:"State"`
	TableName            string           `position:"Query" name:"TableName"`
	SqlType              string           `position:"Query" name:"SqlType"`
	MinConsume           requests.Integer `position:"Query" name:"MinConsume"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	QueryKeyword         string           `position:"Query" name:"QueryKeyword"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MaxConsume           requests.Integer `position:"Query" name:"MaxConsume"`
	ThreadID             string           `position:"Query" name:"ThreadID"`
	LogicalOperator      string           `position:"Query" name:"LogicalOperator"`
	ChildDBInstanceIDs   string           `position:"Query" name:"ChildDBInstanceIDs"`
	DBName               string           `position:"Query" name:"DBName"`
	MaxScanRows          requests.Integer `position:"Query" name:"MaxScanRows"`
}

// StartSqlLogDetailArchiveResponse is the response struct for api StartSqlLogDetailArchive
type StartSqlLogDetailArchiveResponse struct {
	*responses.BaseResponse
	EndTime        string `json:"EndTime" xml:"EndTime"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceID   int    `json:"DBInstanceID" xml:"DBInstanceID"`
	ArchiveJobID   string `json:"ArchiveJobID" xml:"ArchiveJobID"`
	StartTime      string `json:"StartTime" xml:"StartTime"`
	OssTableName   string `json:"OssTableName" xml:"OssTableName"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
}

// CreateStartSqlLogDetailArchiveRequest creates a request to invoke StartSqlLogDetailArchive API
func CreateStartSqlLogDetailArchiveRequest() (request *StartSqlLogDetailArchiveRequest) {
	request = &StartSqlLogDetailArchiveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "StartSqlLogDetailArchive", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartSqlLogDetailArchiveResponse creates a response to parse from StartSqlLogDetailArchive response
func CreateStartSqlLogDetailArchiveResponse() (response *StartSqlLogDetailArchiveResponse) {
	response = &StartSqlLogDetailArchiveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
