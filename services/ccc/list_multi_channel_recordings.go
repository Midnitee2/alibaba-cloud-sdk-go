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

// ListMultiChannelRecordings invokes the ccc.ListMultiChannelRecordings API synchronously
func (client *Client) ListMultiChannelRecordings(request *ListMultiChannelRecordingsRequest) (response *ListMultiChannelRecordingsResponse, err error) {
	response = CreateListMultiChannelRecordingsResponse()
	err = client.DoAction(request, response)
	return
}

// ListMultiChannelRecordingsWithChan invokes the ccc.ListMultiChannelRecordings API asynchronously
func (client *Client) ListMultiChannelRecordingsWithChan(request *ListMultiChannelRecordingsRequest) (<-chan *ListMultiChannelRecordingsResponse, <-chan error) {
	responseChan := make(chan *ListMultiChannelRecordingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMultiChannelRecordings(request)
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

// ListMultiChannelRecordingsWithCallback invokes the ccc.ListMultiChannelRecordings API asynchronously
func (client *Client) ListMultiChannelRecordingsWithCallback(request *ListMultiChannelRecordingsRequest, callback func(response *ListMultiChannelRecordingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMultiChannelRecordingsResponse
		var err error
		defer close(result)
		response, err = client.ListMultiChannelRecordings(request)
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

// ListMultiChannelRecordingsRequest is the request struct for api ListMultiChannelRecordings
type ListMultiChannelRecordingsRequest struct {
	*requests.RpcRequest
	ContactId  string `position:"Query" name:"ContactId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListMultiChannelRecordingsResponse is the response struct for api ListMultiChannelRecordings
type ListMultiChannelRecordingsResponse struct {
	*responses.BaseResponse
	Code           string                                     `json:"Code" xml:"Code"`
	HttpStatusCode int                                        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                                     `json:"Message" xml:"Message"`
	RequestId      string                                     `json:"RequestId" xml:"RequestId"`
	Data           []RecordingDTOInListMultiChannelRecordings `json:"Data" xml:"Data"`
}

// CreateListMultiChannelRecordingsRequest creates a request to invoke ListMultiChannelRecordings API
func CreateListMultiChannelRecordingsRequest() (request *ListMultiChannelRecordingsRequest) {
	request = &ListMultiChannelRecordingsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListMultiChannelRecordings", "", "")
	request.Method = requests.POST
	return
}

// CreateListMultiChannelRecordingsResponse creates a response to parse from ListMultiChannelRecordings response
func CreateListMultiChannelRecordingsResponse() (response *ListMultiChannelRecordingsResponse) {
	response = &ListMultiChannelRecordingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
