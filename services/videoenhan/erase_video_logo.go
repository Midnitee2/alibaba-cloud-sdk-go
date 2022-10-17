package videoenhan

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

// EraseVideoLogo invokes the videoenhan.EraseVideoLogo API synchronously
func (client *Client) EraseVideoLogo(request *EraseVideoLogoRequest) (response *EraseVideoLogoResponse, err error) {
	response = CreateEraseVideoLogoResponse()
	err = client.DoAction(request, response)
	return
}

// EraseVideoLogoWithChan invokes the videoenhan.EraseVideoLogo API asynchronously
func (client *Client) EraseVideoLogoWithChan(request *EraseVideoLogoRequest) (<-chan *EraseVideoLogoResponse, <-chan error) {
	responseChan := make(chan *EraseVideoLogoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EraseVideoLogo(request)
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

// EraseVideoLogoWithCallback invokes the videoenhan.EraseVideoLogo API asynchronously
func (client *Client) EraseVideoLogoWithCallback(request *EraseVideoLogoRequest, callback func(response *EraseVideoLogoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EraseVideoLogoResponse
		var err error
		defer close(result)
		response, err = client.EraseVideoLogo(request)
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

// EraseVideoLogoRequest is the request struct for api EraseVideoLogo
type EraseVideoLogoRequest struct {
	*requests.RpcRequest
	Boxes    *[]EraseVideoLogoBoxes `position:"Body" name:"Boxes"  type:"Repeated"`
	Async    requests.Boolean       `position:"Body" name:"Async"`
	VideoUrl string                 `position:"Body" name:"VideoUrl"`
}

// EraseVideoLogoBoxes is a repeated param struct in EraseVideoLogoRequest
type EraseVideoLogoBoxes struct {
	W string `name:"W"`
	H string `name:"H"`
	X string `name:"X"`
	Y string `name:"Y"`
}

// EraseVideoLogoResponse is the response struct for api EraseVideoLogo
type EraseVideoLogoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEraseVideoLogoRequest creates a request to invoke EraseVideoLogo API
func CreateEraseVideoLogoRequest() (request *EraseVideoLogoRequest) {
	request = &EraseVideoLogoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videoenhan", "2020-03-20", "EraseVideoLogo", "videoenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEraseVideoLogoResponse creates a response to parse from EraseVideoLogo response
func CreateEraseVideoLogoResponse() (response *EraseVideoLogoResponse) {
	response = &EraseVideoLogoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
