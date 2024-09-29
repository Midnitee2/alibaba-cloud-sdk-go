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

// RulesItem is a nested struct in dataworks_public response
type RulesItem struct {
	BlockType                int    `json:"BlockType" xml:"BlockType"`
	OnDutyAccountName        string `json:"OnDutyAccountName" xml:"OnDutyAccountName"`
	Property                 string `json:"Property" xml:"Property"`
	WarningThreshold         string `json:"WarningThreshold" xml:"WarningThreshold"`
	TableName                string `json:"TableName" xml:"TableName"`
	OnDuty                   string `json:"OnDuty" xml:"OnDuty"`
	Comment                  string `json:"Comment" xml:"Comment"`
	RuleCheckerRelationId    int64  `json:"RuleCheckerRelationId" xml:"RuleCheckerRelationId"`
	FixCheck                 bool   `json:"FixCheck" xml:"FixCheck"`
	MethodId                 int    `json:"MethodId" xml:"MethodId"`
	TemplateName             string `json:"TemplateName" xml:"TemplateName"`
	Trend                    string `json:"Trend" xml:"Trend"`
	HistoryWarningThreshold  string `json:"HistoryWarningThreshold" xml:"HistoryWarningThreshold"`
	RuleType                 int    `json:"RuleType" xml:"RuleType"`
	MatchExpression          string `json:"MatchExpression" xml:"MatchExpression"`
	ProjectName              string `json:"ProjectName" xml:"ProjectName"`
	PropertyKey              string `json:"PropertyKey" xml:"PropertyKey"`
	CriticalThreshold        string `json:"CriticalThreshold" xml:"CriticalThreshold"`
	HistoryCriticalThreshold string `json:"HistoryCriticalThreshold" xml:"HistoryCriticalThreshold"`
	MethodName               string `json:"MethodName" xml:"MethodName"`
	CheckerId                int    `json:"CheckerId" xml:"CheckerId"`
	EntityId                 int64  `json:"EntityId" xml:"EntityId"`
	ExpectValue              string `json:"ExpectValue" xml:"ExpectValue"`
	TemplateId               int    `json:"TemplateId" xml:"TemplateId"`
	Id                       int64  `json:"Id" xml:"Id"`
	RuleName                 string `json:"RuleName" xml:"RuleName"`
}
