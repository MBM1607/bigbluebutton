// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: common/meeting_info.proto

package common

import (
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

type MeetingInfo struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	MeetingName      string                 `protobuf:"bytes,1,opt,name=meeting_name,json=meetingName,proto3" json:"meeting_name,omitempty"`
	MeetingExtId     string                 `protobuf:"bytes,2,opt,name=meeting_ext_id,json=meetingExtId,proto3" json:"meeting_ext_id,omitempty"`
	MeetingIntId     string                 `protobuf:"bytes,3,opt,name=meeting_int_id,json=meetingIntId,proto3" json:"meeting_int_id,omitempty"`
	VoiceBridge      string                 `protobuf:"bytes,6,opt,name=voice_bridge,json=voiceBridge,proto3" json:"voice_bridge,omitempty"`
	DialNumber       string                 `protobuf:"bytes,7,opt,name=dial_number,json=dialNumber,proto3" json:"dial_number,omitempty"`
	AttendeePw       string                 `protobuf:"bytes,8,opt,name=attendee_pw,json=attendeePw,proto3" json:"attendee_pw,omitempty"`
	ModeratorPw      string                 `protobuf:"bytes,9,opt,name=moderator_pw,json=moderatorPw,proto3" json:"moderator_pw,omitempty"`
	Recording        bool                   `protobuf:"varint,10,opt,name=recording,proto3" json:"recording,omitempty"`
	Users            []*User                `protobuf:"bytes,11,rep,name=users,proto3" json:"users,omitempty"`
	Metadata         map[string]string      `protobuf:"bytes,12,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BreakoutRooms    []string               `protobuf:"bytes,13,rep,name=breakout_rooms,json=breakoutRooms,proto3" json:"breakout_rooms,omitempty"`
	DurationInfo     *DurationInfo          `protobuf:"bytes,14,opt,name=duration_info,json=durationInfo,proto3" json:"duration_info,omitempty"`
	ParticipantInfo  *ParticipantInfo       `protobuf:"bytes,15,opt,name=participant_info,json=participantInfo,proto3" json:"participant_info,omitempty"`
	BreakoutInfo     *BreakoutInfo          `protobuf:"bytes,16,opt,name=breakout_info,json=breakoutInfo,proto3" json:"breakout_info,omitempty"`
	DisabledFeatures []string               `protobuf:"bytes,17,rep,name=disabled_features,json=disabledFeatures,proto3" json:"disabled_features,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *MeetingInfo) Reset() {
	*x = MeetingInfo{}
	mi := &file_common_meeting_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MeetingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeetingInfo) ProtoMessage() {}

func (x *MeetingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_meeting_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeetingInfo.ProtoReflect.Descriptor instead.
func (*MeetingInfo) Descriptor() ([]byte, []int) {
	return file_common_meeting_info_proto_rawDescGZIP(), []int{0}
}

func (x *MeetingInfo) GetMeetingName() string {
	if x != nil {
		return x.MeetingName
	}
	return ""
}

func (x *MeetingInfo) GetMeetingExtId() string {
	if x != nil {
		return x.MeetingExtId
	}
	return ""
}

func (x *MeetingInfo) GetMeetingIntId() string {
	if x != nil {
		return x.MeetingIntId
	}
	return ""
}

func (x *MeetingInfo) GetVoiceBridge() string {
	if x != nil {
		return x.VoiceBridge
	}
	return ""
}

func (x *MeetingInfo) GetDialNumber() string {
	if x != nil {
		return x.DialNumber
	}
	return ""
}

func (x *MeetingInfo) GetAttendeePw() string {
	if x != nil {
		return x.AttendeePw
	}
	return ""
}

func (x *MeetingInfo) GetModeratorPw() string {
	if x != nil {
		return x.ModeratorPw
	}
	return ""
}

func (x *MeetingInfo) GetRecording() bool {
	if x != nil {
		return x.Recording
	}
	return false
}

func (x *MeetingInfo) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *MeetingInfo) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MeetingInfo) GetBreakoutRooms() []string {
	if x != nil {
		return x.BreakoutRooms
	}
	return nil
}

