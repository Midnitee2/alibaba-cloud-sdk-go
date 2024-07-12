package r_kvstore

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

// DescribeParameterGroup invokes the r_kvstore.DescribeParameterGroup API synchronously
func (client *Client) DescribeParameterGroup(request *DescribeParameterGroupRequest) (response *DescribeParameterGroupResponse, err error) {
	response = CreateDescribeParameterGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeParameterGroupWithChan invokes the r_kvstore.DescribeParameterGroup API asynchronously
func (client *Client) DescribeParameterGroupWithChan(request *DescribeParameterGroupRequest) (<-chan *DescribeParameterGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeParameterGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeParameterGroup(request)
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

// DescribeParameterGroupWithCallback invokes the r_kvstore.DescribeParameterGroup API asynchronously
func (client *Client) DescribeParameterGroupWithCallback(request *DescribeParameterGroupRequest, callback func(response *DescribeParameterGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeParameterGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeParameterGroup(request)
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

// DescribeParameterGroupRequest is the request struct for api DescribeParameterGroup
type DescribeParameterGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ParameterGroupId     string           `position:"Query" name:"ParameterGroupId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeParameterGroupResponse is the response struct for api DescribeParameterGroup
type DescribeParameterGroupResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	ParameterGroup ParameterGroup `json:"ParameterGroup" xml:"ParameterGroup"`
}

// CreateDescribeParameterGroupRequest creates a request to invoke DescribeParameterGroup API
func CreateDescribeParameterGroupRequest() (request *DescribeParameterGroupRequest) {
	request = &DescribeParameterGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeParameterGroup", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeParameterGroupResponse creates a response to parse from DescribeParameterGroup response
func CreateDescribeParameterGroupResponse() (response *DescribeParameterGroupResponse) {
	response = &DescribeParameterGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
