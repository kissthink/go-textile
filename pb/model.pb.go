// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Type controls read (R), annotate (A), and write (W) access
type Thread_Type int32

const (
	Thread_Private  Thread_Type = 0
	Thread_ReadOnly Thread_Type = 1
	Thread_Public   Thread_Type = 2
	Thread_Open     Thread_Type = 3
)

var Thread_Type_name = map[int32]string{
	0: "Private",
	1: "ReadOnly",
	2: "Public",
	3: "Open",
}
var Thread_Type_value = map[string]int32{
	"Private":  0,
	"ReadOnly": 1,
	"Public":   2,
	"Open":     3,
}

func (x Thread_Type) String() string {
	return proto.EnumName(Thread_Type_name, int32(x))
}
func (Thread_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_f2c8651a2e5bd178, []int{1, 0}
}

// Sharing controls if (Y/N) a thread can be shared
type Thread_Sharing int32

const (
	Thread_NotShared  Thread_Sharing = 0
	Thread_InviteOnly Thread_Sharing = 1
	Thread_Shared     Thread_Sharing = 2
)

var Thread_Sharing_name = map[int32]string{
	0: "NotShared",
	1: "InviteOnly",
	2: "Shared",
}
var Thread_Sharing_value = map[string]int32{
	"NotShared":  0,
	"InviteOnly": 1,
	"Shared":     2,
}

func (x Thread_Sharing) String() string {
	return proto.EnumName(Thread_Sharing_name, int32(x))
}
func (Thread_Sharing) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_f2c8651a2e5bd178, []int{1, 1}
}

// State indicates the loading state
type Thread_State int32

const (
	Thread_LoadingBehind Thread_State = 0
	Thread_Loaded        Thread_State = 1
	Thread_LoadingAhead  Thread_State = 2
)

var Thread_State_name = map[int32]string{
	0: "LoadingBehind",
	1: "Loaded",
	2: "LoadingAhead",
}
var Thread_State_value = map[string]int32{
	"LoadingBehind": 0,
	"Loaded":        1,
	"LoadingAhead":  2,
}

func (x Thread_State) String() string {
	return proto.EnumName(Thread_State_name, int32(x))
}
func (Thread_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_f2c8651a2e5bd178, []int{1, 2}
}

type Block_BlockType int32

const (
	Block_MERGE    Block_BlockType = 0
	Block_IGNORE   Block_BlockType = 1
	Block_FLAG     Block_BlockType = 2
	Block_JOIN     Block_BlockType = 3
	Block_ANNOUNCE Block_BlockType = 4
	Block_LEAVE    Block_BlockType = 5
	Block_MESSAGE  Block_BlockType = 6
	Block_FILES    Block_BlockType = 7
	Block_COMMENT  Block_BlockType = 8
	Block_LIKE     Block_BlockType = 9
	Block_INVITE   Block_BlockType = 50
)

var Block_BlockType_name = map[int32]string{
	0:  "MERGE",
	1:  "IGNORE",
	2:  "FLAG",
	3:  "JOIN",
	4:  "ANNOUNCE",
	5:  "LEAVE",
	6:  "MESSAGE",
	7:  "FILES",
	8:  "COMMENT",
	9:  "LIKE",
	50: "INVITE",
}
var Block_BlockType_value = map[string]int32{
	"MERGE":    0,
	"IGNORE":   1,
	"FLAG":     2,
	"JOIN":     3,
	"ANNOUNCE": 4,
	"LEAVE":    5,
	"MESSAGE":  6,
	"FILES":    7,
	"COMMENT":  8,
	"LIKE":     9,
	"INVITE":   50,
}

func (x Block_BlockType) String() string {
	return proto.EnumName(Block_BlockType_name, int32(x))
}
func (Block_BlockType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_f2c8651a2e5bd178, []int{2, 0}
}

