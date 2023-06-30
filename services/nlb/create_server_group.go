package nlb

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

// CreateServerGroup invokes the nlb.CreateServerGroup API synchronously
func (client *Client) CreateServerGroup(request *CreateServerGroupRequest) (response *CreateServerGroupResponse, err error) {
	response = CreateCreateServerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateServerGroupWithChan invokes the nlb.CreateServerGroup API asynchronously
func (client *Client) CreateServerGroupWithChan(request *CreateServerGroupRequest) (<-chan *CreateServerGroupResponse, <-chan error) {
	responseChan := make(chan *CreateServerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateServerGroup(request)
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

// CreateServerGroupWithCallback invokes the nlb.CreateServerGroup API asynchronously
func (client *Client) CreateServerGroupWithCallback(request *CreateServerGroupRequest, callback func(response *CreateServerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateServerGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateServerGroup(request)
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

// CreateServerGroupRequest is the request struct for api CreateServerGroup
type CreateServerGroupRequest struct {
	*requests.RpcRequest
	ServerGroupName         string                             `position:"Body" name:"ServerGroupName"`
	ClientToken             string                             `position:"Body" name:"ClientToken"`
	PreserveClientIpEnabled requests.Boolean                   `position:"Body" name:"PreserveClientIpEnabled"`
	HealthCheckConfig       CreateServerGroupHealthCheckConfig `position:"Body" name:"HealthCheckConfig"  type:"Struct"`
	AddressIPVersion        string                             `position:"Body" name:"AddressIPVersion"`
	Scheduler               string                             `position:"Body" name:"Scheduler"`
	ResourceGroupId         string                             `position:"Body" name:"ResourceGroupId"`
	Protocol                string                             `position:"Body" name:"Protocol"`
	PersistenceEnabled      requests.Boolean                   `position:"Body" name:"PersistenceEnabled"`
	PersistenceTimeout      requests.Integer                   `position:"Body" name:"PersistenceTimeout"`
	Tag                     *[]CreateServerGroupTag            `position:"Body" name:"Tag"  type:"Repeated"`
	DryRun                  requests.Boolean                   `position:"Body" name:"DryRun"`
	ConnectionDrainEnabled  requests.Boolean                   `position:"Body" name:"ConnectionDrainEnabled"`
	ConnectionDrainTimeout  requests.Integer                   `position:"Body" name:"ConnectionDrainTimeout"`
	AnyPortEnabled          requests.Boolean                   `position:"Body" name:"AnyPortEnabled"`
	ServerGroupType         string                             `position:"Body" name:"ServerGroupType"`
	VpcId                   string                             `position:"Body" name:"VpcId"`
}

// CreateServerGroupHealthCheckConfig is a repeated param struct in CreateServerGroupRequest
type CreateServerGroupHealthCheckConfig struct {
	HealthCheckEnabled        string    `name:"HealthCheckEnabled"`
	HealthCheckType           string    `name:"HealthCheckType"`
	HealthCheckConnectPort    string    `name:"HealthCheckConnectPort"`
	HealthyThreshold          string    `name:"HealthyThreshold"`
	UnhealthyThreshold        string    `name:"UnhealthyThreshold"`
	HealthCheckConnectTimeout string    `name:"HealthCheckConnectTimeout"`
	HealthCheckInterval       string    `name:"HealthCheckInterval"`
	HealthCheckDomain         string    `name:"HealthCheckDomain"`
	HealthCheckUrl            string    `name:"HealthCheckUrl"`
	HealthCheckHttpCode       *[]string `name:"HealthCheckHttpCode" type:"Repeated"`
	HttpCheckMethod           string    `name:"HttpCheckMethod"`
}

// CreateServerGroupTag is a repeated param struct in CreateServerGroupRequest
type CreateServerGroupTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateServerGroupResponse is the response struct for api CreateServerGroup
type CreateServerGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	ServerGroupId  string `json:"ServerGroupId" xml:"ServerGroupId"`
	JobId          string `json:"JobId" xml:"JobId"`
}

// CreateCreateServerGroupRequest creates a request to invoke CreateServerGroup API
func CreateCreateServerGroupRequest() (request *CreateServerGroupRequest) {
	request = &CreateServerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "CreateServerGroup", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateServerGroupResponse creates a response to parse from CreateServerGroup response
func CreateCreateServerGroupResponse() (response *CreateServerGroupResponse) {
	response = &CreateServerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
