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

// DescribeDBClusterMigration invokes the polardb.DescribeDBClusterMigration API synchronously
func (client *Client) DescribeDBClusterMigration(request *DescribeDBClusterMigrationRequest) (response *DescribeDBClusterMigrationResponse, err error) {
	response = CreateDescribeDBClusterMigrationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterMigrationWithChan invokes the polardb.DescribeDBClusterMigration API asynchronously
func (client *Client) DescribeDBClusterMigrationWithChan(request *DescribeDBClusterMigrationRequest) (<-chan *DescribeDBClusterMigrationResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterMigrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterMigration(request)
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

// DescribeDBClusterMigrationWithCallback invokes the polardb.DescribeDBClusterMigration API asynchronously
func (client *Client) DescribeDBClusterMigrationWithCallback(request *DescribeDBClusterMigrationRequest, callback func(response *DescribeDBClusterMigrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterMigrationResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterMigration(request)
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

// DescribeDBClusterMigrationRequest is the request struct for api DescribeDBClusterMigration
type DescribeDBClusterMigrationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterMigrationResponse is the response struct for api DescribeDBClusterMigration
type DescribeDBClusterMigrationResponse struct {
	*responses.BaseResponse
	Comment                string              `json:"Comment" xml:"Comment"`
	RequestId              string              `json:"RequestId" xml:"RequestId"`
	ExpiredTime            string              `json:"ExpiredTime" xml:"ExpiredTime"`
	DBClusterId            string              `json:"DBClusterId" xml:"DBClusterId"`
	Topologies             string              `json:"Topologies" xml:"Topologies"`
	RdsReadWriteMode       string              `json:"RdsReadWriteMode" xml:"RdsReadWriteMode"`
	SourceRDSDBInstanceId  string              `json:"SourceRDSDBInstanceId" xml:"SourceRDSDBInstanceId"`
	DBClusterReadWriteMode string              `json:"DBClusterReadWriteMode" xml:"DBClusterReadWriteMode"`
	DelayedSeconds         int                 `json:"DelayedSeconds" xml:"DelayedSeconds"`
	MigrationStatus        string              `json:"MigrationStatus" xml:"MigrationStatus"`
	DtsInstanceId          string              `json:"DtsInstanceId" xml:"DtsInstanceId"`
	DBClusterEndpointList  []DBClusterEndpoint `json:"DBClusterEndpointList" xml:"DBClusterEndpointList"`
	RdsEndpointList        []RdsEndpoint       `json:"RdsEndpointList" xml:"RdsEndpointList"`
}

// CreateDescribeDBClusterMigrationRequest creates a request to invoke DescribeDBClusterMigration API
func CreateDescribeDBClusterMigrationRequest() (request *DescribeDBClusterMigrationRequest) {
	request = &DescribeDBClusterMigrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterMigration", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterMigrationResponse creates a response to parse from DescribeDBClusterMigration response
func CreateDescribeDBClusterMigrationResponse() (response *DescribeDBClusterMigrationResponse) {
	response = &DescribeDBClusterMigrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
