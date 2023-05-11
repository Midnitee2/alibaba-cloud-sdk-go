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

// BatchCheckVehicleDevice invokes the iot.BatchCheckVehicleDevice API synchronously
func (client *Client) BatchCheckVehicleDevice(request *BatchCheckVehicleDeviceRequest) (response *BatchCheckVehicleDeviceResponse, err error) {
	response = CreateBatchCheckVehicleDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// BatchCheckVehicleDeviceWithChan invokes the iot.BatchCheckVehicleDevice API asynchronously
func (client *Client) BatchCheckVehicleDeviceWithChan(request *BatchCheckVehicleDeviceRequest) (<-chan *BatchCheckVehicleDeviceResponse, <-chan error) {
	responseChan := make(chan *BatchCheckVehicleDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchCheckVehicleDevice(request)
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

// BatchCheckVehicleDeviceWithCallback invokes the iot.BatchCheckVehicleDevice API asynchronously
func (client *Client) BatchCheckVehicleDeviceWithCallback(request *BatchCheckVehicleDeviceRequest, callback func(response *BatchCheckVehicleDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchCheckVehicleDeviceResponse
		var err error
		defer close(result)
		response, err = client.BatchCheckVehicleDevice(request)
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

// BatchCheckVehicleDeviceRequest is the request struct for api BatchCheckVehicleDevice
type BatchCheckVehicleDeviceRequest struct {
	*requests.RpcRequest
	IotInstanceId string                               `position:"Query" name:"IotInstanceId"`
	ProductKey    string                               `position:"Query" name:"ProductKey"`
	DeviceList    *[]BatchCheckVehicleDeviceDeviceList `position:"Body" name:"DeviceList"  type:"Repeated"`
	ApiProduct    string                               `position:"Body" name:"ApiProduct"`
	ApiRevision   string                               `position:"Body" name:"ApiRevision"`
}

// BatchCheckVehicleDeviceDeviceList is a repeated param struct in BatchCheckVehicleDeviceRequest
type BatchCheckVehicleDeviceDeviceList struct {
	DeviceName   string `name:"DeviceName"`
	DeviceId     string `name:"DeviceId"`
	Manufacturer string `name:"Manufacturer"`
	DeviceModel  string `name:"DeviceModel"`
}

// BatchCheckVehicleDeviceResponse is the response struct for api BatchCheckVehicleDevice
type BatchCheckVehicleDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	Success      bool                          `json:"Success" xml:"Success"`
	Code         string                        `json:"Code" xml:"Code"`
	ErrorMessage string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInBatchCheckVehicleDevice `json:"Data" xml:"Data"`
}

// CreateBatchCheckVehicleDeviceRequest creates a request to invoke BatchCheckVehicleDevice API
func CreateBatchCheckVehicleDeviceRequest() (request *BatchCheckVehicleDeviceRequest) {
	request = &BatchCheckVehicleDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchCheckVehicleDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchCheckVehicleDeviceResponse creates a response to parse from BatchCheckVehicleDevice response
func CreateBatchCheckVehicleDeviceResponse() (response *BatchCheckVehicleDeviceResponse) {
	response = &BatchCheckVehicleDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
