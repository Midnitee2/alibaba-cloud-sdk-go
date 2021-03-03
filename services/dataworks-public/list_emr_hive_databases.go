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

// ListEmrHiveDatabases invokes the dataworks_public.ListEmrHiveDatabases API synchronously
func (client *Client) ListEmrHiveDatabases(request *ListEmrHiveDatabasesRequest) (response *ListEmrHiveDatabasesResponse, err error) {
	response = CreateListEmrHiveDatabasesResponse()
	err = client.DoAction(request, response)
	return
}

// ListEmrHiveDatabasesWithChan invokes the dataworks_public.ListEmrHiveDatabases API asynchronously
func (client *Client) ListEmrHiveDatabasesWithChan(request *ListEmrHiveDatabasesRequest) (<-chan *ListEmrHiveDatabasesResponse, <-chan error) {
	responseChan := make(chan *ListEmrHiveDatabasesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEmrHiveDatabases(request)
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

// ListEmrHiveDatabasesWithCallback invokes the dataworks_public.ListEmrHiveDatabases API asynchronously
func (client *Client) ListEmrHiveDatabasesWithCallback(request *ListEmrHiveDatabasesRequest, callback func(response *ListEmrHiveDatabasesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEmrHiveDatabasesResponse
		var err error
		defer close(result)
		response, err = client.ListEmrHiveDatabases(request)
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

// ListEmrHiveDatabasesRequest is the request struct for api ListEmrHiveDatabases
type ListEmrHiveDatabasesRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// ListEmrHiveDatabasesResponse is the response struct for api ListEmrHiveDatabases
type ListEmrHiveDatabasesResponse struct {
	*responses.BaseResponse
	ErrorCode    string         `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string         `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string         `json:"RequestId" xml:"RequestId"`
	Data         []HiveDatabase `json:"Data" xml:"Data"`
}

// CreateListEmrHiveDatabasesRequest creates a request to invoke ListEmrHiveDatabases API
func CreateListEmrHiveDatabasesRequest() (request *ListEmrHiveDatabasesRequest) {
	request = &ListEmrHiveDatabasesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2018-06-01", "ListEmrHiveDatabases", "", "")
	request.Method = requests.POST
	return
}

// CreateListEmrHiveDatabasesResponse creates a response to parse from ListEmrHiveDatabases response
func CreateListEmrHiveDatabasesResponse() (response *ListEmrHiveDatabasesResponse) {
	response = &ListEmrHiveDatabasesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}