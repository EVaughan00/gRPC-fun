// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/minimega/MinimegaModule.proto

package minimega

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VlanSpecs      *VlanSpecs      `protobuf:"bytes,1,opt,name=vlanSpecs,proto3" json:"vlanSpecs,omitempty"`
	TapSpecs       *TapSpecs       `protobuf:"bytes,2,opt,name=tapSpecs,proto3" json:"tapSpecs,omitempty"`
	Orchestrations *Orchestrations `protobuf:"bytes,3,opt,name=orchestrations,proto3" json:"orchestrations,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_proto_minimega_MinimegaModule_proto_rawDescGZIP(), []int{0}
}

func (x *Configuration) GetVlanSpecs() *VlanSpecs {
	if x != nil {
		return x.VlanSpecs
	}
	return nil
}

func (x *Configuration) GetTapSpecs() *TapSpecs {
	if x != nil {
		return x.TapSpecs
	}
	return nil
}

func (x *Configuration) GetOrchestrations() *Orchestrations {
	if x != nil {
		return x.Orchestrations
	}
	return nil
}

type VlanSpecs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagementVLAN string            `protobuf:"bytes,1,opt,name=managementVLAN,proto3" json:"managementVLAN,omitempty"`
	SnifferVLANs   map[string]string `protobuf:"bytes,2,rep,name=snifferVLANs,proto3" json:"snifferVLANs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HilVLANs       map[string]string `protobuf:"bytes,3,rep,name=hilVLANs,proto3" json:"hilVLANs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VlanSpecs) Reset() {
	*x = VlanSpecs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VlanSpecs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VlanSpecs) ProtoMessage() {}

func (x *VlanSpecs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VlanSpecs.ProtoReflect.Descriptor instead.
func (*VlanSpecs) Descriptor() ([]byte, []int) {
	return file_proto_minimega_MinimegaModule_proto_rawDescGZIP(), []int{1}
}

func (x *VlanSpecs) GetManagementVLAN() string {
	if x != nil {
		return x.ManagementVLAN
	}
	return ""
}

func (x *VlanSpecs) GetSnifferVLANs() map[string]string {
	if x != nil {
		return x.SnifferVLANs
	}
	return nil
}

func (x *VlanSpecs) GetHilVLANs() map[string]string {
	if x != nil {
		return x.HilVLANs
	}
	return nil
}

type TapSpecs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetflowTapPort   string `protobuf:"bytes,1,opt,name=netflowTapPort,proto3" json:"netflowTapPort,omitempty"`
	NetflowTapIP     string `protobuf:"bytes,2,opt,name=netflowTapIP,proto3" json:"netflowTapIP,omitempty"`
	PowerTapPort     string `protobuf:"bytes,3,opt,name=powerTapPort,proto3" json:"powerTapPort,omitempty"`
	SnifferTapName   string `protobuf:"bytes,4,opt,name=snifferTapName,proto3" json:"snifferTapName,omitempty"`
	PublisherTapName string `protobuf:"bytes,5,opt,name=publisherTapName,proto3" json:"publisherTapName,omitempty"`
	PublisherTapIP   string `protobuf:"bytes,6,opt,name=publisherTapIP,proto3" json:"publisherTapIP,omitempty"`
}

func (x *TapSpecs) Reset() {
	*x = TapSpecs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TapSpecs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TapSpecs) ProtoMessage() {}

