package quickbi_public

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

// QueryComponentPerformance invokes the quickbi_public.QueryComponentPerformance API synchronously
func (client *Client) QueryComponentPerformance(request *QueryComponentPerformanceRequest) (response *QueryComponentPerformanceResponse, err error) {
	response = CreateQueryComponentPerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryComponentPerformanceWithChan invokes the quickbi_public.QueryComponentPerformance API asynchronously
func (client *Client) QueryComponentPerformanceWithChan(request *QueryComponentPerformanceRequest) (<-chan *QueryComponentPerformanceResponse, <-chan error) {
	responseChan := make(chan *QueryComponentPerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryComponentPerformance(request)
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

// QueryComponentPerformanceWithCallback invokes the quickbi_public.QueryComponentPerformance API asynchronously
func (client *Client) QueryComponentPerformanceWithCallback(request *QueryComponentPerformanceRequest, callback func(response *QueryComponentPerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryComponentPerformanceResponse
		var err error
		defer close(result)
		response, err = client.QueryComponentPerformance(request)
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

// QueryComponentPerformanceRequest is the request struct for api QueryComponentPerformance
type QueryComponentPerformanceRequest struct {
	*requests.RpcRequest
	ReportId       string           `position:"Query" name:"ReportId"`
	AccessPoint    string           `position:"Query" name:"AccessPoint"`
	PageNum        requests.Integer `position:"Query" name:"PageNum"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	CostTimeAvgMin requests.Integer `position:"Query" name:"CostTimeAvgMin"`
	QueryType      string           `position:"Query" name:"QueryType"`
	SignType       string           `position:"Query" name:"SignType"`
	ResourceType   string           `position:"Query" name:"ResourceType"`
	WorkspaceId    string           `position:"Query" name:"WorkspaceId"`
}

// QueryComponentPerformanceResponse is the response struct for api QueryComponentPerformance
type QueryComponentPerformanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    []Data `json:"Result" xml:"Result"`
}

// CreateQueryComponentPerformanceRequest creates a request to invoke QueryComponentPerformance API
func CreateQueryComponentPerformanceRequest() (request *QueryComponentPerformanceRequest) {
	request = &QueryComponentPerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryComponentPerformance", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryComponentPerformanceResponse creates a response to parse from QueryComponentPerformance response
func CreateQueryComponentPerformanceResponse() (response *QueryComponentPerformanceResponse) {
	response = &QueryComponentPerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
