package retailadvqa_public

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

// CreateExportTask invokes the retailadvqa_public.CreateExportTask API synchronously
func (client *Client) CreateExportTask(request *CreateExportTaskRequest) (response *CreateExportTaskResponse, err error) {
	response = CreateCreateExportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateExportTaskWithChan invokes the retailadvqa_public.CreateExportTask API asynchronously
func (client *Client) CreateExportTaskWithChan(request *CreateExportTaskRequest) (<-chan *CreateExportTaskResponse, <-chan error) {
	responseChan := make(chan *CreateExportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateExportTask(request)
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

// CreateExportTaskWithCallback invokes the retailadvqa_public.CreateExportTask API asynchronously
func (client *Client) CreateExportTaskWithCallback(request *CreateExportTaskRequest, callback func(response *CreateExportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateExportTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateExportTask(request)
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

// CreateExportTaskRequest is the request struct for api CreateExportTask
type CreateExportTaskRequest struct {
	*requests.RpcRequest
	AccessId   string `position:"Query" name:"AccessId"`
	TenantId   string `position:"Query" name:"TenantId"`
	AudienceId string `position:"Query" name:"AudienceId"`
}

// CreateExportTaskResponse is the response struct for api CreateExportTask
type CreateExportTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateExportTaskRequest creates a request to invoke CreateExportTask API
func CreateCreateExportTaskRequest() (request *CreateExportTaskRequest) {
	request = &CreateExportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "CreateExportTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateExportTaskResponse creates a response to parse from CreateExportTask response
func CreateCreateExportTaskResponse() (response *CreateExportTaskResponse) {
	response = &CreateExportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