func (x *MeetingInfo) GetDurationInfo() *DurationInfo {
	if x != nil {
		return x.DurationInfo
	}
	return nil
}

func (x *MeetingInfo) GetParticipantInfo() *ParticipantInfo {
	if x != nil {
		return x.ParticipantInfo
	}
	return nil
}

func (x *MeetingInfo) GetBreakoutInfo() *BreakoutInfo {
	if x != nil {
		return x.BreakoutInfo
	}
	return nil
}

func (x *MeetingInfo) GetDisabledFeatures() []string {
	if x != nil {
		return x.DisabledFeatures
	}
	return nil
}

var File_common_meeting_info_proto protoreflect.FileDescriptor

var file_common_meeting_info_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x72, 0x67,
	0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x72, 0x65, 0x61,
	0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaa, 0x06, 0x0a, 0x0b, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x65, 0x5f, 0x70,
	0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x65, 0x50, 0x77, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x70, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x50, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75,
	0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x4f, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x62,
	0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x12, 0x4b, 0x0a, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x54, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4b, 0x0a, 0x0d, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x1a,
	0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_meeting_info_proto_rawDescOnce sync.Once
	file_common_meeting_info_proto_rawDescData = file_common_meeting_info_proto_rawDesc
)

func file_common_meeting_info_proto_rawDescGZIP() []byte {
	file_common_meeting_info_proto_rawDescOnce.Do(func() {
		file_common_meeting_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_meeting_info_proto_rawDescData)
	})
	return file_common_meeting_info_proto_rawDescData
}

var file_common_meeting_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_meeting_info_proto_goTypes = []any{
	(*MeetingInfo)(nil),     // 0: org.bigbluebutton.protos.MeetingInfo
	nil,                     // 1: org.bigbluebutton.protos.MeetingInfo.MetadataEntry
	(*User)(nil),            // 2: org.bigbluebutton.protos.User
	(*DurationInfo)(nil),    // 3: org.bigbluebutton.protos.DurationInfo
	(*ParticipantInfo)(nil), // 4: org.bigbluebutton.protos.ParticipantInfo
	(*BreakoutInfo)(nil),    // 5: org.bigbluebutton.protos.BreakoutInfo
}
var file_common_meeting_info_proto_depIdxs = []int32{
	2, // 0: org.bigbluebutton.protos.MeetingInfo.users:type_name -> org.bigbluebutton.protos.User
	1, // 1: org.bigbluebutton.protos.MeetingInfo.metadata:type_name -> org.bigbluebutton.protos.MeetingInfo.MetadataEntry
	3, // 2: org.bigbluebutton.protos.MeetingInfo.duration_info:type_name -> org.bigbluebutton.protos.DurationInfo
	4, // 3: org.bigbluebutton.protos.MeetingInfo.participant_info:type_name -> org.bigbluebutton.protos.ParticipantInfo
	5, // 4: org.bigbluebutton.protos.MeetingInfo.breakout_info:type_name -> org.bigbluebutton.protos.BreakoutInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_meeting_info_proto_init() }
func file_common_meeting_info_proto_init() {
	if File_common_meeting_info_proto != nil {
		return
	}
	file_common_user_proto_init()
	file_common_duration_info_proto_init()
	file_common_participant_info_proto_init()
	file_common_breakout_info_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_meeting_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_meeting_info_proto_goTypes,
		DependencyIndexes: file_common_meeting_info_proto_depIdxs,
		MessageInfos:      file_common_meeting_info_proto_msgTypes,
	}.Build()
	File_common_meeting_info_proto = out.File
	file_common_meeting_info_proto_rawDesc = nil
	file_common_meeting_info_proto_goTypes = nil
	file_common_meeting_info_proto_depIdxs = nil
}
