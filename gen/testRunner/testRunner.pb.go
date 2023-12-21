// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: testRunner/testRunner.proto

package transformation

import (
	graph "github.com/MyWeHub/plugin-sdk/gen/graph"
	pluginrunner "github.com/MyWeHub/plugin-sdk/gen/pluginrunner"
	_ "github.com/MyWeHub/plugin-sdk/gen/schema"
	_ "github.com/amsokol/mongo-go-driver-protobuf/pmongo"
	_ "github.com/amsokol/protoc-gen-gotag/tagger"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RunTestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*RunTestData_StartRunTest
	//	*RunTestData_NodeResult
	Action isRunTestData_Action `protobuf_oneof:"action"`
}

func (x *RunTestData) Reset() {
	*x = RunTestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testRunner_testRunner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunTestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTestData) ProtoMessage() {}

func (x *RunTestData) ProtoReflect() protoreflect.Message {
	mi := &file_testRunner_testRunner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTestData.ProtoReflect.Descriptor instead.
func (*RunTestData) Descriptor() ([]byte, []int) {
	return file_testRunner_testRunner_proto_rawDescGZIP(), []int{0}
}

func (m *RunTestData) GetAction() isRunTestData_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *RunTestData) GetStartRunTest() *StartRunTest {
	if x, ok := x.GetAction().(*RunTestData_StartRunTest); ok {
		return x.StartRunTest
	}
	return nil
}

func (x *RunTestData) GetNodeResult() *NodeResult {
	if x, ok := x.GetAction().(*RunTestData_NodeResult); ok {
		return x.NodeResult
	}
	return nil
}

type isRunTestData_Action interface {
	isRunTestData_Action()
}

type RunTestData_StartRunTest struct {
	StartRunTest *StartRunTest `protobuf:"bytes,1,opt,name=startRunTest,proto3,oneof"`
}

type RunTestData_NodeResult struct {
	NodeResult *NodeResult `protobuf:"bytes,2,opt,name=nodeResult,proto3,oneof"`
}

func (*RunTestData_StartRunTest) isRunTestData_Action() {}

func (*RunTestData_NodeResult) isRunTestData_Action() {}

type StartRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NodeId  string                              `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Data    []*pluginrunner.TransformationField `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Jwt     string                              `protobuf:"bytes,4,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Payload *structpb.Struct                    `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	Events  *structpb.Struct                    `protobuf:"bytes,6,opt,name=events,proto3" json:"events,omitempty"`
}

func (x *StartRun) Reset() {
	*x = StartRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testRunner_testRunner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRun) ProtoMessage() {}

func (x *StartRun) ProtoReflect() protoreflect.Message {
	mi := &file_testRunner_testRunner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRun.ProtoReflect.Descriptor instead.
func (*StartRun) Descriptor() ([]byte, []int) {
	return file_testRunner_testRunner_proto_rawDescGZIP(), []int{1}
}

func (x *StartRun) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StartRun) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *StartRun) GetData() []*pluginrunner.TransformationField {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StartRun) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *StartRun) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *StartRun) GetEvents() *structpb.Struct {
	if x != nil {
		return x.Events
	}
	return nil
}

type StartRunTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GraphDefinition *graph.Graph                        `protobuf:"bytes,2,opt,name=graphDefinition,proto3" json:"graphDefinition,omitempty"`
	Testdata        []*pluginrunner.TransformationField `protobuf:"bytes,3,rep,name=testdata,proto3" json:"testdata,omitempty"`
	Jwt             string                              `protobuf:"bytes,4,opt,name=jwt,proto3" json:"jwt,omitempty"`
	NodeId          string                              `protobuf:"bytes,5,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Payload         *structpb.Struct                    `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	Events          *structpb.Struct                    `protobuf:"bytes,7,opt,name=events,proto3" json:"events,omitempty"`
}

func (x *StartRunTest) Reset() {
	*x = StartRunTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testRunner_testRunner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRunTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRunTest) ProtoMessage() {}

func (x *StartRunTest) ProtoReflect() protoreflect.Message {
	mi := &file_testRunner_testRunner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRunTest.ProtoReflect.Descriptor instead.
func (*StartRunTest) Descriptor() ([]byte, []int) {
	return file_testRunner_testRunner_proto_rawDescGZIP(), []int{2}
}

func (x *StartRunTest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StartRunTest) GetGraphDefinition() *graph.Graph {
	if x != nil {
		return x.GraphDefinition
	}
	return nil
}

func (x *StartRunTest) GetTestdata() []*pluginrunner.TransformationField {
	if x != nil {
		return x.Testdata
	}
	return nil
}

func (x *StartRunTest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *StartRunTest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *StartRunTest) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *StartRunTest) GetEvents() *structpb.Struct {
	if x != nil {
		return x.Events
	}
	return nil
}

type NodeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId     string                              `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Outputs    []*pluginrunner.TransformationField `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	NewOutputs *structpb.Struct                    `protobuf:"bytes,3,opt,name=newOutputs,proto3" json:"newOutputs,omitempty"`
	Events     *structpb.Struct                    `protobuf:"bytes,4,opt,name=events,proto3" json:"events,omitempty"`
}

