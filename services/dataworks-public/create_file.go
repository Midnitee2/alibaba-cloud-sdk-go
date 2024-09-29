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

// CreateFile invokes the dataworks_public.CreateFile API synchronously
func (client *Client) CreateFile(request *CreateFileRequest) (response *CreateFileResponse, err error) {
	response = CreateCreateFileResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFileWithChan invokes the dataworks_public.CreateFile API asynchronously
func (client *Client) CreateFileWithChan(request *CreateFileRequest) (<-chan *CreateFileResponse, <-chan error) {
	responseChan := make(chan *CreateFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFile(request)
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

// CreateFileWithCallback invokes the dataworks_public.CreateFile API asynchronously
func (client *Client) CreateFileWithCallback(request *CreateFileRequest, callback func(response *CreateFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFileResponse
		var err error
		defer close(result)
		response, err = client.CreateFile(request)
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

// CreateFileRequest is the request struct for api CreateFile
type CreateFileRequest struct {
	*requests.RpcRequest
	FileType                        requests.Integer `position:"Body" name:"FileType"`
	DependentNodeIdList             string           `position:"Body" name:"DependentNodeIdList"`
	Content                         string           `position:"Body" name:"Content"`
	ProjectIdentifier               string           `position:"Body" name:"ProjectIdentifier"`
	ResourceGroupId                 requests.Integer `position:"Body" name:"ResourceGroupId"`
	StartImmediately                requests.Boolean `position:"Body" name:"StartImmediately"`
	ProjectId                       requests.Integer `position:"Body" name:"ProjectId"`
	AdvancedSettings                string           `position:"Body" name:"AdvancedSettings"`
	StartEffectDate                 requests.Integer `position:"Body" name:"StartEffectDate"`
	CycleType                       string           `position:"Body" name:"CycleType"`
	Owner                           string           `position:"Body" name:"Owner"`
	AutoRerunIntervalMillis         requests.Integer `position:"Body" name:"AutoRerunIntervalMillis"`
	InputList                       string           `position:"Body" name:"InputList"`
	CreateFolderIfNotExists         requests.Boolean `position:"Body" name:"CreateFolderIfNotExists"`
	ApplyScheduleImmediately        requests.Boolean `position:"Body" name:"ApplyScheduleImmediately"`
	RerunMode                       string           `position:"Body" name:"RerunMode"`
	ConnectionName                  string           `position:"Body" name:"ConnectionName"`
	OutputParameters                string           `position:"Body" name:"OutputParameters"`
	ParaValue                       string           `position:"Body" name:"ParaValue"`
	ResourceGroupIdentifier         string           `position:"Body" name:"ResourceGroupIdentifier"`
	AutoRerunTimes                  requests.Integer `position:"Body" name:"AutoRerunTimes"`
	CronExpress                     string           `position:"Body" name:"CronExpress"`
	IgnoreParentSkipRunningProperty requests.Boolean `position:"Body" name:"IgnoreParentSkipRunningProperty"`
	EndEffectDate                   requests.Integer `position:"Body" name:"EndEffectDate"`
	FileName                        string           `position:"Body" name:"FileName"`
	InputParameters                 string           `position:"Body" name:"InputParameters"`
	Stop                            requests.Boolean `position:"Body" name:"Stop"`
	DependentType                   string           `position:"Body" name:"DependentType"`
	FileFolderPath                  string           `position:"Body" name:"FileFolderPath"`
	FileDescription                 string           `position:"Body" name:"FileDescription"`
	AutoParsing                     requests.Boolean `position:"Body" name:"AutoParsing"`
	SchedulerType                   string           `position:"Body" name:"SchedulerType"`
}

// CreateFileResponse is the response struct for api CreateFile
type CreateFileResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           int64  `json:"Data" xml:"Data"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateCreateFileRequest creates a request to invoke CreateFile API
func CreateCreateFileRequest() (request *CreateFileRequest) {
	request = &CreateFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateFile", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateFileResponse creates a response to parse from CreateFile response
func CreateCreateFileResponse() (response *CreateFileResponse) {
	response = &CreateFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
