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

// Caster is a nested struct in live response
type Caster struct {
	CasterId        string                `json:"CasterId" xml:"CasterId"`
	CasterName      string                `json:"CasterName" xml:"CasterName"`
	CasterTemplate  string                `json:"CasterTemplate" xml:"CasterTemplate"`
	ChannelEnable   int                   `json:"ChannelEnable" xml:"ChannelEnable"`
	ChargeType      string                `json:"ChargeType" xml:"ChargeType"`
	CreateTime      string                `json:"CreateTime" xml:"CreateTime"`
	Duration        string                `json:"Duration" xml:"Duration"`
	ExpireTime      string                `json:"ExpireTime" xml:"ExpireTime"`
	LastModified    string                `json:"LastModified" xml:"LastModified"`
	NormType        int                   `json:"NormType" xml:"NormType"`
	PurchaseTime    string                `json:"PurchaseTime" xml:"PurchaseTime"`
	RoomId          int                   `json:"RoomId" xml:"RoomId"`
	StartTime       string                `json:"StartTime" xml:"StartTime"`
	Status          int                   `json:"Status" xml:"Status"`
	ResourceGroupId int                   `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ClientTokenId   string                `json:"ClientTokenId" xml:"ClientTokenId"`
	Tags            TagsInDescribeCasters `json:"Tags" xml:"Tags"`
}
