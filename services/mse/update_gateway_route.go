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

// UpdateGatewayRoute invokes the mse.UpdateGatewayRoute API synchronously
func (client *Client) UpdateGatewayRoute(request *UpdateGatewayRouteRequest) (response *UpdateGatewayRouteResponse, err error) {
	response = CreateUpdateGatewayRouteResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewayRouteWithChan invokes the mse.UpdateGatewayRoute API asynchronously
func (client *Client) UpdateGatewayRouteWithChan(request *UpdateGatewayRouteRequest) (<-chan *UpdateGatewayRouteResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewayRouteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewayRoute(request)
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

// UpdateGatewayRouteWithCallback invokes the mse.UpdateGatewayRoute API asynchronously
func (client *Client) UpdateGatewayRouteWithCallback(request *UpdateGatewayRouteRequest, callback func(response *UpdateGatewayRouteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewayRouteResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewayRoute(request)
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

// UpdateGatewayRouteRequest is the request struct for api UpdateGatewayRoute
type UpdateGatewayRouteRequest struct {
	*requests.RpcRequest
	MseSessionId       string                                `position:"Query" name:"MseSessionId"`
	DomainIdListJSON   string                                `position:"Query" name:"DomainIdListJSON"`
	Id                 requests.Integer                      `position:"Query" name:"Id"`
	GatewayId          requests.Integer                      `position:"Query" name:"GatewayId"`
	EnableWaf          requests.Boolean                      `position:"Query" name:"EnableWaf"`
	Predicates         UpdateGatewayRoutePredicates          `position:"Query" name:"Predicates"  type:"Struct"`
	DirectResponseJSON UpdateGatewayRouteDirectResponseJSON  `position:"Query" name:"DirectResponseJSON"  type:"Struct"`
	Name               string                                `position:"Query" name:"Name"`
	FallbackServices   *[]UpdateGatewayRouteFallbackServices `position:"Query" name:"FallbackServices"  type:"Json"`
	Fallback           requests.Boolean                      `position:"Query" name:"Fallback"`
	GatewayUniqueId    string                                `position:"Query" name:"GatewayUniqueId"`
	DestinationType    string                                `position:"Query" name:"DestinationType"`
	RouteOrder         requests.Integer                      `position:"Query" name:"RouteOrder"`
	Services           *[]UpdateGatewayRouteServices         `position:"Query" name:"Services"  type:"Json"`
	RedirectJSON       UpdateGatewayRouteRedirectJSON        `position:"Query" name:"RedirectJSON"  type:"Struct"`
	AcceptLanguage     string                                `position:"Query" name:"AcceptLanguage"`
}

// UpdateGatewayRouteFallbackServices is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRouteFallbackServices struct {
	AgreementType string `name:"AgreementType"`
	Name          string `name:"Name"`
	Namespace     string `name:"Namespace"`
	SourceType    string `name:"SourceType"`
	ServiceId     string `name:"ServiceId"`
	Percent       string `name:"Percent"`
	Version       string `name:"Version"`
	GroupName     string `name:"GroupName"`
	ServicePort   string `name:"ServicePort"`
}

// UpdateGatewayRouteServices is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRouteServices struct {
	HttpDubboTranscoder UpdateGatewayRouteServicesHttpDubboTranscoder `name:"HttpDubboTranscoder" type:"Struct"`
	AgreementType       string                                        `name:"AgreementType"`
	Name                string                                        `name:"Name"`
	Namespace           string                                        `name:"Namespace"`
	SourceType          string                                        `name:"SourceType"`
	ServiceId           string                                        `name:"ServiceId"`
	Percent             string                                        `name:"Percent"`
	Version             string                                        `name:"Version"`
	GroupName           string                                        `name:"GroupName"`
	ServicePort         string                                        `name:"ServicePort"`
}

// UpdateGatewayRouteServicesHttpDubboTranscoder is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRouteServicesHttpDubboTranscoder struct {
	DubboServiceName    string                                                            `name:"DubboServiceName"`
	MothedMapList       *[]UpdateGatewayRouteServicesHttpDubboTranscoderMothedMapListItem `name:"MothedMapList" type:"Repeated"`
	DubboServiceVersion string                                                            `name:"DubboServiceVersion"`
	DubboServiceGroup   string                                                            `name:"DubboServiceGroup"`
}

// UpdateGatewayRouteServicesHttpDubboTranscoderMothedMapListItem is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRouteServicesHttpDubboTranscoderMothedMapListItem struct {
	HttpMothed            string                                                                             `name:"HttpMothed"`
	ParamMapsList         *[]UpdateGatewayRouteServicesHttpDubboTranscoderMothedMapListItemParamMapsListItem `name:"ParamMapsList" type:"Repeated"`
	Mothedpath            string                                                                             `name:"Mothedpath"`
	DubboMothedName       string                                                                             `name:"DubboMothedName"`
	PassThroughAllHeaders string                                                                             `name:"PassThroughAllHeaders"`
	PassThroughList       *[]string                                                                          `name:"PassThroughList" type:"Repeated"`
}

// UpdateGatewayRouteServicesHttpDubboTranscoderMothedMapListItemParamMapsListItem is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRouteServicesHttpDubboTranscoderMothedMapListItemParamMapsListItem struct {
	ExtractKeySpec string `name:"ExtractKeySpec"`
	ExtractKey     string `name:"ExtractKey"`
	MappingType    string `name:"MappingType"`
}

// UpdateGatewayRoutePredicates is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRoutePredicates struct {
	PathPredicates   UpdateGatewayRoutePredicatesPathPredicates          `name:"PathPredicates" type:"Struct"`
	MethodPredicates *[]string                                           `name:"MethodPredicates" type:"Repeated"`
	HeaderPredicates *[]UpdateGatewayRoutePredicatesHeaderPredicatesItem `name:"HeaderPredicates" type:"Repeated"`
	QueryPredicates  *[]UpdateGatewayRoutePredicatesQueryPredicatesItem  `name:"QueryPredicates" type:"Repeated"`
}

// UpdateGatewayRouteDirectResponseJSON is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRouteDirectResponseJSON struct {
	Code string `name:"Code"`
	Body string `name:"Body"`
}

// UpdateGatewayRouteRedirectJSON is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRouteRedirectJSON struct {
	Path string `name:"Path"`
	Code string `name:"Code"`
	Host string `name:"Host"`
}

// UpdateGatewayRoutePredicatesPathPredicates is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRoutePredicatesPathPredicates struct {
	Path       string `name:"Path"`
	IgnoreCase string `name:"IgnoreCase"`
	Type       string `name:"Type"`
}

// UpdateGatewayRoutePredicatesHeaderPredicatesItem is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRoutePredicatesHeaderPredicatesItem struct {
	Type  string `name:"Type"`
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// UpdateGatewayRoutePredicatesQueryPredicatesItem is a repeated param struct in UpdateGatewayRouteRequest
type UpdateGatewayRoutePredicatesQueryPredicatesItem struct {
	Type  string `name:"Type"`
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// UpdateGatewayRouteResponse is the response struct for api UpdateGatewayRoute
type UpdateGatewayRouteResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           int64  `json:"Data" xml:"Data"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateUpdateGatewayRouteRequest creates a request to invoke UpdateGatewayRoute API
func CreateUpdateGatewayRouteRequest() (request *UpdateGatewayRouteRequest) {
	request = &UpdateGatewayRouteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewayRoute", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewayRouteResponse creates a response to parse from UpdateGatewayRoute response
func CreateUpdateGatewayRouteResponse() (response *UpdateGatewayRouteResponse) {
	response = &UpdateGatewayRouteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
