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

// Package datacatalog aliases all exported identifiers in package
// "cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb".
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package datacatalog

import (
	src "cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
const (
	EntryType_DATA_STREAM                           = src.EntryType_DATA_STREAM
	EntryType_ENTRY_TYPE_UNSPECIFIED                = src.EntryType_ENTRY_TYPE_UNSPECIFIED
	EntryType_FILESET                               = src.EntryType_FILESET
	EntryType_MODEL                                 = src.EntryType_MODEL
	EntryType_TABLE                                 = src.EntryType_TABLE
	FieldType_BOOL                                  = src.FieldType_BOOL
	FieldType_DOUBLE                                = src.FieldType_DOUBLE
	FieldType_PRIMITIVE_TYPE_UNSPECIFIED            = src.FieldType_PRIMITIVE_TYPE_UNSPECIFIED
	FieldType_STRING                                = src.FieldType_STRING
	FieldType_TIMESTAMP                             = src.FieldType_TIMESTAMP
	IntegratedSystem_BIGQUERY                       = src.IntegratedSystem_BIGQUERY
	IntegratedSystem_CLOUD_PUBSUB                   = src.IntegratedSystem_CLOUD_PUBSUB
	IntegratedSystem_INTEGRATED_SYSTEM_UNSPECIFIED  = src.IntegratedSystem_INTEGRATED_SYSTEM_UNSPECIFIED
	SearchResultType_ENTRY                          = src.SearchResultType_ENTRY
	SearchResultType_ENTRY_GROUP                    = src.SearchResultType_ENTRY_GROUP
	SearchResultType_SEARCH_RESULT_TYPE_UNSPECIFIED = src.SearchResultType_SEARCH_RESULT_TYPE_UNSPECIFIED
	SearchResultType_TAG_TEMPLATE                   = src.SearchResultType_TAG_TEMPLATE
	TableSourceType_BIGQUERY_TABLE                  = src.TableSourceType_BIGQUERY_TABLE
	TableSourceType_BIGQUERY_VIEW                   = src.TableSourceType_BIGQUERY_VIEW
	TableSourceType_TABLE_SOURCE_TYPE_UNSPECIFIED   = src.TableSourceType_TABLE_SOURCE_TYPE_UNSPECIFIED
	Taxonomy_FINE_GRAINED_ACCESS_CONTROL            = src.Taxonomy_FINE_GRAINED_ACCESS_CONTROL
	Taxonomy_POLICY_TYPE_UNSPECIFIED                = src.Taxonomy_POLICY_TYPE_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
var (
	EntryType_name                                                            = src.EntryType_name
	EntryType_value                                                           = src.EntryType_value
	FieldType_PrimitiveType_name                                              = src.FieldType_PrimitiveType_name
	FieldType_PrimitiveType_value                                             = src.FieldType_PrimitiveType_value
	File_google_cloud_datacatalog_v1beta1_common_proto                        = src.File_google_cloud_datacatalog_v1beta1_common_proto
	File_google_cloud_datacatalog_v1beta1_datacatalog_proto                   = src.File_google_cloud_datacatalog_v1beta1_datacatalog_proto
	File_google_cloud_datacatalog_v1beta1_gcs_fileset_spec_proto              = src.File_google_cloud_datacatalog_v1beta1_gcs_fileset_spec_proto
	File_google_cloud_datacatalog_v1beta1_policytagmanager_proto              = src.File_google_cloud_datacatalog_v1beta1_policytagmanager_proto
	File_google_cloud_datacatalog_v1beta1_policytagmanagerserialization_proto = src.File_google_cloud_datacatalog_v1beta1_policytagmanagerserialization_proto
	File_google_cloud_datacatalog_v1beta1_schema_proto                        = src.File_google_cloud_datacatalog_v1beta1_schema_proto
	File_google_cloud_datacatalog_v1beta1_search_proto                        = src.File_google_cloud_datacatalog_v1beta1_search_proto
	File_google_cloud_datacatalog_v1beta1_table_spec_proto                    = src.File_google_cloud_datacatalog_v1beta1_table_spec_proto
	File_google_cloud_datacatalog_v1beta1_tags_proto                          = src.File_google_cloud_datacatalog_v1beta1_tags_proto
	File_google_cloud_datacatalog_v1beta1_timestamps_proto                    = src.File_google_cloud_datacatalog_v1beta1_timestamps_proto
	IntegratedSystem_name                                                     = src.IntegratedSystem_name
	IntegratedSystem_value                                                    = src.IntegratedSystem_value
	SearchResultType_name                                                     = src.SearchResultType_name
	SearchResultType_value                                                    = src.SearchResultType_value
	TableSourceType_name                                                      = src.TableSourceType_name
	TableSourceType_value                                                     = src.TableSourceType_value
	Taxonomy_PolicyType_name                                                  = src.Taxonomy_PolicyType_name
	Taxonomy_PolicyType_value                                                 = src.Taxonomy_PolicyType_value
)

// Spec for a group of BigQuery tables with name pattern `[prefix]YYYYMMDD`.
// Context:
// https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type BigQueryDateShardedSpec = src.BigQueryDateShardedSpec

// Describes a BigQuery table.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type BigQueryTableSpec = src.BigQueryTableSpec
type BigQueryTableSpec_TableSpec = src.BigQueryTableSpec_TableSpec
type BigQueryTableSpec_ViewSpec = src.BigQueryTableSpec_ViewSpec

// Representation of a column within a schema. Columns could be nested inside
// other columns.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ColumnSchema = src.ColumnSchema

// Request message for
// [CreateEntryGroup][google.cloud.datacatalog.v1beta1.DataCatalog.CreateEntryGroup].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type CreateEntryGroupRequest = src.CreateEntryGroupRequest

// Request message for
// [CreateEntry][google.cloud.datacatalog.v1beta1.DataCatalog.CreateEntry].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type CreateEntryRequest = src.CreateEntryRequest

// Request message for
// [CreatePolicyTag][google.cloud.datacatalog.v1beta1.PolicyTagManager.CreatePolicyTag].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type CreatePolicyTagRequest = src.CreatePolicyTagRequest

// Request message for
// [CreateTag][google.cloud.datacatalog.v1beta1.DataCatalog.CreateTag].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type CreateTagRequest = src.CreateTagRequest

// Request message for
// [CreateTagTemplateField][google.cloud.datacatalog.v1beta1.DataCatalog.CreateTagTemplateField].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type CreateTagTemplateFieldRequest = src.CreateTagTemplateFieldRequest

// Request message for
// [CreateTagTemplate][google.cloud.datacatalog.v1beta1.DataCatalog.CreateTagTemplate].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type CreateTagTemplateRequest = src.CreateTagTemplateRequest

// Request message for
// [CreateTaxonomy][google.cloud.datacatalog.v1beta1.PolicyTagManager.CreateTaxonomy].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type CreateTaxonomyRequest = src.CreateTaxonomyRequest

// DataCatalogClient is the client API for DataCatalog service. For semantics
// around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type DataCatalogClient = src.DataCatalogClient

// DataCatalogServer is the server API for DataCatalog service.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type DataCatalogServer = src.DataCatalogServer

// Request message for
// [DeleteEntryGroup][google.cloud.datacatalog.v1beta1.DataCatalog.DeleteEntryGroup].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type DeleteEntryGroupRequest = src.DeleteEntryGroupRequest

// Request message for
// [DeleteEntry][google.cloud.datacatalog.v1beta1.DataCatalog.DeleteEntry].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type DeleteEntryRequest = src.DeleteEntryRequest

// Request message for
// [DeletePolicyTag][google.cloud.datacatalog.v1beta1.PolicyTagManager.DeletePolicyTag].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type DeletePolicyTagRequest = src.DeletePolicyTagRequest

// Request message for
// [DeleteTag][google.cloud.datacatalog.v1beta1.DataCatalog.DeleteTag].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type DeleteTagRequest = src.DeleteTagRequest

// Request message for
// [DeleteTagTemplateField][google.cloud.datacatalog.v1beta1.DataCatalog.DeleteTagTemplateField].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type DeleteTagTemplateFieldRequest = src.DeleteTagTemplateFieldRequest

// Request message for
// [DeleteTagTemplate][google.cloud.datacatalog.v1beta1.DataCatalog.DeleteTagTemplate].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type DeleteTagTemplateRequest = src.DeleteTagTemplateRequest

// Request message for
// [DeleteTaxonomy][google.cloud.datacatalog.v1beta1.PolicyTagManager.DeleteTaxonomy].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type DeleteTaxonomyRequest = src.DeleteTaxonomyRequest

// Entry Metadata. A Data Catalog Entry resource represents another resource
// in Google Cloud Platform (such as a BigQuery dataset or a Pub/Sub topic), or
// outside of Google Cloud Platform. Clients can use the `linked_resource`
// field in the Entry resource to refer to the original resource ID of the
// source system. An Entry resource contains resource details, such as its
// schema. An Entry can also be used to attach flexible metadata, such as a
// [Tag][google.cloud.datacatalog.v1beta1.Tag].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type Entry = src.Entry

// EntryGroup Metadata. An EntryGroup resource represents a logical grouping
// of zero or more Data Catalog [Entry][google.cloud.datacatalog.v1beta1.Entry]
// resources.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type EntryGroup = src.EntryGroup

// Entry resources in Data Catalog can be of different types e.g. a BigQuery
// Table entry is of type `TABLE`. This enum describes all the possible types
// Data Catalog contains.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type EntryType = src.EntryType
type Entry_BigqueryDateShardedSpec = src.Entry_BigqueryDateShardedSpec
type Entry_BigqueryTableSpec = src.Entry_BigqueryTableSpec
type Entry_GcsFilesetSpec = src.Entry_GcsFilesetSpec
type Entry_IntegratedSystem = src.Entry_IntegratedSystem
type Entry_Type = src.Entry_Type
type Entry_UserSpecifiedSystem = src.Entry_UserSpecifiedSystem
type Entry_UserSpecifiedType = src.Entry_UserSpecifiedType

// Request message for
// [ExportTaxonomies][google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization.ExportTaxonomies].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ExportTaxonomiesRequest = src.ExportTaxonomiesRequest
type ExportTaxonomiesRequest_SerializedTaxonomies = src.ExportTaxonomiesRequest_SerializedTaxonomies

// Response message for
// [ExportTaxonomies][google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization.ExportTaxonomies].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ExportTaxonomiesResponse = src.ExportTaxonomiesResponse
type FieldType = src.FieldType
type FieldType_EnumType = src.FieldType_EnumType
type FieldType_EnumType_ = src.FieldType_EnumType_
type FieldType_EnumType_EnumValue = src.FieldType_EnumType_EnumValue
type FieldType_PrimitiveType = src.FieldType_PrimitiveType
type FieldType_PrimitiveType_ = src.FieldType_PrimitiveType_

// Specifications of a single file in Cloud Storage.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type GcsFileSpec = src.GcsFileSpec

// Describes a Cloud Storage fileset entry.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type GcsFilesetSpec = src.GcsFilesetSpec

// Request message for
// [GetEntryGroup][google.cloud.datacatalog.v1beta1.DataCatalog.GetEntryGroup].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type GetEntryGroupRequest = src.GetEntryGroupRequest

// Request message for
// [GetEntry][google.cloud.datacatalog.v1beta1.DataCatalog.GetEntry].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type GetEntryRequest = src.GetEntryRequest

// Request message for
// [GetPolicyTag][google.cloud.datacatalog.v1beta1.PolicyTagManager.GetPolicyTag].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type GetPolicyTagRequest = src.GetPolicyTagRequest

// Request message for
// [GetTagTemplate][google.cloud.datacatalog.v1beta1.DataCatalog.GetTagTemplate].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type GetTagTemplateRequest = src.GetTagTemplateRequest

// Request message for
// [GetTaxonomy][google.cloud.datacatalog.v1beta1.PolicyTagManager.GetTaxonomy].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type GetTaxonomyRequest = src.GetTaxonomyRequest

// Request message for
// [ImportTaxonomies][google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization.ImportTaxonomies].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ImportTaxonomiesRequest = src.ImportTaxonomiesRequest
type ImportTaxonomiesRequest_InlineSource = src.ImportTaxonomiesRequest_InlineSource

// Response message for
// [ImportTaxonomies][google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization.ImportTaxonomies].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ImportTaxonomiesResponse = src.ImportTaxonomiesResponse

// Inline source used for taxonomies import.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type InlineSource = src.InlineSource

// This enum describes all the possible systems that Data Catalog integrates
// with.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type IntegratedSystem = src.IntegratedSystem

// Request message for
// [ListEntries][google.cloud.datacatalog.v1beta1.DataCatalog.ListEntries].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ListEntriesRequest = src.ListEntriesRequest

// Response message for
// [ListEntries][google.cloud.datacatalog.v1beta1.DataCatalog.ListEntries].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ListEntriesResponse = src.ListEntriesResponse

// Request message for
// [ListEntryGroups][google.cloud.datacatalog.v1beta1.DataCatalog.ListEntryGroups].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ListEntryGroupsRequest = src.ListEntryGroupsRequest

// Response message for
// [ListEntryGroups][google.cloud.datacatalog.v1beta1.DataCatalog.ListEntryGroups].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ListEntryGroupsResponse = src.ListEntryGroupsResponse

// Request message for
// [ListPolicyTags][google.cloud.datacatalog.v1beta1.PolicyTagManager.ListPolicyTags].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ListPolicyTagsRequest = src.ListPolicyTagsRequest

// Response message for
// [ListPolicyTags][google.cloud.datacatalog.v1beta1.PolicyTagManager.ListPolicyTags].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ListPolicyTagsResponse = src.ListPolicyTagsResponse

// Request message for
// [ListTags][google.cloud.datacatalog.v1beta1.DataCatalog.ListTags].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ListTagsRequest = src.ListTagsRequest

// Response message for
// [ListTags][google.cloud.datacatalog.v1beta1.DataCatalog.ListTags].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ListTagsResponse = src.ListTagsResponse

// Request message for
// [ListTaxonomies][google.cloud.datacatalog.v1beta1.PolicyTagManager.ListTaxonomies].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ListTaxonomiesRequest = src.ListTaxonomiesRequest

// Response message for
// [ListTaxonomies][google.cloud.datacatalog.v1beta1.PolicyTagManager.ListTaxonomies].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ListTaxonomiesResponse = src.ListTaxonomiesResponse

// Request message for
// [LookupEntry][google.cloud.datacatalog.v1beta1.DataCatalog.LookupEntry].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type LookupEntryRequest = src.LookupEntryRequest
type LookupEntryRequest_LinkedResource = src.LookupEntryRequest_LinkedResource
type LookupEntryRequest_SqlResource = src.LookupEntryRequest_SqlResource

// Denotes one policy tag in a taxonomy (e.g. ssn). Policy Tags can be defined
// in a hierarchy. For example, consider the following hierarchy: Geolocation
// -&gt; (LatLong, City, ZipCode). PolicyTag "Geolocation" contains three child
// policy tags: "LatLong", "City", and "ZipCode".
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type PolicyTag = src.PolicyTag

// PolicyTagManagerClient is the client API for PolicyTagManager service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type PolicyTagManagerClient = src.PolicyTagManagerClient

// PolicyTagManagerSerializationClient is the client API for
// PolicyTagManagerSerialization service. For semantics around ctx use and
// closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type PolicyTagManagerSerializationClient = src.PolicyTagManagerSerializationClient

// PolicyTagManagerSerializationServer is the server API for
// PolicyTagManagerSerialization service.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type PolicyTagManagerSerializationServer = src.PolicyTagManagerSerializationServer

// PolicyTagManagerServer is the server API for PolicyTagManager service.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type PolicyTagManagerServer = src.PolicyTagManagerServer

// Request message for
// [RenameTagTemplateField][google.cloud.datacatalog.v1beta1.DataCatalog.RenameTagTemplateField].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type RenameTagTemplateFieldRequest = src.RenameTagTemplateFieldRequest

// Represents a schema (e.g. BigQuery, GoogleSQL, Avro schema).
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type Schema = src.Schema

// Request message for
// [SearchCatalog][google.cloud.datacatalog.v1beta1.DataCatalog.SearchCatalog].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type SearchCatalogRequest = src.SearchCatalogRequest

// The criteria that select the subspace used for query matching.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type SearchCatalogRequest_Scope = src.SearchCatalogRequest_Scope

// Response message for
// [SearchCatalog][google.cloud.datacatalog.v1beta1.DataCatalog.SearchCatalog].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type SearchCatalogResponse = src.SearchCatalogResponse

// A result that appears in the response of a search request. Each result
// captures details of one entry that matches the search.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type SearchCatalogResult = src.SearchCatalogResult

// The different types of resources that can be returned in search.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type SearchResultType = src.SearchResultType

// Message representing one policy tag when exported as a nested proto.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type SerializedPolicyTag = src.SerializedPolicyTag

// Message capturing a taxonomy and its policy tag hierarchy as a nested
// proto. Used for taxonomy import/export and mutation.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type SerializedTaxonomy = src.SerializedTaxonomy

// Timestamps about this resource according to a particular system.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type SystemTimestamps = src.SystemTimestamps

// Table source type.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type TableSourceType = src.TableSourceType

// Normal BigQuery table spec.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type TableSpec = src.TableSpec

// Tags are used to attach custom metadata to Data Catalog resources. Tags
// conform to the specifications within their tag template. See [Data Catalog
// IAM](https://cloud.google.com/data-catalog/docs/concepts/iam) for
// information on the permissions needed to create or view tags.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type Tag = src.Tag

// Contains the value and supporting information for a field within a
// [Tag][google.cloud.datacatalog.v1beta1.Tag].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type TagField = src.TagField
type TagField_BoolValue = src.TagField_BoolValue
type TagField_DoubleValue = src.TagField_DoubleValue

// Holds an enum value.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type TagField_EnumValue = src.TagField_EnumValue
type TagField_EnumValue_ = src.TagField_EnumValue_
type TagField_StringValue = src.TagField_StringValue
type TagField_TimestampValue = src.TagField_TimestampValue

// A tag template defines a tag, which can have one or more typed fields. The
// template is used to create and attach the tag to GCP resources. [Tag
// template
// roles](https://cloud.google.com/iam/docs/understanding-roles#data-catalog-roles)
// provide permissions to create, edit, and use the template. See, for example,
// the [TagTemplate
// User](https://cloud.google.com/data-catalog/docs/how-to/template-user) role,
// which includes permission to use the tag template to tag resources.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type TagTemplate = src.TagTemplate

// The template for an individual field within a tag template.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type TagTemplateField = src.TagTemplateField
type Tag_Column = src.Tag_Column

// A taxonomy is a collection of policy tags that classify data along a common
// axis. For instance a data *sensitivity* taxonomy could contain policy tags
// denoting PII such as age, zipcode, and SSN. A data *origin* taxonomy could
// contain policy tags to distinguish user data, employee data, partner data,
// public data.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type Taxonomy = src.Taxonomy

// Defines policy types where policy tag can be used for.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type Taxonomy_PolicyType = src.Taxonomy_PolicyType

// UnimplementedDataCatalogServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type UnimplementedDataCatalogServer = src.UnimplementedDataCatalogServer

// UnimplementedPolicyTagManagerSerializationServer can be embedded to have
// forward compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type UnimplementedPolicyTagManagerSerializationServer = src.UnimplementedPolicyTagManagerSerializationServer

// UnimplementedPolicyTagManagerServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type UnimplementedPolicyTagManagerServer = src.UnimplementedPolicyTagManagerServer

// Request message for
// [UpdateEntryGroup][google.cloud.datacatalog.v1beta1.DataCatalog.UpdateEntryGroup].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type UpdateEntryGroupRequest = src.UpdateEntryGroupRequest

// Request message for
// [UpdateEntry][google.cloud.datacatalog.v1beta1.DataCatalog.UpdateEntry].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type UpdateEntryRequest = src.UpdateEntryRequest

// Request message for
// [UpdatePolicyTag][google.cloud.datacatalog.v1beta1.PolicyTagManager.UpdatePolicyTag].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type UpdatePolicyTagRequest = src.UpdatePolicyTagRequest

// Request message for
// [UpdateTag][google.cloud.datacatalog.v1beta1.DataCatalog.UpdateTag].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type UpdateTagRequest = src.UpdateTagRequest

// Request message for
// [UpdateTagTemplateField][google.cloud.datacatalog.v1beta1.DataCatalog.UpdateTagTemplateField].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type UpdateTagTemplateFieldRequest = src.UpdateTagTemplateFieldRequest

// Request message for
// [UpdateTagTemplate][google.cloud.datacatalog.v1beta1.DataCatalog.UpdateTagTemplate].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type UpdateTagTemplateRequest = src.UpdateTagTemplateRequest

// Request message for
// [UpdateTaxonomy][google.cloud.datacatalog.v1beta1.PolicyTagManager.UpdateTaxonomy].
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type UpdateTaxonomyRequest = src.UpdateTaxonomyRequest

// Table view specification.
//
// Deprecated: Please use types in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
type ViewSpec = src.ViewSpec

// Deprecated: Please use funcs in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
func NewDataCatalogClient(cc grpc.ClientConnInterface) DataCatalogClient {
	return src.NewDataCatalogClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
func NewPolicyTagManagerClient(cc grpc.ClientConnInterface) PolicyTagManagerClient {
	return src.NewPolicyTagManagerClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
func NewPolicyTagManagerSerializationClient(cc grpc.ClientConnInterface) PolicyTagManagerSerializationClient {
	return src.NewPolicyTagManagerSerializationClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
func RegisterDataCatalogServer(s *grpc.Server, srv DataCatalogServer) {
	src.RegisterDataCatalogServer(s, srv)
}

// Deprecated: Please use funcs in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
func RegisterPolicyTagManagerSerializationServer(s *grpc.Server, srv PolicyTagManagerSerializationServer) {
	src.RegisterPolicyTagManagerSerializationServer(s, srv)
}

// Deprecated: Please use funcs in: cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb
func RegisterPolicyTagManagerServer(s *grpc.Server, srv PolicyTagManagerServer) {
	src.RegisterPolicyTagManagerServer(s, srv)
}
