// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/protobuf/compiler/plugin.proto

package plugin_go

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	descriptor "github.com/cosmos/gogoproto/protoc-gen-gogo/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Sync with code_generator.h.
type CodeGeneratorResponse_Feature int32

const (
	CodeGeneratorResponse_FEATURE_NONE            CodeGeneratorResponse_Feature = 0
	CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL CodeGeneratorResponse_Feature = 1
)

var CodeGeneratorResponse_Feature_name = map[int32]string{
	0: "FEATURE_NONE",
	1: "FEATURE_PROTO3_OPTIONAL",
}

var CodeGeneratorResponse_Feature_value = map[string]int32{
	"FEATURE_NONE":            0,
	"FEATURE_PROTO3_OPTIONAL": 1,
}

func (x CodeGeneratorResponse_Feature) Enum() *CodeGeneratorResponse_Feature {
	p := new(CodeGeneratorResponse_Feature)
	*p = x
	return p
}

func (x CodeGeneratorResponse_Feature) String() string {
	return proto.EnumName(CodeGeneratorResponse_Feature_name, int32(x))
}

func (x *CodeGeneratorResponse_Feature) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CodeGeneratorResponse_Feature_value, data, "CodeGeneratorResponse_Feature")
	if err != nil {
		return err
	}
	*x = CodeGeneratorResponse_Feature(value)
	return nil
}

func (CodeGeneratorResponse_Feature) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3562add825dafed5, []int{2, 0}
}

