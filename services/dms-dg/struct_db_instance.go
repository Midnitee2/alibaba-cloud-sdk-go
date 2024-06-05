package dms_dg

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

// DbInstance is a nested struct in dms_dg response
type DbInstance struct {
	Port           int    `json:"Port" xml:"Port"`
	VpcInstanceId  string `json:"VpcInstanceId" xml:"VpcInstanceId"`
	ParentId       string `json:"ParentId" xml:"ParentId"`
	GatewayName    string `json:"GatewayName" xml:"GatewayName"`
	GatewayId      string `json:"GatewayId" xml:"GatewayId"`
	ConnectHost    string `json:"ConnectHost" xml:"ConnectHost"`
	ConnectPort    int    `json:"ConnectPort" xml:"ConnectPort"`
	DbType         string `json:"DbType" xml:"DbType"`
	GmtCreate      int64  `json:"GmtCreate" xml:"GmtCreate"`
	Host           string `json:"Host" xml:"Host"`
	NodeId         string `json:"NodeId" xml:"NodeId"`
	RegionId       string `json:"RegionId" xml:"RegionId"`
	InstanceId     string `json:"InstanceId" xml:"InstanceId"`
	UserId         string `json:"UserId" xml:"UserId"`
	InstanceStatus string `json:"InstanceStatus" xml:"InstanceStatus"`
	DbDescription  string `json:"DbDescription" xml:"DbDescription"`
	NetworkType    string `json:"NetworkType" xml:"NetworkType"`
	VpcId          string `json:"VpcId" xml:"VpcId"`
	ServiceType    string `json:"ServiceType" xml:"ServiceType"`
}
