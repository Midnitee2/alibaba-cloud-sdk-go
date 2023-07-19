package dts

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

// ModifyDtsJobDedicatedCluster invokes the dts.ModifyDtsJobDedicatedCluster API synchronously
func (client *Client) ModifyDtsJobDedicatedCluster(request *ModifyDtsJobDedicatedClusterRequest) (response *ModifyDtsJobDedicatedClusterResponse, err error) {
	response = CreateModifyDtsJobDedicatedClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDtsJobDedicatedClusterWithChan invokes the dts.ModifyDtsJobDedicatedCluster API asynchronously
func (client *Client) ModifyDtsJobDedicatedClusterWithChan(request *ModifyDtsJobDedicatedClusterRequest) (<-chan *ModifyDtsJobDedicatedClusterResponse, <-chan error) {
	responseChan := make(chan *ModifyDtsJobDedicatedClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDtsJobDedicatedCluster(request)
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

// ModifyDtsJobDedicatedClusterWithCallback invokes the dts.ModifyDtsJobDedicatedCluster API asynchronously
func (client *Client) ModifyDtsJobDedicatedClusterWithCallback(request *ModifyDtsJobDedicatedClusterRequest, callback func(response *ModifyDtsJobDedicatedClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDtsJobDedicatedClusterResponse
		var err error
		defer close(result)
		response, err = client.ModifyDtsJobDedicatedCluster(request)
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

// ModifyDtsJobDedicatedClusterRequest is the request struct for api ModifyDtsJobDedicatedCluster
type ModifyDtsJobDedicatedClusterRequest struct {
	*requests.RpcRequest
	DtsJobIds          string `position:"Query" name:"DtsJobIds"`
	DedicatedClusterId string `position:"Query" name:"DedicatedClusterId"`
	OwnerId            string `position:"Query" name:"OwnerId"`
}

// ModifyDtsJobDedicatedClusterResponse is the response struct for api ModifyDtsJobDedicatedCluster
type ModifyDtsJobDedicatedClusterResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
}

// CreateModifyDtsJobDedicatedClusterRequest creates a request to invoke ModifyDtsJobDedicatedCluster API
func CreateModifyDtsJobDedicatedClusterRequest() (request *ModifyDtsJobDedicatedClusterRequest) {
	request = &ModifyDtsJobDedicatedClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "ModifyDtsJobDedicatedCluster", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDtsJobDedicatedClusterResponse creates a response to parse from ModifyDtsJobDedicatedCluster response
func CreateModifyDtsJobDedicatedClusterResponse() (response *ModifyDtsJobDedicatedClusterResponse) {
	response = &ModifyDtsJobDedicatedClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
