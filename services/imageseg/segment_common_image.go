package imageseg

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

// SegmentCommonImage invokes the imageseg.SegmentCommonImage API synchronously
func (client *Client) SegmentCommonImage(request *SegmentCommonImageRequest) (response *SegmentCommonImageResponse, err error) {
	response = CreateSegmentCommonImageResponse()
	err = client.DoAction(request, response)
	return
}

// SegmentCommonImageWithChan invokes the imageseg.SegmentCommonImage API asynchronously
func (client *Client) SegmentCommonImageWithChan(request *SegmentCommonImageRequest) (<-chan *SegmentCommonImageResponse, <-chan error) {
	responseChan := make(chan *SegmentCommonImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SegmentCommonImage(request)
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

// SegmentCommonImageWithCallback invokes the imageseg.SegmentCommonImage API asynchronously
func (client *Client) SegmentCommonImageWithCallback(request *SegmentCommonImageRequest, callback func(response *SegmentCommonImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SegmentCommonImageResponse
		var err error
		defer close(result)
		response, err = client.SegmentCommonImage(request)
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

// SegmentCommonImageRequest is the request struct for api SegmentCommonImage
type SegmentCommonImageRequest struct {
	*requests.RpcRequest
	ReturnForm     string `position:"Query" name:"ReturnForm"`
	OssFile        string `position:"Query" name:"OssFile"`
	RequestProxyBy string `position:"Query" name:"RequestProxyBy"`
	ImageURL       string `position:"Query" name:"ImageURL"`
}

// SegmentCommonImageResponse is the response struct for api SegmentCommonImage
type SegmentCommonImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSegmentCommonImageRequest creates a request to invoke SegmentCommonImage API
func CreateSegmentCommonImageRequest() (request *SegmentCommonImageRequest) {
	request = &SegmentCommonImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageseg", "2019-12-30", "SegmentCommonImage", "imageseg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSegmentCommonImageResponse creates a response to parse from SegmentCommonImage response
func CreateSegmentCommonImageResponse() (response *SegmentCommonImageResponse) {
	response = &SegmentCommonImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
