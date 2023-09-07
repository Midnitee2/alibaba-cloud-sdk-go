package iot

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

// QueryTask invokes the iot.QueryTask API synchronously
func (client *Client) QueryTask(request *QueryTaskRequest) (response *QueryTaskResponse, err error) {
	response = CreateQueryTaskResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTaskWithChan invokes the iot.QueryTask API asynchronously
func (client *Client) QueryTaskWithChan(request *QueryTaskRequest) (<-chan *QueryTaskResponse, <-chan error) {
	responseChan := make(chan *QueryTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTask(request)
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

// QueryTaskWithCallback invokes the iot.QueryTask API asynchronously
func (client *Client) QueryTaskWithCallback(request *QueryTaskRequest, callback func(response *QueryTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTaskResponse
		var err error
		defer close(result)
		response, err = client.QueryTask(request)
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

// QueryTaskRequest is the request struct for api QueryTask
type QueryTaskRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	TaskId            string `position:"Query" name:"TaskId"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
}

// QueryTaskResponse is the response struct for api QueryTask
type QueryTaskResponse struct {
	*responses.BaseResponse
	RequestId    string          `json:"RequestId" xml:"RequestId"`
	Success      bool            `json:"Success" xml:"Success"`
	Code         string          `json:"Code" xml:"Code"`
	ErrorMessage string          `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryTask `json:"Data" xml:"Data"`
}

// CreateQueryTaskRequest creates a request to invoke QueryTask API
func CreateQueryTaskRequest() (request *QueryTaskRequest) {
	request = &QueryTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryTask", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryTaskResponse creates a response to parse from QueryTask response
func CreateQueryTaskResponse() (response *QueryTaskResponse) {
	response = &QueryTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