func (x *NodeResult) Reset() {
	*x = NodeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testRunner_testRunner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeResult) ProtoMessage() {}

func (x *NodeResult) ProtoReflect() protoreflect.Message {
	mi := &file_testRunner_testRunner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeResult.ProtoReflect.Descriptor instead.
func (*NodeResult) Descriptor() ([]byte, []int) {
	return file_testRunner_testRunner_proto_rawDescGZIP(), []int{3}
}

func (x *NodeResult) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeResult) GetOutputs() []*pluginrunner.TransformationField {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *NodeResult) GetNewOutputs() *structpb.Struct {
	if x != nil {
		return x.NewOutputs
	}
	return nil
}

func (x *NodeResult) GetEvents() *structpb.Struct {
	if x != nil {
		return x.Events
	}
	return nil
}

type RunTestInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*RunTestInstruction_RunNode
	//	*RunTestInstruction_RunTestResult
	Action isRunTestInstruction_Action `protobuf_oneof:"action"`
}

func (x *RunTestInstruction) Reset() {
	*x = RunTestInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testRunner_testRunner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunTestInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTestInstruction) ProtoMessage() {}

func (x *RunTestInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_testRunner_testRunner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTestInstruction.ProtoReflect.Descriptor instead.
func (*RunTestInstruction) Descriptor() ([]byte, []int) {
	return file_testRunner_testRunner_proto_rawDescGZIP(), []int{4}
}

func (m *RunTestInstruction) GetAction() isRunTestInstruction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *RunTestInstruction) GetRunNode() *RunNode {
	if x, ok := x.GetAction().(*RunTestInstruction_RunNode); ok {
		return x.RunNode
	}
	return nil
}

func (x *RunTestInstruction) GetRunTestResult() *RunTestResult {
	if x, ok := x.GetAction().(*RunTestInstruction_RunTestResult); ok {
		return x.RunTestResult
	}
	return nil
}

type isRunTestInstruction_Action interface {
	isRunTestInstruction_Action()
}

type RunTestInstruction_RunNode struct {
	RunNode *RunNode `protobuf:"bytes,1,opt,name=runNode,proto3,oneof"`
}

type RunTestInstruction_RunTestResult struct {
	RunTestResult *RunTestResult `protobuf:"bytes,2,opt,name=runTestResult,proto3,oneof"`
}

func (*RunTestInstruction_RunNode) isRunTestInstruction_Action() {}

func (*RunTestInstruction_RunTestResult) isRunTestInstruction_Action() {}

type RunTestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event   string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Output  []byte `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	NodeId  string `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RunTestResult) Reset() {
	*x = RunTestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testRunner_testRunner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunTestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTestResult) ProtoMessage() {}

