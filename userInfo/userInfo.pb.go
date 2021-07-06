// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: userInfo.proto

//default go package name

package userInfo

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

type CardLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsGroupHeader bool   `protobuf:"varint,1,opt,name=isGroupHeader,proto3" json:"isGroupHeader,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ThumbnailUrl  string `protobuf:"bytes,3,opt,name=thumbnailUrl,proto3" json:"thumbnailUrl,omitempty"`
	Link          string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	Css           string `protobuf:"bytes,5,opt,name=css,proto3" json:"css,omitempty"`
}

func (x *CardLink) Reset() {
	*x = CardLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardLink) ProtoMessage() {}

func (x *CardLink) ProtoReflect() protoreflect.Message {
	mi := &file_userInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardLink.ProtoReflect.Descriptor instead.
func (*CardLink) Descriptor() ([]byte, []int) {
	return file_userInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CardLink) GetIsGroupHeader() bool {
	if x != nil {
		return x.IsGroupHeader
	}
	return false
}

func (x *CardLink) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CardLink) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *CardLink) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *CardLink) GetCss() string {
	if x != nil {
		return x.Css
	}
	return ""
}

type SocialLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IconSVG string `protobuf:"bytes,1,opt,name=iconSVG,proto3" json:"iconSVG,omitempty"`
	Link    string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *SocialLink) Reset() {
	*x = SocialLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialLink) ProtoMessage() {}

func (x *SocialLink) ProtoReflect() protoreflect.Message {
	mi := &file_userInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialLink.ProtoReflect.Descriptor instead.
func (*SocialLink) Descriptor() ([]byte, []int) {
	return file_userInfo_proto_rawDescGZIP(), []int{1}
}

func (x *SocialLink) GetIconSVG() string {
	if x != nil {
		return x.IconSVG
	}
	return ""
}

func (x *SocialLink) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GeneralCss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardCss string `protobuf:"bytes,1,opt,name=cardCss,proto3" json:"cardCss,omitempty"`
	FontCss string `protobuf:"bytes,2,opt,name=fontCss,proto3" json:"fontCss,omitempty"`
}

func (x *GeneralCss) Reset() {
	*x = GeneralCss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userInfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralCss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralCss) ProtoMessage() {}

func (x *GeneralCss) ProtoReflect() protoreflect.Message {
	mi := &file_userInfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralCss.ProtoReflect.Descriptor instead.
func (*GeneralCss) Descriptor() ([]byte, []int) {
	return file_userInfo_proto_rawDescGZIP(), []int{2}
}

func (x *GeneralCss) GetCardCss() string {
	if x != nil {
		return x.CardCss
	}
	return ""
}

func (x *GeneralCss) GetFontCss() string {
	if x != nil {
		return x.FontCss
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID          string        `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Email           string        `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Mobile          string        `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Salt            string        `protobuf:"bytes,4,opt,name=salt,proto3" json:"salt,omitempty"`
	HashedPasswords string        `protobuf:"bytes,5,opt,name=hashedPasswords,proto3" json:"hashedPasswords,omitempty"`
	ProfileTitle    string        `protobuf:"bytes,6,opt,name=profileTitle,proto3" json:"profileTitle,omitempty"`
	AvatarUrl       string        `protobuf:"bytes,7,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
	BgImageUrl      string        `protobuf:"bytes,8,opt,name=bgImageUrl,proto3" json:"bgImageUrl,omitempty"`
	CardLinks       []*CardLink   `protobuf:"bytes,9,rep,name=cardLinks,proto3" json:"cardLinks,omitempty"`
	SocialLinks     []*SocialLink `protobuf:"bytes,10,rep,name=socialLinks,proto3" json:"socialLinks,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userInfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_userInfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_userInfo_proto_rawDescGZIP(), []int{3}
}

func (x *UserInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserInfo) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

func (x *UserInfo) GetHashedPasswords() string {
	if x != nil {
		return x.HashedPasswords
	}
	return ""
}

func (x *UserInfo) GetProfileTitle() string {
	if x != nil {
		return x.ProfileTitle
	}
	return ""
}

func (x *UserInfo) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UserInfo) GetBgImageUrl() string {
	if x != nil {
		return x.BgImageUrl
	}
	return ""
}

func (x *UserInfo) GetCardLinks() []*CardLink {
	if x != nil {
		return x.CardLinks
	}
	return nil
}

func (x *UserInfo) GetSocialLinks() []*SocialLink {
	if x != nil {
		return x.SocialLinks
	}
	return nil
}

var File_userInfo_proto protoreflect.FileDescriptor

var file_userInfo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x08, 0x63,
	0x61, 0x72, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x73, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x69, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x73, 0x73, 0x22, 0x3a, 0x0a,
	0x0a, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x63, 0x6f, 0x6e, 0x53, 0x56, 0x47, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63,
	0x6f, 0x6e, 0x53, 0x56, 0x47, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x40, 0x0a, 0x0a, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x43, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x43,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x64, 0x43, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x6e, 0x74, 0x43, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x6e, 0x74, 0x43, 0x73, 0x73, 0x22, 0xda, 0x02, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x61,
	0x6c, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x68, 0x61, 0x73,
	0x68, 0x65, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x30,
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x63, 0x61, 0x72,
	0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x09, 0x63, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x12, 0x36, 0x0a, 0x0b, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0b, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x54, 0x72, 0x65,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userInfo_proto_rawDescOnce sync.Once
	file_userInfo_proto_rawDescData = file_userInfo_proto_rawDesc
)

func file_userInfo_proto_rawDescGZIP() []byte {
	file_userInfo_proto_rawDescOnce.Do(func() {
		file_userInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_userInfo_proto_rawDescData)
	})
	return file_userInfo_proto_rawDescData
}

var file_userInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_userInfo_proto_goTypes = []interface{}{
	(*CardLink)(nil),   // 0: userInfo.cardLink
	(*SocialLink)(nil), // 1: userInfo.socialLink
	(*GeneralCss)(nil), // 2: userInfo.GeneralCss
	(*UserInfo)(nil),   // 3: userInfo.UserInfo
}
var file_userInfo_proto_depIdxs = []int32{
	0, // 0: userInfo.UserInfo.cardLinks:type_name -> userInfo.cardLink
	1, // 1: userInfo.UserInfo.socialLinks:type_name -> userInfo.socialLink
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_userInfo_proto_init() }
func file_userInfo_proto_init() {
	if File_userInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardLink); i {
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
		file_userInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialLink); i {
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
		file_userInfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralCss); i {
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
		file_userInfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
			RawDescriptor: file_userInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_userInfo_proto_goTypes,
		DependencyIndexes: file_userInfo_proto_depIdxs,
		MessageInfos:      file_userInfo_proto_msgTypes,
	}.Build()
	File_userInfo_proto = out.File
	file_userInfo_proto_rawDesc = nil
	file_userInfo_proto_goTypes = nil
	file_userInfo_proto_depIdxs = nil
}
