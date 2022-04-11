package vod

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

// VideoInGetVideoList is a nested struct in vod response
type VideoInGetVideoList struct {
	StorageLocation  string                  `json:"StorageLocation" xml:"StorageLocation"`
	Status           string                  `json:"Status" xml:"Status"`
	CreationTime     string                  `json:"CreationTime" xml:"CreationTime"`
	CateId           int64                   `json:"CateId" xml:"CateId"`
	VideoId          string                  `json:"VideoId" xml:"VideoId"`
	CreateTime       string                  `json:"CreateTime" xml:"CreateTime"`
	Tags             string                  `json:"Tags" xml:"Tags"`
	ModificationTime string                  `json:"ModificationTime" xml:"ModificationTime"`
	CateName         string                  `json:"CateName" xml:"CateName"`
	Description      string                  `json:"Description" xml:"Description"`
	AppId            string                  `json:"AppId" xml:"AppId"`
	Size             int64                   `json:"Size" xml:"Size"`
	CoverURL         string                  `json:"CoverURL" xml:"CoverURL"`
	Duration         float64                 `json:"Duration" xml:"Duration"`
	Title            string                  `json:"Title" xml:"Title"`
	ModifyTime       string                  `json:"ModifyTime" xml:"ModifyTime"`
	Snapshots        SnapshotsInGetVideoList `json:"Snapshots" xml:"Snapshots"`
}
