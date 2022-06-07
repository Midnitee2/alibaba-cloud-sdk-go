package adb

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

// DownloadDiagnosisRecords invokes the adb.DownloadDiagnosisRecords API synchronously
func (client *Client) DownloadDiagnosisRecords(request *DownloadDiagnosisRecordsRequest) (response *DownloadDiagnosisRecordsResponse, err error) {
	response = CreateDownloadDiagnosisRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadDiagnosisRecordsWithChan invokes the adb.DownloadDiagnosisRecords API asynchronously
func (client *Client) DownloadDiagnosisRecordsWithChan(request *DownloadDiagnosisRecordsRequest) (<-chan *DownloadDiagnosisRecordsResponse, <-chan error) {
	responseChan := make(chan *DownloadDiagnosisRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadDiagnosisRecords(request)
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

// DownloadDiagnosisRecordsWithCallback invokes the adb.DownloadDiagnosisRecords API asynchronously
func (client *Client) DownloadDiagnosisRecordsWithCallback(request *DownloadDiagnosisRecordsRequest, callback func(response *DownloadDiagnosisRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadDiagnosisRecordsResponse
		var err error
		defer close(result)
		response, err = client.DownloadDiagnosisRecords(request)
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

// DownloadDiagnosisRecordsRequest is the request struct for api DownloadDiagnosisRecords
type DownloadDiagnosisRecordsRequest struct {
	*requests.RpcRequest
	MaxScanSize    requests.Integer `position:"Query" name:"MaxScanSize"`
	ResourceGroup  string           `position:"Query" name:"ResourceGroup"`
	DBClusterId    string           `position:"Query" name:"DBClusterId"`
	QueryCondition string           `position:"Query" name:"QueryCondition"`
	EndTime        string           `position:"Query" name:"EndTime"`
	StartTime      string           `position:"Query" name:"StartTime"`
	RawStartTime   string           `position:"Query" name:"RawStartTime"`
	MinPeakMemory  requests.Integer `position:"Query" name:"MinPeakMemory"`
	RawEndTime     string           `position:"Query" name:"RawEndTime"`
	MinScanSize    requests.Integer `position:"Query" name:"MinScanSize"`
	Database       string           `position:"Query" name:"Database"`
	ClientIp       string           `position:"Query" name:"ClientIp"`
	MaxPeakMemory  requests.Integer `position:"Query" name:"MaxPeakMemory"`
	Keyword        string           `position:"Query" name:"Keyword"`
	Lang           string           `position:"Query" name:"Lang"`
	UserName       string           `position:"Query" name:"UserName"`
}

// DownloadDiagnosisRecordsResponse is the response struct for api DownloadDiagnosisRecords
type DownloadDiagnosisRecordsResponse struct {
	*responses.BaseResponse
	DownloadId int    `json:"DownloadId" xml:"DownloadId"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateDownloadDiagnosisRecordsRequest creates a request to invoke DownloadDiagnosisRecords API
func CreateDownloadDiagnosisRecordsRequest() (request *DownloadDiagnosisRecordsRequest) {
	request = &DownloadDiagnosisRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "DownloadDiagnosisRecords", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDownloadDiagnosisRecordsResponse creates a response to parse from DownloadDiagnosisRecords response
func CreateDownloadDiagnosisRecordsResponse() (response *DownloadDiagnosisRecordsResponse) {
	response = &DownloadDiagnosisRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
