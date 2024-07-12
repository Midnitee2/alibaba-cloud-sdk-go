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

// DescribeParameterGroupSupportParam invokes the r_kvstore.DescribeParameterGroupSupportParam API synchronously
func (client *Client) DescribeParameterGroupSupportParam(request *DescribeParameterGroupSupportParamRequest) (response *DescribeParameterGroupSupportParamResponse, err error) {
	response = CreateDescribeParameterGroupSupportParamResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeParameterGroupSupportParamWithChan invokes the r_kvstore.DescribeParameterGroupSupportParam API asynchronously
func (client *Client) DescribeParameterGroupSupportParamWithChan(request *DescribeParameterGroupSupportParamRequest) (<-chan *DescribeParameterGroupSupportParamResponse, <-chan error) {
	responseChan := make(chan *DescribeParameterGroupSupportParamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeParameterGroupSupportParam(request)
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

// DescribeParameterGroupSupportParamWithCallback invokes the r_kvstore.DescribeParameterGroupSupportParam API asynchronously
func (client *Client) DescribeParameterGroupSupportParamWithCallback(request *DescribeParameterGroupSupportParamRequest, callback func(response *DescribeParameterGroupSupportParamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeParameterGroupSupportParamResponse
		var err error
		defer close(result)
		response, err = client.DescribeParameterGroupSupportParam(request)
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

// DescribeParameterGroupSupportParamRequest is the request struct for api DescribeParameterGroupSupportParam
type DescribeParameterGroupSupportParamRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	EngineType           string           `position:"Query" name:"EngineType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Category             string           `position:"Query" name:"Category"`
}

// DescribeParameterGroupSupportParamResponse is the response struct for api DescribeParameterGroupSupportParam
type DescribeParameterGroupSupportParamResponse struct {
	*responses.BaseResponse
	RequestId    string             `json:"RequestId" xml:"RequestId"`
	ResourceList []ResourceListItem `json:"ResourceList" xml:"ResourceList"`
}

// CreateDescribeParameterGroupSupportParamRequest creates a request to invoke DescribeParameterGroupSupportParam API
func CreateDescribeParameterGroupSupportParamRequest() (request *DescribeParameterGroupSupportParamRequest) {
	request = &DescribeParameterGroupSupportParamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeParameterGroupSupportParam", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeParameterGroupSupportParamResponse creates a response to parse from DescribeParameterGroupSupportParam response
func CreateDescribeParameterGroupSupportParamResponse() (response *DescribeParameterGroupSupportParamResponse) {
	response = &DescribeParameterGroupSupportParamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