func (x *TapSpecs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TapSpecs.ProtoReflect.Descriptor instead.
func (*TapSpecs) Descriptor() ([]byte, []int) {
	return file_proto_minimega_MinimegaModule_proto_rawDescGZIP(), []int{2}
}

func (x *TapSpecs) GetNetflowTapPort() string {
	if x != nil {
		return x.NetflowTapPort
	}
	return ""
}

func (x *TapSpecs) GetNetflowTapIP() string {
	if x != nil {
		return x.NetflowTapIP
	}
	return ""
}

func (x *TapSpecs) GetPowerTapPort() string {
	if x != nil {
		return x.PowerTapPort
	}
	return ""
}

func (x *TapSpecs) GetSnifferTapName() string {
	if x != nil {
		return x.SnifferTapName
	}
	return ""
}

func (x *TapSpecs) GetPublisherTapName() string {
	if x != nil {
		return x.PublisherTapName
	}
	return ""
}

func (x *TapSpecs) GetPublisherTapIP() string {
	if x != nil {
		return x.PublisherTapIP
	}
	return ""
}

type Orchestrations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Orchestrations) Reset() {
	*x = Orchestrations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orchestrations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orchestrations) ProtoMessage() {}

func (x *Orchestrations) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orchestrations.ProtoReflect.Descriptor instead.
func (*Orchestrations) Descriptor() ([]byte, []int) {
	return file_proto_minimega_MinimegaModule_proto_rawDescGZIP(), []int{3}
}

func (x *Orchestrations) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_minimega_MinimegaModule_proto_rawDescGZIP(), []int{4}
}

func (x *StatusRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DetailedEvaluation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinimegaProcessIsRunning bool `protobuf:"varint,1,opt,name=minimegaProcessIsRunning,proto3" json:"minimegaProcessIsRunning,omitempty"`
	MiniwebProcessIsRunning  bool `protobuf:"varint,2,opt,name=miniwebProcessIsRunning,proto3" json:"miniwebProcessIsRunning,omitempty"`
}

func (x *DetailedEvaluation) Reset() {
	*x = DetailedEvaluation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailedEvaluation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailedEvaluation) ProtoMessage() {}

func (x *DetailedEvaluation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailedEvaluation.ProtoReflect.Descriptor instead.
func (*DetailedEvaluation) Descriptor() ([]byte, []int) {
	return file_proto_minimega_MinimegaModule_proto_rawDescGZIP(), []int{5}
}

func (x *DetailedEvaluation) GetMinimegaProcessIsRunning() bool {
	if x != nil {
		return x.MinimegaProcessIsRunning
	}
	return false
}

func (x *DetailedEvaluation) GetMiniwebProcessIsRunning() bool {
	if x != nil {
		return x.MiniwebProcessIsRunning
	}
	return false
}

type ConfigurationUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath  string `protobuf:"bytes,1,opt,name=filePath,proto3" json:"filePath,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ConfigurationUpdate) Reset() {
	*x = ConfigurationUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurationUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurationUpdate) ProtoMessage() {}

func (x *ConfigurationUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurationUpdate.ProtoReflect.Descriptor instead.
func (*ConfigurationUpdate) Descriptor() ([]byte, []int) {
	return file_proto_minimega_MinimegaModule_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigurationUpdate) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *ConfigurationUpdate) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type CustomCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CustomCommand) Reset() {
	*x = CustomCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomCommand) ProtoMessage() {}

func (x *CustomCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomCommand.ProtoReflect.Descriptor instead.
func (*CustomCommand) Descriptor() ([]byte, []int) {
	return file_proto_minimega_MinimegaModule_proto_rawDescGZIP(), []int{7}
}

func (x *CustomCommand) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Completed    bool   `protobuf:"varint,1,opt,name=completed,proto3" json:"completed,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minimega_MinimegaModule_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_proto_minimega_MinimegaModule_proto_rawDescGZIP(), []int{8}
}

func (x *Confirmation) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *Confirmation) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_proto_minimega_MinimegaModule_proto protoreflect.FileDescriptor

