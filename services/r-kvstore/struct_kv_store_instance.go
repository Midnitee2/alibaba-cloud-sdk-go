package r_kvstore

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

// KVStoreInstance is a nested struct in r_kvstore response
type KVStoreInstance struct {
	InstanceClass       string                  `json:"InstanceClass" xml:"InstanceClass"`
	PackageType         string                  `json:"PackageType" xml:"PackageType"`
	ChargeType          string                  `json:"ChargeType" xml:"ChargeType"`
	ConnectionDomain    string                  `json:"ConnectionDomain" xml:"ConnectionDomain"`
	CreateTime          string                  `json:"CreateTime" xml:"CreateTime"`
	EditionType         string                  `json:"EditionType" xml:"EditionType"`
	InstanceType        string                  `json:"InstanceType" xml:"InstanceType"`
	VpcCloudInstanceId  string                  `json:"VpcCloudInstanceId" xml:"VpcCloudInstanceId"`
	GlobalInstanceId    string                  `json:"GlobalInstanceId" xml:"GlobalInstanceId"`
	DestroyTime         string                  `json:"DestroyTime" xml:"DestroyTime"`
	RegionId            string                  `json:"RegionId" xml:"RegionId"`
	ResourceGroupId     string                  `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ComputingType       string                  `json:"ComputingType" xml:"ComputingType"`
	CloudType           string                  `json:"CloudType" xml:"CloudType"`
	PrivateIp           string                  `json:"PrivateIp" xml:"PrivateIp"`
	InstanceId          string                  `json:"InstanceId" xml:"InstanceId"`
	InstanceStatus      string                  `json:"InstanceStatus" xml:"InstanceStatus"`
	Bandwidth           int64                   `json:"Bandwidth" xml:"Bandwidth"`
	VpcId               string                  `json:"VpcId" xml:"VpcId"`
	NetworkType         string                  `json:"NetworkType" xml:"NetworkType"`
	HasRenewChangeOrder bool                    `json:"HasRenewChangeOrder" xml:"HasRenewChangeOrder"`
	ReadOnlyCount       string                  `json:"ReadOnlyCount" xml:"ReadOnlyCount"`
	NodeType            string                  `json:"NodeType" xml:"NodeType"`
	Connections         int64                   `json:"Connections" xml:"Connections"`
	ArchitectureType    string                  `json:"ArchitectureType" xml:"ArchitectureType"`
	ReplacateId         string                  `json:"ReplacateId" xml:"ReplacateId"`
	EngineVersion       string                  `json:"EngineVersion" xml:"EngineVersion"`
	ProxyCount          int                     `json:"ProxyCount" xml:"ProxyCount"`
	ShardClass          string                  `json:"ShardClass" xml:"ShardClass"`
	Capacity            int64                   `json:"Capacity" xml:"Capacity"`
	VSwitchId           string                  `json:"VSwitchId" xml:"VSwitchId"`
	InstanceName        string                  `json:"InstanceName" xml:"InstanceName"`
	SecondaryZoneId     string                  `json:"SecondaryZoneId" xml:"SecondaryZoneId"`
	Port                int64                   `json:"Port" xml:"Port"`
	EndTime             string                  `json:"EndTime" xml:"EndTime"`
	ZoneId              string                  `json:"ZoneId" xml:"ZoneId"`
	ShardCount          int                     `json:"ShardCount" xml:"ShardCount"`
	QPS                 int64                   `json:"QPS" xml:"QPS"`
	UserName            string                  `json:"UserName" xml:"UserName"`
	IsRds               bool                    `json:"IsRds" xml:"IsRds"`
	Config              string                  `json:"Config" xml:"Config"`
	ConnectionMode      string                  `json:"ConnectionMode" xml:"ConnectionMode"`
	Tags                TagsInDescribeInstances `json:"Tags" xml:"Tags"`
}
