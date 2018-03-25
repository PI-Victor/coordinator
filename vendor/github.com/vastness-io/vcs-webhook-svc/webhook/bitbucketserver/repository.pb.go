// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repository.proto

package bitbucketserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Repository struct {
	Forkable      bool     `protobuf:"varint,1,opt,name=forkable" json:"forkable,omitempty"`
	Id            int64    `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Name          string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Project       *Project `protobuf:"bytes,4,opt,name=project" json:"project,omitempty"`
	Public        bool     `protobuf:"varint,5,opt,name=public" json:"public,omitempty"`
	ScmId         string   `protobuf:"bytes,6,opt,name=scmId" json:"scmId,omitempty"`
	Slug          string   `protobuf:"bytes,7,opt,name=slug" json:"slug,omitempty"`
	State         string   `protobuf:"bytes,8,opt,name=state" json:"state,omitempty"`
	StatusMessage string   `protobuf:"bytes,9,opt,name=statusMessage" json:"statusMessage,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *Repository) GetForkable() bool {
	if m != nil {
		return m.Forkable
	}
	return false
}

func (m *Repository) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Repository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Repository) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *Repository) GetPublic() bool {
	if m != nil {
		return m.Public
	}
	return false
}

func (m *Repository) GetScmId() string {
	if m != nil {
		return m.ScmId
	}
	return ""
}

func (m *Repository) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Repository) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Repository) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Repository)(nil), "bitbucketserver.Repository")
}

func init() { proto.RegisterFile("repository.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xd9, 0xb4, 0x4d, 0xd3, 0x91, 0xaa, 0x0c, 0x22, 0x43, 0x4f, 0x41, 0x3c, 0xe4, 0x94,
	0x43, 0xfd, 0x14, 0x1e, 0x04, 0xd9, 0x6f, 0x90, 0x4d, 0xc6, 0xb2, 0x36, 0x75, 0xc3, 0xfe, 0x11,
	0xfc, 0xf2, 0x22, 0xbb, 0x1b, 0x03, 0x7a, 0x7b, 0x6f, 0xdf, 0xdb, 0xf9, 0x0d, 0x03, 0xb7, 0x96,
	0x27, 0xe3, 0xb4, 0x37, 0xf6, 0xab, 0x9d, 0xac, 0xf1, 0x06, 0x6f, 0x94, 0xf6, 0x2a, 0xf4, 0x67,
	0xf6, 0x8e, 0xed, 0x27, 0xdb, 0xc3, 0x7e, 0xb2, 0xe6, 0x9d, 0x7b, 0x9f, 0xf3, 0x87, 0x6f, 0x01,
	0x20, 0x97, 0x4f, 0x78, 0x80, 0xea, 0xcd, 0xd8, 0x73, 0xa7, 0x46, 0x26, 0x51, 0x8b, 0xa6, 0x92,
	0x8b, 0xc7, 0x6b, 0x28, 0xf4, 0x40, 0x45, 0x2d, 0x9a, 0x95, 0x2c, 0xf4, 0x80, 0x08, 0xeb, 0x8f,
	0xee, 0xc2, 0xb4, 0xaa, 0x45, 0xb3, 0x93, 0x49, 0xe3, 0x11, 0xb6, 0xf3, 0x7c, 0x5a, 0xd7, 0xa2,
	0xb9, 0x3a, 0x52, 0xfb, 0x6f, 0x81, 0xf6, 0x35, 0xe7, 0xf2, 0xb7, 0x88, 0xf7, 0x50, 0x4e, 0x41,
	0x8d, 0xba, 0xa7, 0x4d, 0x22, 0xce, 0x0e, 0xef, 0x60, 0xe3, 0xfa, 0xcb, 0xf3, 0x40, 0x65, 0x02,
	0x64, 0x13, 0xa9, 0x6e, 0x0c, 0x27, 0xda, 0x66, 0x6a, 0xd4, 0xa9, 0xe9, 0x3b, 0xcf, 0x54, 0xcd,
	0xcd, 0x68, 0xf0, 0x11, 0xf6, 0x51, 0x04, 0xf7, 0xc2, 0xce, 0x75, 0x27, 0xa6, 0x5d, 0x4a, 0xff,
	0x3e, 0xaa, 0x32, 0xdd, 0xe1, 0xe9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xac, 0xbb, 0xb9, 0x3b,
	0x01, 0x00, 0x00,
}