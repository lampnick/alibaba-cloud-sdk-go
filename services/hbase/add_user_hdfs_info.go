package hbase

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

// AddUserHdfsInfo invokes the hbase.AddUserHdfsInfo API synchronously
func (client *Client) AddUserHdfsInfo(request *AddUserHdfsInfoRequest) (response *AddUserHdfsInfoResponse, err error) {
	response = CreateAddUserHdfsInfoResponse()
	err = client.DoAction(request, response)
	return
}

// AddUserHdfsInfoWithChan invokes the hbase.AddUserHdfsInfo API asynchronously
func (client *Client) AddUserHdfsInfoWithChan(request *AddUserHdfsInfoRequest) (<-chan *AddUserHdfsInfoResponse, <-chan error) {
	responseChan := make(chan *AddUserHdfsInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddUserHdfsInfo(request)
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

// AddUserHdfsInfoWithCallback invokes the hbase.AddUserHdfsInfo API asynchronously
func (client *Client) AddUserHdfsInfoWithCallback(request *AddUserHdfsInfoRequest, callback func(response *AddUserHdfsInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddUserHdfsInfoResponse
		var err error
		defer close(result)
		response, err = client.AddUserHdfsInfo(request)
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

// AddUserHdfsInfoRequest is the request struct for api AddUserHdfsInfo
type AddUserHdfsInfoRequest struct {
	*requests.RpcRequest
	ExtInfo   string `position:"Query" name:"ExtInfo"`
	ClusterId string `position:"Query" name:"ClusterId"`
}

// AddUserHdfsInfoResponse is the response struct for api AddUserHdfsInfo
type AddUserHdfsInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddUserHdfsInfoRequest creates a request to invoke AddUserHdfsInfo API
func CreateAddUserHdfsInfoRequest() (request *AddUserHdfsInfoRequest) {
	request = &AddUserHdfsInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "AddUserHdfsInfo", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddUserHdfsInfoResponse creates a response to parse from AddUserHdfsInfo response
func CreateAddUserHdfsInfoResponse() (response *AddUserHdfsInfoResponse) {
	response = &AddUserHdfsInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
