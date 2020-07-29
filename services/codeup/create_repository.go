package codeup

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

// CreateRepository invokes the codeup.CreateRepository API synchronously
// api document: https://help.aliyun.com/api/codeup/createrepository.html
func (client *Client) CreateRepository(request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
	response = CreateCreateRepositoryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRepositoryWithChan invokes the codeup.CreateRepository API asynchronously
// api document: https://help.aliyun.com/api/codeup/createrepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRepositoryWithChan(request *CreateRepositoryRequest) (<-chan *CreateRepositoryResponse, <-chan error) {
	responseChan := make(chan *CreateRepositoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRepository(request)
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

// CreateRepositoryWithCallback invokes the codeup.CreateRepository API asynchronously
// api document: https://help.aliyun.com/api/codeup/createrepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRepositoryWithCallback(request *CreateRepositoryRequest, callback func(response *CreateRepositoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRepositoryResponse
		var err error
		defer close(result)
		response, err = client.CreateRepository(request)
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

// CreateRepositoryRequest is the request struct for api CreateRepository
type CreateRepositoryRequest struct {
	*requests.RoaRequest
	OrganizationId   string           `position:"Query" name:"OrganizationId"`
	SubUserId        string           `position:"Query" name:"SubUserId"`
	ClientToken      string           `position:"Query" name:"ClientToken"`
	AccessToken      string           `position:"Query" name:"AccessToken"`
	Sync             requests.Boolean `position:"Query" name:"Sync"`
	CreateParentPath requests.Boolean `position:"Query" name:"CreateParentPath"`
}

// CreateRepositoryResponse is the response struct for api CreateRepository
type CreateRepositoryResponse struct {
	*responses.BaseResponse
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateCreateRepositoryRequest creates a request to invoke CreateRepository API
func CreateCreateRepositoryRequest() (request *CreateRepositoryRequest) {
	request = &CreateRepositoryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "CreateRepository", "/api/v3/projects", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateRepositoryResponse creates a response to parse from CreateRepository response
func CreateCreateRepositoryResponse() (response *CreateRepositoryResponse) {
	response = &CreateRepositoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}