// The version number of protocol compiler.
type Version struct {
	Major *int32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor *int32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Patch *int32 `protobuf:"varint,3,opt,name=patch" json:"patch,omitempty"`
	// A suffix for alpha, beta or rc release, e.g., "alpha-1", "rc2". It should
	// be empty for mainline stable releases.
	Suffix               *string  `protobuf:"bytes,4,opt,name=suffix" json:"suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_3562add825dafed5, []int{0}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() int32 {
	if m != nil && m.Major != nil {
		return *m.Major
	}
	return 0
}

func (m *Version) GetMinor() int32 {
	if m != nil && m.Minor != nil {
		return *m.Minor
	}
	return 0
}

func (m *Version) GetPatch() int32 {
	if m != nil && m.Patch != nil {
		return *m.Patch
	}
	return 0
}

func (m *Version) GetSuffix() string {
	if m != nil && m.Suffix != nil {
		return *m.Suffix
	}
	return ""
}

// An encoded CodeGeneratorRequest is written to the plugin's stdin.
type CodeGeneratorRequest struct {
	// The .proto files that were explicitly listed on the command-line.  The
	// code generator should generate code only for these files.  Each file's
	// descriptor will be included in proto_file, below.
	FileToGenerate []string `protobuf:"bytes,1,rep,name=file_to_generate,json=fileToGenerate" json:"file_to_generate,omitempty"`
	// The generator parameter passed on the command-line.
	Parameter *string `protobuf:"bytes,2,opt,name=parameter" json:"parameter,omitempty"`
	// FileDescriptorProtos for all files in files_to_generate and everything
	// they import.  The files will appear in topological order, so each file
	// appears before any file that imports it.
	//
	// protoc guarantees that all proto_files will be written after
	// the fields above, even though this is not technically guaranteed by the
	// protobuf wire format.  This theoretically could allow a plugin to stream
	// in the FileDescriptorProtos and handle them one by one rather than read
	// the entire set into memory at once.  However, as of this writing, this
	// is not similarly optimized on protoc's end -- it will store all fields in
	// memory at once before sending them to the plugin.
	//
	// Type names of fields and extensions in the FileDescriptorProto are always
	// fully qualified.
	ProtoFile []*descriptor.FileDescriptorProto `protobuf:"bytes,15,rep,name=proto_file,json=protoFile" json:"proto_file,omitempty"`
	// The version number of protocol compiler.
	CompilerVersion      *Version `protobuf:"bytes,3,opt,name=compiler_version,json=compilerVersion" json:"compiler_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeGeneratorRequest) Reset()         { *m = CodeGeneratorRequest{} }
func (m *CodeGeneratorRequest) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorRequest) ProtoMessage()    {}
func (*CodeGeneratorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3562add825dafed5, []int{1}
}
func (m *CodeGeneratorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorRequest.Unmarshal(m, b)
}
func (m *CodeGeneratorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorRequest.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorRequest.Merge(m, src)
}
func (m *CodeGeneratorRequest) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorRequest.Size(m)
}
func (m *CodeGeneratorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorRequest proto.InternalMessageInfo

func (m *CodeGeneratorRequest) GetFileToGenerate() []string {
	if m != nil {
		return m.FileToGenerate
	}
	return nil
}

func (m *CodeGeneratorRequest) GetParameter() string {
	if m != nil && m.Parameter != nil {
		return *m.Parameter
	}
	return ""
}

func (m *CodeGeneratorRequest) GetProtoFile() []*descriptor.FileDescriptorProto {
	if m != nil {
		return m.ProtoFile
	}
	return nil
}

func (m *CodeGeneratorRequest) GetCompilerVersion() *Version {
	if m != nil {
		return m.CompilerVersion
	}
	return nil
}

// The plugin writes an encoded CodeGeneratorResponse to stdout.
type CodeGeneratorResponse struct {
	// Error message.  If non-empty, code generation failed.  The plugin process
	// should exit with status code zero even if it reports an error in this way.
	//
	// This should be used to indicate errors in .proto files which prevent the
	// code generator from generating correct code.  Errors which indicate a
	// problem in protoc itself -- such as the input CodeGeneratorRequest being
	// unparseable -- should be reported by writing a message to stderr and
	// exiting with a non-zero status code.
	Error *string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	// A bitmask of supported features that the code generator supports.
	// This is a bitwise "or" of values from the Feature enum.
	SupportedFeatures    *uint64                       `protobuf:"varint,2,opt,name=supported_features,json=supportedFeatures" json:"supported_features,omitempty"`
	File                 []*CodeGeneratorResponse_File `protobuf:"bytes,15,rep,name=file" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CodeGeneratorResponse) Reset()         { *m = CodeGeneratorResponse{} }
func (m *CodeGeneratorResponse) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorResponse) ProtoMessage()    {}
func (*CodeGeneratorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3562add825dafed5, []int{2}
}
func (m *CodeGeneratorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorResponse.Unmarshal(m, b)
}
func (m *CodeGeneratorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorResponse.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorResponse.Merge(m, src)
}
func (m *CodeGeneratorResponse) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorResponse.Size(m)
}
func (m *CodeGeneratorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorResponse proto.InternalMessageInfo

func (m *CodeGeneratorResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *CodeGeneratorResponse) GetSupportedFeatures() uint64 {
	if m != nil && m.SupportedFeatures != nil {
		return *m.SupportedFeatures
	}
	return 0
}

func (m *CodeGeneratorResponse) GetFile() []*CodeGeneratorResponse_File {
	if m != nil {
		return m.File
	}
	return nil
}

// Represents a single generated file.
type CodeGeneratorResponse_File struct {
	// The file name, relative to the output directory.  The name must not
	// contain "." or ".." components and must be relative, not be absolute (so,
	// the file cannot lie outside the output directory).  "/" must be used as
	// the path separator, not "\".
	//
	// If the name is omitted, the content will be appended to the previous
	// file.  This allows the generator to break large files into small chunks,
	// and allows the generated text to be streamed back to protoc so that large
	// files need not reside completely in memory at one time.  Note that as of
	// this writing protoc does not optimize for this -- it will read the entire
	// CodeGeneratorResponse before writing files to disk.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// If non-empty, indicates that the named file should already exist, and the
	// content here is to be inserted into that file at a defined insertion
	// point.  This feature allows a code generator to extend the output
	// produced by another code generator.  The original generator may provide
	// insertion points by placing special annotations in the file that look
	// like:
	//
	//	@@protoc_insertion_point(NAME)
	//
	// The annotation can have arbitrary text before and after it on the line,
	// which allows it to be placed in a comment.  NAME should be replaced with
	// an identifier naming the point -- this is what other generators will use
	// as the insertion_point.  Code inserted at this point will be placed
	// immediately above the line containing the insertion point (thus multiple
	// insertions to the same point will come out in the order they were added).
	// The double-@ is intended to make it unlikely that the generated code
	// could contain things that look like insertion points by accident.
	//
	// For example, the C++ code generator places the following line in the
	// .pb.h files that it generates:
	//
	//	// @@protoc_insertion_point(namespace_scope)
	//
	// This line appears within the scope of the file's package namespace, but
	// outside of any particular class.  Another plugin can then specify the
	// insertion_point "namespace_scope" to generate additional classes or
	// other declarations that should be placed in this scope.
	//
	// Note that if the line containing the insertion point begins with
	// whitespace, the same whitespace will be added to every line of the
	// inserted text.  This is useful for languages like Python, where
	// indentation matters.  In these languages, the insertion point comment
	// should be indented the same amount as any inserted code will need to be
	// in order to work correctly in that context.
	//
	// The code generator that generates the initial file and the one which
	// inserts into it must both run as part of a single invocation of protoc.
	// Code generators are executed in the order in which they appear on the
	// command line.
	//
	// If |insertion_point| is present, |name| must also be present.
	InsertionPoint *string `protobuf:"bytes,2,opt,name=insertion_point,json=insertionPoint" json:"insertion_point,omitempty"`
	// The file contents.
	Content *string `protobuf:"bytes,15,opt,name=content" json:"content,omitempty"`
	// Information describing the file content being inserted. If an insertion
	// point is used, this information will be appropriately offset and inserted
	// into the code generation metadata for the generated files.
	GeneratedCodeInfo    *descriptor.GeneratedCodeInfo `protobuf:"bytes,16,opt,name=generated_code_info,json=generatedCodeInfo" json:"generated_code_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CodeGeneratorResponse_File) Reset()         { *m = CodeGeneratorResponse_File{} }
func (m *CodeGeneratorResponse_File) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorResponse_File) ProtoMessage()    {}
func (*CodeGeneratorResponse_File) Descriptor() ([]byte, []int) {
	return fileDescriptor_3562add825dafed5, []int{2, 0}
}
func (m *CodeGeneratorResponse_File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorResponse_File.Unmarshal(m, b)
}
func (m *CodeGeneratorResponse_File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorResponse_File.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorResponse_File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorResponse_File.Merge(m, src)
}
func (m *CodeGeneratorResponse_File) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorResponse_File.Size(m)
}
func (m *CodeGeneratorResponse_File) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorResponse_File.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorResponse_File proto.InternalMessageInfo

