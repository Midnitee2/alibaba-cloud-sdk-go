package ens

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

// RescaleApplication invokes the ens.RescaleApplication API synchronously
func (client *Client) RescaleApplication(request *RescaleApplicationRequest) (response *RescaleApplicationResponse, err error) {
	response = CreateRescaleApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// RescaleApplicationWithChan invokes the ens.RescaleApplication API asynchronously
func (client *Client) RescaleApplicationWithChan(request *RescaleApplicationRequest) (<-chan *RescaleApplicationResponse, <-chan error) {
	responseChan := make(chan *RescaleApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RescaleApplication(request)
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

// RescaleApplicationWithCallback invokes the ens.RescaleApplication API asynchronously
func (client *Client) RescaleApplicationWithCallback(request *RescaleApplicationRequest, callback func(response *RescaleApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RescaleApplicationResponse
		var err error
		defer close(result)
		response, err = client.RescaleApplication(request)
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

// RescaleApplicationRequest is the request struct for api RescaleApplication
type RescaleApplicationRequest struct {
	*requests.RpcRequest
	RescaleType      string           `position:"Query" name:"RescaleType"`
	ResourceSelector string           `position:"Query" name:"ResourceSelector"`
	Timeout          requests.Integer `position:"Query" name:"Timeout"`
	RescaleLevel     string           `position:"Query" name:"RescaleLevel"`
	AppId            string           `position:"Query" name:"AppId"`
	ToAppVersion     string           `position:"Query" name:"ToAppVersion"`
}

// RescaleApplicationResponse is the response struct for api RescaleApplication
type RescaleApplicationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRescaleApplicationRequest creates a request to invoke RescaleApplication API
func CreateRescaleApplicationRequest() (request *RescaleApplicationRequest) {
	request = &RescaleApplicationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "RescaleApplication", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRescaleApplicationResponse creates a response to parse from RescaleApplication response
func CreateRescaleApplicationResponse() (response *RescaleApplicationResponse) {
	response = &RescaleApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
