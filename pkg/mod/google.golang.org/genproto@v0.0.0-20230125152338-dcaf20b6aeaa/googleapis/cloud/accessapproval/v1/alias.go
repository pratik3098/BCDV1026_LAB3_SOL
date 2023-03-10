// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package accessapproval aliases all exported identifiers in package
// "cloud.google.com/go/accessapproval/apiv1/accessapprovalpb".
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package accessapproval

import (
	src "cloud.google.com/go/accessapproval/apiv1/accessapprovalpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
const (
	AccessReason_CUSTOMER_INITIATED_SUPPORT          = src.AccessReason_CUSTOMER_INITIATED_SUPPORT
	AccessReason_GOOGLE_INITIATED_REVIEW             = src.AccessReason_GOOGLE_INITIATED_REVIEW
	AccessReason_GOOGLE_INITIATED_SERVICE            = src.AccessReason_GOOGLE_INITIATED_SERVICE
	AccessReason_GOOGLE_RESPONSE_TO_PRODUCTION_ALERT = src.AccessReason_GOOGLE_RESPONSE_TO_PRODUCTION_ALERT
	AccessReason_THIRD_PARTY_DATA_REQUEST            = src.AccessReason_THIRD_PARTY_DATA_REQUEST
	AccessReason_TYPE_UNSPECIFIED                    = src.AccessReason_TYPE_UNSPECIFIED
	EnrollmentLevel_BLOCK_ALL                        = src.EnrollmentLevel_BLOCK_ALL
	EnrollmentLevel_ENROLLMENT_LEVEL_UNSPECIFIED     = src.EnrollmentLevel_ENROLLMENT_LEVEL_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
var (
	AccessReason_Type_name                                   = src.AccessReason_Type_name
	AccessReason_Type_value                                  = src.AccessReason_Type_value
	EnrollmentLevel_name                                     = src.EnrollmentLevel_name
	EnrollmentLevel_value                                    = src.EnrollmentLevel_value
	File_google_cloud_accessapproval_v1_accessapproval_proto = src.File_google_cloud_accessapproval_v1_accessapproval_proto
)

// AccessApprovalClient is the client API for AccessApproval service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type AccessApprovalClient = src.AccessApprovalClient

// AccessApprovalServer is the server API for AccessApproval service.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type AccessApprovalServer = src.AccessApprovalServer

// Access Approval service account related to a project/folder/organization.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type AccessApprovalServiceAccount = src.AccessApprovalServiceAccount

// Settings on a Project/Folder/Organization related to Access Approval.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type AccessApprovalSettings = src.AccessApprovalSettings

// Home office and physical location of the principal.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type AccessLocations = src.AccessLocations
type AccessReason = src.AccessReason

// Type of access justification.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type AccessReason_Type = src.AccessReason_Type

// A request for the customer to approve access to a resource.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type ApprovalRequest = src.ApprovalRequest
type ApprovalRequest_Approve = src.ApprovalRequest_Approve
type ApprovalRequest_Dismiss = src.ApprovalRequest_Dismiss

// Request to approve an ApprovalRequest.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type ApproveApprovalRequestMessage = src.ApproveApprovalRequestMessage

// A decision that has been made to approve access to a resource.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type ApproveDecision = src.ApproveDecision

// Request to delete access approval settings.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type DeleteAccessApprovalSettingsMessage = src.DeleteAccessApprovalSettingsMessage

// Request to dismiss an approval request.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type DismissApprovalRequestMessage = src.DismissApprovalRequestMessage

// A decision that has been made to dismiss an approval request.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type DismissDecision = src.DismissDecision

// Represents the enrollment of a cloud resource into a specific service.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type EnrolledService = src.EnrolledService

// Represents the type of enrollment for a given service to Access Approval.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type EnrollmentLevel = src.EnrollmentLevel

// Request to get an Access Approval service account.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type GetAccessApprovalServiceAccountMessage = src.GetAccessApprovalServiceAccountMessage

// Request to get access approval settings.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type GetAccessApprovalSettingsMessage = src.GetAccessApprovalSettingsMessage

// Request to get an approval request.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type GetApprovalRequestMessage = src.GetApprovalRequestMessage

// Request to invalidate an existing approval.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type InvalidateApprovalRequestMessage = src.InvalidateApprovalRequestMessage

// Request to list approval requests.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type ListApprovalRequestsMessage = src.ListApprovalRequestsMessage

// Response to listing of ApprovalRequest objects.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type ListApprovalRequestsResponse = src.ListApprovalRequestsResponse

// The properties associated with the resource of the request.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type ResourceProperties = src.ResourceProperties

// Information about the digital signature of the resource.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type SignatureInfo = src.SignatureInfo
type SignatureInfo_CustomerKmsKeyVersion = src.SignatureInfo_CustomerKmsKeyVersion
type SignatureInfo_GooglePublicKeyPem = src.SignatureInfo_GooglePublicKeyPem

// UnimplementedAccessApprovalServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type UnimplementedAccessApprovalServer = src.UnimplementedAccessApprovalServer

// Request to update access approval settings.
//
// Deprecated: Please use types in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
type UpdateAccessApprovalSettingsMessage = src.UpdateAccessApprovalSettingsMessage

// Deprecated: Please use funcs in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
func NewAccessApprovalClient(cc grpc.ClientConnInterface) AccessApprovalClient {
	return src.NewAccessApprovalClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/accessapproval/apiv1/accessapprovalpb
func RegisterAccessApprovalServer(s *grpc.Server, srv AccessApprovalServer) {
	src.RegisterAccessApprovalServer(s, srv)
}
