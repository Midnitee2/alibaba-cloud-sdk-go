package cdrs

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

// Data is a nested struct in cdrs response
type Data struct {
	TotalPage    int                `json:"TotalPage" xml:"TotalPage"`
	PageSize     int                `json:"PageSize" xml:"PageSize"`
	TaskId       string             `json:"TaskId" xml:"TaskId"`
	PageNo       int                `json:"PageNo" xml:"PageNo"`
	TotalCount   int                `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int                `json:"PageNumber" xml:"PageNumber"`
	BodyList     []BodyListItem     `json:"BodyList" xml:"BodyList"`
	MotorList    []MotorListItem    `json:"MotorList" xml:"MotorList"`
	NonMotorList []NonMotorListItem `json:"NonMotorList" xml:"NonMotorList"`
	FaceList     []FaceListItem     `json:"FaceList" xml:"FaceList"`
	Records      []Record           `json:"Records" xml:"Records"`
}
