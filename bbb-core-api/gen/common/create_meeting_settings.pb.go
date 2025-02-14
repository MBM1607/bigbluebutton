// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: common/create_meeting_settings.proto

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

type CreateMeetingSettings struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	MeetingSettings        *MeetingSettings       `protobuf:"bytes,1,opt,name=meeting_settings,json=meetingSettings,proto3" json:"meeting_settings,omitempty"`
	BreakoutSettings       *BreakoutSettings      `protobuf:"bytes,2,opt,name=breakout_settings,json=breakoutSettings,proto3" json:"breakout_settings,omitempty"`
	DurationSettings       *DurationSettings      `protobuf:"bytes,3,opt,name=duration_settings,json=durationSettings,proto3" json:"duration_settings,omitempty"`
	PasswordSettings       *PasswordSettings      `protobuf:"bytes,4,opt,name=password_settings,json=passwordSettings,proto3" json:"password_settings,omitempty"`
	RecordSettings         *RecordSettings        `protobuf:"bytes,5,opt,name=record_settings,json=recordSettings,proto3" json:"record_settings,omitempty"`
	WelcomeSettings        *WelcomeSettings       `protobuf:"bytes,6,opt,name=welcome_settings,json=welcomeSettings,proto3" json:"welcome_settings,omitempty"`
	VoiceSettings          *VoiceSettings         `protobuf:"bytes,7,opt,name=voice_settings,json=voiceSettings,proto3" json:"voice_settings,omitempty"`
	UserSettings           *UserSettings          `protobuf:"bytes,8,opt,name=user_settings,json=userSettings,proto3" json:"user_settings,omitempty"`
	MetadataSettings       *MetadataSettings      `protobuf:"bytes,9,opt,name=metadata_settings,json=metadataSettings,proto3" json:"metadata_settings,omitempty"`
	LockSettings           *LockSettings          `protobuf:"bytes,10,opt,name=lock_settings,json=lockSettings,proto3" json:"lock_settings,omitempty"`
	SystemSettings         *SystemSettings        `protobuf:"bytes,11,opt,name=system_settings,json=systemSettings,proto3" json:"system_settings,omitempty"`
	GroupSettings          []*GroupSettings       `protobuf:"bytes,12,rep,name=group_settings,json=groupSettings,proto3" json:"group_settings,omitempty"`
	OverrideClientSettings string                 `protobuf:"bytes,13,opt,name=override_client_settings,json=overrideClientSettings,proto3" json:"override_client_settings,omitempty"`
	PluginSettings         *PluginSettings        `protobuf:"bytes,14,opt,name=plugin_settings,json=pluginSettings,proto3" json:"plugin_settings,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *CreateMeetingSettings) Reset() {
	*x = CreateMeetingSettings{}
	mi := &file_common_create_meeting_settings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMeetingSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMeetingSettings) ProtoMessage() {}

func (x *CreateMeetingSettings) ProtoReflect() protoreflect.Message {
	mi := &file_common_create_meeting_settings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMeetingSettings.ProtoReflect.Descriptor instead.
func (*CreateMeetingSettings) Descriptor() ([]byte, []int) {
	return file_common_create_meeting_settings_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMeetingSettings) GetMeetingSettings() *MeetingSettings {
	if x != nil {
		return x.MeetingSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetBreakoutSettings() *BreakoutSettings {
	if x != nil {
		return x.BreakoutSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetDurationSettings() *DurationSettings {
	if x != nil {
		return x.DurationSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetPasswordSettings() *PasswordSettings {
	if x != nil {
		return x.PasswordSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetRecordSettings() *RecordSettings {
	if x != nil {
		return x.RecordSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetWelcomeSettings() *WelcomeSettings {
	if x != nil {
		return x.WelcomeSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetVoiceSettings() *VoiceSettings {
	if x != nil {
		return x.VoiceSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetUserSettings() *UserSettings {
	if x != nil {
		return x.UserSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetMetadataSettings() *MetadataSettings {
	if x != nil {
		return x.MetadataSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetLockSettings() *LockSettings {
	if x != nil {
		return x.LockSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetSystemSettings() *SystemSettings {
	if x != nil {
		return x.SystemSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetGroupSettings() []*GroupSettings {
	if x != nil {
		return x.GroupSettings
	}
	return nil
}

func (x *CreateMeetingSettings) GetOverrideClientSettings() string {
	if x != nil {
		return x.OverrideClientSettings
	}
	return ""
}

func (x *CreateMeetingSettings) GetPluginSettings() *PluginSettings {
	if x != nil {
		return x.PluginSettings
	}
	return nil
}

var File_common_create_meeting_settings_proto protoreflect.FileDescriptor

var file_common_create_meeting_settings_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62,
	0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x1a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x09, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x54, 0x0a, 0x10, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0f, 0x6d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x57, 0x0a, 0x11, 0x62, 0x72,
	0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62,
	0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x10, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x57, 0x0a, 0x11, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x10, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x57, 0x0a, 0x11,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69,
	0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x51, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x54, 0x0a, 0x10, 0x77, 0x65, 0x6c, 0x63,
	0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65,
	0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x57, 0x65,
	0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0f, 0x77,
	0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4e,
	0x0a, 0x0e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67,
	0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x0d, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4b,
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62,
	0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x57, 0x0a, 0x11, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67,
	0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x4b, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x51, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x4e, 0x0a, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x51,
	0x0a, 0x0f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69,
	0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x0e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_create_meeting_settings_proto_rawDescOnce sync.Once
	file_common_create_meeting_settings_proto_rawDescData = file_common_create_meeting_settings_proto_rawDesc
)

func file_common_create_meeting_settings_proto_rawDescGZIP() []byte {
	file_common_create_meeting_settings_proto_rawDescOnce.Do(func() {
		file_common_create_meeting_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_create_meeting_settings_proto_rawDescData)
	})
	return file_common_create_meeting_settings_proto_rawDescData
}

var file_common_create_meeting_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_create_meeting_settings_proto_goTypes = []any{
	(*CreateMeetingSettings)(nil), // 0: org.bigbluebutton.protos.CreateMeetingSettings
	(*MeetingSettings)(nil),       // 1: org.bigbluebutton.protos.MeetingSettings
	(*BreakoutSettings)(nil),      // 2: org.bigbluebutton.protos.BreakoutSettings
	(*DurationSettings)(nil),      // 3: org.bigbluebutton.protos.DurationSettings
	(*PasswordSettings)(nil),      // 4: org.bigbluebutton.protos.PasswordSettings
	(*RecordSettings)(nil),        // 5: org.bigbluebutton.protos.RecordSettings
	(*WelcomeSettings)(nil),       // 6: org.bigbluebutton.protos.WelcomeSettings
	(*VoiceSettings)(nil),         // 7: org.bigbluebutton.protos.VoiceSettings
	(*UserSettings)(nil),          // 8: org.bigbluebutton.protos.UserSettings
	(*MetadataSettings)(nil),      // 9: org.bigbluebutton.protos.MetadataSettings
	(*LockSettings)(nil),          // 10: org.bigbluebutton.protos.LockSettings
	(*SystemSettings)(nil),        // 11: org.bigbluebutton.protos.SystemSettings
	(*GroupSettings)(nil),         // 12: org.bigbluebutton.protos.GroupSettings
	(*PluginSettings)(nil),        // 13: org.bigbluebutton.protos.PluginSettings
}
var file_common_create_meeting_settings_proto_depIdxs = []int32{
	1,  // 0: org.bigbluebutton.protos.CreateMeetingSettings.meeting_settings:type_name -> org.bigbluebutton.protos.MeetingSettings
	2,  // 1: org.bigbluebutton.protos.CreateMeetingSettings.breakout_settings:type_name -> org.bigbluebutton.protos.BreakoutSettings
	3,  // 2: org.bigbluebutton.protos.CreateMeetingSettings.duration_settings:type_name -> org.bigbluebutton.protos.DurationSettings
	4,  // 3: org.bigbluebutton.protos.CreateMeetingSettings.password_settings:type_name -> org.bigbluebutton.protos.PasswordSettings
	5,  // 4: org.bigbluebutton.protos.CreateMeetingSettings.record_settings:type_name -> org.bigbluebutton.protos.RecordSettings
	6,  // 5: org.bigbluebutton.protos.CreateMeetingSettings.welcome_settings:type_name -> org.bigbluebutton.protos.WelcomeSettings
	7,  // 6: org.bigbluebutton.protos.CreateMeetingSettings.voice_settings:type_name -> org.bigbluebutton.protos.VoiceSettings
	8,  // 7: org.bigbluebutton.protos.CreateMeetingSettings.user_settings:type_name -> org.bigbluebutton.protos.UserSettings
	9,  // 8: org.bigbluebutton.protos.CreateMeetingSettings.metadata_settings:type_name -> org.bigbluebutton.protos.MetadataSettings
	10, // 9: org.bigbluebutton.protos.CreateMeetingSettings.lock_settings:type_name -> org.bigbluebutton.protos.LockSettings
	11, // 10: org.bigbluebutton.protos.CreateMeetingSettings.system_settings:type_name -> org.bigbluebutton.protos.SystemSettings
	12, // 11: org.bigbluebutton.protos.CreateMeetingSettings.group_settings:type_name -> org.bigbluebutton.protos.GroupSettings
	13, // 12: org.bigbluebutton.protos.CreateMeetingSettings.plugin_settings:type_name -> org.bigbluebutton.protos.PluginSettings
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_common_create_meeting_settings_proto_init() }
func file_common_create_meeting_settings_proto_init() {
	if File_common_create_meeting_settings_proto != nil {
		return
	}
	file_common_meeting_settings_proto_init()
	file_common_breakout_settings_proto_init()
	file_common_duration_settings_proto_init()
	file_common_password_settings_proto_init()
	file_common_record_settings_proto_init()
	file_common_welcome_settings_proto_init()
	file_common_voice_settings_proto_init()
	file_common_user_settings_proto_init()
	file_common_metadata_settings_proto_init()
	file_common_lock_settings_proto_init()
	file_common_system_settings_proto_init()
	file_common_group_settings_proto_init()
	file_common_plugin_settings_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_create_meeting_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_create_meeting_settings_proto_goTypes,
		DependencyIndexes: file_common_create_meeting_settings_proto_depIdxs,
		MessageInfos:      file_common_create_meeting_settings_proto_msgTypes,
	}.Build()
	File_common_create_meeting_settings_proto = out.File
	file_common_create_meeting_settings_proto_rawDesc = nil
	file_common_create_meeting_settings_proto_goTypes = nil
	file_common_create_meeting_settings_proto_depIdxs = nil
}
