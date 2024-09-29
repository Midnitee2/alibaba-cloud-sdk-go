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

// Datas is a nested struct in dataworks_public response
type Datas struct {
	RuleName        string     `json:"RuleName" xml:"RuleName"`
	SceneName       string     `json:"SceneName" xml:"SceneName"`
	Desc            string     `json:"Desc" xml:"Desc"`
	SceneLevel      int        `json:"SceneLevel" xml:"SceneLevel"`
	AccountType     int        `json:"AccountType" xml:"AccountType"`
	ParentAccountId string     `json:"ParentAccountId" xml:"ParentAccountId"`
	GmtCreate       string     `json:"GmtCreate" xml:"GmtCreate"`
	DataType        string     `json:"DataType" xml:"DataType"`
	DesensWay       string     `json:"DesensWay" xml:"DesensWay"`
	Project         string     `json:"Project" xml:"Project"`
	YunAccount      string     `json:"YunAccount" xml:"YunAccount"`
	CheckWatermark  bool       `json:"CheckWatermark" xml:"CheckWatermark"`
	BaseId          string     `json:"BaseId" xml:"BaseId"`
	ClusterId       string     `json:"ClusterId" xml:"ClusterId"`
	Status          int        `json:"Status" xml:"Status"`
	GmtModified     string     `json:"GmtModified" xml:"GmtModified"`
	Id              int64      `json:"Id" xml:"Id"`
	Owner           string     `json:"Owner" xml:"Owner"`
	SceneCode       string     `json:"SceneCode" xml:"SceneCode"`
	UserGroups      string     `json:"UserGroups" xml:"UserGroups"`
	AccountName     string     `json:"AccountName" xml:"AccountName"`
	DesensRule      string     `json:"DesensRule" xml:"DesensRule"`
	DesenMode       string     `json:"DesenMode" xml:"DesenMode"`
	Children        []string   `json:"Children" xml:"Children"`
	DesensPlan      DesensPlan `json:"DesensPlan" xml:"DesensPlan"`
	Projects        []Project  `json:"Projects" xml:"Projects"`
}
