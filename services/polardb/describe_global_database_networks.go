package polardb

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

// DescribeGlobalDatabaseNetworks invokes the polardb.DescribeGlobalDatabaseNetworks API synchronously
func (client *Client) DescribeGlobalDatabaseNetworks(request *DescribeGlobalDatabaseNetworksRequest) (response *DescribeGlobalDatabaseNetworksResponse, err error) {
	response = CreateDescribeGlobalDatabaseNetworksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGlobalDatabaseNetworksWithChan invokes the polardb.DescribeGlobalDatabaseNetworks API asynchronously
func (client *Client) DescribeGlobalDatabaseNetworksWithChan(request *DescribeGlobalDatabaseNetworksRequest) (<-chan *DescribeGlobalDatabaseNetworksResponse, <-chan error) {
	responseChan := make(chan *DescribeGlobalDatabaseNetworksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGlobalDatabaseNetworks(request)
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

// DescribeGlobalDatabaseNetworksWithCallback invokes the polardb.DescribeGlobalDatabaseNetworks API asynchronously
func (client *Client) DescribeGlobalDatabaseNetworksWithCallback(request *DescribeGlobalDatabaseNetworksRequest, callback func(response *DescribeGlobalDatabaseNetworksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGlobalDatabaseNetworksResponse
		var err error
		defer close(result)
		response, err = client.DescribeGlobalDatabaseNetworks(request)
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

// DescribeGlobalDatabaseNetworksRequest is the request struct for api DescribeGlobalDatabaseNetworks
type DescribeGlobalDatabaseNetworksRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	GDNId                string           `position:"Query" name:"GDNId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	GDNDescription       string           `position:"Query" name:"GDNDescription"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeGlobalDatabaseNetworksResponse is the response struct for api DescribeGlobalDatabaseNetworks
type DescribeGlobalDatabaseNetworksResponse struct {
	*responses.BaseResponse
	TotalRecordCount int                     `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int                     `json:"PageRecordCount" xml:"PageRecordCount"`
	RequestId        string                  `json:"RequestId" xml:"RequestId"`
	PageNumber       int                     `json:"PageNumber" xml:"PageNumber"`
	Items            []GlobalDatabaseNetwork `json:"Items" xml:"Items"`
}

// CreateDescribeGlobalDatabaseNetworksRequest creates a request to invoke DescribeGlobalDatabaseNetworks API
func CreateDescribeGlobalDatabaseNetworksRequest() (request *DescribeGlobalDatabaseNetworksRequest) {
	request = &DescribeGlobalDatabaseNetworksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeGlobalDatabaseNetworks", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGlobalDatabaseNetworksResponse creates a response to parse from DescribeGlobalDatabaseNetworks response
func CreateDescribeGlobalDatabaseNetworksResponse() (response *DescribeGlobalDatabaseNetworksResponse) {
	response = &DescribeGlobalDatabaseNetworksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
