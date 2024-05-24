package live

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

// DataItem is a nested struct in live response
type DataItem struct {
	AudioDuration int64  `json:"AudioDuration" xml:"AudioDuration"`
	SingleAudio   int64  `json:"Single_Audio" xml:"Single_Audio"`
	V480Duration  int64  `json:"V480Duration" xml:"V480Duration"`
	TotalDuration int64  `json:"TotalDuration" xml:"TotalDuration"`
	Timestamp     string `json:"Timestamp" xml:"Timestamp"`
	V720Duration  int64  `json:"V720Duration" xml:"V720Duration"`
	SingleVideo   int64  `json:"Single_Video" xml:"Single_Video"`
	V1080Duration int64  `json:"V1080Duration" xml:"V1080Duration"`
}
