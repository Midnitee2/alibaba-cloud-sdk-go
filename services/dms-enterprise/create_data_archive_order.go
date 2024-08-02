package dms_enterprise

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateDataArchiveOrder invokes the dms_enterprise.CreateDataArchiveOrder API synchronously
func (client *Client) CreateDataArchiveOrder(request *CreateDataArchiveOrderRequest) (response *CreateDataArchiveOrderResponse, err error) {
	response = CreateCreateDataArchiveOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataArchiveOrderWithChan invokes the dms_enterprise.CreateDataArchiveOrder API asynchronously
func (client *Client) CreateDataArchiveOrderWithChan(request *CreateDataArchiveOrderRequest) (<-chan *CreateDataArchiveOrderResponse, <-chan error) {
	responseChan := make(chan *CreateDataArchiveOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataArchiveOrder(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateDataArchiveOrderWithCallback invokes the dms_enterprise.CreateDataArchiveOrder API asynchronously
func (client *Client) CreateDataArchiveOrderWithCallback(request *CreateDataArchiveOrderRequest, callback func(response *CreateDataArchiveOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataArchiveOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateDataArchiveOrder(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateDataArchiveOrderRequest is the request struct for api CreateDataArchiveOrder
type CreateDataArchiveOrderRequest struct {
	*requests.RpcRequest
	Tid             requests.Integer            `position:"Query" name:"Tid"`
	ParentId        requests.Integer            `position:"Query" name:"ParentId"`
	PluginType      string                      `position:"Query" name:"PluginType"`
	Param           CreateDataArchiveOrderParam `position:"Query" name:"Param"  type:"Struct"`
	RelatedUserList *[]string                   `position:"Query" name:"RelatedUserList"  type:"Json"`
	Comment         string                      `position:"Query" name:"Comment"`
}

// CreateDataArchiveOrderParam is a repeated param struct in CreateDataArchiveOrderRequest
type CreateDataArchiveOrderParam struct {
	Variables          *[]CreateDataArchiveOrderParamVariablesItem     `name:"Variables" type:"Repeated"`
	SourceInstanceName string                                          `name:"SourceInstanceName"`
	CronStr            string                                          `name:"CronStr"`
	TableMapping       *[]string                                       `name:"TableMapping" type:"Repeated"`
	OrderAfter         *[]string                                       `name:"OrderAfter" type:"Repeated"`
	TargetInstanceHost string                                          `name:"TargetInstanceHost"`
	TableIncludes      *[]CreateDataArchiveOrderParamTableIncludesItem `name:"TableIncludes" type:"Repeated"`
	SourceCatalogName  string                                          `name:"SourceCatalogName"`
	RunMethod          string                                          `name:"RunMethod"`
	Logic              string                                          `name:"Logic"`
	SourceSchemaName   string                                          `name:"SourceSchemaName"`
	ArchiveMethod      string                                          `name:"ArchiveMethod"`
	DatabaseId         string                                          `name:"DatabaseId"`
}

// CreateDataArchiveOrderParamVariablesItem is a repeated param struct in CreateDataArchiveOrderRequest
type CreateDataArchiveOrderParamVariablesItem struct {
	Name    string `name:"Name"`
	Pattern string `name:"Pattern"`
}

// CreateDataArchiveOrderParamTableIncludesItem is a repeated param struct in CreateDataArchiveOrderRequest
type CreateDataArchiveOrderParamTableIncludesItem struct {
	TableWhere string `name:"TableWhere"`
	TableName  string `name:"TableName"`
}

// CreateDataArchiveOrderResponse is the response struct for api CreateDataArchiveOrder
type CreateDataArchiveOrderResponse struct {
	*responses.BaseResponse
	RequestId         string  `json:"RequestId" xml:"RequestId"`
	Success           bool    `json:"Success" xml:"Success"`
	ErrorMessage      string  `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode         string  `json:"ErrorCode" xml:"ErrorCode"`
	CreateOrderResult []int64 `json:"CreateOrderResult" xml:"CreateOrderResult"`
}

// CreateCreateDataArchiveOrderRequest creates a request to invoke CreateDataArchiveOrder API
func CreateCreateDataArchiveOrderRequest() (request *CreateDataArchiveOrderRequest) {
	request = &CreateDataArchiveOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateDataArchiveOrder", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDataArchiveOrderResponse creates a response to parse from CreateDataArchiveOrder response
func CreateCreateDataArchiveOrderResponse() (response *CreateDataArchiveOrderResponse) {
	response = &CreateDataArchiveOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
