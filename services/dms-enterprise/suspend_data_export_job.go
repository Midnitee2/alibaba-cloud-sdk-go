package dms_enterprise

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

// SuspendDataExportJob invokes the dms_enterprise.SuspendDataExportJob API synchronously
func (client *Client) SuspendDataExportJob(request *SuspendDataExportJobRequest) (response *SuspendDataExportJobResponse, err error) {
	response = CreateSuspendDataExportJobResponse()
	err = client.DoAction(request, response)
	return
}

// SuspendDataExportJobWithChan invokes the dms_enterprise.SuspendDataExportJob API asynchronously
func (client *Client) SuspendDataExportJobWithChan(request *SuspendDataExportJobRequest) (<-chan *SuspendDataExportJobResponse, <-chan error) {
	responseChan := make(chan *SuspendDataExportJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SuspendDataExportJob(request)
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

// SuspendDataExportJobWithCallback invokes the dms_enterprise.SuspendDataExportJob API asynchronously
func (client *Client) SuspendDataExportJobWithCallback(request *SuspendDataExportJobRequest, callback func(response *SuspendDataExportJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SuspendDataExportJobResponse
		var err error
		defer close(result)
		response, err = client.SuspendDataExportJob(request)
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

// SuspendDataExportJobRequest is the request struct for api SuspendDataExportJob
type SuspendDataExportJobRequest struct {
	*requests.RpcRequest
	Tid     requests.Integer `position:"Query" name:"Tid"`
	JobId   requests.Integer `position:"Query" name:"JobId"`
	OrderId requests.Integer `position:"Query" name:"OrderId"`
}

// SuspendDataExportJobResponse is the response struct for api SuspendDataExportJob
type SuspendDataExportJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateSuspendDataExportJobRequest creates a request to invoke SuspendDataExportJob API
func CreateSuspendDataExportJobRequest() (request *SuspendDataExportJobRequest) {
	request = &SuspendDataExportJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "SuspendDataExportJob", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSuspendDataExportJobResponse creates a response to parse from SuspendDataExportJob response
func CreateSuspendDataExportJobResponse() (response *SuspendDataExportJobResponse) {
	response = &SuspendDataExportJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
