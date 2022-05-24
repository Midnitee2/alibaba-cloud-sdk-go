package rds

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

// DescribeModifyPGHbaConfigLog invokes the rds.DescribeModifyPGHbaConfigLog API synchronously
func (client *Client) DescribeModifyPGHbaConfigLog(request *DescribeModifyPGHbaConfigLogRequest) (response *DescribeModifyPGHbaConfigLogResponse, err error) {
	response = CreateDescribeModifyPGHbaConfigLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeModifyPGHbaConfigLogWithChan invokes the rds.DescribeModifyPGHbaConfigLog API asynchronously
func (client *Client) DescribeModifyPGHbaConfigLogWithChan(request *DescribeModifyPGHbaConfigLogRequest) (<-chan *DescribeModifyPGHbaConfigLogResponse, <-chan error) {
	responseChan := make(chan *DescribeModifyPGHbaConfigLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeModifyPGHbaConfigLog(request)
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

// DescribeModifyPGHbaConfigLogWithCallback invokes the rds.DescribeModifyPGHbaConfigLog API asynchronously
func (client *Client) DescribeModifyPGHbaConfigLogWithCallback(request *DescribeModifyPGHbaConfigLogRequest, callback func(response *DescribeModifyPGHbaConfigLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeModifyPGHbaConfigLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeModifyPGHbaConfigLog(request)
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

// DescribeModifyPGHbaConfigLogRequest is the request struct for api DescribeModifyPGHbaConfigLog
type DescribeModifyPGHbaConfigLogRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	StartTime            string           `position:"Query" name:"StartTime"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeModifyPGHbaConfigLogResponse is the response struct for api DescribeModifyPGHbaConfigLog
type DescribeModifyPGHbaConfigLogResponse struct {
	*responses.BaseResponse
	DBInstanceId string      `json:"DBInstanceId" xml:"DBInstanceId"`
	RequestId    string      `json:"RequestId" xml:"RequestId"`
	LogItemCount int         `json:"LogItemCount" xml:"LogItemCount"`
	HbaLogItems  HbaLogItems `json:"HbaLogItems" xml:"HbaLogItems"`
}

// CreateDescribeModifyPGHbaConfigLogRequest creates a request to invoke DescribeModifyPGHbaConfigLog API
func CreateDescribeModifyPGHbaConfigLogRequest() (request *DescribeModifyPGHbaConfigLogRequest) {
	request = &DescribeModifyPGHbaConfigLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeModifyPGHbaConfigLog", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeModifyPGHbaConfigLogResponse creates a response to parse from DescribeModifyPGHbaConfigLog response
func CreateDescribeModifyPGHbaConfigLogResponse() (response *DescribeModifyPGHbaConfigLogResponse) {
	response = &DescribeModifyPGHbaConfigLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
