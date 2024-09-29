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

// GetAlertMessage invokes the dataworks_public.GetAlertMessage API synchronously
func (client *Client) GetAlertMessage(request *GetAlertMessageRequest) (response *GetAlertMessageResponse, err error) {
	response = CreateGetAlertMessageResponse()
	err = client.DoAction(request, response)
	return
}

// GetAlertMessageWithChan invokes the dataworks_public.GetAlertMessage API asynchronously
func (client *Client) GetAlertMessageWithChan(request *GetAlertMessageRequest) (<-chan *GetAlertMessageResponse, <-chan error) {
	responseChan := make(chan *GetAlertMessageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAlertMessage(request)
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

// GetAlertMessageWithCallback invokes the dataworks_public.GetAlertMessage API asynchronously
func (client *Client) GetAlertMessageWithCallback(request *GetAlertMessageRequest, callback func(response *GetAlertMessageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAlertMessageResponse
		var err error
		defer close(result)
		response, err = client.GetAlertMessage(request)
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

// GetAlertMessageRequest is the request struct for api GetAlertMessage
type GetAlertMessageRequest struct {
	*requests.RpcRequest
	AlertId string `position:"Body" name:"AlertId"`
}

// GetAlertMessageResponse is the response struct for api GetAlertMessage
type GetAlertMessageResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetAlertMessageRequest creates a request to invoke GetAlertMessage API
func CreateGetAlertMessageRequest() (request *GetAlertMessageRequest) {
	request = &GetAlertMessageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetAlertMessage", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAlertMessageResponse creates a response to parse from GetAlertMessage response
func CreateGetAlertMessageResponse() (response *GetAlertMessageResponse) {
	response = &GetAlertMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
