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

// Package cloudtasks aliases all exported identifiers in package
// "cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb".
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package cloudtasks

import (
	src "cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
const (
	HttpMethod_DELETE                  = src.HttpMethod_DELETE
	HttpMethod_GET                     = src.HttpMethod_GET
	HttpMethod_HEAD                    = src.HttpMethod_HEAD
	HttpMethod_HTTP_METHOD_UNSPECIFIED = src.HttpMethod_HTTP_METHOD_UNSPECIFIED
	HttpMethod_OPTIONS                 = src.HttpMethod_OPTIONS
	HttpMethod_PATCH                   = src.HttpMethod_PATCH
	HttpMethod_POST                    = src.HttpMethod_POST
	HttpMethod_PUT                     = src.HttpMethod_PUT
	Queue_DISABLED                     = src.Queue_DISABLED
	Queue_PAUSED                       = src.Queue_PAUSED
	Queue_PULL                         = src.Queue_PULL
	Queue_PUSH                         = src.Queue_PUSH
	Queue_RUNNING                      = src.Queue_RUNNING
	Queue_STATE_UNSPECIFIED            = src.Queue_STATE_UNSPECIFIED
	Queue_TYPE_UNSPECIFIED             = src.Queue_TYPE_UNSPECIFIED
	Task_BASIC                         = src.Task_BASIC
	Task_FULL                          = src.Task_FULL
	Task_VIEW_UNSPECIFIED              = src.Task_VIEW_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
var (
	File_google_cloud_tasks_v2beta3_cloudtasks_proto = src.File_google_cloud_tasks_v2beta3_cloudtasks_proto
	File_google_cloud_tasks_v2beta3_queue_proto      = src.File_google_cloud_tasks_v2beta3_queue_proto
	File_google_cloud_tasks_v2beta3_target_proto     = src.File_google_cloud_tasks_v2beta3_target_proto
	File_google_cloud_tasks_v2beta3_task_proto       = src.File_google_cloud_tasks_v2beta3_task_proto
	HttpMethod_name                                  = src.HttpMethod_name
	HttpMethod_value                                 = src.HttpMethod_value
	Queue_State_name                                 = src.Queue_State_name
	Queue_State_value                                = src.Queue_State_value
	Queue_Type_name                                  = src.Queue_Type_name
	Queue_Type_value                                 = src.Queue_Type_value
	Task_View_name                                   = src.Task_View_name
	Task_View_value                                  = src.Task_View_value
)

// App Engine HTTP queue. The task will be delivered to the App Engine
// application hostname specified by its
// [AppEngineHttpQueue][google.cloud.tasks.v2beta3.AppEngineHttpQueue] and
// [AppEngineHttpRequest][google.cloud.tasks.v2beta3.AppEngineHttpRequest]. The
// documentation for
// [AppEngineHttpRequest][google.cloud.tasks.v2beta3.AppEngineHttpRequest]
// explains how the task's host URL is constructed. Using
// [AppEngineHttpQueue][google.cloud.tasks.v2beta3.AppEngineHttpQueue] requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project and the following scope:
// `https://www.googleapis.com/auth/cloud-platform`
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type AppEngineHttpQueue = src.AppEngineHttpQueue

// App Engine HTTP request. The message defines the HTTP request that is sent
// to an App Engine app when the task is dispatched. Using
// [AppEngineHttpRequest][google.cloud.tasks.v2beta3.AppEngineHttpRequest]
// requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project and the following scope:
// `https://www.googleapis.com/auth/cloud-platform` The task will be delivered
// to the App Engine app which belongs to the same project as the queue. For
// more information, see [How Requests are
// Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
// and how routing is affected by [dispatch
// files](https://cloud.google.com/appengine/docs/python/config/dispatchref).
// Traffic is encrypted during transport and never leaves Google datacenters.
// Because this traffic is carried over a communication mechanism internal to
// Google, you cannot explicitly set the protocol (for example, HTTP or HTTPS).
// The request to the handler, however, will appear to have used the HTTP
// protocol. The
// [AppEngineRouting][google.cloud.tasks.v2beta3.AppEngineRouting] used to
// construct the URL that the task is delivered to can be set at the
// queue-level or task-level: - If set,
// [app_engine_routing_override][google.cloud.tasks.v2beta3.AppEngineHttpQueue.app_engine_routing_override]
// is used for all tasks in the queue, no matter what the setting is for the
// [task-level
// app_engine_routing][google.cloud.tasks.v2beta3.AppEngineHttpRequest.app_engine_routing].
// The `url` that the task will be sent to is: - `url =`
// [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] `+`
// [relative_uri][google.cloud.tasks.v2beta3.AppEngineHttpRequest.relative_uri]
// Tasks can be dispatched to secure app handlers, unsecure app handlers, and
// URIs restricted with [`login:
// admin`](https://cloud.google.com/appengine/docs/standard/python/config/appref).
// Because tasks are not run as any user, they cannot be dispatched to URIs
// restricted with [`login:
// required`](https://cloud.google.com/appengine/docs/standard/python/config/appref)
// Task dispatches also do not follow redirects. The task attempt has succeeded
// if the app's request handler returns an HTTP response code in the range
// [`200` - `299`]. The task attempt has failed if the app's handler returns a
// non-2xx response code or Cloud Tasks does not receive response before the
// [deadline][google.cloud.tasks.v2beta3.Task.dispatch_deadline]. Failed tasks
// will be retried according to the [retry
// configuration][google.cloud.tasks.v2beta3.Queue.retry_config]. `503`
// (Service Unavailable) is considered an App Engine system error instead of an
// application error and will cause Cloud Tasks' traffic congestion control to
// temporarily throttle the queue's dispatches. Unlike other types of task
// targets, a `429` (Too Many Requests) response from an app handler does not
// cause traffic congestion control to throttle the queue.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type AppEngineHttpRequest = src.AppEngineHttpRequest

// App Engine Routing. Defines routing characteristics specific to App Engine
// - service, version, and instance. For more information about services,
// versions, and instances see [An Overview of App
// Engine](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine),
// [Microservices Architecture on Google App
// Engine](https://cloud.google.com/appengine/docs/python/microservices-on-app-engine),
// [App Engine Standard request
// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed),
// and [App Engine Flex request
// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type AppEngineRouting = src.AppEngineRouting

// The status of a task attempt.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type Attempt = src.Attempt

// CloudTasksClient is the client API for CloudTasks service. For semantics
// around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type CloudTasksClient = src.CloudTasksClient

// CloudTasksServer is the server API for CloudTasks service.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type CloudTasksServer = src.CloudTasksServer

// Request message for
// [CreateQueue][google.cloud.tasks.v2beta3.CloudTasks.CreateQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type CreateQueueRequest = src.CreateQueueRequest

// Request message for
// [CreateTask][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type CreateTaskRequest = src.CreateTaskRequest

// Request message for
// [DeleteQueue][google.cloud.tasks.v2beta3.CloudTasks.DeleteQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type DeleteQueueRequest = src.DeleteQueueRequest

// Request message for deleting a task using
// [DeleteTask][google.cloud.tasks.v2beta3.CloudTasks.DeleteTask].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type DeleteTaskRequest = src.DeleteTaskRequest

// Request message for
// [GetQueue][google.cloud.tasks.v2beta3.CloudTasks.GetQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type GetQueueRequest = src.GetQueueRequest

// Request message for getting a task using
// [GetTask][google.cloud.tasks.v2beta3.CloudTasks.GetTask].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type GetTaskRequest = src.GetTaskRequest

// The HTTP method used to execute the task.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type HttpMethod = src.HttpMethod

// HTTP request. The task will be pushed to the worker as an HTTP request. If
// the worker or the redirected worker acknowledges the task by returning a
// successful HTTP response code ([`200` - `299`]), the task will be removed
// from the queue. If any other HTTP response code is returned or no response
// is received, the task will be retried according to the following: -
// User-specified throttling: [retry
// configuration][google.cloud.tasks.v2beta3.Queue.retry_config], [rate
// limits][google.cloud.tasks.v2beta3.Queue.rate_limits], and the [queue's
// state][google.cloud.tasks.v2beta3.Queue.state]. - System throttling: To
// prevent the worker from overloading, Cloud Tasks may temporarily reduce the
// queue's effective rate. User-specified settings will not be changed. System
// throttling happens because: - Cloud Tasks backs off on all errors. Normally
// the backoff specified in [rate
// limits][google.cloud.tasks.v2beta3.Queue.rate_limits] will be used. But if
// the worker returns `429` (Too Many Requests), `503` (Service Unavailable),
// or the rate of errors is high, Cloud Tasks will use a higher backoff rate.
// The retry specified in the `Retry-After` HTTP response header is considered.
// - To prevent traffic spikes and to smooth sudden increases in traffic,
// dispatches ramp up slowly when the queue is newly created or idle and if
// large numbers of tasks suddenly become available to dispatch (due to spikes
// in create task rates, the queue being unpaused, or many tasks that are
// scheduled at the same time).
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type HttpRequest = src.HttpRequest
type HttpRequest_OauthToken = src.HttpRequest_OauthToken
type HttpRequest_OidcToken = src.HttpRequest_OidcToken

// Request message for
// [ListQueues][google.cloud.tasks.v2beta3.CloudTasks.ListQueues].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type ListQueuesRequest = src.ListQueuesRequest

// Response message for
// [ListQueues][google.cloud.tasks.v2beta3.CloudTasks.ListQueues].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type ListQueuesResponse = src.ListQueuesResponse

// Request message for listing tasks using
// [ListTasks][google.cloud.tasks.v2beta3.CloudTasks.ListTasks].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type ListTasksRequest = src.ListTasksRequest

// Response message for listing tasks using
// [ListTasks][google.cloud.tasks.v2beta3.CloudTasks.ListTasks].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type ListTasksResponse = src.ListTasksResponse

// Contains information needed for generating an [OAuth
// token](https://developers.google.com/identity/protocols/OAuth2). This type
// of authorization should generally only be used when calling Google APIs
// hosted on *.googleapis.com.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type OAuthToken = src.OAuthToken

// Contains information needed for generating an [OpenID Connect
// token](https://developers.google.com/identity/protocols/OpenIDConnect). This
// type of authorization can be used for many scenarios, including calling
// Cloud Run, or endpoints where you intend to validate the token yourself.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type OidcToken = src.OidcToken

// Request message for
// [PauseQueue][google.cloud.tasks.v2beta3.CloudTasks.PauseQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type PauseQueueRequest = src.PauseQueueRequest

// Pull Message. This proto can only be used for tasks in a queue which has
// [PULL][google.cloud.tasks.v2beta3.Queue.type] type. It currently exists for
// backwards compatibility with the App Engine Task Queue SDK. This message
// type maybe returned with methods
// [list][google.cloud.tasks.v2beta3.CloudTask.ListTasks] and
// [get][google.cloud.tasks.v2beta3.CloudTask.ListTasks], when the response
// view is [FULL][google.cloud.tasks.v2beta3.Task.View.Full].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type PullMessage = src.PullMessage

// Request message for
// [PurgeQueue][google.cloud.tasks.v2beta3.CloudTasks.PurgeQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type PurgeQueueRequest = src.PurgeQueueRequest

// A queue is a container of related tasks. Queues are configured to manage
// how those tasks are dispatched. Configurable properties include rate limits,
// retry options, queue types, and others.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type Queue = src.Queue

// Statistics for a queue.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type QueueStats = src.QueueStats
type Queue_AppEngineHttpQueue = src.Queue_AppEngineHttpQueue

// State of the queue.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type Queue_State = src.Queue_State

// The type of the queue.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type Queue_Type = src.Queue_Type

// Rate limits. This message determines the maximum rate that tasks can be
// dispatched by a queue, regardless of whether the dispatch is a first task
// attempt or a retry. Note: The debugging command,
// [RunTask][google.cloud.tasks.v2beta3.CloudTasks.RunTask], will run a task
// even if the queue has reached its
// [RateLimits][google.cloud.tasks.v2beta3.RateLimits].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type RateLimits = src.RateLimits

// Request message for
// [ResumeQueue][google.cloud.tasks.v2beta3.CloudTasks.ResumeQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type ResumeQueueRequest = src.ResumeQueueRequest

// Retry config. These settings determine when a failed task attempt is
// retried.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type RetryConfig = src.RetryConfig

// Request message for forcing a task to run now using
// [RunTask][google.cloud.tasks.v2beta3.CloudTasks.RunTask].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type RunTaskRequest = src.RunTaskRequest

// Configuration options for writing logs to [Stackdriver
// Logging](https://cloud.google.com/logging/docs/).
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type StackdriverLoggingConfig = src.StackdriverLoggingConfig

// A unit of scheduled work.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type Task = src.Task
type Task_AppEngineHttpRequest = src.Task_AppEngineHttpRequest
type Task_HttpRequest = src.Task_HttpRequest
type Task_PullMessage = src.Task_PullMessage

// The view specifies a subset of [Task][google.cloud.tasks.v2beta3.Task]
// data. When a task is returned in a response, not all information is
// retrieved by default because some data, such as payloads, might be desirable
// to return only when needed because of its large size or because of the
// sensitivity of data that it contains.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type Task_View = src.Task_View

// UnimplementedCloudTasksServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type UnimplementedCloudTasksServer = src.UnimplementedCloudTasksServer

// Request message for
// [UpdateQueue][google.cloud.tasks.v2beta3.CloudTasks.UpdateQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
type UpdateQueueRequest = src.UpdateQueueRequest

// Deprecated: Please use funcs in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
func NewCloudTasksClient(cc grpc.ClientConnInterface) CloudTasksClient {
	return src.NewCloudTasksClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb
func RegisterCloudTasksServer(s *grpc.Server, srv CloudTasksServer) {
	src.RegisterCloudTasksServer(s, srv)
}
