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

// QuerySensClassification invokes the dataworks_public.QuerySensClassification API synchronously
func (client *Client) QuerySensClassification(request *QuerySensClassificationRequest) (response *QuerySensClassificationResponse, err error) {
	response = CreateQuerySensClassificationResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySensClassificationWithChan invokes the dataworks_public.QuerySensClassification API asynchronously
func (client *Client) QuerySensClassificationWithChan(request *QuerySensClassificationRequest) (<-chan *QuerySensClassificationResponse, <-chan error) {
	responseChan := make(chan *QuerySensClassificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySensClassification(request)
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

// QuerySensClassificationWithCallback invokes the dataworks_public.QuerySensClassification API asynchronously
func (client *Client) QuerySensClassificationWithCallback(request *QuerySensClassificationRequest, callback func(response *QuerySensClassificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySensClassificationResponse
		var err error
		defer close(result)
		response, err = client.QuerySensClassification(request)
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

// QuerySensClassificationRequest is the request struct for api QuerySensClassification
type QuerySensClassificationRequest struct {
	*requests.RpcRequest
	TenantId   string `position:"Body" name:"TenantId"`
	TemplateId string `position:"Body" name:"TemplateId"`
}

// QuerySensClassificationResponse is the response struct for api QuerySensClassification
type QuerySensClassificationResponse struct {
	*responses.BaseResponse
}

// CreateQuerySensClassificationRequest creates a request to invoke QuerySensClassification API
func CreateQuerySensClassificationRequest() (request *QuerySensClassificationRequest) {
	request = &QuerySensClassificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "QuerySensClassification", "", "")
	request.Method = requests.POST
	return
}

// CreateQuerySensClassificationResponse creates a response to parse from QuerySensClassification response
func CreateQuerySensClassificationResponse() (response *QuerySensClassificationResponse) {
	response = &QuerySensClassificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