func (m *CodeGeneratorResponse_File) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetInsertionPoint() string {
	if m != nil && m.InsertionPoint != nil {
		return *m.InsertionPoint
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetGeneratedCodeInfo() *descriptor.GeneratedCodeInfo {
	if m != nil {
		return m.GeneratedCodeInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.protobuf.compiler.CodeGeneratorResponse_Feature", CodeGeneratorResponse_Feature_name, CodeGeneratorResponse_Feature_value)
	proto.RegisterType((*Version)(nil), "google.protobuf.compiler.Version")
	proto.RegisterType((*CodeGeneratorRequest)(nil), "google.protobuf.compiler.CodeGeneratorRequest")
	proto.RegisterType((*CodeGeneratorResponse)(nil), "google.protobuf.compiler.CodeGeneratorResponse")
	proto.RegisterType((*CodeGeneratorResponse_File)(nil), "google.protobuf.compiler.CodeGeneratorResponse.File")
}

func init() {
	proto.RegisterFile("google/protobuf/compiler/plugin.proto", fileDescriptor_3562add825dafed5)
}

var fileDescriptor_3562add825dafed5 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x80, 0x7f, 0xff, 0x49, 0xa9, 0x3c, 0xad, 0x1a, 0x77, 0x29, 0x60, 0x95, 0x1e, 0x42, 0x04,
	0x22, 0x1c, 0x70, 0xa5, 0xc2, 0x81, 0x6b, 0x5b, 0x12, 0xa8, 0x54, 0x25, 0xd1, 0x2a, 0x80, 0xc4,
	0xc5, 0x72, 0xed, 0xb1, 0x59, 0xe4, 0xec, 0x2c, 0xeb, 0x35, 0x82, 0xf7, 0xe0, 0x45, 0x78, 0x34,
	0xde, 0x00, 0x79, 0xbd, 0x4e, 0x51, 0x20, 0x37, 0xcf, 0x37, 0xb3, 0xab, 0x99, 0x6f, 0x3d, 0xf0,
	0xa4, 0x20, 0x2a, 0x4a, 0x3c, 0x55, 0x9a, 0x0c, 0xdd, 0xd4, 0xf9, 0x69, 0x4a, 0x2b, 0x25, 0x4a,
	0xd4, 0xa7, 0xaa, 0xac, 0x0b, 0x21, 0x23, 0x9b, 0x60, 0x61, 0x5b, 0x16, 0x75, 0x65, 0x51, 0x57,
	0x76, 0x3c, 0xdc, 0xbc, 0x20, 0xc3, 0x2a, 0xd5, 0x42, 0x19, 0xd2, 0x6d, 0xf5, 0x28, 0x85, 0xdd,
	0xf7, 0xa8, 0x2b, 0x41, 0x92, 0x1d, 0xc1, 0xce, 0x2a, 0xf9, 0x4c, 0x3a, 0xf4, 0x86, 0xde, 0x78,
	0x87, 0xb7, 0x81, 0xa5, 0x42, 0x92, 0x0e, 0xff, 0x77, 0xb4, 0x09, 0x1a, 0xaa, 0x12, 0x93, 0x7e,
	0x0a, 0x7b, 0x2d, 0xb5, 0x01, 0xbb, 0x0f, 0x77, 0xaa, 0x3a, 0xcf, 0xc5, 0xb7, 0xb0, 0x3f, 0xf4,
	0xc6, 0x3e, 0x77, 0xd1, 0xe8, 0x97, 0x07, 0x47, 0x97, 0x94, 0xe1, 0x1b, 0x94, 0xa8, 0x13, 0x43,
	0x9a, 0xe3, 0x97, 0x1a, 0x2b, 0xc3, 0xc6, 0x10, 0xe4, 0xa2, 0xc4, 0xd8, 0x50, 0x5c, 0xb4, 0x39,
	0x0c, 0xbd, 0x61, 0x6f, 0xec, 0xf3, 0x83, 0x86, 0x2f, 0xc9, 0x9d, 0x40, 0x76, 0x02, 0xbe, 0x4a,
	0x74, 0xb2, 0x42, 0x83, 0x6d, 0x2b, 0x3e, 0xbf, 0x05, 0xec, 0x12, 0xc0, 0x8e, 0x13, 0x37, 0xa7,
	0xc2, 0xc1, 0xb0, 0x37, 0xde, 0x3b, 0x7b, 0x1c, 0x6d, 0x6a, 0x99, 0x8a, 0x12, 0x5f, 0xaf, 0x05,
	0x2c, 0x1a, 0xcc, 0x7d, 0x9b, 0x6d, 0x32, 0xec, 0x1a, 0x82, 0x4e, 0x5c, 0xfc, 0xb5, 0x75, 0x62,
	0xc7, 0xdb, 0x3b, 0x7b, 0x14, 0x6d, 0x33, 0x1c, 0x39, 0x79, 0x7c, 0xd0, 0x11, 0x07, 0x46, 0x3f,
	0x7a, 0x70, 0x6f, 0x63, 0xe6, 0x4a, 0x91, 0xac, 0xb0, 0x71, 0x87, 0x5a, 0x3b, 0xcf, 0x3e, 0x6f,
	0x03, 0xf6, 0x1c, 0x58, 0x55, 0x2b, 0x45, 0xda, 0x60, 0x16, 0xe7, 0x98, 0x98, 0x5a, 0x63, 0x65,
	0x27, 0xed, 0xf3, 0xc3, 0x75, 0x66, 0xea, 0x12, 0xec, 0x2d, 0xf4, 0xff, 0x98, 0xf5, 0xe5, 0xf6,
	0x06, 0xff, 0xd9, 0x83, 0x55, 0xc1, 0xed, 0x0d, 0xc7, 0x3f, 0x3d, 0xe8, 0xdb, 0xf9, 0x19, 0xf4,
	0x65, 0xb2, 0x42, 0xd7, 0x96, 0xfd, 0x66, 0x4f, 0x61, 0x20, 0x64, 0x85, 0xda, 0x08, 0x92, 0xb1,
	0x22, 0x21, 0x8d, 0x93, 0x7f, 0xb0, 0xc6, 0x8b, 0x86, 0xb2, 0x10, 0x76, 0x53, 0x92, 0x06, 0xa5,
	0x09, 0x07, 0xb6, 0xa0, 0x0b, 0x19, 0x87, 0xbb, 0xdd, 0xdb, 0x66, 0x71, 0x4a, 0x19, 0xc6, 0x42,
	0xe6, 0x14, 0x06, 0xd6, 0xec, 0xe8, 0xaf, 0xc6, 0xbb, 0x17, 0xcf, 0x9a, 0xc6, 0xaf, 0x64, 0x4e,
	0xfc, 0xb0, 0xd8, 0x44, 0xa3, 0x57, 0xb0, 0xeb, 0x4c, 0xb0, 0x00, 0xf6, 0xa7, 0x93, 0xf3, 0xe5,
	0x3b, 0x3e, 0x89, 0x67, 0xf3, 0xd9, 0x24, 0xf8, 0x8f, 0x3d, 0x84, 0x07, 0x1d, 0x59, 0xf0, 0xf9,
	0x72, 0xfe, 0x22, 0x9e, 0x2f, 0x96, 0x57, 0xf3, 0xd9, 0xf9, 0x75, 0xe0, 0x5d, 0x7c, 0x80, 0x93,
	0x94, 0x56, 0x5b, 0x75, 0x5d, 0xec, 0x2f, 0xec, 0x66, 0xd9, 0x9f, 0xa3, 0xfa, 0xf8, 0xcc, 0xd5,
	0x15, 0x54, 0x26, 0xb2, 0x88, 0x48, 0x17, 0xb7, 0xab, 0x64, 0xbe, 0x2b, 0xac, 0xdc, 0x22, 0xaa,
	0x9b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x49, 0x12, 0x8f, 0x4e, 0xac, 0x03, 0x00, 0x00,
}