type Contact struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string               `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Username             string               `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Avatar               string               `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Inboxes              []*Cafe              `protobuf:"bytes,5,rep,name=inboxes,proto3" json:"inboxes,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_f2c8651a2e5bd178, []int{0}
}
func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (dst *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(dst, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Contact) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Contact) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Contact) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Contact) GetInboxes() []*Cafe {
	if m != nil {
		return m.Inboxes
	}
	return nil
}

func (m *Contact) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Contact) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

type Thread struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Sk                   []byte         `protobuf:"bytes,3,opt,name=sk,proto3" json:"sk,omitempty"`
	Name                 string         `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Schema               string         `protobuf:"bytes,5,opt,name=schema,proto3" json:"schema,omitempty"`
	Initiator            string         `protobuf:"bytes,6,opt,name=initiator,proto3" json:"initiator,omitempty"`
	Type                 Thread_Type    `protobuf:"varint,7,opt,name=type,proto3,enum=Thread_Type" json:"type,omitempty"`
	Sharing              Thread_Sharing `protobuf:"varint,8,opt,name=sharing,proto3,enum=Thread_Sharing" json:"sharing,omitempty"`
	Members              []string       `protobuf:"bytes,9,rep,name=members,proto3" json:"members,omitempty"`
	State                Thread_State   `protobuf:"varint,10,opt,name=state,proto3,enum=Thread_State" json:"state,omitempty"`
	Head                 string         `protobuf:"bytes,11,opt,name=head,proto3" json:"head,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Thread) Reset()         { *m = Thread{} }
func (m *Thread) String() string { return proto.CompactTextString(m) }
func (*Thread) ProtoMessage()    {}
func (*Thread) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_f2c8651a2e5bd178, []int{1}
}
func (m *Thread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Thread.Unmarshal(m, b)
}
func (m *Thread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Thread.Marshal(b, m, deterministic)
}
func (dst *Thread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Thread.Merge(dst, src)
}
func (m *Thread) XXX_Size() int {
	return xxx_messageInfo_Thread.Size(m)
}
func (m *Thread) XXX_DiscardUnknown() {
	xxx_messageInfo_Thread.DiscardUnknown(m)
}

var xxx_messageInfo_Thread proto.InternalMessageInfo

func (m *Thread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Thread) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Thread) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *Thread) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Thread) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *Thread) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *Thread) GetType() Thread_Type {
	if m != nil {
		return m.Type
	}
	return Thread_Private
}

func (m *Thread) GetSharing() Thread_Sharing {
	if m != nil {
		return m.Sharing
	}
	return Thread_NotShared
}

func (m *Thread) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Thread) GetState() Thread_State {
	if m != nil {
		return m.State
	}
	return Thread_LoadingBehind
}

func (m *Thread) GetHead() string {
	if m != nil {
		return m.Head
	}
	return ""
}

