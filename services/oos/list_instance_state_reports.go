package oos

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

// ListInstanceStateReports invokes the oos.ListInstanceStateReports API synchronously
func (client *Client) ListInstanceStateReports(request *ListInstanceStateReportsRequest) (response *ListInstanceStateReportsResponse, err error) {
	response = CreateListInstanceStateReportsResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstanceStateReportsWithChan invokes the oos.ListInstanceStateReports API asynchronously
func (client *Client) ListInstanceStateReportsWithChan(request *ListInstanceStateReportsRequest) (<-chan *ListInstanceStateReportsResponse, <-chan error) {
	responseChan := make(chan *ListInstanceStateReportsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstanceStateReports(request)
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

// ListInstanceStateReportsWithCallback invokes the oos.ListInstanceStateReports API asynchronously
func (client *Client) ListInstanceStateReportsWithCallback(request *ListInstanceStateReportsRequest, callback func(response *ListInstanceStateReportsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstanceStateReportsResponse
		var err error
		defer close(result)
		response, err = client.ListInstanceStateReports(request)
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

// ListInstanceStateReportsRequest is the request struct for api ListInstanceStateReports
type ListInstanceStateReportsRequest struct {
	*requests.RpcRequest
	InstanceId           string           `position:"Query" name:"InstanceId"`
	NextToken            string           `position:"Query" name:"NextToken"`
	MaxResults           requests.Integer `position:"Query" name:"MaxResults"`
	StateConfigurationId string           `position:"Query" name:"StateConfigurationId"`
}

// ListInstanceStateReportsResponse is the response struct for api ListInstanceStateReports
type ListInstanceStateReportsResponse struct {
	*responses.BaseResponse
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	MaxResults   int           `json:"MaxResults" xml:"MaxResults"`
	NextToken    string        `json:"NextToken" xml:"NextToken"`
	StateReports []StateReport `json:"StateReports" xml:"StateReports"`
}

// CreateListInstanceStateReportsRequest creates a request to invoke ListInstanceStateReports API
func CreateListInstanceStateReportsRequest() (request *ListInstanceStateReportsRequest) {
	request = &ListInstanceStateReportsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "ListInstanceStateReports", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInstanceStateReportsResponse creates a response to parse from ListInstanceStateReports response
func CreateListInstanceStateReportsResponse() (response *ListInstanceStateReportsResponse) {
	response = &ListInstanceStateReportsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}