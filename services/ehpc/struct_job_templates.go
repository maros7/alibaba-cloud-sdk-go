package ehpc

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

// JobTemplates is a nested struct in ehpc response
type JobTemplates struct {
	StdoutRedirectPath string `json:"StdoutRedirectPath" xml:"StdoutRedirectPath"`
	Variables          string `json:"Variables" xml:"Variables"`
	CommandLine        string `json:"CommandLine" xml:"CommandLine"`
	PackagePath        string `json:"PackagePath" xml:"PackagePath"`
	Priority           int    `json:"Priority" xml:"Priority"`
	ReRunable          bool   `json:"ReRunable" xml:"ReRunable"`
	Name               string `json:"Name" xml:"Name"`
	ArrayRequest       string `json:"ArrayRequest" xml:"ArrayRequest"`
	Id                 string `json:"Id" xml:"Id"`
	StderrRedirectPath string `json:"StderrRedirectPath" xml:"StderrRedirectPath"`
	RunasUser          string `json:"RunasUser" xml:"RunasUser"`
}
