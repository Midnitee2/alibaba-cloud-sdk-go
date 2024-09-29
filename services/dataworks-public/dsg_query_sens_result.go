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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DsgQuerySensResult invokes the dataworks_public.DsgQuerySensResult API synchronously
func (client *Client) DsgQuerySensResult(request *DsgQuerySensResultRequest) (response *DsgQuerySensResultResponse, err error) {
	response = CreateDsgQuerySensResultResponse()
	err = client.DoAction(request, response)
	return
}

// DsgQuerySensResultWithChan invokes the dataworks_public.DsgQuerySensResult API asynchronously
func (client *Client) DsgQuerySensResultWithChan(request *DsgQuerySensResultRequest) (<-chan *DsgQuerySensResultResponse, <-chan error) {
	responseChan := make(chan *DsgQuerySensResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DsgQuerySensResult(request)
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

// DsgQuerySensResultWithCallback invokes the dataworks_public.DsgQuerySensResult API asynchronously
func (client *Client) DsgQuerySensResultWithCallback(request *DsgQuerySensResultRequest, callback func(response *DsgQuerySensResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DsgQuerySensResultResponse
		var err error
		defer close(result)
		response, err = client.DsgQuerySensResult(request)
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

// DsgQuerySensResultRequest is the request struct for api DsgQuerySensResult
type DsgQuerySensResultRequest struct {
	*requests.RpcRequest
	Col           string           `position:"Body" name:"Col"`
	ProjectName   string           `position:"Body" name:"ProjectName"`
	SchemaName    string           `position:"Body" name:"SchemaName"`
	Level         string           `position:"Body" name:"Level"`
	SensStatus    string           `position:"Body" name:"SensStatus"`
	NodeName      string           `position:"Body" name:"NodeName"`
	SensitiveId   string           `position:"Body" name:"SensitiveId"`
	PageNo        requests.Integer `position:"Body" name:"PageNo"`
	TenantId      string           `position:"Body" name:"TenantId"`
	DbType        string           `position:"Body" name:"DbType"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	SensitiveName string           `position:"Body" name:"SensitiveName"`
	Table         string           `position:"Body" name:"Table"`
	Order         string           `position:"Body" name:"Order"`
	OrderField    string           `position:"Body" name:"OrderField"`
}

// DsgQuerySensResultResponse is the response struct for api DsgQuerySensResult
type DsgQuerySensResultResponse struct {
	*responses.BaseResponse
}

// CreateDsgQuerySensResultRequest creates a request to invoke DsgQuerySensResult API
func CreateDsgQuerySensResultRequest() (request *DsgQuerySensResultRequest) {
	request = &DsgQuerySensResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DsgQuerySensResult", "", "")
	request.Method = requests.POST
	return
}

// CreateDsgQuerySensResultResponse creates a response to parse from DsgQuerySensResult response
func CreateDsgQuerySensResultResponse() (response *DsgQuerySensResultResponse) {
	response = &DsgQuerySensResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