type Block struct {
	Id      string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Thread  string               `protobuf:"bytes,2,opt,name=thread,proto3" json:"thread,omitempty"`
	Author  string               `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Type    Block_BlockType      `protobuf:"varint,4,opt,name=type,proto3,enum=Block_BlockType" json:"type,omitempty"`
	Date    *timestamp.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Parents []string             `protobuf:"bytes,6,rep,name=parents,proto3" json:"parents,omitempty"`
	Target  string               `protobuf:"bytes,7,opt,name=target,proto3" json:"target,omitempty"`
	Body    string               `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	// view info
	Username             string   `protobuf:"bytes,101,opt,name=username,proto3" json:"username,omitempty"`
	Avatar               string   `protobuf:"bytes,102,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_f2c8651a2e5bd178, []int{2}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Block) GetThread() string {
	if m != nil {
		return m.Thread
	}
	return ""
}

func (m *Block) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Block) GetType() Block_BlockType {
	if m != nil {
		return m.Type
	}
	return Block_MERGE
}

func (m *Block) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Block) GetParents() []string {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *Block) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Block) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Block) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Block) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type FileIndex struct {
	Mill                 string               `protobuf:"bytes,1,opt,name=mill,proto3" json:"mill,omitempty"`
	Checksum             string               `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Source               string               `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Opts                 string               `protobuf:"bytes,4,opt,name=opts,proto3" json:"opts,omitempty"`
	Hash                 string               `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	Key                  string               `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Media                string               `protobuf:"bytes,7,opt,name=media,proto3" json:"media,omitempty"`
	Name                 string               `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64                `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	Added                *timestamp.Timestamp `protobuf:"bytes,10,opt,name=added,proto3" json:"added,omitempty"`
	Meta                 *_struct.Struct      `protobuf:"bytes,11,opt,name=meta,proto3" json:"meta,omitempty"`
	Targets              []string             `protobuf:"bytes,12,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FileIndex) Reset()         { *m = FileIndex{} }
func (m *FileIndex) String() string { return proto.CompactTextString(m) }
func (*FileIndex) ProtoMessage()    {}
func (*FileIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_f2c8651a2e5bd178, []int{3}
}
func (m *FileIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileIndex.Unmarshal(m, b)
}
func (m *FileIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileIndex.Marshal(b, m, deterministic)
}
func (dst *FileIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileIndex.Merge(dst, src)
}
func (m *FileIndex) XXX_Size() int {
	return xxx_messageInfo_FileIndex.Size(m)
}
func (m *FileIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_FileIndex.DiscardUnknown(m)
}

var xxx_messageInfo_FileIndex proto.InternalMessageInfo

func (m *FileIndex) GetMill() string {
	if m != nil {
		return m.Mill
	}
	return ""
}

func (m *FileIndex) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *FileIndex) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *FileIndex) GetOpts() string {
	if m != nil {
		return m.Opts
	}
	return ""
}

func (m *FileIndex) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *FileIndex) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FileIndex) GetMedia() string {
	if m != nil {
		return m.Media
	}
	return ""
}

func (m *FileIndex) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileIndex) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileIndex) GetAdded() *timestamp.Timestamp {
	if m != nil {
		return m.Added
	}
	return nil
}

func (m *FileIndex) GetMeta() *_struct.Struct {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *FileIndex) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type Cafe struct {
	Peer                 string   `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Api                  string   `protobuf:"bytes,3,opt,name=api,proto3" json:"api,omitempty"`
	Protocol             string   `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Node                 string   `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
	Url                  string   `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Swarm                []string `protobuf:"bytes,7,rep,name=swarm,proto3" json:"swarm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cafe) Reset()         { *m = Cafe{} }
func (m *Cafe) String() string { return proto.CompactTextString(m) }
func (*Cafe) ProtoMessage()    {}
func (*Cafe) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_f2c8651a2e5bd178, []int{4}
}
func (m *Cafe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cafe.Unmarshal(m, b)
}
func (m *Cafe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cafe.Marshal(b, m, deterministic)
}
func (dst *Cafe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cafe.Merge(dst, src)
}
func (m *Cafe) XXX_Size() int {
	return xxx_messageInfo_Cafe.Size(m)
}
func (m *Cafe) XXX_DiscardUnknown() {
	xxx_messageInfo_Cafe.DiscardUnknown(m)
}

var xxx_messageInfo_Cafe proto.InternalMessageInfo

func (m *Cafe) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *Cafe) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Cafe) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *Cafe) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Cafe) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Cafe) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Cafe) GetSwarm() []string {
	if m != nil {
		return m.Swarm
	}
	return nil
}

type CafeSession struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Access               string               `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
	Exp                  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=exp,proto3" json:"exp,omitempty"`
	Refresh              string               `protobuf:"bytes,4,opt,name=refresh,proto3" json:"refresh,omitempty"`
	Rexp                 *timestamp.Timestamp `protobuf:"bytes,5,opt,name=rexp,proto3" json:"rexp,omitempty"`
	Subject              string               `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject,omitempty"`
	Type                 string               `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Cafe                 *Cafe                `protobuf:"bytes,8,opt,name=cafe,proto3" json:"cafe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CafeSession) Reset()         { *m = CafeSession{} }
func (m *CafeSession) String() string { return proto.CompactTextString(m) }
func (*CafeSession) ProtoMessage()    {}
func (*CafeSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_f2c8651a2e5bd178, []int{5}
}
func (m *CafeSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeSession.Unmarshal(m, b)
}
func (m *CafeSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeSession.Marshal(b, m, deterministic)
}
func (dst *CafeSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeSession.Merge(dst, src)
}
func (m *CafeSession) XXX_Size() int {
	return xxx_messageInfo_CafeSession.Size(m)
}
func (m *CafeSession) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeSession.DiscardUnknown(m)
}

