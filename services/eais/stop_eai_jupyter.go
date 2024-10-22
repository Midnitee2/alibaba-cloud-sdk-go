package eais

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

// StopEaiJupyter invokes the eais.StopEaiJupyter API synchronously
func (client *Client) StopEaiJupyter(request *StopEaiJupyterRequest) (response *StopEaiJupyterResponse, err error) {
	response = CreateStopEaiJupyterResponse()
	err = client.DoAction(request, response)
	return
}

// StopEaiJupyterWithChan invokes the eais.StopEaiJupyter API asynchronously
func (client *Client) StopEaiJupyterWithChan(request *StopEaiJupyterRequest) (<-chan *StopEaiJupyterResponse, <-chan error) {
	responseChan := make(chan *StopEaiJupyterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopEaiJupyter(request)
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

// StopEaiJupyterWithCallback invokes the eais.StopEaiJupyter API asynchronously
func (client *Client) StopEaiJupyterWithCallback(request *StopEaiJupyterRequest, callback func(response *StopEaiJupyterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopEaiJupyterResponse
		var err error
		defer close(result)
		response, err = client.StopEaiJupyter(request)
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

// StopEaiJupyterRequest is the request struct for api StopEaiJupyter
type StopEaiJupyterRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// StopEaiJupyterResponse is the response struct for api StopEaiJupyter
type StopEaiJupyterResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	AccessDeniedDetail AccessDeniedDetail `json:"AccessDeniedDetail" xml:"AccessDeniedDetail"`
}

// CreateStopEaiJupyterRequest creates a request to invoke StopEaiJupyter API
func CreateStopEaiJupyterRequest() (request *StopEaiJupyterRequest) {
	request = &StopEaiJupyterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eais", "2019-06-24", "StopEaiJupyter", "eais", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopEaiJupyterResponse creates a response to parse from StopEaiJupyter response
func CreateStopEaiJupyterResponse() (response *StopEaiJupyterResponse) {
	response = &StopEaiJupyterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
