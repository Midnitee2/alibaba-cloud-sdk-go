package outboundbot

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

// CreateJobGroup invokes the outboundbot.CreateJobGroup API synchronously
func (client *Client) CreateJobGroup(request *CreateJobGroupRequest) (response *CreateJobGroupResponse, err error) {
	response = CreateCreateJobGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateJobGroupWithChan invokes the outboundbot.CreateJobGroup API asynchronously
func (client *Client) CreateJobGroupWithChan(request *CreateJobGroupRequest) (<-chan *CreateJobGroupResponse, <-chan error) {
	responseChan := make(chan *CreateJobGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateJobGroup(request)
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

// CreateJobGroupWithCallback invokes the outboundbot.CreateJobGroup API asynchronously
func (client *Client) CreateJobGroupWithCallback(request *CreateJobGroupRequest, callback func(response *CreateJobGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateJobGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateJobGroup(request)
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

// CreateJobGroupRequest is the request struct for api CreateJobGroup
type CreateJobGroupRequest struct {
	*requests.RpcRequest
	RecallStrategyJson  string           `position:"Query" name:"RecallStrategyJson"`
	JobGroupName        string           `position:"Query" name:"JobGroupName"`
	ScriptId            string           `position:"Query" name:"ScriptId"`
	StrategyJson        string           `position:"Query" name:"StrategyJson"`
	RingingDuration     requests.Integer `position:"Query" name:"RingingDuration"`
	ScenarioId          string           `position:"Query" name:"ScenarioId"`
	Priority            string           `position:"Query" name:"Priority"`
	JobGroupDescription string           `position:"Query" name:"JobGroupDescription"`
	CallingNumber       *[]string        `position:"Query" name:"CallingNumber"  type:"Repeated"`
	InstanceId          string           `position:"Query" name:"InstanceId"`
	MinConcurrency      requests.Integer `position:"Query" name:"MinConcurrency"`
}

// CreateJobGroupResponse is the response struct for api CreateJobGroup
type CreateJobGroupResponse struct {
	*responses.BaseResponse
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	JobGroup       JobGroup `json:"JobGroup" xml:"JobGroup"`
}

// CreateCreateJobGroupRequest creates a request to invoke CreateJobGroup API
func CreateCreateJobGroupRequest() (request *CreateJobGroupRequest) {
	request = &CreateJobGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "CreateJobGroup", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateJobGroupResponse creates a response to parse from CreateJobGroup response
func CreateCreateJobGroupResponse() (response *CreateJobGroupResponse) {
	response = &CreateJobGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
