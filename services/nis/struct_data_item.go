package nis

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

// DataItem is a nested struct in nis response
type DataItem struct {
	BeginTime        string `json:"BeginTime" xml:"BeginTime"`
	Direction        string `json:"Direction" xml:"Direction"`
	InstanceId       string `json:"InstanceId" xml:"InstanceId"`
	AccessRegion     string `json:"AccessRegion" xml:"AccessRegion"`
	CloudIp          string `json:"CloudIp" xml:"CloudIp"`
	CloudPort        string `json:"CloudPort" xml:"CloudPort"`
	OtherIp          string `json:"OtherIp" xml:"OtherIp"`
	OtherPort        string `json:"OtherPort" xml:"OtherPort"`
	Protocol         string `json:"Protocol" xml:"Protocol"`
	CloudCountry     string `json:"CloudCountry" xml:"CloudCountry"`
	CloudProvince    string `json:"CloudProvince" xml:"CloudProvince"`
	CloudCity        string `json:"CloudCity" xml:"CloudCity"`
	CloudIsp         string `json:"CloudIsp" xml:"CloudIsp"`
	CloudProduct     string `json:"CloudProduct" xml:"CloudProduct"`
	OtherCountry     string `json:"OtherCountry" xml:"OtherCountry"`
	OtherProvince    string `json:"OtherProvince" xml:"OtherProvince"`
	OtherCity        string `json:"OtherCity" xml:"OtherCity"`
	OtherIsp         string `json:"OtherIsp" xml:"OtherIsp"`
	OtherProduct     string `json:"OtherProduct" xml:"OtherProduct"`
	Rtt              string `json:"Rtt" xml:"Rtt"`
	ByteCount        string `json:"ByteCount" xml:"ByteCount"`
	PacketCount      string `json:"PacketCount" xml:"PacketCount"`
	OutOrderCount    string `json:"OutOrderCount" xml:"OutOrderCount"`
	RetranCount      string `json:"RetranCount" xml:"RetranCount"`
	InByteCount      string `json:"InByteCount" xml:"InByteCount"`
	OutByteCount     string `json:"OutByteCount" xml:"OutByteCount"`
	InPacketCount    string `json:"InPacketCount" xml:"InPacketCount"`
	OutPacketCount   string `json:"OutPacketCount" xml:"OutPacketCount"`
	InOutOrderCount  string `json:"InOutOrderCount" xml:"InOutOrderCount"`
	OutOutOrderCount string `json:"OutOutOrderCount" xml:"OutOutOrderCount"`
	InRetranCount    string `json:"InRetranCount" xml:"InRetranCount"`
	OutRetranCount   string `json:"OutRetranCount" xml:"OutRetranCount"`
	RetransmitRate   string `json:"RetransmitRate" xml:"RetransmitRate"`
}
