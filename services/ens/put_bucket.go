package ens

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

// PutBucket invokes the ens.PutBucket API synchronously
func (client *Client) PutBucket(request *PutBucketRequest) (response *PutBucketResponse, err error) {
	response = CreatePutBucketResponse()
	err = client.DoAction(request, response)
	return
}

// PutBucketWithChan invokes the ens.PutBucket API asynchronously
func (client *Client) PutBucketWithChan(request *PutBucketRequest) (<-chan *PutBucketResponse, <-chan error) {
	responseChan := make(chan *PutBucketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutBucket(request)
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

// PutBucketWithCallback invokes the ens.PutBucket API asynchronously
func (client *Client) PutBucketWithCallback(request *PutBucketRequest, callback func(response *PutBucketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutBucketResponse
		var err error
		defer close(result)
		response, err = client.PutBucket(request)
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

// PutBucketRequest is the request struct for api PutBucket
type PutBucketRequest struct {
	*requests.RpcRequest
	EnsRegionId        string `position:"Body" name:"EnsRegionId"`
	LogicalBucketType  string `position:"Body" name:"LogicalBucketType"`
	City               string `position:"Body" name:"City"`
	DataRedundancyType string `position:"Body" name:"DataRedundancyType"`
	Endpoint           string `position:"Body" name:"Endpoint"`
	BucketName         string `position:"Body" name:"BucketName"`
	StorageDomainId    string `position:"Body" name:"StorageDomainId"`
	EngineId           string `position:"Body" name:"EngineId"`
	BucketAcl          string `position:"Body" name:"BucketAcl"`
	DispatcherType     string `position:"Body" name:"DispatcherType"`
	ResourceType       string `position:"Body" name:"ResourceType"`
	StorageClass       string `position:"Body" name:"StorageClass"`
	DispatchScope      string `position:"Body" name:"DispatchScope"`
	Comment            string `position:"Body" name:"Comment"`
}

// PutBucketResponse is the response struct for api PutBucket
type PutBucketResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePutBucketRequest creates a request to invoke PutBucket API
func CreatePutBucketRequest() (request *PutBucketRequest) {
	request = &PutBucketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "PutBucket", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutBucketResponse creates a response to parse from PutBucket response
func CreatePutBucketResponse() (response *PutBucketResponse) {
	response = &PutBucketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
