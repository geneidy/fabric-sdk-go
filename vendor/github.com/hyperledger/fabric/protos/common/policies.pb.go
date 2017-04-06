// Code generated by protoc-gen-go.
// source: common/policies.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Policy_PolicyType int32

const (
	Policy_UNKNOWN       Policy_PolicyType = 0
	Policy_SIGNATURE     Policy_PolicyType = 1
	Policy_MSP           Policy_PolicyType = 2
	Policy_IMPLICIT_META Policy_PolicyType = 3
)

var Policy_PolicyType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SIGNATURE",
	2: "MSP",
	3: "IMPLICIT_META",
}
var Policy_PolicyType_value = map[string]int32{
	"UNKNOWN":       0,
	"SIGNATURE":     1,
	"MSP":           2,
	"IMPLICIT_META": 3,
}

func (x Policy_PolicyType) String() string {
	return proto.EnumName(Policy_PolicyType_name, int32(x))
}
func (Policy_PolicyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0, 0} }

type ImplicitMetaPolicy_Rule int32

const (
	ImplicitMetaPolicy_ANY      ImplicitMetaPolicy_Rule = 0
	ImplicitMetaPolicy_ALL      ImplicitMetaPolicy_Rule = 1
	ImplicitMetaPolicy_MAJORITY ImplicitMetaPolicy_Rule = 2
)

var ImplicitMetaPolicy_Rule_name = map[int32]string{
	0: "ANY",
	1: "ALL",
	2: "MAJORITY",
}
var ImplicitMetaPolicy_Rule_value = map[string]int32{
	"ANY":      0,
	"ALL":      1,
	"MAJORITY": 2,
}

func (x ImplicitMetaPolicy_Rule) String() string {
	return proto.EnumName(ImplicitMetaPolicy_Rule_name, int32(x))
}
func (ImplicitMetaPolicy_Rule) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{3, 0} }

