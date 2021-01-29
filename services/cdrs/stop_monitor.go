package cdrs

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

// StopMonitor invokes the cdrs.StopMonitor API synchronously
func (client *Client) StopMonitor(request *StopMonitorRequest) (response *StopMonitorResponse, err error) {
	response = CreateStopMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// StopMonitorWithChan invokes the cdrs.StopMonitor API asynchronously
func (client *Client) StopMonitorWithChan(request *StopMonitorRequest) (<-chan *StopMonitorResponse, <-chan error) {
	responseChan := make(chan *StopMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopMonitor(request)
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

// StopMonitorWithCallback invokes the cdrs.StopMonitor API asynchronously
func (client *Client) StopMonitorWithCallback(request *StopMonitorRequest, callback func(response *StopMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopMonitorResponse
		var err error
		defer close(result)
		response, err = client.StopMonitor(request)
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

// StopMonitorRequest is the request struct for api StopMonitor
type StopMonitorRequest struct {
	*requests.RpcRequest
	CorpId          string `position:"Body" name:"CorpId"`
	BizId           string `position:"Body" name:"BizId"`
	AlgorithmVendor string `position:"Body" name:"AlgorithmVendor"`
	TaskId          string `position:"Body" name:"TaskId"`
}

// StopMonitorResponse is the response struct for api StopMonitor
type StopMonitorResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopMonitorRequest creates a request to invoke StopMonitor API
func CreateStopMonitorRequest() (request *StopMonitorRequest) {
	request = &StopMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "StopMonitor", "", "")
	request.Method = requests.POST
	return
}

// CreateStopMonitorResponse creates a response to parse from StopMonitor response
func CreateStopMonitorResponse() (response *StopMonitorResponse) {
	response = &StopMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