func (x *RunTestResult) ProtoReflect() protoreflect.Message {
	mi := &file_testRunner_testRunner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTestResult.ProtoReflect.Descriptor instead.
func (*RunTestResult) Descriptor() ([]byte, []int) {
	return file_testRunner_testRunner_proto_rawDescGZIP(), []int{5}
}

func (x *RunTestResult) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *RunTestResult) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *RunTestResult) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *RunTestResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RunNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId        string                              `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Event         string                              `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	NodeType      string                              `protobuf:"bytes,3,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	Configuration []byte                              `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Inputs        []*pluginrunner.TransformationField `protobuf:"bytes,5,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Skipped       []string                            `protobuf:"bytes,6,rep,name=skipped,proto3" json:"skipped,omitempty"`
	NewInputs     *structpb.Struct                    `protobuf:"bytes,7,opt,name=newInputs,proto3" json:"newInputs,omitempty"`
}

func (x *RunNode) Reset() {
	*x = RunNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testRunner_testRunner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunNode) ProtoMessage() {}

func (x *RunNode) ProtoReflect() protoreflect.Message {
	mi := &file_testRunner_testRunner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunNode.ProtoReflect.Descriptor instead.
func (*RunNode) Descriptor() ([]byte, []int) {
	return file_testRunner_testRunner_proto_rawDescGZIP(), []int{6}
}

func (x *RunNode) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *RunNode) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *RunNode) GetNodeType() string {
	if x != nil {
		return x.NodeType
	}
	return ""
}

func (x *RunNode) GetConfiguration() []byte {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *RunNode) GetInputs() []*pluginrunner.TransformationField {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *RunNode) GetSkipped() []string {
	if x != nil {
		return x.Skipped
	}
	return nil
}

func (x *RunNode) GetNewInputs() *structpb.Struct {
	if x != nil {
		return x.NewInputs
	}
	return nil
}

var File_testRunner_testRunner_proto protoreflect.FileDescriptor

var file_testRunner_testRunner_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74,
	0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x74,
	0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x70, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x54,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xdf, 0x01, 0x0a, 0x08,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2f, 0x0a, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xa6, 0x02,
	0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39,
	0x0a, 0x0f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x0f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x08,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x3b,
	0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x6e,
	0x65, 0x77, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x12, 0x52, 0x75, 0x6e, 0x54, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x07,
	0x72, 0x75, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x4e, 0x6f,
	0x64, 0x65, 0x48, 0x00, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x41, 0x0a,
	0x0d, 0x72, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x52, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48,
	0x00, 0x52, 0x0d, 0x72, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x70, 0x0a, 0x0d, 0x52, 0x75,
	0x6e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x87, 0x02, 0x0a,
	0x07, 0x52, 0x75, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x06, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x12,
	0x35, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x09, 0x6e, 0x65, 0x77,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x42, 0x1e, 0x5a, 0x1c, 0x77, 0x65, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testRunner_testRunner_proto_rawDescOnce sync.Once
	file_testRunner_testRunner_proto_rawDescData = file_testRunner_testRunner_proto_rawDesc
)

func file_testRunner_testRunner_proto_rawDescGZIP() []byte {
	file_testRunner_testRunner_proto_rawDescOnce.Do(func() {
		file_testRunner_testRunner_proto_rawDescData = protoimpl.X.CompressGZIP(file_testRunner_testRunner_proto_rawDescData)
	})
	return file_testRunner_testRunner_proto_rawDescData
}

var file_testRunner_testRunner_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_testRunner_testRunner_proto_goTypes = []interface{}{
	(*RunTestData)(nil),                      // 0: testRunner.RunTestData
	(*StartRun)(nil),                         // 1: testRunner.StartRun
	(*StartRunTest)(nil),                     // 2: testRunner.StartRunTest
	(*NodeResult)(nil),                       // 3: testRunner.NodeResult
	(*RunTestInstruction)(nil),               // 4: testRunner.RunTestInstruction
	(*RunTestResult)(nil),                    // 5: testRunner.RunTestResult
	(*RunNode)(nil),                          // 6: testRunner.RunNode
	(*pluginrunner.TransformationField)(nil), // 7: pluginrunner.TransformationField
	(*structpb.Struct)(nil),                  // 8: google.protobuf.Struct
	(*graph.Graph)(nil),                      // 9: workflow.Graph
}
var file_testRunner_testRunner_proto_depIdxs = []int32{
	2,  // 0: testRunner.RunTestData.startRunTest:type_name -> testRunner.StartRunTest
	3,  // 1: testRunner.RunTestData.nodeResult:type_name -> testRunner.NodeResult
	7,  // 2: testRunner.StartRun.data:type_name -> pluginrunner.TransformationField
	8,  // 3: testRunner.StartRun.payload:type_name -> google.protobuf.Struct
	8,  // 4: testRunner.StartRun.events:type_name -> google.protobuf.Struct
	9,  // 5: testRunner.StartRunTest.graphDefinition:type_name -> workflow.Graph
	7,  // 6: testRunner.StartRunTest.testdata:type_name -> pluginrunner.TransformationField
	8,  // 7: testRunner.StartRunTest.payload:type_name -> google.protobuf.Struct
	8,  // 8: testRunner.StartRunTest.events:type_name -> google.protobuf.Struct
	7,  // 9: testRunner.NodeResult.outputs:type_name -> pluginrunner.TransformationField
	8,  // 10: testRunner.NodeResult.newOutputs:type_name -> google.protobuf.Struct
	8,  // 11: testRunner.NodeResult.events:type_name -> google.protobuf.Struct
	6,  // 12: testRunner.RunTestInstruction.runNode:type_name -> testRunner.RunNode
	5,  // 13: testRunner.RunTestInstruction.runTestResult:type_name -> testRunner.RunTestResult
	7,  // 14: testRunner.RunNode.inputs:type_name -> pluginrunner.TransformationField
	8,  // 15: testRunner.RunNode.newInputs:type_name -> google.protobuf.Struct
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_testRunner_testRunner_proto_init() }
func file_testRunner_testRunner_proto_init() {
	if File_testRunner_testRunner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testRunner_testRunner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunTestData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testRunner_testRunner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRun); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testRunner_testRunner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRunTest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testRunner_testRunner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testRunner_testRunner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunTestInstruction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testRunner_testRunner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunTestResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testRunner_testRunner_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunNode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_testRunner_testRunner_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RunTestData_StartRunTest)(nil),
		(*RunTestData_NodeResult)(nil),
	}
	file_testRunner_testRunner_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*RunTestInstruction_RunNode)(nil),
		(*RunTestInstruction_RunTestResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testRunner_testRunner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testRunner_testRunner_proto_goTypes,
		DependencyIndexes: file_testRunner_testRunner_proto_depIdxs,
		MessageInfos:      file_testRunner_testRunner_proto_msgTypes,
	}.Build()
	File_testRunner_testRunner_proto = out.File
	file_testRunner_testRunner_proto_rawDesc = nil
	file_testRunner_testRunner_proto_goTypes = nil
	file_testRunner_testRunner_proto_depIdxs = nil
}