var xxx_messageInfo_CafeSession proto.InternalMessageInfo

func (m *CafeSession) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeSession) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *CafeSession) GetExp() *timestamp.Timestamp {
	if m != nil {
		return m.Exp
	}
	return nil
}

func (m *CafeSession) GetRefresh() string {
	if m != nil {
		return m.Refresh
	}
	return ""
}

func (m *CafeSession) GetRexp() *timestamp.Timestamp {
	if m != nil {
		return m.Rexp
	}
	return nil
}

func (m *CafeSession) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *CafeSession) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CafeSession) GetCafe() *Cafe {
	if m != nil {
		return m.Cafe
	}
	return nil
}

func init() {
	proto.RegisterType((*Contact)(nil), "Contact")
	proto.RegisterType((*Thread)(nil), "Thread")
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterType((*FileIndex)(nil), "FileIndex")
	proto.RegisterType((*Cafe)(nil), "Cafe")
	proto.RegisterType((*CafeSession)(nil), "CafeSession")
	proto.RegisterEnum("Thread_Type", Thread_Type_name, Thread_Type_value)
	proto.RegisterEnum("Thread_Sharing", Thread_Sharing_name, Thread_Sharing_value)
	proto.RegisterEnum("Thread_State", Thread_State_name, Thread_State_value)
	proto.RegisterEnum("Block_BlockType", Block_BlockType_name, Block_BlockType_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_model_f2c8651a2e5bd178) }

