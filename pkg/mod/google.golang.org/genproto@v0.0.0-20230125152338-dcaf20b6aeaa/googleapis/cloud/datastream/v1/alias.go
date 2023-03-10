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

// Package datastream aliases all exported identifiers in package
// "cloud.google.com/go/datastream/apiv1/datastreampb".
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package datastream

import (
	src "cloud.google.com/go/datastream/apiv1/datastreampb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/datastream/apiv1/datastreampb
const (
	BackfillJob_ACTIVE                            = src.BackfillJob_ACTIVE
	BackfillJob_AUTOMATIC                         = src.BackfillJob_AUTOMATIC
	BackfillJob_COMPLETED                         = src.BackfillJob_COMPLETED
	BackfillJob_FAILED                            = src.BackfillJob_FAILED
	BackfillJob_MANUAL                            = src.BackfillJob_MANUAL
	BackfillJob_NOT_STARTED                       = src.BackfillJob_NOT_STARTED
	BackfillJob_PENDING                           = src.BackfillJob_PENDING
	BackfillJob_STATE_UNSPECIFIED                 = src.BackfillJob_STATE_UNSPECIFIED
	BackfillJob_STOPPED                           = src.BackfillJob_STOPPED
	BackfillJob_TRIGGER_UNSPECIFIED               = src.BackfillJob_TRIGGER_UNSPECIFIED
	BackfillJob_UNSUPPORTED                       = src.BackfillJob_UNSUPPORTED
	JsonFileFormat_AVRO_SCHEMA_FILE               = src.JsonFileFormat_AVRO_SCHEMA_FILE
	JsonFileFormat_GZIP                           = src.JsonFileFormat_GZIP
	JsonFileFormat_JSON_COMPRESSION_UNSPECIFIED   = src.JsonFileFormat_JSON_COMPRESSION_UNSPECIFIED
	JsonFileFormat_NO_COMPRESSION                 = src.JsonFileFormat_NO_COMPRESSION
	JsonFileFormat_NO_SCHEMA_FILE                 = src.JsonFileFormat_NO_SCHEMA_FILE
	JsonFileFormat_SCHEMA_FILE_FORMAT_UNSPECIFIED = src.JsonFileFormat_SCHEMA_FILE_FORMAT_UNSPECIFIED
	PrivateConnection_CREATED                     = src.PrivateConnection_CREATED
	PrivateConnection_CREATING                    = src.PrivateConnection_CREATING
	PrivateConnection_DELETING                    = src.PrivateConnection_DELETING
	PrivateConnection_FAILED                      = src.PrivateConnection_FAILED
	PrivateConnection_FAILED_TO_DELETE            = src.PrivateConnection_FAILED_TO_DELETE
	PrivateConnection_STATE_UNSPECIFIED           = src.PrivateConnection_STATE_UNSPECIFIED
	Stream_DRAINING                               = src.Stream_DRAINING
	Stream_FAILED                                 = src.Stream_FAILED
	Stream_FAILED_PERMANENTLY                     = src.Stream_FAILED_PERMANENTLY
	Stream_MAINTENANCE                            = src.Stream_MAINTENANCE
	Stream_NOT_STARTED                            = src.Stream_NOT_STARTED
	Stream_PAUSED                                 = src.Stream_PAUSED
	Stream_RUNNING                                = src.Stream_RUNNING
	Stream_STARTING                               = src.Stream_STARTING
	Stream_STATE_UNSPECIFIED                      = src.Stream_STATE_UNSPECIFIED
	ValidationMessage_ERROR                       = src.ValidationMessage_ERROR
	ValidationMessage_LEVEL_UNSPECIFIED           = src.ValidationMessage_LEVEL_UNSPECIFIED
	ValidationMessage_WARNING                     = src.ValidationMessage_WARNING
	Validation_FAILED                             = src.Validation_FAILED
	Validation_NOT_EXECUTED                       = src.Validation_NOT_EXECUTED
	Validation_PASSED                             = src.Validation_PASSED
	Validation_STATE_UNSPECIFIED                  = src.Validation_STATE_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/datastream/apiv1/datastreampb
var (
	BackfillJob_State_name                                     = src.BackfillJob_State_name
	BackfillJob_State_value                                    = src.BackfillJob_State_value
	BackfillJob_Trigger_name                                   = src.BackfillJob_Trigger_name
	BackfillJob_Trigger_value                                  = src.BackfillJob_Trigger_value
	File_google_cloud_datastream_v1_datastream_proto           = src.File_google_cloud_datastream_v1_datastream_proto
	File_google_cloud_datastream_v1_datastream_resources_proto = src.File_google_cloud_datastream_v1_datastream_resources_proto
	JsonFileFormat_JsonCompression_name                        = src.JsonFileFormat_JsonCompression_name
	JsonFileFormat_JsonCompression_value                       = src.JsonFileFormat_JsonCompression_value
	JsonFileFormat_SchemaFileFormat_name                       = src.JsonFileFormat_SchemaFileFormat_name
	JsonFileFormat_SchemaFileFormat_value                      = src.JsonFileFormat_SchemaFileFormat_value
	PrivateConnection_State_name                               = src.PrivateConnection_State_name
	PrivateConnection_State_value                              = src.PrivateConnection_State_value
	Stream_State_name                                          = src.Stream_State_name
	Stream_State_value                                         = src.Stream_State_value
	ValidationMessage_Level_name                               = src.ValidationMessage_Level_name
	ValidationMessage_Level_value                              = src.ValidationMessage_Level_value
	Validation_State_name                                      = src.Validation_State_name
	Validation_State_value                                     = src.Validation_State_value
)

// AVRO file format configuration.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type AvroFileFormat = src.AvroFileFormat

// Represents a backfill job on a specific stream object.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type BackfillJob = src.BackfillJob

// State of the stream object's backfill job.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type BackfillJob_State = src.BackfillJob_State

// Triggering reason for a backfill job.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type BackfillJob_Trigger = src.BackfillJob_Trigger
type BigQueryDestinationConfig = src.BigQueryDestinationConfig

// A single target dataset to which all data will be streamed.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type BigQueryDestinationConfig_SingleTargetDataset = src.BigQueryDestinationConfig_SingleTargetDataset
type BigQueryDestinationConfig_SingleTargetDataset_ = src.BigQueryDestinationConfig_SingleTargetDataset_

// Destination datasets are created so that hierarchy of the destination data
// objects matches the source hierarchy.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type BigQueryDestinationConfig_SourceHierarchyDatasets = src.BigQueryDestinationConfig_SourceHierarchyDatasets
type BigQueryDestinationConfig_SourceHierarchyDatasets_ = src.BigQueryDestinationConfig_SourceHierarchyDatasets_

// Dataset template used for dynamic dataset creation.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate = src.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate

// BigQuery warehouse profile.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type BigQueryProfile = src.BigQueryProfile

// A set of reusable connection configurations to be used as a source or
// destination for a stream.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ConnectionProfile = src.ConnectionProfile
type ConnectionProfile_BigqueryProfile = src.ConnectionProfile_BigqueryProfile
type ConnectionProfile_ForwardSshConnectivity = src.ConnectionProfile_ForwardSshConnectivity
type ConnectionProfile_GcsProfile = src.ConnectionProfile_GcsProfile
type ConnectionProfile_MysqlProfile = src.ConnectionProfile_MysqlProfile
type ConnectionProfile_OracleProfile = src.ConnectionProfile_OracleProfile
type ConnectionProfile_PostgresqlProfile = src.ConnectionProfile_PostgresqlProfile
type ConnectionProfile_PrivateConnectivity = src.ConnectionProfile_PrivateConnectivity
type ConnectionProfile_StaticServiceIpConnectivity = src.ConnectionProfile_StaticServiceIpConnectivity

// Request message for creating a connection profile.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type CreateConnectionProfileRequest = src.CreateConnectionProfileRequest

// Request for creating a private connection.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type CreatePrivateConnectionRequest = src.CreatePrivateConnectionRequest

// Route creation request.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type CreateRouteRequest = src.CreateRouteRequest

// Request message for creating a stream.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type CreateStreamRequest = src.CreateStreamRequest

// DatastreamClient is the client API for Datastream service. For semantics
// around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type DatastreamClient = src.DatastreamClient

// DatastreamServer is the server API for Datastream service.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type DatastreamServer = src.DatastreamServer

// Request message for deleting a connection profile.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type DeleteConnectionProfileRequest = src.DeleteConnectionProfileRequest

// Request to delete a private connection.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type DeletePrivateConnectionRequest = src.DeletePrivateConnectionRequest

// Route deletion request.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type DeleteRouteRequest = src.DeleteRouteRequest

// Request message for deleting a stream.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type DeleteStreamRequest = src.DeleteStreamRequest

// The configuration of the stream destination.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type DestinationConfig = src.DestinationConfig
type DestinationConfig_BigqueryDestinationConfig = src.DestinationConfig_BigqueryDestinationConfig
type DestinationConfig_GcsDestinationConfig = src.DestinationConfig_GcsDestinationConfig

// Request message for 'discover' ConnectionProfile request.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type DiscoverConnectionProfileRequest = src.DiscoverConnectionProfileRequest
type DiscoverConnectionProfileRequest_ConnectionProfile = src.DiscoverConnectionProfileRequest_ConnectionProfile
type DiscoverConnectionProfileRequest_ConnectionProfileName = src.DiscoverConnectionProfileRequest_ConnectionProfileName
type DiscoverConnectionProfileRequest_FullHierarchy = src.DiscoverConnectionProfileRequest_FullHierarchy
type DiscoverConnectionProfileRequest_HierarchyDepth = src.DiscoverConnectionProfileRequest_HierarchyDepth
type DiscoverConnectionProfileRequest_MysqlRdbms = src.DiscoverConnectionProfileRequest_MysqlRdbms
type DiscoverConnectionProfileRequest_OracleRdbms = src.DiscoverConnectionProfileRequest_OracleRdbms
type DiscoverConnectionProfileRequest_PostgresqlRdbms = src.DiscoverConnectionProfileRequest_PostgresqlRdbms

// Response from a discover request.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type DiscoverConnectionProfileResponse = src.DiscoverConnectionProfileResponse
type DiscoverConnectionProfileResponse_MysqlRdbms = src.DiscoverConnectionProfileResponse_MysqlRdbms
type DiscoverConnectionProfileResponse_OracleRdbms = src.DiscoverConnectionProfileResponse_OracleRdbms
type DiscoverConnectionProfileResponse_PostgresqlRdbms = src.DiscoverConnectionProfileResponse_PostgresqlRdbms

// Represent a user-facing Error.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type Error = src.Error

// Request message for 'FetchStaticIps' request.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type FetchStaticIpsRequest = src.FetchStaticIpsRequest

// Response message for a 'FetchStaticIps' response.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type FetchStaticIpsResponse = src.FetchStaticIpsResponse

// Forward SSH Tunnel connectivity.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ForwardSshTunnelConnectivity = src.ForwardSshTunnelConnectivity
type ForwardSshTunnelConnectivity_Password = src.ForwardSshTunnelConnectivity_Password
type ForwardSshTunnelConnectivity_PrivateKey = src.ForwardSshTunnelConnectivity_PrivateKey

// Google Cloud Storage destination configuration
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type GcsDestinationConfig = src.GcsDestinationConfig
type GcsDestinationConfig_AvroFileFormat = src.GcsDestinationConfig_AvroFileFormat
type GcsDestinationConfig_JsonFileFormat = src.GcsDestinationConfig_JsonFileFormat

// Cloud Storage bucket profile.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type GcsProfile = src.GcsProfile

// Request message for getting a connection profile.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type GetConnectionProfileRequest = src.GetConnectionProfileRequest

// Request to get a private connection configuration.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type GetPrivateConnectionRequest = src.GetPrivateConnectionRequest

// Route get request.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type GetRouteRequest = src.GetRouteRequest

// Request for fetching a specific stream object.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type GetStreamObjectRequest = src.GetStreamObjectRequest

// Request message for getting a stream.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type GetStreamRequest = src.GetStreamRequest

// JSON file format configuration.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type JsonFileFormat = src.JsonFileFormat

// Json file compression.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type JsonFileFormat_JsonCompression = src.JsonFileFormat_JsonCompression

// Schema file format.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type JsonFileFormat_SchemaFileFormat = src.JsonFileFormat_SchemaFileFormat

// Request message for listing connection profiles.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ListConnectionProfilesRequest = src.ListConnectionProfilesRequest

// Response message for listing connection profiles.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ListConnectionProfilesResponse = src.ListConnectionProfilesResponse

// Request for listing private connections.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ListPrivateConnectionsRequest = src.ListPrivateConnectionsRequest

// Response containing a list of private connection configurations.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ListPrivateConnectionsResponse = src.ListPrivateConnectionsResponse

// Route list request.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ListRoutesRequest = src.ListRoutesRequest

// Route list response.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ListRoutesResponse = src.ListRoutesResponse

// Request for listing all objects for a specific stream.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ListStreamObjectsRequest = src.ListStreamObjectsRequest

// Response containing the objects for a stream.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ListStreamObjectsResponse = src.ListStreamObjectsResponse

// Request message for listing streams.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ListStreamsRequest = src.ListStreamsRequest

// Response message for listing streams.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ListStreamsResponse = src.ListStreamsResponse

// Request for looking up a specific stream object by its source object
// identifier.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type LookupStreamObjectRequest = src.LookupStreamObjectRequest

// MySQL Column.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type MysqlColumn = src.MysqlColumn

// MySQL database.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type MysqlDatabase = src.MysqlDatabase

// MySQL database profile.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type MysqlProfile = src.MysqlProfile

// MySQL database structure
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type MysqlRdbms = src.MysqlRdbms

// MySQL source configuration
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type MysqlSourceConfig = src.MysqlSourceConfig

// MySQL SSL configuration information.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type MysqlSslConfig = src.MysqlSslConfig

// MySQL table.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type MysqlTable = src.MysqlTable

// Represents the metadata of the long-running operation.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type OperationMetadata = src.OperationMetadata

// Oracle Column.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type OracleColumn = src.OracleColumn

// Oracle database profile.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type OracleProfile = src.OracleProfile

// Oracle database structure.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type OracleRdbms = src.OracleRdbms

// Oracle schema.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type OracleSchema = src.OracleSchema

// Oracle data source configuration
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type OracleSourceConfig = src.OracleSourceConfig

// Configuration to drop large object values.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type OracleSourceConfig_DropLargeObjects = src.OracleSourceConfig_DropLargeObjects
type OracleSourceConfig_DropLargeObjects_ = src.OracleSourceConfig_DropLargeObjects_

// Configuration to stream large object values.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type OracleSourceConfig_StreamLargeObjects = src.OracleSourceConfig_StreamLargeObjects
type OracleSourceConfig_StreamLargeObjects_ = src.OracleSourceConfig_StreamLargeObjects_

// Oracle table.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type OracleTable = src.OracleTable

// PostgreSQL Column.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type PostgresqlColumn = src.PostgresqlColumn

// PostgreSQL database profile.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type PostgresqlProfile = src.PostgresqlProfile

// PostgreSQL database structure.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type PostgresqlRdbms = src.PostgresqlRdbms

// PostgreSQL schema.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type PostgresqlSchema = src.PostgresqlSchema

// PostgreSQL data source configuration
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type PostgresqlSourceConfig = src.PostgresqlSourceConfig

// PostgreSQL table.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type PostgresqlTable = src.PostgresqlTable

// The PrivateConnection resource is used to establish private connectivity
// between Datastream and a customer's network.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type PrivateConnection = src.PrivateConnection

// Private Connection state.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type PrivateConnection_State = src.PrivateConnection_State

// Private Connectivity
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type PrivateConnectivity = src.PrivateConnectivity

// The route resource is the child of the private connection resource, used
// for defining a route for a private connection.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type Route = src.Route

// The configuration of the stream source.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type SourceConfig = src.SourceConfig
type SourceConfig_MysqlSourceConfig = src.SourceConfig_MysqlSourceConfig
type SourceConfig_OracleSourceConfig = src.SourceConfig_OracleSourceConfig
type SourceConfig_PostgresqlSourceConfig = src.SourceConfig_PostgresqlSourceConfig

// Represents an identifier of an object in the data source.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type SourceObjectIdentifier = src.SourceObjectIdentifier
type SourceObjectIdentifier_MysqlIdentifier = src.SourceObjectIdentifier_MysqlIdentifier

// Mysql data source object identifier.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type SourceObjectIdentifier_MysqlObjectIdentifier = src.SourceObjectIdentifier_MysqlObjectIdentifier
type SourceObjectIdentifier_OracleIdentifier = src.SourceObjectIdentifier_OracleIdentifier

// Oracle data source object identifier.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type SourceObjectIdentifier_OracleObjectIdentifier = src.SourceObjectIdentifier_OracleObjectIdentifier
type SourceObjectIdentifier_PostgresqlIdentifier = src.SourceObjectIdentifier_PostgresqlIdentifier

// PostgreSQL data source object identifier.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type SourceObjectIdentifier_PostgresqlObjectIdentifier = src.SourceObjectIdentifier_PostgresqlObjectIdentifier

// Request for manually initiating a backfill job for a specific stream
// object.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type StartBackfillJobRequest = src.StartBackfillJobRequest

// Response for manually initiating a backfill job for a specific stream
// object.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type StartBackfillJobResponse = src.StartBackfillJobResponse

// Static IP address connectivity.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type StaticServiceIpConnectivity = src.StaticServiceIpConnectivity

// Request for manually stopping a running backfill job for a specific stream
// object.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type StopBackfillJobRequest = src.StopBackfillJobRequest

// Response for manually stop a backfill job for a specific stream object.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type StopBackfillJobResponse = src.StopBackfillJobResponse

// A resource representing streaming data from a source to a destination.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type Stream = src.Stream

// A specific stream object (e.g a specific DB table).
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type StreamObject = src.StreamObject
type Stream_BackfillAll = src.Stream_BackfillAll

// Backfill strategy to automatically backfill the Stream's objects. Specific
// objects can be excluded.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type Stream_BackfillAllStrategy = src.Stream_BackfillAllStrategy
type Stream_BackfillAllStrategy_MysqlExcludedObjects = src.Stream_BackfillAllStrategy_MysqlExcludedObjects
type Stream_BackfillAllStrategy_OracleExcludedObjects = src.Stream_BackfillAllStrategy_OracleExcludedObjects
type Stream_BackfillAllStrategy_PostgresqlExcludedObjects = src.Stream_BackfillAllStrategy_PostgresqlExcludedObjects
type Stream_BackfillNone = src.Stream_BackfillNone

// Backfill strategy to disable automatic backfill for the Stream's objects.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type Stream_BackfillNoneStrategy = src.Stream_BackfillNoneStrategy

// Stream state.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type Stream_State = src.Stream_State

// UnimplementedDatastreamServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type UnimplementedDatastreamServer = src.UnimplementedDatastreamServer

// Connection profile update message.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type UpdateConnectionProfileRequest = src.UpdateConnectionProfileRequest

// Request message for updating a stream.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type UpdateStreamRequest = src.UpdateStreamRequest

// A validation to perform on a stream.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type Validation = src.Validation

// Represent user-facing validation result message.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ValidationMessage = src.ValidationMessage

// Validation message level.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ValidationMessage_Level = src.ValidationMessage_Level

// Contains the current validation results.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type ValidationResult = src.ValidationResult

// Validation execution state.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type Validation_State = src.Validation_State

// The VPC Peering configuration is used to create VPC peering between
// Datastream and the consumer's VPC.
//
// Deprecated: Please use types in: cloud.google.com/go/datastream/apiv1/datastreampb
type VpcPeeringConfig = src.VpcPeeringConfig

// Deprecated: Please use funcs in: cloud.google.com/go/datastream/apiv1/datastreampb
func NewDatastreamClient(cc grpc.ClientConnInterface) DatastreamClient {
	return src.NewDatastreamClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/datastream/apiv1/datastreampb
func RegisterDatastreamServer(s *grpc.Server, srv DatastreamServer) {
	src.RegisterDatastreamServer(s, srv)
}
