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

// Package debugger aliases all exported identifiers in package
// "cloud.google.com/go/debugger/apiv2/debuggerpb".
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package debugger

import (
	src "cloud.google.com/go/debugger/apiv2/debuggerpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/debugger/apiv2/debuggerpb
const (
	Breakpoint_CAPTURE                       = src.Breakpoint_CAPTURE
	Breakpoint_ERROR                         = src.Breakpoint_ERROR
	Breakpoint_INFO                          = src.Breakpoint_INFO
	Breakpoint_LOG                           = src.Breakpoint_LOG
	Breakpoint_WARNING                       = src.Breakpoint_WARNING
	StatusMessage_BREAKPOINT_AGE             = src.StatusMessage_BREAKPOINT_AGE
	StatusMessage_BREAKPOINT_CONDITION       = src.StatusMessage_BREAKPOINT_CONDITION
	StatusMessage_BREAKPOINT_EXPRESSION      = src.StatusMessage_BREAKPOINT_EXPRESSION
	StatusMessage_BREAKPOINT_SOURCE_LOCATION = src.StatusMessage_BREAKPOINT_SOURCE_LOCATION
	StatusMessage_UNSPECIFIED                = src.StatusMessage_UNSPECIFIED
	StatusMessage_VARIABLE_NAME              = src.StatusMessage_VARIABLE_NAME
	StatusMessage_VARIABLE_VALUE             = src.StatusMessage_VARIABLE_VALUE
)

// Deprecated: Please use vars in: cloud.google.com/go/debugger/apiv2/debuggerpb
var (
	Breakpoint_Action_name                                 = src.Breakpoint_Action_name
	Breakpoint_Action_value                                = src.Breakpoint_Action_value
	Breakpoint_LogLevel_name                               = src.Breakpoint_LogLevel_name
	Breakpoint_LogLevel_value                              = src.Breakpoint_LogLevel_value
	File_google_devtools_clouddebugger_v2_controller_proto = src.File_google_devtools_clouddebugger_v2_controller_proto
	File_google_devtools_clouddebugger_v2_data_proto       = src.File_google_devtools_clouddebugger_v2_data_proto
	File_google_devtools_clouddebugger_v2_debugger_proto   = src.File_google_devtools_clouddebugger_v2_debugger_proto
	StatusMessage_Reference_name                           = src.StatusMessage_Reference_name
	StatusMessage_Reference_value                          = src.StatusMessage_Reference_value
)

// Represents the breakpoint specification, status and results.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type Breakpoint = src.Breakpoint

// Actions that can be taken when a breakpoint hits. Agents should reject
// breakpoints with unsupported or unknown action values.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type Breakpoint_Action = src.Breakpoint_Action

// Log severity levels.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type Breakpoint_LogLevel = src.Breakpoint_LogLevel

// Controller2Client is the client API for Controller2 service. For semantics
// around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type Controller2Client = src.Controller2Client

// Controller2Server is the server API for Controller2 service.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type Controller2Server = src.Controller2Server

// Represents the debugged application. The application may include one or
// more replicated processes executing the same code. Each of these processes
// is attached with a debugger agent, carrying out the debugging commands.
// Agents attached to the same debuggee identify themselves as such by using
// exactly the same Debuggee message value when registering.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type Debuggee = src.Debuggee

// Debugger2Client is the client API for Debugger2 service. For semantics
// around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type Debugger2Client = src.Debugger2Client

// Debugger2Server is the server API for Debugger2 service.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type Debugger2Server = src.Debugger2Server

// Request to delete a breakpoint.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type DeleteBreakpointRequest = src.DeleteBreakpointRequest

// Represents a message with parameters.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type FormatMessage = src.FormatMessage

// Request to get breakpoint information.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type GetBreakpointRequest = src.GetBreakpointRequest

// Response for getting breakpoint information.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type GetBreakpointResponse = src.GetBreakpointResponse

// Request to list active breakpoints.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type ListActiveBreakpointsRequest = src.ListActiveBreakpointsRequest

// Response for listing active breakpoints.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type ListActiveBreakpointsResponse = src.ListActiveBreakpointsResponse

// Request to list breakpoints.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type ListBreakpointsRequest = src.ListBreakpointsRequest

// Wrapper message for `Breakpoint.Action`. Defines a filter on the action
// field of breakpoints.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type ListBreakpointsRequest_BreakpointActionValue = src.ListBreakpointsRequest_BreakpointActionValue

// Response for listing breakpoints.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type ListBreakpointsResponse = src.ListBreakpointsResponse

// Request to list debuggees.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type ListDebuggeesRequest = src.ListDebuggeesRequest

// Response for listing debuggees.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type ListDebuggeesResponse = src.ListDebuggeesResponse

// Request to register a debuggee.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type RegisterDebuggeeRequest = src.RegisterDebuggeeRequest

// Response for registering a debuggee.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type RegisterDebuggeeResponse = src.RegisterDebuggeeResponse

// Request to set a breakpoint
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type SetBreakpointRequest = src.SetBreakpointRequest

// Response for setting a breakpoint.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type SetBreakpointResponse = src.SetBreakpointResponse

// Represents a location in the source code.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type SourceLocation = src.SourceLocation

// Represents a stack frame context.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type StackFrame = src.StackFrame

// Represents a contextual status message. The message can indicate an error
// or informational status, and refer to specific parts of the containing
// object. For example, the `Breakpoint.status` field can indicate an error
// referring to the `BREAKPOINT_SOURCE_LOCATION` with the message `Location not
// found`.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type StatusMessage = src.StatusMessage

// Enumerates references to which the message applies.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type StatusMessage_Reference = src.StatusMessage_Reference

// UnimplementedController2Server can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type UnimplementedController2Server = src.UnimplementedController2Server

// UnimplementedDebugger2Server can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type UnimplementedDebugger2Server = src.UnimplementedDebugger2Server

// Request to update an active breakpoint.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type UpdateActiveBreakpointRequest = src.UpdateActiveBreakpointRequest

// Response for updating an active breakpoint. The message is defined to allow
// future extensions.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type UpdateActiveBreakpointResponse = src.UpdateActiveBreakpointResponse

// Represents a variable or an argument possibly of a compound object type.
// Note how the following variables are represented: 1) A simple variable: int
// x = 5 { name: "x", value: "5", type: "int" } // Captured variable 2) A
// compound object: struct T { int m1; int m2; }; T x = { 3, 7 }; { // Captured
// variable name: "x", type: "T", members { name: "m1", value: "3", type: "int"
// }, members { name: "m2", value: "7", type: "int" } } 3) A pointer where the
// pointee was captured: T x = { 3, 7 }; T* p = &x; { // Captured variable
// name: "p", type: "T*", value: "0x00500500", members { name: "m1", value:
// "3", type: "int" }, members { name: "m2", value: "7", type: "int" } } 4) A
// pointer where the pointee was not captured: T* p = new T; { // Captured
// variable name: "p", type: "T*", value: "0x00400400" status { is_error: true,
// description { format: "unavailable" } } } The status should describe the
// reason for the missing value, such as `<optimized out>`, `<inaccessible>`,
// `<pointers limit reached>`. Note that a null pointer should not have
// members. 5) An unnamed value: int* p = new int(7); { // Captured variable
// name: "p", value: "0x00500500", type: "int*", members { value: "7", type:
// "int" } } 6) An unnamed pointer where the pointee was not captured: int* p =
// new int(7); int** pp = &p; { // Captured variable name: "pp", value:
// "0x00500500", type: "int**", members { value: "0x00400400", type: "int*"
// status { is_error: true, description: { format: "unavailable" } } } } } To
// optimize computation, memory and network traffic, variables that repeat in
// the output multiple times can be stored once in a shared variable table and
// be referenced using the `var_table_index` field. The variables stored in the
// shared table are nameless and are essentially a partition of the complete
// variable. To reconstruct the complete variable, merge the referencing
// variable with the referenced variable. When using the shared variable table,
// the following variables: T x = { 3, 7 }; T* p = &x; T& r = x; { name: "x",
// var_table_index: 3, type: "T" } // Captured variables { name: "p", value
// "0x00500500", type="T*", var_table_index: 3 } { name: "r", type="T&",
// var_table_index: 3 } { // Shared variable table entry #3: members { name:
// "m1", value: "3", type: "int" }, members { name: "m2", value: "7", type:
// "int" } } Note that the pointer address is stored with the referencing
// variable and not with the referenced variable. This allows the referenced
// variable to be shared between pointers and references. The type field is
// optional. The debugger agent may or may not support it.
//
// Deprecated: Please use types in: cloud.google.com/go/debugger/apiv2/debuggerpb
type Variable = src.Variable

// Deprecated: Please use funcs in: cloud.google.com/go/debugger/apiv2/debuggerpb
func NewController2Client(cc grpc.ClientConnInterface) Controller2Client {
	return src.NewController2Client(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/debugger/apiv2/debuggerpb
func NewDebugger2Client(cc grpc.ClientConnInterface) Debugger2Client {
	return src.NewDebugger2Client(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/debugger/apiv2/debuggerpb
func RegisterController2Server(s *grpc.Server, srv Controller2Server) {
	src.RegisterController2Server(s, srv)
}

// Deprecated: Please use funcs in: cloud.google.com/go/debugger/apiv2/debuggerpb
func RegisterDebugger2Server(s *grpc.Server, srv Debugger2Server) {
	src.RegisterDebugger2Server(s, srv)
}
