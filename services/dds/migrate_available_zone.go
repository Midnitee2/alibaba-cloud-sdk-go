package dds

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

// MigrateAvailableZone invokes the dds.MigrateAvailableZone API synchronously
func (client *Client) MigrateAvailableZone(request *MigrateAvailableZoneRequest) (response *MigrateAvailableZoneResponse, err error) {
	response = CreateMigrateAvailableZoneResponse()
	err = client.DoAction(request, response)
	return
}

// MigrateAvailableZoneWithChan invokes the dds.MigrateAvailableZone API asynchronously
func (client *Client) MigrateAvailableZoneWithChan(request *MigrateAvailableZoneRequest) (<-chan *MigrateAvailableZoneResponse, <-chan error) {
	responseChan := make(chan *MigrateAvailableZoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MigrateAvailableZone(request)
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

// MigrateAvailableZoneWithCallback invokes the dds.MigrateAvailableZone API asynchronously
func (client *Client) MigrateAvailableZoneWithCallback(request *MigrateAvailableZoneRequest, callback func(response *MigrateAvailableZoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MigrateAvailableZoneResponse
		var err error
		defer close(result)
		response, err = client.MigrateAvailableZone(request)
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

// MigrateAvailableZoneRequest is the request struct for api MigrateAvailableZone
type MigrateAvailableZoneRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecondaryZoneId      string           `position:"Query" name:"SecondaryZoneId"`
	EffectiveTime        string           `position:"Query" name:"EffectiveTime"`
	HiddenZoneId         string           `position:"Query" name:"HiddenZoneId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Vswitch              string           `position:"Query" name:"Vswitch"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// MigrateAvailableZoneResponse is the response struct for api MigrateAvailableZone
type MigrateAvailableZoneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMigrateAvailableZoneRequest creates a request to invoke MigrateAvailableZone API
func CreateMigrateAvailableZoneRequest() (request *MigrateAvailableZoneRequest) {
	request = &MigrateAvailableZoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "MigrateAvailableZone", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMigrateAvailableZoneResponse creates a response to parse from MigrateAvailableZone response
func CreateMigrateAvailableZoneResponse() (response *MigrateAvailableZoneResponse) {
	response = &MigrateAvailableZoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
