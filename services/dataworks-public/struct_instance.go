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

// Instance is a nested struct in dataworks_public response
type Instance struct {
	NodeId            int64  `json:"NodeId" xml:"NodeId"`
	InstanceId        int64  `json:"InstanceId" xml:"InstanceId"`
	DagId             int64  `json:"DagId" xml:"DagId"`
	DagType           string `json:"DagType" xml:"DagType"`
	Status            string `json:"Status" xml:"Status"`
	Bizdate           int64  `json:"Bizdate" xml:"Bizdate"`
	CycTime           int64  `json:"CycTime" xml:"CycTime"`
	CreateTime        int64  `json:"CreateTime" xml:"CreateTime"`
	ModifyTime        int64  `json:"ModifyTime" xml:"ModifyTime"`
	NodeName          string `json:"NodeName" xml:"NodeName"`
	BeginWaitTimeTime int64  `json:"BeginWaitTimeTime" xml:"BeginWaitTimeTime"`
	BeginWaitResTime  int64  `json:"BeginWaitResTime" xml:"BeginWaitResTime"`
	BeginRunningTime  int64  `json:"BeginRunningTime" xml:"BeginRunningTime"`
	ParamValues       string `json:"ParamValues" xml:"ParamValues"`
	FinishTime        int64  `json:"FinishTime" xml:"FinishTime"`
	Priority          int    `json:"Priority" xml:"Priority"`
	BaselineId        int64  `json:"BaselineId" xml:"BaselineId"`
	Repeatability     bool   `json:"Repeatability" xml:"Repeatability"`
	RepeatInterval    int64  `json:"RepeatInterval" xml:"RepeatInterval"`
	Connection        string `json:"Connection" xml:"Connection"`
	DqcType           int    `json:"DqcType" xml:"DqcType"`
	DqcDescription    string `json:"DqcDescription" xml:"DqcDescription"`
	ErrorMessage      string `json:"ErrorMessage" xml:"ErrorMessage"`
	RelatedFlowId     int64  `json:"RelatedFlowId" xml:"RelatedFlowId"`
	TaskType          string `json:"TaskType" xml:"TaskType"`
	TaskRerunTime     int    `json:"TaskRerunTime" xml:"TaskRerunTime"`
}
