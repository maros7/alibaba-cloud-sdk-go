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

// DataInGetGateway is a nested struct in mse response
type DataInGetGateway struct {
	Id               int64            `json:"Id" xml:"Id"`
	Name             string           `json:"Name" xml:"Name"`
	GatewayUniqueId  string           `json:"GatewayUniqueId" xml:"GatewayUniqueId"`
	Region           string           `json:"Region" xml:"Region"`
	PrimaryUser      string           `json:"PrimaryUser" xml:"PrimaryUser"`
	Status           int              `json:"Status" xml:"Status"`
	Vpc              string           `json:"Vpc" xml:"Vpc"`
	Vswitch          string           `json:"Vswitch" xml:"Vswitch"`
	SecurityGroup    string           `json:"SecurityGroup" xml:"SecurityGroup"`
	Spec             string           `json:"Spec" xml:"Spec"`
	Replica          int              `json:"Replica" xml:"Replica"`
	GmtCreate        string           `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified      string           `json:"GmtModified" xml:"GmtModified"`
	Vswitch2         string           `json:"Vswitch2" xml:"Vswitch2"`
	InstanceId       string           `json:"InstanceId" xml:"InstanceId"`
	ChargeType       string           `json:"ChargeType" xml:"ChargeType"`
	EndDate          string           `json:"EndDate" xml:"EndDate"`
	StatusDesc       string           `json:"StatusDesc" xml:"StatusDesc"`
	MseTag           string           `json:"MseTag" xml:"MseTag"`
	ResourceGroupId  string           `json:"ResourceGroupId" xml:"ResourceGroupId"`
	TotalReplica     int              `json:"TotalReplica" xml:"TotalReplica"`
	Elastic          bool             `json:"Elastic" xml:"Elastic"`
	ElasticReplica   int              `json:"ElasticReplica" xml:"ElasticReplica"`
	ElasticType      string           `json:"ElasticType" xml:"ElasticType"`
	XtraceDetails    XtraceDetails    `json:"XtraceDetails" xml:"XtraceDetails"`
	LogConfigDetails LogConfigDetails `json:"LogConfigDetails" xml:"LogConfigDetails"`
	ElasticPolicy    ElasticPolicy    `json:"ElasticPolicy" xml:"ElasticPolicy"`
}
