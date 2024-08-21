package ecs_workbench

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

// SetInstanceRecordConfig invokes the ecs_workbench.SetInstanceRecordConfig API synchronously
func (client *Client) SetInstanceRecordConfig(request *SetInstanceRecordConfigRequest) (response *SetInstanceRecordConfigResponse, err error) {
	response = CreateSetInstanceRecordConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetInstanceRecordConfigWithChan invokes the ecs_workbench.SetInstanceRecordConfig API asynchronously
func (client *Client) SetInstanceRecordConfigWithChan(request *SetInstanceRecordConfigRequest) (<-chan *SetInstanceRecordConfigResponse, <-chan error) {
	responseChan := make(chan *SetInstanceRecordConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetInstanceRecordConfig(request)
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

// SetInstanceRecordConfigWithCallback invokes the ecs_workbench.SetInstanceRecordConfig API asynchronously
func (client *Client) SetInstanceRecordConfigWithCallback(request *SetInstanceRecordConfigRequest, callback func(response *SetInstanceRecordConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetInstanceRecordConfigResponse
		var err error
		defer close(result)
		response, err = client.SetInstanceRecordConfig(request)
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

// SetInstanceRecordConfigRequest is the request struct for api SetInstanceRecordConfig
type SetInstanceRecordConfigRequest struct {
	*requests.RpcRequest
	Enabled             requests.Boolean `position:"Body" name:"Enabled"`
	RecordStorageTarget string           `position:"Body" name:"RecordStorageTarget"`
	InstanceId          string           `position:"Body" name:"InstanceId"`
	ExpirationDays      requests.Integer `position:"Body" name:"ExpirationDays"`
}

// SetInstanceRecordConfigResponse is the response struct for api SetInstanceRecordConfig
type SetInstanceRecordConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Root      bool   `json:"Root" xml:"Root"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSetInstanceRecordConfigRequest creates a request to invoke SetInstanceRecordConfig API
func CreateSetInstanceRecordConfigRequest() (request *SetInstanceRecordConfigRequest) {
	request = &SetInstanceRecordConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecs-workbench", "2022-02-20", "SetInstanceRecordConfig", "ecs-workbench", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetInstanceRecordConfigResponse creates a response to parse from SetInstanceRecordConfig response
func CreateSetInstanceRecordConfigResponse() (response *SetInstanceRecordConfigResponse) {
	response = &SetInstanceRecordConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
