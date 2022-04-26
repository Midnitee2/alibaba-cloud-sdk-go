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

// CreateCallTags invokes the ccc.CreateCallTags API synchronously
func (client *Client) CreateCallTags(request *CreateCallTagsRequest) (response *CreateCallTagsResponse, err error) {
	response = CreateCreateCallTagsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCallTagsWithChan invokes the ccc.CreateCallTags API asynchronously
func (client *Client) CreateCallTagsWithChan(request *CreateCallTagsRequest) (<-chan *CreateCallTagsResponse, <-chan error) {
	responseChan := make(chan *CreateCallTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCallTags(request)
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

// CreateCallTagsWithCallback invokes the ccc.CreateCallTags API asynchronously
func (client *Client) CreateCallTagsWithCallback(request *CreateCallTagsRequest, callback func(response *CreateCallTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCallTagsResponse
		var err error
		defer close(result)
		response, err = client.CreateCallTags(request)
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

// CreateCallTagsRequest is the request struct for api CreateCallTags
type CreateCallTagsRequest struct {
	*requests.RpcRequest
	InstanceId      string `position:"Query" name:"InstanceId"`
	CallTagNameList string `position:"Query" name:"CallTagNameList"`
}

// CreateCallTagsResponse is the response struct for api CreateCallTags
type CreateCallTagsResponse struct {
	*responses.BaseResponse
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string        `json:"Code" xml:"Code"`
	Message        string        `json:"Message" xml:"Message"`
	Data           []FailureItem `json:"Data" xml:"Data"`
}

// CreateCreateCallTagsRequest creates a request to invoke CreateCallTags API
func CreateCreateCallTagsRequest() (request *CreateCallTagsRequest) {
	request = &CreateCallTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "CreateCallTags", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCallTagsResponse creates a response to parse from CreateCallTags response
func CreateCreateCallTagsResponse() (response *CreateCallTagsResponse) {
	response = &CreateCallTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
