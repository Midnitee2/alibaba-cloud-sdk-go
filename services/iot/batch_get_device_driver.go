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

// BatchGetDeviceDriver invokes the iot.BatchGetDeviceDriver API synchronously
// api document: https://help.aliyun.com/api/iot/batchgetdevicedriver.html
func (client *Client) BatchGetDeviceDriver(request *BatchGetDeviceDriverRequest) (response *BatchGetDeviceDriverResponse, err error) {
	response = CreateBatchGetDeviceDriverResponse()
	err = client.DoAction(request, response)
	return
}

// BatchGetDeviceDriverWithChan invokes the iot.BatchGetDeviceDriver API asynchronously
// api document: https://help.aliyun.com/api/iot/batchgetdevicedriver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetDeviceDriverWithChan(request *BatchGetDeviceDriverRequest) (<-chan *BatchGetDeviceDriverResponse, <-chan error) {
	responseChan := make(chan *BatchGetDeviceDriverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchGetDeviceDriver(request)
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

// BatchGetDeviceDriverWithCallback invokes the iot.BatchGetDeviceDriver API asynchronously
// api document: https://help.aliyun.com/api/iot/batchgetdevicedriver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetDeviceDriverWithCallback(request *BatchGetDeviceDriverRequest, callback func(response *BatchGetDeviceDriverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchGetDeviceDriverResponse
		var err error
		defer close(result)
		response, err = client.BatchGetDeviceDriver(request)
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

// BatchGetDeviceDriverRequest is the request struct for api BatchGetDeviceDriver
type BatchGetDeviceDriverRequest struct {
	*requests.RpcRequest
	InstanceId    string    `position:"Query" name:"InstanceId"`
	IotIds        *[]string `position:"Query" name:"IotIds"  type:"Repeated"`
	IotInstanceId string    `position:"Query" name:"IotInstanceId"`
}

// BatchGetDeviceDriverResponse is the response struct for api BatchGetDeviceDriver
type BatchGetDeviceDriverResponse struct {
	*responses.BaseResponse
	RequestId        string         `json:"RequestId" xml:"RequestId"`
	Success          bool           `json:"Success" xml:"Success"`
	Code             string         `json:"Code" xml:"Code"`
	ErrorMessage     string         `json:"ErrorMessage" xml:"ErrorMessage"`
	DeviceDriverList []DeviceDriver `json:"DeviceDriverList" xml:"DeviceDriverList"`
}

// CreateBatchGetDeviceDriverRequest creates a request to invoke BatchGetDeviceDriver API
func CreateBatchGetDeviceDriverRequest() (request *BatchGetDeviceDriverRequest) {
	request = &BatchGetDeviceDriverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchGetDeviceDriver", "Iot", "openAPI")
	return
}

// CreateBatchGetDeviceDriverResponse creates a response to parse from BatchGetDeviceDriver response
func CreateBatchGetDeviceDriverResponse() (response *BatchGetDeviceDriverResponse) {
	response = &BatchGetDeviceDriverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