var file_proto_minimega_MinimegaModule_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61,
	0x2f, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x22, 0xc6, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x76, 0x6c, 0x61, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x6c, 0x61, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x73, 0x52, 0x09, 0x76, 0x6c, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x73,
	0x12, 0x34, 0x0a, 0x08, 0x74, 0x61, 0x70, 0x53, 0x70, 0x65, 0x63, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x70, 0x53, 0x70, 0x65, 0x63, 0x73, 0x52, 0x08, 0x74, 0x61,
	0x70, 0x53, 0x70, 0x65, 0x63, 0x73, 0x12, 0x46, 0x0a, 0x0e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0e,
	0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc7,
	0x02, 0x0a, 0x09, 0x56, 0x6c, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x4c, 0x41, 0x4e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x56, 0x4c, 0x41, 0x4e, 0x12, 0x4f, 0x0a, 0x0c, 0x73, 0x6e, 0x69, 0x66, 0x66, 0x65, 0x72, 0x56,
	0x4c, 0x41, 0x4e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x6c, 0x61, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x53, 0x6e, 0x69, 0x66, 0x66, 0x65, 0x72, 0x56, 0x4c, 0x41,
	0x4e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x6e, 0x69, 0x66, 0x66, 0x65, 0x72,
	0x56, 0x4c, 0x41, 0x4e, 0x73, 0x12, 0x43, 0x0a, 0x08, 0x68, 0x69, 0x6c, 0x56, 0x4c, 0x41, 0x4e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65,
	0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x6c, 0x61, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x73, 0x2e, 0x48, 0x69, 0x6c, 0x56, 0x4c, 0x41, 0x4e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x68, 0x69, 0x6c, 0x56, 0x4c, 0x41, 0x4e, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x53, 0x6e,
	0x69, 0x66, 0x66, 0x65, 0x72, 0x56, 0x4c, 0x41, 0x4e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x48,
	0x69, 0x6c, 0x56, 0x4c, 0x41, 0x4e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf6, 0x01, 0x0a, 0x08, 0x54, 0x61, 0x70,
	0x53, 0x70, 0x65, 0x63, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77,
	0x54, 0x61, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e,
	0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x61, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x61, 0x70, 0x49, 0x50, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x61, 0x70, 0x49,
	0x50, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x61, 0x70, 0x50, 0x6f, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x61,
	0x70, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6e, 0x69, 0x66, 0x66, 0x65, 0x72,
	0x54, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x6e, 0x69, 0x66, 0x66, 0x65, 0x72, 0x54, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a,
	0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x54, 0x61, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x54, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x54, 0x61, 0x70, 0x49, 0x50, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x54, 0x61, 0x70, 0x49,
	0x50, 0x22, 0x2c, 0x0a, 0x0e, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x29, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3a, 0x0a, 0x18, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x18, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a,
	0x17, 0x6d, 0x69, 0x6e, 0x69, 0x77, 0x65, 0x62, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17,
	0x6d, 0x69, 0x6e, 0x69, 0x77, 0x65, 0x62, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x73,
	0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x4f, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x50, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0xef, 0x02, 0x0a, 0x0f, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x12, 0x1d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1c, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12,
	0x61, 0x0a, 0x1a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1d, 0x2e,
	0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d,
	0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x5a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x6d, 0x69, 0x6e, 0x69,
	0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x1c,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x10, 0x52, 0x75, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x1d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x1a, 0x1c, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x67, 0x61, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x67, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_minimega_MinimegaModule_proto_rawDescOnce sync.Once
	file_proto_minimega_MinimegaModule_proto_rawDescData = file_proto_minimega_MinimegaModule_proto_rawDesc
)

func file_proto_minimega_MinimegaModule_proto_rawDescGZIP() []byte {
	file_proto_minimega_MinimegaModule_proto_rawDescOnce.Do(func() {
		file_proto_minimega_MinimegaModule_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_minimega_MinimegaModule_proto_rawDescData)
	})
	return file_proto_minimega_MinimegaModule_proto_rawDescData
}

