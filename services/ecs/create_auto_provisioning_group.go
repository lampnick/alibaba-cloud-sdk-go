package ecs

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

// CreateAutoProvisioningGroup invokes the ecs.CreateAutoProvisioningGroup API synchronously
func (client *Client) CreateAutoProvisioningGroup(request *CreateAutoProvisioningGroupRequest) (response *CreateAutoProvisioningGroupResponse, err error) {
	response = CreateCreateAutoProvisioningGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAutoProvisioningGroupWithChan invokes the ecs.CreateAutoProvisioningGroup API asynchronously
func (client *Client) CreateAutoProvisioningGroupWithChan(request *CreateAutoProvisioningGroupRequest) (<-chan *CreateAutoProvisioningGroupResponse, <-chan error) {
	responseChan := make(chan *CreateAutoProvisioningGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAutoProvisioningGroup(request)
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

// CreateAutoProvisioningGroupWithCallback invokes the ecs.CreateAutoProvisioningGroup API asynchronously
func (client *Client) CreateAutoProvisioningGroupWithCallback(request *CreateAutoProvisioningGroupRequest, callback func(response *CreateAutoProvisioningGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAutoProvisioningGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateAutoProvisioningGroup(request)
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

// CreateAutoProvisioningGroupRequest is the request struct for api CreateAutoProvisioningGroup
type CreateAutoProvisioningGroupRequest struct {
	*requests.RpcRequest
	LaunchConfigurationDataDisk                    *[]CreateAutoProvisioningGroupLaunchConfigurationDataDisk `position:"Query" name:"LaunchConfiguration.DataDisk"  type:"Repeated"`
	ResourceOwnerId                                requests.Integer                                          `position:"Query" name:"ResourceOwnerId"`
	LaunchConfigurationSystemDiskCategory          string                                                    `position:"Query" name:"LaunchConfiguration.SystemDiskCategory"`
	AutoProvisioningGroupType                      string                                                    `position:"Query" name:"AutoProvisioningGroupType"`
	ResourceGroupId                                string                                                    `position:"Query" name:"ResourceGroupId"`
	LaunchConfigurationImageId                     string                                                    `position:"Query" name:"LaunchConfiguration.ImageId"`
	LaunchConfigurationResourceGroupId             string                                                    `position:"Query" name:"LaunchConfiguration.ResourceGroupId"`
	PayAsYouGoAllocationStrategy                   string                                                    `position:"Query" name:"PayAsYouGoAllocationStrategy"`
	DefaultTargetCapacityType                      string                                                    `position:"Query" name:"DefaultTargetCapacityType"`
	LaunchConfigurationKeyPairName                 string                                                    `position:"Query" name:"LaunchConfiguration.KeyPairName"`
	SystemDiskConfig                               *[]CreateAutoProvisioningGroupSystemDiskConfig            `position:"Query" name:"SystemDiskConfig"  type:"Repeated"`
	DataDiskConfig                                 *[]CreateAutoProvisioningGroupDataDiskConfig              `position:"Query" name:"DataDiskConfig"  type:"Repeated"`
	ValidUntil                                     string                                                    `position:"Query" name:"ValidUntil"`
	LaunchTemplateId                               string                                                    `position:"Query" name:"LaunchTemplateId"`
	OwnerId                                        requests.Integer                                          `position:"Query" name:"OwnerId"`
	LaunchConfigurationSystemDiskSize              requests.Integer                                          `position:"Query" name:"LaunchConfiguration.SystemDiskSize"`
	LaunchConfigurationInternetMaxBandwidthOut     requests.Integer                                          `position:"Query" name:"LaunchConfiguration.InternetMaxBandwidthOut"`
	LaunchConfigurationHostName                    string                                                    `position:"Query" name:"LaunchConfiguration.HostName"`
	MaxSpotPrice                                   requests.Float                                            `position:"Query" name:"MaxSpotPrice"`
	LaunchConfigurationPasswordInherit             requests.Boolean                                          `position:"Query" name:"LaunchConfiguration.PasswordInherit"`
	LaunchConfigurationSecurityGroupId             string                                                    `position:"Query" name:"LaunchConfiguration.SecurityGroupId"`
	Description                                    string                                                    `position:"Query" name:"Description"`
	TerminateInstancesWithExpiration               requests.Boolean                                          `position:"Query" name:"TerminateInstancesWithExpiration"`
	LaunchConfigurationUserData                    string                                                    `position:"Query" name:"LaunchConfiguration.UserData"`
	LaunchConfigurationInstanceName                string                                                    `position:"Query" name:"LaunchConfiguration.InstanceName"`
	LaunchConfigurationInstanceDescription         string                                                    `position:"Query" name:"LaunchConfiguration.InstanceDescription"`
	SpotAllocationStrategy                         string                                                    `position:"Query" name:"SpotAllocationStrategy"`
	TerminateInstances                             requests.Boolean                                          `position:"Query" name:"TerminateInstances"`
	LaunchConfigurationSystemDiskName              string                                                    `position:"Query" name:"LaunchConfiguration.SystemDiskName"`
	LaunchConfigurationSystemDiskDescription       string                                                    `position:"Query" name:"LaunchConfiguration.SystemDiskDescription"`
	ExcessCapacityTerminationPolicy                string                                                    `position:"Query" name:"ExcessCapacityTerminationPolicy"`
	LaunchTemplateConfig                           *[]CreateAutoProvisioningGroupLaunchTemplateConfig        `position:"Query" name:"LaunchTemplateConfig"  type:"Repeated"`
	LaunchConfigurationRamRoleName                 string                                                    `position:"Query" name:"LaunchConfiguration.RamRoleName"`
	LaunchConfigurationInternetMaxBandwidthIn      requests.Integer                                          `position:"Query" name:"LaunchConfiguration.InternetMaxBandwidthIn"`
	SpotInstanceInterruptionBehavior               string                                                    `position:"Query" name:"SpotInstanceInterruptionBehavior"`
	LaunchConfigurationSecurityEnhancementStrategy string                                                    `position:"Query" name:"LaunchConfiguration.SecurityEnhancementStrategy"`
	LaunchConfigurationTag                         *[]CreateAutoProvisioningGroupLaunchConfigurationTag      `position:"Query" name:"LaunchConfiguration.Tag"  type:"Repeated"`
	ResourceOwnerAccount                           string                                                    `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                                   string                                                    `position:"Query" name:"OwnerAccount"`
	SpotInstancePoolsToUseCount                    requests.Integer                                          `position:"Query" name:"SpotInstancePoolsToUseCount"`
	LaunchConfigurationInternetChargeType          string                                                    `position:"Query" name:"LaunchConfiguration.InternetChargeType"`
	LaunchTemplateVersion                          string                                                    `position:"Query" name:"LaunchTemplateVersion"`
	LaunchConfigurationIoOptimized                 string                                                    `position:"Query" name:"LaunchConfiguration.IoOptimized"`
	PayAsYouGoTargetCapacity                       string                                                    `position:"Query" name:"PayAsYouGoTargetCapacity"`
	TotalTargetCapacity                            string                                                    `position:"Query" name:"TotalTargetCapacity"`
	SpotTargetCapacity                             string                                                    `position:"Query" name:"SpotTargetCapacity"`
	ValidFrom                                      string                                                    `position:"Query" name:"ValidFrom"`
	AutoProvisioningGroupName                      string                                                    `position:"Query" name:"AutoProvisioningGroupName"`
}

// CreateAutoProvisioningGroupLaunchConfiguration.DataDisk is a repeated param struct in CreateAutoProvisioningGroupRequest
type CreateAutoProvisioningGroupLaunchConfigurationDataDisk struct {
	Size     string `name:"Size"`
	Category string `name:"Category"`
}

// CreateAutoProvisioningGroupSystemDiskConfig is a repeated param struct in CreateAutoProvisioningGroupRequest
type CreateAutoProvisioningGroupSystemDiskConfig struct {
	DiskCategory string `name:"DiskCategory"`
}

// CreateAutoProvisioningGroupDataDiskConfig is a repeated param struct in CreateAutoProvisioningGroupRequest
type CreateAutoProvisioningGroupDataDiskConfig struct {
	DiskCategory string `name:"DiskCategory"`
}

// CreateAutoProvisioningGroupLaunchTemplateConfig is a repeated param struct in CreateAutoProvisioningGroupRequest
type CreateAutoProvisioningGroupLaunchTemplateConfig struct {
	InstanceType     string `name:"InstanceType"`
	MaxPrice         string `name:"MaxPrice"`
	VSwitchId        string `name:"VSwitchId"`
	WeightedCapacity string `name:"WeightedCapacity"`
	Priority         string `name:"Priority"`
}

// CreateAutoProvisioningGroupLaunchConfiguration.Tag is a repeated param struct in CreateAutoProvisioningGroupRequest
type CreateAutoProvisioningGroupLaunchConfigurationTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateAutoProvisioningGroupResponse is the response struct for api CreateAutoProvisioningGroup
type CreateAutoProvisioningGroupResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	AutoProvisioningGroupId string `json:"AutoProvisioningGroupId" xml:"AutoProvisioningGroupId"`
}

// CreateCreateAutoProvisioningGroupRequest creates a request to invoke CreateAutoProvisioningGroup API
func CreateCreateAutoProvisioningGroupRequest() (request *CreateAutoProvisioningGroupRequest) {
	request = &CreateAutoProvisioningGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateAutoProvisioningGroup", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAutoProvisioningGroupResponse creates a response to parse from CreateAutoProvisioningGroup response
func CreateCreateAutoProvisioningGroupResponse() (response *CreateAutoProvisioningGroupResponse) {
	response = &CreateAutoProvisioningGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
