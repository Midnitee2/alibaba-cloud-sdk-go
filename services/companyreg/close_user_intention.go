package companyreg

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

// CloseUserIntention invokes the companyreg.CloseUserIntention API synchronously
func (client *Client) CloseUserIntention(request *CloseUserIntentionRequest) (response *CloseUserIntentionResponse, err error) {
	response = CreateCloseUserIntentionResponse()
	err = client.DoAction(request, response)
	return
}

// CloseUserIntentionWithChan invokes the companyreg.CloseUserIntention API asynchronously
func (client *Client) CloseUserIntentionWithChan(request *CloseUserIntentionRequest) (<-chan *CloseUserIntentionResponse, <-chan error) {
	responseChan := make(chan *CloseUserIntentionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CloseUserIntention(request)
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

// CloseUserIntentionWithCallback invokes the companyreg.CloseUserIntention API asynchronously
func (client *Client) CloseUserIntentionWithCallback(request *CloseUserIntentionRequest, callback func(response *CloseUserIntentionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CloseUserIntentionResponse
		var err error
		defer close(result)
		response, err = client.CloseUserIntention(request)
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

// CloseUserIntentionRequest is the request struct for api CloseUserIntention
type CloseUserIntentionRequest struct {
	*requests.RpcRequest
	BizType        string `position:"Query" name:"BizType"`
	Note           string `position:"Query" name:"Note"`
	IntentionBizId string `position:"Query" name:"IntentionBizId"`
}

// CloseUserIntentionResponse is the response struct for api CloseUserIntention
type CloseUserIntentionResponse struct {
	*responses.BaseResponse
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCloseUserIntentionRequest creates a request to invoke CloseUserIntention API
func CreateCloseUserIntentionRequest() (request *CloseUserIntentionRequest) {
	request = &CloseUserIntentionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "CloseUserIntention", "", "")
	request.Method = requests.POST
	return
}

// CreateCloseUserIntentionResponse creates a response to parse from CloseUserIntention response
func CreateCloseUserIntentionResponse() (response *CloseUserIntentionResponse) {
	response = &CloseUserIntentionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