var file_proto_minimega_MinimegaModule_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_minimega_MinimegaModule_proto_goTypes = []interface{}{
	(*Configuration)(nil),       // 0: minimegamodule.Configuration
	(*VlanSpecs)(nil),           // 1: minimegamodule.VlanSpecs
	(*TapSpecs)(nil),            // 2: minimegamodule.TapSpecs
	(*Orchestrations)(nil),      // 3: minimegamodule.Orchestrations
	(*StatusRequest)(nil),       // 4: minimegamodule.StatusRequest
	(*DetailedEvaluation)(nil),  // 5: minimegamodule.DetailedEvaluation
	(*ConfigurationUpdate)(nil), // 6: minimegamodule.ConfigurationUpdate
	(*CustomCommand)(nil),       // 7: minimegamodule.CustomCommand
	(*Confirmation)(nil),        // 8: minimegamodule.Confirmation
	nil,                         // 9: minimegamodule.VlanSpecs.SnifferVLANsEntry
	nil,                         // 10: minimegamodule.VlanSpecs.HilVLANsEntry
}
var file_proto_minimega_MinimegaModule_proto_depIdxs = []int32{
	1,  // 0: minimegamodule.Configuration.vlanSpecs:type_name -> minimegamodule.VlanSpecs
	2,  // 1: minimegamodule.Configuration.tapSpecs:type_name -> minimegamodule.TapSpecs
	3,  // 2: minimegamodule.Configuration.orchestrations:type_name -> minimegamodule.Orchestrations
	9,  // 3: minimegamodule.VlanSpecs.snifferVLANs:type_name -> minimegamodule.VlanSpecs.SnifferVLANsEntry
	10, // 4: minimegamodule.VlanSpecs.hilVLANs:type_name -> minimegamodule.VlanSpecs.HilVLANsEntry
	0,  // 5: minimegamodule.MinimegaService.Configure:input_type -> minimegamodule.Configuration
	4,  // 6: minimegamodule.MinimegaService.RequestDetailedHealthCheck:input_type -> minimegamodule.StatusRequest
	6,  // 7: minimegamodule.MinimegaService.UpdateConfiguration:input_type -> minimegamodule.ConfigurationUpdate
	7,  // 8: minimegamodule.MinimegaService.RunCustomCommand:input_type -> minimegamodule.CustomCommand
	8,  // 9: minimegamodule.MinimegaService.Configure:output_type -> minimegamodule.Confirmation
	5,  // 10: minimegamodule.MinimegaService.RequestDetailedHealthCheck:output_type -> minimegamodule.DetailedEvaluation
	8,  // 11: minimegamodule.MinimegaService.UpdateConfiguration:output_type -> minimegamodule.Confirmation
	8,  // 12: minimegamodule.MinimegaService.RunCustomCommand:output_type -> minimegamodule.Confirmation
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_minimega_MinimegaModule_proto_init() }
func file_proto_minimega_MinimegaModule_proto_init() {
	if File_proto_minimega_MinimegaModule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_minimega_MinimegaModule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
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
		file_proto_minimega_MinimegaModule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VlanSpecs); i {
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
		file_proto_minimega_MinimegaModule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TapSpecs); i {
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
		file_proto_minimega_MinimegaModule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orchestrations); i {
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
		file_proto_minimega_MinimegaModule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_proto_minimega_MinimegaModule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailedEvaluation); i {
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
		file_proto_minimega_MinimegaModule_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurationUpdate); i {
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
		file_proto_minimega_MinimegaModule_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomCommand); i {
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
		file_proto_minimega_MinimegaModule_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmation); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_minimega_MinimegaModule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_minimega_MinimegaModule_proto_goTypes,
		DependencyIndexes: file_proto_minimega_MinimegaModule_proto_depIdxs,
		MessageInfos:      file_proto_minimega_MinimegaModule_proto_msgTypes,
	}.Build()
	File_proto_minimega_MinimegaModule_proto = out.File
	file_proto_minimega_MinimegaModule_proto_rawDesc = nil
	file_proto_minimega_MinimegaModule_proto_goTypes = nil
	file_proto_minimega_MinimegaModule_proto_depIdxs = nil
}
