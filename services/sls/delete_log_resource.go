package sls

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

// DeleteLogResource invokes the sls.DeleteLogResource API synchronously
func (client *Client) DeleteLogResource(request *DeleteLogResourceRequest) (response *DeleteLogResourceResponse, err error) {
	response = CreateDeleteLogResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLogResourceWithChan invokes the sls.DeleteLogResource API asynchronously
func (client *Client) DeleteLogResourceWithChan(request *DeleteLogResourceRequest) (<-chan *DeleteLogResourceResponse, <-chan error) {
	responseChan := make(chan *DeleteLogResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLogResource(request)
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

// DeleteLogResourceWithCallback invokes the sls.DeleteLogResource API asynchronously
func (client *Client) DeleteLogResourceWithCallback(request *DeleteLogResourceRequest, callback func(response *DeleteLogResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLogResourceResponse
		var err error
		defer close(result)
		response, err = client.DeleteLogResource(request)
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

// DeleteLogResourceRequest is the request struct for api DeleteLogResource
type DeleteLogResourceRequest struct {
	*requests.RpcRequest
	Project  string `position:"Query" name:"Project"`
	Region   string `position:"Query" name:"Region"`
	Logstore string `position:"Query" name:"Logstore"`
}

// DeleteLogResourceResponse is the response struct for api DeleteLogResource
type DeleteLogResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   string `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeleteLogResourceRequest creates a request to invoke DeleteLogResource API
func CreateDeleteLogResourceRequest() (request *DeleteLogResourceRequest) {
	request = &DeleteLogResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2018-06-13", "DeleteLogResource", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteLogResourceResponse creates a response to parse from DeleteLogResource response
func CreateDeleteLogResourceResponse() (response *DeleteLogResourceResponse) {
	response = &DeleteLogResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
