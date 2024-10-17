package domain

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

// Module is a nested struct in domain response
type Module struct {
	Domain                      string                       `json:"Domain" xml:"Domain"`
	ProductId                   string                       `json:"ProductId" xml:"ProductId"`
	BizName                     string                       `json:"BizName" xml:"BizName"`
	SaleId                      string                       `json:"SaleId" xml:"SaleId"`
	DownloadUrl                 string                       `json:"DownloadUrl" xml:"DownloadUrl"`
	CreateTime                  int64                        `json:"CreateTime" xml:"CreateTime"`
	PageSize                    int                          `json:"PageSize" xml:"PageSize"`
	EndTime                     int64                        `json:"EndTime" xml:"EndTime"`
	OrderNo                     string                       `json:"OrderNo" xml:"OrderNo"`
	OrderId                     string                       `json:"OrderId" xml:"OrderId"`
	Premium                     bool                         `json:"Premium" xml:"Premium"`
	GmtCreate                   string                       `json:"GmtCreate" xml:"GmtCreate"`
	BizType                     string                       `json:"BizType" xml:"BizType"`
	UserId                      string                       `json:"UserId" xml:"UserId"`
	PayPrice                    int64                        `json:"PayPrice" xml:"PayPrice"`
	TotalPageNum                int                          `json:"TotalPageNum" xml:"TotalPageNum"`
	DeadDate                    int64                        `json:"DeadDate" xml:"DeadDate"`
	CurrentPageNum              int                          `json:"CurrentPageNum" xml:"CurrentPageNum"`
	BizStatus                   string                       `json:"BizStatus" xml:"BizStatus"`
	Status                      int                          `json:"Status" xml:"Status"`
	UpdateTime                  int64                        `json:"UpdateTime" xml:"UpdateTime"`
	DomainName                  string                       `json:"DomainName" xml:"DomainName"`
	GmtModified                 string                       `json:"GmtModified" xml:"GmtModified"`
	Id                          int64                        `json:"Id" xml:"Id"`
	Currency                    string                       `json:"Currency" xml:"Currency"`
	PayUrl                      string                       `json:"PayUrl" xml:"PayUrl"`
	RegDate                     int64                        `json:"RegDate" xml:"RegDate"`
	AuditMsg                    string                       `json:"AuditMsg" xml:"AuditMsg"`
	BizNo                       string                       `json:"BizNo" xml:"BizNo"`
	TotalItemNum                int                          `json:"TotalItemNum" xml:"TotalItemNum"`
	StatusDesc                  string                       `json:"StatusDesc" xml:"StatusDesc"`
	Price                       int64                        `json:"Price" xml:"Price"`
	DomainSpecialBizContact     DomainSpecialBizContact      `json:"DomainSpecialBizContact" xml:"DomainSpecialBizContact"`
	DomainSpecialOrderResult    DomainSpecialOrderResult     `json:"DomainSpecialOrderResult" xml:"DomainSpecialOrderResult"`
	DomainSpecialBizCredentials []DomainSpecialBizCredential `json:"DomainSpecialBizCredentials" xml:"DomainSpecialBizCredentials"`
	Data                        []OrderList                  `json:"Data" xml:"Data"`
}
