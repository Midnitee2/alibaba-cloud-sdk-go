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

// ApproveOrder invokes the dms_enterprise.ApproveOrder API synchronously
func (client *Client) ApproveOrder(request *ApproveOrderRequest) (response *ApproveOrderResponse, err error) {
	response = CreateApproveOrderResponse()
	err = client.DoAction(request, response)
	return
}

// ApproveOrderWithChan invokes the dms_enterprise.ApproveOrder API asynchronously
func (client *Client) ApproveOrderWithChan(request *ApproveOrderRequest) (<-chan *ApproveOrderResponse, <-chan error) {
	responseChan := make(chan *ApproveOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApproveOrder(request)
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

// ApproveOrderWithCallback invokes the dms_enterprise.ApproveOrder API asynchronously
func (client *Client) ApproveOrderWithCallback(request *ApproveOrderRequest, callback func(response *ApproveOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApproveOrderResponse
		var err error
		defer close(result)
		response, err = client.ApproveOrder(request)
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

// ApproveOrderRequest is the request struct for api ApproveOrder
type ApproveOrderRequest struct {
	*requests.RpcRequest
	Tid                requests.Integer `position:"Query" name:"Tid"`
	WorkflowInstanceId requests.Integer `position:"Query" name:"WorkflowInstanceId"`
	ApprovalType       string           `position:"Query" name:"ApprovalType"`
	NewApprover        requests.Integer `position:"Query" name:"NewApprover"`
	ApprovalNodeId     requests.Integer `position:"Query" name:"ApprovalNodeId"`
	OldApprover        requests.Integer `position:"Query" name:"OldApprover"`
	RealLoginUserUid   string           `position:"Query" name:"RealLoginUserUid"`
	Comment            string           `position:"Query" name:"Comment"`
	ApprovalNodePos    string           `position:"Query" name:"ApprovalNodePos"`
	NewApproverList    string           `position:"Query" name:"NewApproverList"`
}

// ApproveOrderResponse is the response struct for api ApproveOrder
type ApproveOrderResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateApproveOrderRequest creates a request to invoke ApproveOrder API
func CreateApproveOrderRequest() (request *ApproveOrderRequest) {
	request = &ApproveOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ApproveOrder", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateApproveOrderResponse creates a response to parse from ApproveOrder response
func CreateApproveOrderResponse() (response *ApproveOrderResponse) {
	response = &ApproveOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
