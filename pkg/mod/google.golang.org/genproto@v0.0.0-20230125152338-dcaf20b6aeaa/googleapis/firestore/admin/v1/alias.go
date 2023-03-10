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

// Package admin aliases all exported identifiers in package
// "cloud.google.com/go/firestore/apiv1/admin/adminpb".
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package admin

import (
	src "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/firestore/apiv1/admin/adminpb
const (
	Database_APP_ENGINE_INTEGRATION_MODE_UNSPECIFIED                = src.Database_APP_ENGINE_INTEGRATION_MODE_UNSPECIFIED
	Database_CONCURRENCY_MODE_UNSPECIFIED                           = src.Database_CONCURRENCY_MODE_UNSPECIFIED
	Database_DATABASE_TYPE_UNSPECIFIED                              = src.Database_DATABASE_TYPE_UNSPECIFIED
	Database_DATASTORE_MODE                                         = src.Database_DATASTORE_MODE
	Database_DISABLED                                               = src.Database_DISABLED
	Database_ENABLED                                                = src.Database_ENABLED
	Database_FIRESTORE_NATIVE                                       = src.Database_FIRESTORE_NATIVE
	Database_OPTIMISTIC                                             = src.Database_OPTIMISTIC
	Database_OPTIMISTIC_WITH_ENTITY_GROUPS                          = src.Database_OPTIMISTIC_WITH_ENTITY_GROUPS
	Database_PESSIMISTIC                                            = src.Database_PESSIMISTIC
	FieldOperationMetadata_IndexConfigDelta_ADD                     = src.FieldOperationMetadata_IndexConfigDelta_ADD
	FieldOperationMetadata_IndexConfigDelta_CHANGE_TYPE_UNSPECIFIED = src.FieldOperationMetadata_IndexConfigDelta_CHANGE_TYPE_UNSPECIFIED
	FieldOperationMetadata_IndexConfigDelta_REMOVE                  = src.FieldOperationMetadata_IndexConfigDelta_REMOVE
	FieldOperationMetadata_TtlConfigDelta_ADD                       = src.FieldOperationMetadata_TtlConfigDelta_ADD
	FieldOperationMetadata_TtlConfigDelta_CHANGE_TYPE_UNSPECIFIED   = src.FieldOperationMetadata_TtlConfigDelta_CHANGE_TYPE_UNSPECIFIED
	FieldOperationMetadata_TtlConfigDelta_REMOVE                    = src.FieldOperationMetadata_TtlConfigDelta_REMOVE
	Field_TtlConfig_ACTIVE                                          = src.Field_TtlConfig_ACTIVE
	Field_TtlConfig_CREATING                                        = src.Field_TtlConfig_CREATING
	Field_TtlConfig_NEEDS_REPAIR                                    = src.Field_TtlConfig_NEEDS_REPAIR
	Field_TtlConfig_STATE_UNSPECIFIED                               = src.Field_TtlConfig_STATE_UNSPECIFIED
	Index_COLLECTION                                                = src.Index_COLLECTION
	Index_COLLECTION_GROUP                                          = src.Index_COLLECTION_GROUP
	Index_CREATING                                                  = src.Index_CREATING
	Index_IndexField_ARRAY_CONFIG_UNSPECIFIED                       = src.Index_IndexField_ARRAY_CONFIG_UNSPECIFIED
	Index_IndexField_ASCENDING                                      = src.Index_IndexField_ASCENDING
	Index_IndexField_CONTAINS                                       = src.Index_IndexField_CONTAINS
	Index_IndexField_DESCENDING                                     = src.Index_IndexField_DESCENDING
	Index_IndexField_ORDER_UNSPECIFIED                              = src.Index_IndexField_ORDER_UNSPECIFIED
	Index_NEEDS_REPAIR                                              = src.Index_NEEDS_REPAIR
	Index_QUERY_SCOPE_UNSPECIFIED                                   = src.Index_QUERY_SCOPE_UNSPECIFIED
	Index_READY                                                     = src.Index_READY
	Index_STATE_UNSPECIFIED                                         = src.Index_STATE_UNSPECIFIED
	OperationState_CANCELLED                                        = src.OperationState_CANCELLED
	OperationState_CANCELLING                                       = src.OperationState_CANCELLING
	OperationState_FAILED                                           = src.OperationState_FAILED
	OperationState_FINALIZING                                       = src.OperationState_FINALIZING
	OperationState_INITIALIZING                                     = src.OperationState_INITIALIZING
	OperationState_OPERATION_STATE_UNSPECIFIED                      = src.OperationState_OPERATION_STATE_UNSPECIFIED
	OperationState_PROCESSING                                       = src.OperationState_PROCESSING
	OperationState_SUCCESSFUL                                       = src.OperationState_SUCCESSFUL
)

// Deprecated: Please use vars in: cloud.google.com/go/firestore/apiv1/admin/adminpb
var (
	Database_AppEngineIntegrationMode_name                   = src.Database_AppEngineIntegrationMode_name
	Database_AppEngineIntegrationMode_value                  = src.Database_AppEngineIntegrationMode_value
	Database_ConcurrencyMode_name                            = src.Database_ConcurrencyMode_name
	Database_ConcurrencyMode_value                           = src.Database_ConcurrencyMode_value
	Database_DatabaseType_name                               = src.Database_DatabaseType_name
	Database_DatabaseType_value                              = src.Database_DatabaseType_value
	FieldOperationMetadata_IndexConfigDelta_ChangeType_name  = src.FieldOperationMetadata_IndexConfigDelta_ChangeType_name
	FieldOperationMetadata_IndexConfigDelta_ChangeType_value = src.FieldOperationMetadata_IndexConfigDelta_ChangeType_value
	FieldOperationMetadata_TtlConfigDelta_ChangeType_name    = src.FieldOperationMetadata_TtlConfigDelta_ChangeType_name
	FieldOperationMetadata_TtlConfigDelta_ChangeType_value   = src.FieldOperationMetadata_TtlConfigDelta_ChangeType_value
	Field_TtlConfig_State_name                               = src.Field_TtlConfig_State_name
	Field_TtlConfig_State_value                              = src.Field_TtlConfig_State_value
	File_google_firestore_admin_v1_database_proto            = src.File_google_firestore_admin_v1_database_proto
	File_google_firestore_admin_v1_field_proto               = src.File_google_firestore_admin_v1_field_proto
	File_google_firestore_admin_v1_firestore_admin_proto     = src.File_google_firestore_admin_v1_firestore_admin_proto
	File_google_firestore_admin_v1_index_proto               = src.File_google_firestore_admin_v1_index_proto
	File_google_firestore_admin_v1_location_proto            = src.File_google_firestore_admin_v1_location_proto
	File_google_firestore_admin_v1_operation_proto           = src.File_google_firestore_admin_v1_operation_proto
	Index_IndexField_ArrayConfig_name                        = src.Index_IndexField_ArrayConfig_name
	Index_IndexField_ArrayConfig_value                       = src.Index_IndexField_ArrayConfig_value
	Index_IndexField_Order_name                              = src.Index_IndexField_Order_name
	Index_IndexField_Order_value                             = src.Index_IndexField_Order_value
	Index_QueryScope_name                                    = src.Index_QueryScope_name
	Index_QueryScope_value                                   = src.Index_QueryScope_value
	Index_State_name                                         = src.Index_State_name
	Index_State_value                                        = src.Index_State_value
	OperationState_name                                      = src.OperationState_name
	OperationState_value                                     = src.OperationState_value
)

// The request for
// [FirestoreAdmin.CreateIndex][google.firestore.admin.v1.FirestoreAdmin.CreateIndex].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type CreateIndexRequest = src.CreateIndexRequest

// A Cloud Firestore Database. Currently only one database is allowed per
// cloud project; this database must have a `database_id` of '(default)'.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Database = src.Database

// The type of App Engine integration mode.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Database_AppEngineIntegrationMode = src.Database_AppEngineIntegrationMode

// The type of concurrency control mode for transactions.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Database_ConcurrencyMode = src.Database_ConcurrencyMode

// The type of the database. See
// https://cloud.google.com/datastore/docs/firestore-or-datastore for
// information about how to choose. Mode changes are only allowed if the
// database is empty.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Database_DatabaseType = src.Database_DatabaseType

// The request for
// [FirestoreAdmin.DeleteIndex][google.firestore.admin.v1.FirestoreAdmin.DeleteIndex].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type DeleteIndexRequest = src.DeleteIndexRequest

// Metadata for [google.longrunning.Operation][google.longrunning.Operation]
// results from
// [FirestoreAdmin.ExportDocuments][google.firestore.admin.v1.FirestoreAdmin.ExportDocuments].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ExportDocumentsMetadata = src.ExportDocumentsMetadata

// The request for
// [FirestoreAdmin.ExportDocuments][google.firestore.admin.v1.FirestoreAdmin.ExportDocuments].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ExportDocumentsRequest = src.ExportDocumentsRequest

// Returned in the
// [google.longrunning.Operation][google.longrunning.Operation] response field.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ExportDocumentsResponse = src.ExportDocumentsResponse

// Represents a single field in the database. Fields are grouped by their
// "Collection Group", which represent all collections in the database with the
// same id.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Field = src.Field

// Metadata for [google.longrunning.Operation][google.longrunning.Operation]
// results from
// [FirestoreAdmin.UpdateField][google.firestore.admin.v1.FirestoreAdmin.UpdateField].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type FieldOperationMetadata = src.FieldOperationMetadata

// Information about an index configuration change.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type FieldOperationMetadata_IndexConfigDelta = src.FieldOperationMetadata_IndexConfigDelta

// Specifies how the index is changing.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type FieldOperationMetadata_IndexConfigDelta_ChangeType = src.FieldOperationMetadata_IndexConfigDelta_ChangeType

// Information about an TTL configuration change.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type FieldOperationMetadata_TtlConfigDelta = src.FieldOperationMetadata_TtlConfigDelta

// Specifies how the TTL config is changing.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type FieldOperationMetadata_TtlConfigDelta_ChangeType = src.FieldOperationMetadata_TtlConfigDelta_ChangeType

// The index configuration for this field.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Field_IndexConfig = src.Field_IndexConfig

// The TTL (time-to-live) configuration for documents that have this `Field`
// set. Storing a timestamp value into a TTL-enabled field will be treated as
// the document's absolute expiration time. Using any other data type or
// leaving the field absent will disable the TTL for the individual document.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Field_TtlConfig = src.Field_TtlConfig

// The state of applying the TTL configuration to all documents.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Field_TtlConfig_State = src.Field_TtlConfig_State

// FirestoreAdminClient is the client API for FirestoreAdmin service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type FirestoreAdminClient = src.FirestoreAdminClient

// FirestoreAdminServer is the server API for FirestoreAdmin service.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type FirestoreAdminServer = src.FirestoreAdminServer

// The request for
// [FirestoreAdmin.GetDatabase][google.firestore.admin.v1.FirestoreAdmin.GetDatabase].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type GetDatabaseRequest = src.GetDatabaseRequest

// The request for
// [FirestoreAdmin.GetField][google.firestore.admin.v1.FirestoreAdmin.GetField].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type GetFieldRequest = src.GetFieldRequest

// The request for
// [FirestoreAdmin.GetIndex][google.firestore.admin.v1.FirestoreAdmin.GetIndex].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type GetIndexRequest = src.GetIndexRequest

// Metadata for [google.longrunning.Operation][google.longrunning.Operation]
// results from
// [FirestoreAdmin.ImportDocuments][google.firestore.admin.v1.FirestoreAdmin.ImportDocuments].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ImportDocumentsMetadata = src.ImportDocumentsMetadata

// The request for
// [FirestoreAdmin.ImportDocuments][google.firestore.admin.v1.FirestoreAdmin.ImportDocuments].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ImportDocumentsRequest = src.ImportDocumentsRequest

// Cloud Firestore indexes enable simple and complex queries against documents
// in a database.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Index = src.Index

// Metadata for [google.longrunning.Operation][google.longrunning.Operation]
// results from
// [FirestoreAdmin.CreateIndex][google.firestore.admin.v1.FirestoreAdmin.CreateIndex].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type IndexOperationMetadata = src.IndexOperationMetadata

// A field in an index. The field_path describes which field is indexed, the
// value_mode describes how the field value is indexed.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Index_IndexField = src.Index_IndexField

// The supported array value configurations.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Index_IndexField_ArrayConfig = src.Index_IndexField_ArrayConfig
type Index_IndexField_ArrayConfig_ = src.Index_IndexField_ArrayConfig_

// The supported orderings.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Index_IndexField_Order = src.Index_IndexField_Order
type Index_IndexField_Order_ = src.Index_IndexField_Order_

// Query Scope defines the scope at which a query is run. This is specified on
// a StructuredQuery's `from` field.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Index_QueryScope = src.Index_QueryScope

// The state of an index. During index creation, an index will be in the
// `CREATING` state. If the index is created successfully, it will transition
// to the `READY` state. If the index creation encounters a problem, the index
// will transition to the `NEEDS_REPAIR` state.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Index_State = src.Index_State

// A request to list the Firestore Databases in all locations for a project.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ListDatabasesRequest = src.ListDatabasesRequest

// The list of databases for a project.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ListDatabasesResponse = src.ListDatabasesResponse

// The request for
// [FirestoreAdmin.ListFields][google.firestore.admin.v1.FirestoreAdmin.ListFields].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ListFieldsRequest = src.ListFieldsRequest

// The response for
// [FirestoreAdmin.ListFields][google.firestore.admin.v1.FirestoreAdmin.ListFields].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ListFieldsResponse = src.ListFieldsResponse

// The request for
// [FirestoreAdmin.ListIndexes][google.firestore.admin.v1.FirestoreAdmin.ListIndexes].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ListIndexesRequest = src.ListIndexesRequest

// The response for
// [FirestoreAdmin.ListIndexes][google.firestore.admin.v1.FirestoreAdmin.ListIndexes].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type ListIndexesResponse = src.ListIndexesResponse

// The metadata message for
// [google.cloud.location.Location.metadata][google.cloud.location.Location.metadata].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type LocationMetadata = src.LocationMetadata

// Describes the state of the operation.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type OperationState = src.OperationState

// Describes the progress of the operation. Unit of work is generic and must
// be interpreted based on where [Progress][google.firestore.admin.v1.Progress]
// is used.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type Progress = src.Progress

// UnimplementedFirestoreAdminServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type UnimplementedFirestoreAdminServer = src.UnimplementedFirestoreAdminServer

// Metadata related to the update database operation.
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type UpdateDatabaseMetadata = src.UpdateDatabaseMetadata

// The request for
// [FirestoreAdmin.UpdateDatabase][google.firestore.admin.v1.FirestoreAdmin.UpdateDatabase].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type UpdateDatabaseRequest = src.UpdateDatabaseRequest

// The request for
// [FirestoreAdmin.UpdateField][google.firestore.admin.v1.FirestoreAdmin.UpdateField].
//
// Deprecated: Please use types in: cloud.google.com/go/firestore/apiv1/admin/adminpb
type UpdateFieldRequest = src.UpdateFieldRequest

// Deprecated: Please use funcs in: cloud.google.com/go/firestore/apiv1/admin/adminpb
func NewFirestoreAdminClient(cc grpc.ClientConnInterface) FirestoreAdminClient {
	return src.NewFirestoreAdminClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/firestore/apiv1/admin/adminpb
func RegisterFirestoreAdminServer(s *grpc.Server, srv FirestoreAdminServer) {
	src.RegisterFirestoreAdminServer(s, srv)
}
