package dts

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

// DescribeDataCheckTableDetails invokes the dts.DescribeDataCheckTableDetails API synchronously
func (client *Client) DescribeDataCheckTableDetails(request *DescribeDataCheckTableDetailsRequest) (response *DescribeDataCheckTableDetailsResponse, err error) {
	response = CreateDescribeDataCheckTableDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataCheckTableDetailsWithChan invokes the dts.DescribeDataCheckTableDetails API asynchronously
func (client *Client) DescribeDataCheckTableDetailsWithChan(request *DescribeDataCheckTableDetailsRequest) (<-chan *DescribeDataCheckTableDetailsResponse, <-chan error) {
	responseChan := make(chan *DescribeDataCheckTableDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataCheckTableDetails(request)
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

// DescribeDataCheckTableDetailsWithCallback invokes the dts.DescribeDataCheckTableDetails API asynchronously
func (client *Client) DescribeDataCheckTableDetailsWithCallback(request *DescribeDataCheckTableDetailsRequest, callback func(response *DescribeDataCheckTableDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataCheckTableDetailsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataCheckTableDetails(request)
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

// DescribeDataCheckTableDetailsRequest is the request struct for api DescribeDataCheckTableDetails
type DescribeDataCheckTableDetailsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	CheckType  requests.Integer `position:"Query" name:"CheckType"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	DtsJobId   string           `position:"Query" name:"DtsJobId"`
	TableName  string           `position:"Query" name:"TableName"`
	Status     string           `position:"Query" name:"Status"`
}

// DescribeDataCheckTableDetailsResponse is the response struct for api DescribeDataCheckTableDetails
type DescribeDataCheckTableDetailsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string             `json:"RequestId" xml:"RequestId"`
	ErrCode        string             `json:"ErrCode" xml:"ErrCode"`
	Success        bool               `json:"Success" xml:"Success"`
	ErrMessage     string             `json:"ErrMessage" xml:"ErrMessage"`
	DynamicMessage string             `json:"DynamicMessage" xml:"DynamicMessage"`
	PageNumber     int                `json:"PageNumber" xml:"PageNumber"`
	DynamicCode    string             `json:"DynamicCode" xml:"DynamicCode"`
	TotalCount     int64              `json:"TotalCount" xml:"TotalCount"`
	FinishedCount  int64              `json:"FinishedCount" xml:"FinishedCount"`
	DiffTableCount int64              `json:"DiffTableCount" xml:"DiffTableCount"`
	TableDetails   []TableDetailsItem `json:"TableDetails" xml:"TableDetails"`
}

// CreateDescribeDataCheckTableDetailsRequest creates a request to invoke DescribeDataCheckTableDetails API
func CreateDescribeDataCheckTableDetailsRequest() (request *DescribeDataCheckTableDetailsRequest) {
	request = &DescribeDataCheckTableDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeDataCheckTableDetails", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataCheckTableDetailsResponse creates a response to parse from DescribeDataCheckTableDetails response
func CreateDescribeDataCheckTableDetailsResponse() (response *DescribeDataCheckTableDetailsResponse) {
	response = &DescribeDataCheckTableDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
