package ccc

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

// ListAgentSummaryReportsSinceMidnight invokes the ccc.ListAgentSummaryReportsSinceMidnight API synchronously
func (client *Client) ListAgentSummaryReportsSinceMidnight(request *ListAgentSummaryReportsSinceMidnightRequest) (response *ListAgentSummaryReportsSinceMidnightResponse, err error) {
	response = CreateListAgentSummaryReportsSinceMidnightResponse()
	err = client.DoAction(request, response)
	return
}

// ListAgentSummaryReportsSinceMidnightWithChan invokes the ccc.ListAgentSummaryReportsSinceMidnight API asynchronously
func (client *Client) ListAgentSummaryReportsSinceMidnightWithChan(request *ListAgentSummaryReportsSinceMidnightRequest) (<-chan *ListAgentSummaryReportsSinceMidnightResponse, <-chan error) {
	responseChan := make(chan *ListAgentSummaryReportsSinceMidnightResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAgentSummaryReportsSinceMidnight(request)
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

// ListAgentSummaryReportsSinceMidnightWithCallback invokes the ccc.ListAgentSummaryReportsSinceMidnight API asynchronously
func (client *Client) ListAgentSummaryReportsSinceMidnightWithCallback(request *ListAgentSummaryReportsSinceMidnightRequest, callback func(response *ListAgentSummaryReportsSinceMidnightResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAgentSummaryReportsSinceMidnightResponse
		var err error
		defer close(result)
		response, err = client.ListAgentSummaryReportsSinceMidnight(request)
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

// ListAgentSummaryReportsSinceMidnightRequest is the request struct for api ListAgentSummaryReportsSinceMidnight
type ListAgentSummaryReportsSinceMidnightRequest struct {
	*requests.RpcRequest
	AgentIds     string           `position:"Query" name:"AgentIds"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
}

// ListAgentSummaryReportsSinceMidnightResponse is the response struct for api ListAgentSummaryReportsSinceMidnight
type ListAgentSummaryReportsSinceMidnightResponse struct {
	*responses.BaseResponse
	HttpStatusCode          int                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code                    string                  `json:"Code" xml:"Code"`
	Message                 string                  `json:"Message" xml:"Message"`
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	Success                 bool                    `json:"Success" xml:"Success"`
	PagedAgentSummaryReport PagedAgentSummaryReport `json:"PagedAgentSummaryReport" xml:"PagedAgentSummaryReport"`
}

// CreateListAgentSummaryReportsSinceMidnightRequest creates a request to invoke ListAgentSummaryReportsSinceMidnight API
func CreateListAgentSummaryReportsSinceMidnightRequest() (request *ListAgentSummaryReportsSinceMidnightRequest) {
	request = &ListAgentSummaryReportsSinceMidnightRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListAgentSummaryReportsSinceMidnight", "", "")
	request.Method = requests.GET
	return
}

// CreateListAgentSummaryReportsSinceMidnightResponse creates a response to parse from ListAgentSummaryReportsSinceMidnight response
func CreateListAgentSummaryReportsSinceMidnightResponse() (response *ListAgentSummaryReportsSinceMidnightResponse) {
	response = &ListAgentSummaryReportsSinceMidnightResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