var fileDescriptor_model_f2c8651a2e5bd178 = []byte{
	// 960 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xe4, 0x44,
	0x10, 0x5e, 0x8f, 0x3d, 0x3f, 0xae, 0x49, 0x82, 0x69, 0xa1, 0xc5, 0x44, 0x2b, 0xed, 0xc8, 0x70,
	0x08, 0x02, 0xcd, 0xa2, 0xb0, 0x12, 0x5c, 0x93, 0xc8, 0x89, 0x06, 0x92, 0x99, 0x95, 0x27, 0xec,
	0x81, 0x5b, 0x8f, 0xbb, 0x92, 0x31, 0xf1, 0x9f, 0xba, 0xdb, 0x21, 0xc3, 0x23, 0xf0, 0x10, 0x9c,
	0x78, 0x14, 0xae, 0xbc, 0x0d, 0xe2, 0x8c, 0xaa, 0xdd, 0x9e, 0x04, 0x02, 0x61, 0x2f, 0xa3, 0xfa,
	0xaa, 0xab, 0xcb, 0x5f, 0xfd, 0x7c, 0x3d, 0x30, 0x2e, 0x2a, 0x81, 0xf9, 0xb4, 0x96, 0x95, 0xae,
	0xf6, 0x5f, 0x5e, 0x57, 0xd5, 0x75, 0x8e, 0xaf, 0x0c, 0x5a, 0x35, 0x57, 0xaf, 0x74, 0x56, 0xa0,
	0xd2, 0xbc, 0xa8, 0x6d, 0xc0, 0x8b, 0x7f, 0x06, 0x28, 0x2d, 0x9b, 0x54, 0xb7, 0xa7, 0xd1, 0x1f,
	0x0e, 0x0c, 0x4f, 0xaa, 0x52, 0xf3, 0x54, 0xb3, 0x3d, 0xe8, 0x65, 0x22, 0x74, 0x26, 0xce, 0x81,
	0x9f, 0xf4, 0x32, 0xc1, 0x42, 0x18, 0x72, 0x21, 0x24, 0x2a, 0x15, 0xf6, 0x8c, 0xb3, 0x83, 0x6c,
	0x1f, 0x46, 0x8d, 0x42, 0x59, 0xf2, 0x02, 0x43, 0xd7, 0x1c, 0x6d, 0x31, 0x7b, 0x0e, 0x03, 0x7e,
	0xcb, 0x35, 0x97, 0xa1, 0x67, 0x4e, 0x2c, 0x62, 0x2f, 0x61, 0x98, 0x95, 0xab, 0xea, 0x0e, 0x55,
	0xd8, 0x9f, 0xb8, 0x07, 0xe3, 0xc3, 0xfe, 0xf4, 0x84, 0x5f, 0x61, 0xd2, 0x79, 0xd9, 0x6b, 0x18,
	0xa6, 0x12, 0xb9, 0x46, 0x11, 0x0e, 0x26, 0xce, 0xc1, 0xf8, 0x70, 0x7f, 0xda, 0x52, 0x9f, 0x76,
	0xd4, 0xa7, 0x97, 0x5d, 0x6d, 0x49, 0x17, 0x4a, 0xb7, 0x9a, 0x5a, 0x98, 0x5b, 0xc3, 0xff, 0xbf,
	0x65, 0x43, 0xa3, 0xdf, 0x5d, 0x18, 0x5c, 0xae, 0x25, 0x72, 0xf1, 0xa8, 0xea, 0x00, 0xdc, 0x1b,
	0xdc, 0xd8, 0x8a, 0xc9, 0xa4, 0x08, 0x75, 0x63, 0xea, 0xdc, 0x49, 0x7a, 0xea, 0x86, 0x31, 0xf0,
	0x4c, 0xe5, 0x6d, 0x7d, 0x5e, 0x57, 0xb5, 0x4a, 0xd7, 0x58, 0xf0, 0xb0, 0xdf, 0x56, 0xdd, 0x22,
	0xf6, 0x02, 0xfc, 0xac, 0xcc, 0x74, 0xc6, 0x75, 0x25, 0x4d, 0x59, 0x7e, 0x72, 0xef, 0x60, 0x13,
	0xf0, 0xf4, 0xa6, 0x46, 0xc3, 0x7c, 0xef, 0x70, 0x67, 0xda, 0x52, 0x9a, 0x5e, 0x6e, 0x6a, 0x4c,
	0xcc, 0x09, 0xfb, 0x14, 0x86, 0x6a, 0xcd, 0x65, 0x56, 0x5e, 0x87, 0x23, 0x13, 0xf4, 0x5e, 0x17,
	0xb4, 0x6c, 0xdd, 0x49, 0x77, 0x4e, 0xe3, 0x2a, 0xb0, 0x58, 0xa1, 0x54, 0xa1, 0x3f, 0x71, 0x69,
	0x5c, 0x16, 0xb2, 0x8f, 0xa1, 0xaf, 0x34, 0xd7, 0x18, 0x82, 0x49, 0xb1, 0xbb, 0x4d, 0x41, 0xce,
	0xa4, 0x3d, 0xa3, 0xaa, 0xd6, 0xc8, 0x45, 0x38, 0x6e, 0xab, 0x22, 0x3b, 0xfa, 0x0a, 0x3c, 0xe2,
	0xc2, 0xc6, 0x30, 0x7c, 0x23, 0xb3, 0x5b, 0xae, 0x31, 0x78, 0xc6, 0x76, 0x60, 0x94, 0x20, 0x17,
	0x8b, 0x32, 0xdf, 0x04, 0x0e, 0x03, 0x18, 0xbc, 0x69, 0x56, 0x79, 0x96, 0x06, 0x3d, 0x36, 0x02,
	0x6f, 0x51, 0x63, 0x19, 0xb8, 0xd1, 0x6b, 0x18, 0x5a, 0x7e, 0x6c, 0x17, 0xfc, 0x79, 0xa5, 0x09,
	0xa1, 0x08, 0x9e, 0xb1, 0x3d, 0x80, 0x59, 0x79, 0x9b, 0x69, 0xbc, 0xbf, 0x6f, 0xcf, 0x7a, 0xd1,
	0xd7, 0xd0, 0x37, 0x94, 0xd8, 0xfb, 0xb0, 0x7b, 0x5e, 0x71, 0x91, 0x95, 0xd7, 0xc7, 0xb8, 0xce,
	0x4a, 0xba, 0x07, 0x30, 0x20, 0x17, 0x8a, 0xc0, 0x61, 0x01, 0xec, 0xd8, 0xe3, 0x23, 0xa2, 0x19,
	0xf4, 0xa2, 0x5f, 0x5d, 0xe8, 0x1f, 0xe7, 0x55, 0x7a, 0xf3, 0x68, 0x9c, 0xcf, 0x61, 0xa0, 0x4d,
	0xb5, 0x76, 0xa2, 0x16, 0x99, 0x35, 0x6d, 0xf4, 0xba, 0x92, 0x76, 0x81, 0x2d, 0x62, 0x9f, 0xd8,
	0x91, 0x78, 0xa6, 0x55, 0xc1, 0xd4, 0x64, 0x6d, 0x7f, 0x1f, 0x8c, 0x65, 0x0a, 0x1e, 0x2d, 0x92,
	0x19, 0xf6, 0xd3, 0x2b, 0x67, 0xe2, 0x68, 0x36, 0x35, 0x97, 0x58, 0x6a, 0x15, 0x0e, 0xda, 0xd9,
	0x58, 0x68, 0xf8, 0x71, 0x79, 0x8d, 0xda, 0x2c, 0x01, 0xf1, 0x33, 0x88, 0xc6, 0xb1, 0xaa, 0xc4,
	0xc6, 0x4c, 0xdd, 0x4f, 0x8c, 0xfd, 0x37, 0xd9, 0xe1, 0x7f, 0xca, 0xee, 0xea, 0xa1, 0xec, 0xa2,
	0x9f, 0x1d, 0xf0, 0xb7, 0xec, 0x99, 0x0f, 0xfd, 0x8b, 0x38, 0x39, 0x8b, 0xdb, 0x86, 0xce, 0xce,
	0xe6, 0x8b, 0x24, 0x0e, 0x1c, 0x1a, 0xdc, 0xe9, 0xf9, 0xd1, 0x59, 0x3b, 0xc2, 0x6f, 0x16, 0xb3,
	0x79, 0xe0, 0xd2, 0x98, 0x8f, 0xe6, 0xf3, 0xc5, 0x77, 0xf3, 0x93, 0x38, 0xf0, 0xe8, 0xe2, 0x79,
	0x7c, 0xf4, 0x36, 0x0e, 0xfa, 0xb4, 0x0c, 0x17, 0xf1, 0x72, 0x79, 0x74, 0x16, 0x07, 0x03, 0xf2,
	0x9f, 0xce, 0xce, 0xe3, 0x65, 0x30, 0x24, 0xff, 0xc9, 0xe2, 0xe2, 0x22, 0x9e, 0x5f, 0x06, 0x23,
	0xca, 0x73, 0x3e, 0xfb, 0x36, 0x0e, 0x7c, 0xf3, 0x9d, 0xf9, 0xdb, 0xd9, 0x65, 0x1c, 0x1c, 0x46,
	0xbf, 0xf5, 0xc0, 0x3f, 0xcd, 0x72, 0x9c, 0x95, 0x02, 0xef, 0xa8, 0xc4, 0x22, 0xcb, 0x73, 0x3b,
	0x2c, 0x63, 0x53, 0x89, 0xe9, 0x1a, 0xd3, 0x1b, 0xd5, 0x14, 0x76, 0x60, 0x5b, 0x6c, 0x34, 0x56,
	0x35, 0x32, 0xed, 0xde, 0x1c, 0x8b, 0x28, 0x4f, 0x55, 0x6b, 0xd5, 0xe9, 0x91, 0x6c, 0xb3, 0xcd,
	0x5c, 0xad, 0xad, 0x1a, 0x8d, 0xdd, 0x29, 0x7b, 0x70, 0xaf, 0xec, 0x0f, 0xa0, 0x5f, 0xa0, 0xc8,
	0xb8, 0xed, 0x7d, 0x0b, 0xb6, 0xfa, 0x1e, 0x3d, 0xd0, 0x37, 0x03, 0x4f, 0x65, 0x3f, 0x61, 0xe8,
	0x4f, 0x9c, 0x03, 0x37, 0x31, 0x36, 0xfb, 0x02, 0xfa, 0x5c, 0x08, 0x14, 0x46, 0x56, 0x4f, 0x6f,
	0x41, 0x1b, 0xc8, 0x3e, 0x03, 0xaf, 0x40, 0xcd, 0x8d, 0xc6, 0xc6, 0x87, 0x1f, 0x3e, 0xba, 0xb0,
	0x34, 0x4f, 0x73, 0x62, 0x82, 0x68, 0x67, 0xda, 0x5d, 0x50, 0xe1, 0x4e, 0xbb, 0x33, 0x16, 0x46,
	0xbf, 0x38, 0xe0, 0xd1, 0xdb, 0x49, 0xac, 0x6a, 0x44, 0xd9, 0x75, 0x90, 0xec, 0x27, 0x5e, 0xed,
	0x00, 0x5c, 0x5e, 0x67, 0xb6, 0x79, 0x64, 0x52, 0xb7, 0xcd, 0xb7, 0xd3, 0x2a, 0xb7, 0xdd, 0xdb,
	0x62, 0xd3, 0x85, 0x4a, 0x60, 0xd7, 0x41, 0xb2, 0x29, 0x43, 0x23, 0xf3, 0xae, 0x83, 0x8d, 0xcc,
	0xa9, 0x83, 0xea, 0x47, 0x2e, 0x8b, 0x70, 0x68, 0x28, 0xb6, 0x20, 0xfa, 0xd3, 0x81, 0x31, 0x11,
	0x5c, 0xa2, 0x52, 0x59, 0x55, 0xfe, 0x9b, 0x28, 0x79, 0x9a, 0xde, 0x53, 0xb4, 0x88, 0x7d, 0x0e,
	0x2e, 0xde, 0xd5, 0x86, 0xe1, 0xd3, 0xfd, 0xa4, 0x30, 0xaa, 0x54, 0xe2, 0x95, 0x44, 0xb5, 0xb6,
	0xe4, 0x3b, 0x48, 0xf2, 0x94, 0x94, 0xe8, 0x1d, 0xe4, 0x29, 0x6d, 0x26, 0xd5, 0xac, 0x7e, 0xc0,
	0x54, 0xdb, 0xda, 0x3a, 0x48, 0x5d, 0xd8, 0xbe, 0xd0, 0xbe, 0x15, 0xff, 0x47, 0xe0, 0xa5, 0xfc,
	0xaa, 0xdd, 0x8f, 0xed, 0xdf, 0x98, 0x71, 0x1d, 0x7b, 0xdf, 0xf7, 0xea, 0xd5, 0x6a, 0x60, 0x3e,
	0xf4, 0xe5, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xd7, 0xfc, 0xa1, 0xa9, 0x07, 0x00, 0x00,
}