// Policy expresses a policy which the orderer can evaluate, because there has been some desire expressed to support
// multiple policy engines, this is typed as a oneof for now
type Policy struct {
	Type   int32  `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Policy []byte `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

// SignaturePolicyEnvelope wraps a SignaturePolicy and includes a version for future enhancements
type SignaturePolicyEnvelope struct {
	Version    int32            `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Policy     *SignaturePolicy `protobuf:"bytes,2,opt,name=policy" json:"policy,omitempty"`
	Identities []*MSPPrincipal  `protobuf:"bytes,3,rep,name=identities" json:"identities,omitempty"`
}

func (m *SignaturePolicyEnvelope) Reset()                    { *m = SignaturePolicyEnvelope{} }
func (m *SignaturePolicyEnvelope) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicyEnvelope) ProtoMessage()               {}
func (*SignaturePolicyEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *SignaturePolicyEnvelope) GetPolicy() *SignaturePolicy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *SignaturePolicyEnvelope) GetIdentities() []*MSPPrincipal {
	if m != nil {
		return m.Identities
	}
	return nil
}

// SignaturePolicy is a recursive message structure which defines a featherweight DSL for describing
// policies which are more complicated than 'exactly this signature'.  The NOutOf operator is sufficent
// to express AND as well as OR, as well as of course N out of the following M policies
// SignedBy implies that the signature is from a valid certificate which is signed by the trusted
// authority specified in the bytes.  This will be the certificate itself for a self-signed certificate
// and will be the CA for more traditional certificates
type SignaturePolicy struct {
	// Types that are valid to be assigned to Type:
	//	*SignaturePolicy_SignedBy
	//	*SignaturePolicy_NOutOf_
	Type isSignaturePolicy_Type `protobuf_oneof:"Type"`
}

func (m *SignaturePolicy) Reset()                    { *m = SignaturePolicy{} }
func (m *SignaturePolicy) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicy) ProtoMessage()               {}
func (*SignaturePolicy) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

type isSignaturePolicy_Type interface {
	isSignaturePolicy_Type()
}

type SignaturePolicy_SignedBy struct {
	SignedBy int32 `protobuf:"varint,1,opt,name=signed_by,json=signedBy,oneof"`
}
type SignaturePolicy_NOutOf_ struct {
	NOutOf *SignaturePolicy_NOutOf `protobuf:"bytes,2,opt,name=n_out_of,json=nOutOf,oneof"`
}

func (*SignaturePolicy_SignedBy) isSignaturePolicy_Type() {}
func (*SignaturePolicy_NOutOf_) isSignaturePolicy_Type()  {}

func (m *SignaturePolicy) GetType() isSignaturePolicy_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SignaturePolicy) GetSignedBy() int32 {
	if x, ok := m.GetType().(*SignaturePolicy_SignedBy); ok {
		return x.SignedBy
	}
	return 0
}

func (m *SignaturePolicy) GetNOutOf() *SignaturePolicy_NOutOf {
	if x, ok := m.GetType().(*SignaturePolicy_NOutOf_); ok {
		return x.NOutOf
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SignaturePolicy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SignaturePolicy_OneofMarshaler, _SignaturePolicy_OneofUnmarshaler, _SignaturePolicy_OneofSizer, []interface{}{
		(*SignaturePolicy_SignedBy)(nil),
		(*SignaturePolicy_NOutOf_)(nil),
	}
}

func _SignaturePolicy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SignaturePolicy)
	// Type
	switch x := m.Type.(type) {
	case *SignaturePolicy_SignedBy:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.SignedBy))
	case *SignaturePolicy_NOutOf_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NOutOf); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SignaturePolicy.Type has unexpected type %T", x)
	}
	return nil
}

func _SignaturePolicy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SignaturePolicy)
	switch tag {
	case 1: // Type.signed_by
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &SignaturePolicy_SignedBy{int32(x)}
		return true, err
	case 2: // Type.n_out_of
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SignaturePolicy_NOutOf)
		err := b.DecodeMessage(msg)
		m.Type = &SignaturePolicy_NOutOf_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SignaturePolicy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SignaturePolicy)
	// Type
	switch x := m.Type.(type) {
	case *SignaturePolicy_SignedBy:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.SignedBy))
	case *SignaturePolicy_NOutOf_:
		s := proto.Size(x.NOutOf)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SignaturePolicy_NOutOf struct {
	N        int32              `protobuf:"varint,1,opt,name=N" json:"N,omitempty"`
	Policies []*SignaturePolicy `protobuf:"bytes,2,rep,name=policies" json:"policies,omitempty"`
}

func (m *SignaturePolicy_NOutOf) Reset()                    { *m = SignaturePolicy_NOutOf{} }
func (m *SignaturePolicy_NOutOf) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicy_NOutOf) ProtoMessage()               {}
func (*SignaturePolicy_NOutOf) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2, 0} }

func (m *SignaturePolicy_NOutOf) GetPolicies() []*SignaturePolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

// ImplicitMetaPolicy is a policy type which depends on the hierarchical nature of the configuration
// It is implicit because the rule is generate implicitly based on the number of sub policies
// It is meta because it depends only on the result of other policies
// When evaluated, this policy iterates over all immediate child sub-groups, retrieves the policy
// of name sub_policy, evaluates the collection and applies the rule.
// For example, with 4 sub-groups, and a policy name of "foo", ImplicitMetaPolicy retrieves
// each sub-group, retrieves policy "foo" for each subgroup, evaluates it, and, in the case of ANY
// 1 satisfied is sufficient, ALL would require 4 signatures, and MAJORITY would require 3 signatures.
type ImplicitMetaPolicy struct {
	SubPolicy string                  `protobuf:"bytes,1,opt,name=sub_policy,json=subPolicy" json:"sub_policy,omitempty"`
	Rule      ImplicitMetaPolicy_Rule `protobuf:"varint,2,opt,name=rule,enum=common.ImplicitMetaPolicy_Rule" json:"rule,omitempty"`
}

func (m *ImplicitMetaPolicy) Reset()                    { *m = ImplicitMetaPolicy{} }
func (m *ImplicitMetaPolicy) String() string            { return proto.CompactTextString(m) }
func (*ImplicitMetaPolicy) ProtoMessage()               {}
func (*ImplicitMetaPolicy) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func init() {
	proto.RegisterType((*Policy)(nil), "common.Policy")
	proto.RegisterType((*SignaturePolicyEnvelope)(nil), "common.SignaturePolicyEnvelope")
	proto.RegisterType((*SignaturePolicy)(nil), "common.SignaturePolicy")
	proto.RegisterType((*SignaturePolicy_NOutOf)(nil), "common.SignaturePolicy.NOutOf")
	proto.RegisterType((*ImplicitMetaPolicy)(nil), "common.ImplicitMetaPolicy")
	proto.RegisterEnum("common.Policy_PolicyType", Policy_PolicyType_name, Policy_PolicyType_value)
	proto.RegisterEnum("common.ImplicitMetaPolicy_Rule", ImplicitMetaPolicy_Rule_name, ImplicitMetaPolicy_Rule_value)
}

func init() { proto.RegisterFile("common/policies.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0x51, 0x8f, 0xd2, 0x40,
	0x10, 0x66, 0x01, 0x0b, 0x0c, 0x9c, 0xd6, 0x8d, 0x7a, 0x84, 0x44, 0x25, 0x7d, 0x30, 0x24, 0xc6,
	0x36, 0x01, 0x9f, 0x7c, 0x03, 0x25, 0x5e, 0x3d, 0x5a, 0x9a, 0xa5, 0x17, 0x73, 0xbe, 0x34, 0x14,
	0x16, 0x6e, 0x93, 0xd2, 0x6e, 0xda, 0xed, 0x25, 0xf5, 0x57, 0xf8, 0xec, 0xbf, 0xf1, 0x9f, 0x99,
	0xee, 0xb6, 0x17, 0xbd, 0xcb, 0xbd, 0xcd, 0xcc, 0x7e, 0xf3, 0xcd, 0xf7, 0xcd, 0x2c, 0xbc, 0xdc,
	0x25, 0xa7, 0x53, 0x12, 0x5b, 0x3c, 0x89, 0xd8, 0x8e, 0xd1, 0xcc, 0xe4, 0x69, 0x22, 0x12, 0xac,
	0xa9, 0xf2, 0x68, 0x54, 0x3d, 0x9f, 0x32, 0x1e, 0xf0, 0x94, 0xc5, 0x3b, 0xc6, 0xb7, 0x91, 0xc2,
	0x18, 0x3f, 0x41, 0xf3, 0xca, 0xae, 0x02, 0x63, 0x68, 0x8b, 0x82, 0xd3, 0x21, 0x1a, 0xa3, 0xc9,
	0x13, 0x22, 0x63, 0xfc, 0x0a, 0x34, 0xc9, 0x59, 0x0c, 0x9b, 0x63, 0x34, 0x19, 0x90, 0x2a, 0x33,
	0xbe, 0x00, 0xa8, 0x2e, 0xbf, 0x44, 0xf5, 0xa1, 0x73, 0xe5, 0x5e, 0xba, 0xeb, 0xef, 0xae, 0xde,
	0xc0, 0x67, 0xd0, 0xdb, 0xd8, 0x5f, 0xdd, 0xb9, 0x7f, 0x45, 0x96, 0x3a, 0xc2, 0x1d, 0x68, 0x39,
	0x1b, 0x4f, 0x6f, 0xe2, 0xe7, 0x70, 0x66, 0x3b, 0xde, 0xca, 0xfe, 0x6c, 0xfb, 0x81, 0xb3, 0xf4,
	0xe7, 0x7a, 0xcb, 0xf8, 0x8d, 0xe0, 0x7c, 0xc3, 0x8e, 0xf1, 0x56, 0xe4, 0x29, 0x55, 0x7c, 0xcb,
	0xf8, 0x96, 0x46, 0x09, 0xa7, 0x78, 0x08, 0x9d, 0x5b, 0x9a, 0x66, 0x2c, 0x89, 0x2b, 0x41, 0x75,
	0x8a, 0xad, 0xff, 0x34, 0xf5, 0xa7, 0xe7, 0xa6, 0xb2, 0x67, 0xde, 0xa3, 0xaa, 0xc5, 0xe2, 0x8f,
	0x00, 0x6c, 0x4f, 0x63, 0xc1, 0x04, 0xa3, 0xd9, 0xb0, 0x35, 0x6e, 0x4d, 0xfa, 0xd3, 0x17, 0x75,
	0x93, 0xb3, 0xf1, 0xbc, 0x7a, 0x25, 0xe4, 0x1f, 0x9c, 0xf1, 0x07, 0xc1, 0xb3, 0x7b, 0x8c, 0xf8,
	0x35, 0xf4, 0x32, 0x76, 0x8c, 0xe9, 0x3e, 0x08, 0x0b, 0x25, 0xeb, 0xa2, 0x41, 0xba, 0xaa, 0xb4,
	0x28, 0xf0, 0x27, 0xe8, 0xc6, 0x41, 0x92, 0x8b, 0x20, 0x39, 0x54, 0xda, 0xde, 0x3c, 0xa2, 0xcd,
	0x74, 0xd7, 0xb9, 0x58, 0x1f, 0x2e, 0x1a, 0x44, 0x8b, 0x65, 0x34, 0xba, 0x04, 0x4d, 0xd5, 0xf0,
	0x00, 0x90, 0x5b, 0x79, 0x46, 0x2e, 0x9e, 0x41, 0xb7, 0xbe, 0xea, 0xb0, 0x29, 0xa5, 0x3f, 0xea,
	0xf7, 0x0e, 0xb8, 0xd0, 0xa0, 0x5d, 0x1e, 0xc6, 0xf8, 0x85, 0x00, 0xdb, 0x27, 0x5e, 0x56, 0x85,
	0x43, 0xc5, 0xf6, 0xce, 0x06, 0x64, 0x79, 0x18, 0x54, 0x5b, 0x2c, 0x47, 0xf5, 0x48, 0x2f, 0xcb,
	0xc3, 0xea, 0x79, 0x06, 0xed, 0x34, 0x8f, 0xa8, 0xb4, 0xf0, 0x74, 0xfa, 0xb6, 0x1e, 0xf7, 0x90,
	0xc8, 0x24, 0x79, 0x44, 0x89, 0x04, 0x1b, 0xef, 0xa0, 0x5d, 0x66, 0xe5, 0xbd, 0xe7, 0xee, 0xb5,
	0xde, 0x90, 0xc1, 0x6a, 0xa5, 0x23, 0x3c, 0x80, 0xae, 0x33, 0xff, 0xb6, 0x26, 0xb6, 0x7f, 0xad,
	0x37, 0x17, 0x1f, 0x7e, 0xbc, 0x3f, 0x32, 0x71, 0x93, 0x87, 0x25, 0xad, 0x75, 0x53, 0x70, 0x9a,
	0x46, 0x74, 0x7f, 0xa4, 0xa9, 0x75, 0xd8, 0x86, 0x29, 0xdb, 0x59, 0xf2, 0x5b, 0x66, 0x96, 0x1a,
	0x1a, 0x6a, 0x32, 0x9d, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x29, 0x52, 0xc3, 0xda, 0xe2, 0x02,
	0x00, 0x00,
}