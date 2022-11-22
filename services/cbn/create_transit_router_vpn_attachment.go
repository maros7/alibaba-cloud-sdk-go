package cbn

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

// CreateTransitRouterVpnAttachment invokes the cbn.CreateTransitRouterVpnAttachment API synchronously
func (client *Client) CreateTransitRouterVpnAttachment(request *CreateTransitRouterVpnAttachmentRequest) (response *CreateTransitRouterVpnAttachmentResponse, err error) {
	response = CreateCreateTransitRouterVpnAttachmentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTransitRouterVpnAttachmentWithChan invokes the cbn.CreateTransitRouterVpnAttachment API asynchronously
func (client *Client) CreateTransitRouterVpnAttachmentWithChan(request *CreateTransitRouterVpnAttachmentRequest) (<-chan *CreateTransitRouterVpnAttachmentResponse, <-chan error) {
	responseChan := make(chan *CreateTransitRouterVpnAttachmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTransitRouterVpnAttachment(request)
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

// CreateTransitRouterVpnAttachmentWithCallback invokes the cbn.CreateTransitRouterVpnAttachment API asynchronously
func (client *Client) CreateTransitRouterVpnAttachmentWithCallback(request *CreateTransitRouterVpnAttachmentRequest, callback func(response *CreateTransitRouterVpnAttachmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTransitRouterVpnAttachmentResponse
		var err error
		defer close(result)
		response, err = client.CreateTransitRouterVpnAttachment(request)
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

// CreateTransitRouterVpnAttachmentRequest is the request struct for api CreateTransitRouterVpnAttachment
type CreateTransitRouterVpnAttachmentRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                    requests.Integer                        `position:"Query" name:"ResourceOwnerId"`
	ClientToken                        string                                  `position:"Query" name:"ClientToken"`
	CenId                              string                                  `position:"Query" name:"CenId"`
	RouteTableAssociationEnabled       requests.Boolean                        `position:"Query" name:"RouteTableAssociationEnabled"`
	TransitRouterAttachmentName        string                                  `position:"Query" name:"TransitRouterAttachmentName"`
	Zone                               *[]CreateTransitRouterVpnAttachmentZone `position:"Query" name:"Zone"  type:"Repeated"`
	Tag                                *[]CreateTransitRouterVpnAttachmentTag  `position:"Query" name:"Tag"  type:"Repeated"`
	AutoPublishRouteEnabled            requests.Boolean                        `position:"Query" name:"AutoPublishRouteEnabled"`
	RouteTablePropagationEnabled       requests.Boolean                        `position:"Query" name:"RouteTablePropagationEnabled"`
	DryRun                             requests.Boolean                        `position:"Query" name:"DryRun"`
	ResourceOwnerAccount               string                                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                       string                                  `position:"Query" name:"OwnerAccount"`
	OwnerId                            requests.Integer                        `position:"Query" name:"OwnerId"`
	TransitRouterId                    string                                  `position:"Query" name:"TransitRouterId"`
	ResourceType                       string                                  `position:"Query" name:"ResourceType"`
	TransitRouterAttachmentDescription string                                  `position:"Query" name:"TransitRouterAttachmentDescription"`
	VpnOwnerId                         requests.Integer                        `position:"Query" name:"VpnOwnerId"`
	ChargeType                         string                                  `position:"Query" name:"ChargeType"`
	VpnId                              string                                  `position:"Query" name:"VpnId"`
}

// CreateTransitRouterVpnAttachmentZone is a repeated param struct in CreateTransitRouterVpnAttachmentRequest
type CreateTransitRouterVpnAttachmentZone struct {
	ZoneId string `name:"ZoneId"`
}

// CreateTransitRouterVpnAttachmentTag is a repeated param struct in CreateTransitRouterVpnAttachmentRequest
type CreateTransitRouterVpnAttachmentTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateTransitRouterVpnAttachmentResponse is the response struct for api CreateTransitRouterVpnAttachment
type CreateTransitRouterVpnAttachmentResponse struct {
	*responses.BaseResponse
	TransitRouterAttachmentId string `json:"TransitRouterAttachmentId" xml:"TransitRouterAttachmentId"`
	RequestId                 string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateTransitRouterVpnAttachmentRequest creates a request to invoke CreateTransitRouterVpnAttachment API
func CreateCreateTransitRouterVpnAttachmentRequest() (request *CreateTransitRouterVpnAttachmentRequest) {
	request = &CreateTransitRouterVpnAttachmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "CreateTransitRouterVpnAttachment", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTransitRouterVpnAttachmentResponse creates a response to parse from CreateTransitRouterVpnAttachment response
func CreateCreateTransitRouterVpnAttachmentResponse() (response *CreateTransitRouterVpnAttachmentResponse) {
	response = &CreateTransitRouterVpnAttachmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
