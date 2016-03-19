// Code generated by protoc-gen-gogo.
// source: handler.proto
// DO NOT EDIT!

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DataUpHandlerReq struct {
	Payload  []byte    `protobuf:"bytes,1,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
	Metadata *Metadata `protobuf:"bytes,2,opt,name=Metadata,json=metadata" json:"Metadata,omitempty"`
	AppEUI   []byte    `protobuf:"bytes,3,opt,name=AppEUI,json=appEUI,proto3" json:"AppEUI,omitempty"`
	DevEUI   []byte    `protobuf:"bytes,4,opt,name=DevEUI,json=devEUI,proto3" json:"DevEUI,omitempty"`
	FCnt     uint32    `protobuf:"varint,5,opt,name=FCnt,json=fCnt,proto3" json:"FCnt,omitempty"`
	MType    uint32    `protobuf:"varint,6,opt,name=MType,json=mType,proto3" json:"MType,omitempty"`
}

func (m *DataUpHandlerReq) Reset()                    { *m = DataUpHandlerReq{} }
func (m *DataUpHandlerReq) String() string            { return proto.CompactTextString(m) }
func (*DataUpHandlerReq) ProtoMessage()               {}
func (*DataUpHandlerReq) Descriptor() ([]byte, []int) { return fileDescriptorHandler, []int{0} }

func (m *DataUpHandlerReq) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type DataUpHandlerRes struct {
	Payload  *LoRaWANData `protobuf:"bytes,1,opt,name=Payload,json=payload" json:"Payload,omitempty"`
	Metadata *Metadata    `protobuf:"bytes,2,opt,name=Metadata,json=metadata" json:"Metadata,omitempty"`
}

func (m *DataUpHandlerRes) Reset()                    { *m = DataUpHandlerRes{} }
func (m *DataUpHandlerRes) String() string            { return proto.CompactTextString(m) }
func (*DataUpHandlerRes) ProtoMessage()               {}
func (*DataUpHandlerRes) Descriptor() ([]byte, []int) { return fileDescriptorHandler, []int{1} }

func (m *DataUpHandlerRes) GetPayload() *LoRaWANData {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *DataUpHandlerRes) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type DataDownHandlerReq struct {
	Payload []byte `protobuf:"bytes,1,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
	AppEUI  []byte `protobuf:"bytes,2,opt,name=AppEUI,json=appEUI,proto3" json:"AppEUI,omitempty"`
	DevEUI  []byte `protobuf:"bytes,3,opt,name=DevEUI,json=devEUI,proto3" json:"DevEUI,omitempty"`
}

func (m *DataDownHandlerReq) Reset()                    { *m = DataDownHandlerReq{} }
func (m *DataDownHandlerReq) String() string            { return proto.CompactTextString(m) }
func (*DataDownHandlerReq) ProtoMessage()               {}
func (*DataDownHandlerReq) Descriptor() ([]byte, []int) { return fileDescriptorHandler, []int{2} }

type DataDownHandlerRes struct {
}

func (m *DataDownHandlerRes) Reset()                    { *m = DataDownHandlerRes{} }
func (m *DataDownHandlerRes) String() string            { return proto.CompactTextString(m) }
func (*DataDownHandlerRes) ProtoMessage()               {}
func (*DataDownHandlerRes) Descriptor() ([]byte, []int) { return fileDescriptorHandler, []int{3} }

type ABPSubHandlerReq struct {
	AppEUI  []byte `protobuf:"bytes,1,opt,name=AppEUI,json=appEUI,proto3" json:"AppEUI,omitempty"`
	DevAddr []byte `protobuf:"bytes,2,opt,name=DevAddr,json=devAddr,proto3" json:"DevAddr,omitempty"`
	NwkSKey []byte `protobuf:"bytes,3,opt,name=NwkSKey,json=nwkSKey,proto3" json:"NwkSKey,omitempty"`
	AppSKey []byte `protobuf:"bytes,4,opt,name=AppSKey,json=appSKey,proto3" json:"AppSKey,omitempty"`
}

func (m *ABPSubHandlerReq) Reset()                    { *m = ABPSubHandlerReq{} }
func (m *ABPSubHandlerReq) String() string            { return proto.CompactTextString(m) }
func (*ABPSubHandlerReq) ProtoMessage()               {}
func (*ABPSubHandlerReq) Descriptor() ([]byte, []int) { return fileDescriptorHandler, []int{4} }

type ABPSubHandlerRes struct {
}

func (m *ABPSubHandlerRes) Reset()                    { *m = ABPSubHandlerRes{} }
func (m *ABPSubHandlerRes) String() string            { return proto.CompactTextString(m) }
func (*ABPSubHandlerRes) ProtoMessage()               {}
func (*ABPSubHandlerRes) Descriptor() ([]byte, []int) { return fileDescriptorHandler, []int{5} }

func init() {
	proto.RegisterType((*DataUpHandlerReq)(nil), "core.DataUpHandlerReq")
	proto.RegisterType((*DataUpHandlerRes)(nil), "core.DataUpHandlerRes")
	proto.RegisterType((*DataDownHandlerReq)(nil), "core.DataDownHandlerReq")
	proto.RegisterType((*DataDownHandlerRes)(nil), "core.DataDownHandlerRes")
	proto.RegisterType((*ABPSubHandlerReq)(nil), "core.ABPSubHandlerReq")
	proto.RegisterType((*ABPSubHandlerRes)(nil), "core.ABPSubHandlerRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Handler service

type HandlerClient interface {
	HandleDataUp(ctx context.Context, in *DataUpHandlerReq, opts ...grpc.CallOption) (*DataUpHandlerRes, error)
	HandleDataDown(ctx context.Context, in *DataDownHandlerReq, opts ...grpc.CallOption) (*DataDownHandlerRes, error)
	SubscribePersonalized(ctx context.Context, in *ABPSubHandlerReq, opts ...grpc.CallOption) (*ABPSubHandlerRes, error)
}

type handlerClient struct {
	cc *grpc.ClientConn
}

func NewHandlerClient(cc *grpc.ClientConn) HandlerClient {
	return &handlerClient{cc}
}

func (c *handlerClient) HandleDataUp(ctx context.Context, in *DataUpHandlerReq, opts ...grpc.CallOption) (*DataUpHandlerRes, error) {
	out := new(DataUpHandlerRes)
	err := grpc.Invoke(ctx, "/core.Handler/HandleDataUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerClient) HandleDataDown(ctx context.Context, in *DataDownHandlerReq, opts ...grpc.CallOption) (*DataDownHandlerRes, error) {
	out := new(DataDownHandlerRes)
	err := grpc.Invoke(ctx, "/core.Handler/HandleDataDown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerClient) SubscribePersonalized(ctx context.Context, in *ABPSubHandlerReq, opts ...grpc.CallOption) (*ABPSubHandlerRes, error) {
	out := new(ABPSubHandlerRes)
	err := grpc.Invoke(ctx, "/core.Handler/SubscribePersonalized", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Handler service

type HandlerServer interface {
	HandleDataUp(context.Context, *DataUpHandlerReq) (*DataUpHandlerRes, error)
	HandleDataDown(context.Context, *DataDownHandlerReq) (*DataDownHandlerRes, error)
	SubscribePersonalized(context.Context, *ABPSubHandlerReq) (*ABPSubHandlerRes, error)
}

func RegisterHandlerServer(s *grpc.Server, srv HandlerServer) {
	s.RegisterService(&_Handler_serviceDesc, srv)
}

func _Handler_HandleDataUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DataUpHandlerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HandlerServer).HandleDataUp(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Handler_HandleDataDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DataDownHandlerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HandlerServer).HandleDataDown(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Handler_SubscribePersonalized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ABPSubHandlerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HandlerServer).SubscribePersonalized(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Handler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.Handler",
	HandlerType: (*HandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleDataUp",
			Handler:    _Handler_HandleDataUp_Handler,
		},
		{
			MethodName: "HandleDataDown",
			Handler:    _Handler_HandleDataDown_Handler,
		},
		{
			MethodName: "SubscribePersonalized",
			Handler:    _Handler_SubscribePersonalized_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func (m *DataUpHandlerReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DataUpHandlerReq) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		if len(m.Payload) > 0 {
			data[i] = 0xa
			i++
			i = encodeVarintHandler(data, i, uint64(len(m.Payload)))
			i += copy(data[i:], m.Payload)
		}
	}
	if m.Metadata != nil {
		data[i] = 0x12
		i++
		i = encodeVarintHandler(data, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.AppEUI != nil {
		if len(m.AppEUI) > 0 {
			data[i] = 0x1a
			i++
			i = encodeVarintHandler(data, i, uint64(len(m.AppEUI)))
			i += copy(data[i:], m.AppEUI)
		}
	}
	if m.DevEUI != nil {
		if len(m.DevEUI) > 0 {
			data[i] = 0x22
			i++
			i = encodeVarintHandler(data, i, uint64(len(m.DevEUI)))
			i += copy(data[i:], m.DevEUI)
		}
	}
	if m.FCnt != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintHandler(data, i, uint64(m.FCnt))
	}
	if m.MType != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintHandler(data, i, uint64(m.MType))
	}
	return i, nil
}

func (m *DataUpHandlerRes) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DataUpHandlerRes) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		data[i] = 0xa
		i++
		i = encodeVarintHandler(data, i, uint64(m.Payload.Size()))
		n2, err := m.Payload.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Metadata != nil {
		data[i] = 0x12
		i++
		i = encodeVarintHandler(data, i, uint64(m.Metadata.Size()))
		n3, err := m.Metadata.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *DataDownHandlerReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DataDownHandlerReq) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		if len(m.Payload) > 0 {
			data[i] = 0xa
			i++
			i = encodeVarintHandler(data, i, uint64(len(m.Payload)))
			i += copy(data[i:], m.Payload)
		}
	}
	if m.AppEUI != nil {
		if len(m.AppEUI) > 0 {
			data[i] = 0x12
			i++
			i = encodeVarintHandler(data, i, uint64(len(m.AppEUI)))
			i += copy(data[i:], m.AppEUI)
		}
	}
	if m.DevEUI != nil {
		if len(m.DevEUI) > 0 {
			data[i] = 0x1a
			i++
			i = encodeVarintHandler(data, i, uint64(len(m.DevEUI)))
			i += copy(data[i:], m.DevEUI)
		}
	}
	return i, nil
}

func (m *DataDownHandlerRes) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DataDownHandlerRes) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ABPSubHandlerReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ABPSubHandlerReq) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AppEUI != nil {
		if len(m.AppEUI) > 0 {
			data[i] = 0xa
			i++
			i = encodeVarintHandler(data, i, uint64(len(m.AppEUI)))
			i += copy(data[i:], m.AppEUI)
		}
	}
	if m.DevAddr != nil {
		if len(m.DevAddr) > 0 {
			data[i] = 0x12
			i++
			i = encodeVarintHandler(data, i, uint64(len(m.DevAddr)))
			i += copy(data[i:], m.DevAddr)
		}
	}
	if m.NwkSKey != nil {
		if len(m.NwkSKey) > 0 {
			data[i] = 0x1a
			i++
			i = encodeVarintHandler(data, i, uint64(len(m.NwkSKey)))
			i += copy(data[i:], m.NwkSKey)
		}
	}
	if m.AppSKey != nil {
		if len(m.AppSKey) > 0 {
			data[i] = 0x22
			i++
			i = encodeVarintHandler(data, i, uint64(len(m.AppSKey)))
			i += copy(data[i:], m.AppSKey)
		}
	}
	return i, nil
}

func (m *ABPSubHandlerRes) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ABPSubHandlerRes) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64Handler(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Handler(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHandler(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *DataUpHandlerReq) Size() (n int) {
	var l int
	_ = l
	if m.Payload != nil {
		l = len(m.Payload)
		if l > 0 {
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovHandler(uint64(l))
	}
	if m.AppEUI != nil {
		l = len(m.AppEUI)
		if l > 0 {
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if m.DevEUI != nil {
		l = len(m.DevEUI)
		if l > 0 {
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if m.FCnt != 0 {
		n += 1 + sovHandler(uint64(m.FCnt))
	}
	if m.MType != 0 {
		n += 1 + sovHandler(uint64(m.MType))
	}
	return n
}

func (m *DataUpHandlerRes) Size() (n int) {
	var l int
	_ = l
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovHandler(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovHandler(uint64(l))
	}
	return n
}

func (m *DataDownHandlerReq) Size() (n int) {
	var l int
	_ = l
	if m.Payload != nil {
		l = len(m.Payload)
		if l > 0 {
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if m.AppEUI != nil {
		l = len(m.AppEUI)
		if l > 0 {
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if m.DevEUI != nil {
		l = len(m.DevEUI)
		if l > 0 {
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	return n
}

func (m *DataDownHandlerRes) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ABPSubHandlerReq) Size() (n int) {
	var l int
	_ = l
	if m.AppEUI != nil {
		l = len(m.AppEUI)
		if l > 0 {
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if m.DevAddr != nil {
		l = len(m.DevAddr)
		if l > 0 {
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if m.NwkSKey != nil {
		l = len(m.NwkSKey)
		if l > 0 {
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if m.AppSKey != nil {
		l = len(m.AppSKey)
		if l > 0 {
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	return n
}

func (m *ABPSubHandlerRes) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovHandler(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHandler(x uint64) (n int) {
	return sovHandler(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataUpHandlerReq) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataUpHandlerReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataUpHandlerReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], data[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppEUI = append(m.AppEUI[:0], data[iNdEx:postIndex]...)
			if m.AppEUI == nil {
				m.AppEUI = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DevEUI = append(m.DevEUI[:0], data[iNdEx:postIndex]...)
			if m.DevEUI == nil {
				m.DevEUI = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FCnt", wireType)
			}
			m.FCnt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.FCnt |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MType", wireType)
			}
			m.MType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.MType |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DataUpHandlerRes) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataUpHandlerRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataUpHandlerRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &LoRaWANData{}
			}
			if err := m.Payload.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DataDownHandlerReq) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataDownHandlerReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataDownHandlerReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], data[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppEUI = append(m.AppEUI[:0], data[iNdEx:postIndex]...)
			if m.AppEUI == nil {
				m.AppEUI = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DevEUI = append(m.DevEUI[:0], data[iNdEx:postIndex]...)
			if m.DevEUI == nil {
				m.DevEUI = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DataDownHandlerRes) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataDownHandlerRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataDownHandlerRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ABPSubHandlerReq) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ABPSubHandlerReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ABPSubHandlerReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppEUI = append(m.AppEUI[:0], data[iNdEx:postIndex]...)
			if m.AppEUI == nil {
				m.AppEUI = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DevAddr = append(m.DevAddr[:0], data[iNdEx:postIndex]...)
			if m.DevAddr == nil {
				m.DevAddr = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NwkSKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NwkSKey = append(m.NwkSKey[:0], data[iNdEx:postIndex]...)
			if m.NwkSKey == nil {
				m.NwkSKey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppSKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppSKey = append(m.AppSKey[:0], data[iNdEx:postIndex]...)
			if m.AppSKey == nil {
				m.AppSKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ABPSubHandlerRes) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ABPSubHandlerRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ABPSubHandlerRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHandler(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHandler
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHandler
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHandler(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHandler = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHandler   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorHandler = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xbe, 0x69, 0xd3, 0xa4, 0x4c, 0x7f, 0xe8, 0x1d, 0x7a, 0x4b, 0xc8, 0xa2, 0x5c, 0xb2, 0x12,
	0x85, 0x2e, 0xea, 0x5e, 0x48, 0x8d, 0x7f, 0x68, 0x4b, 0x49, 0x2d, 0xee, 0x84, 0x69, 0x67, 0xc4,
	0xd2, 0x34, 0x13, 0x33, 0xd1, 0x5a, 0x9f, 0xc4, 0xe7, 0xf0, 0x29, 0x5c, 0xba, 0x74, 0x29, 0xfa,
	0x22, 0xce, 0x64, 0x52, 0xd2, 0x5f, 0x10, 0x17, 0x03, 0xe7, 0xfb, 0xce, 0x9c, 0xef, 0x7c, 0xe7,
	0x4c, 0x02, 0x4a, 0xb7, 0xc8, 0xc7, 0x1e, 0x09, 0x1b, 0x41, 0x48, 0x23, 0x0a, 0xd5, 0x21, 0x0d,
	0x89, 0x59, 0xf2, 0x68, 0x88, 0xa6, 0xc8, 0x97, 0xa4, 0x09, 0x04, 0x29, 0x63, 0xeb, 0x45, 0x01,
	0x15, 0x07, 0x45, 0xa8, 0x1f, 0x9c, 0xca, 0x42, 0x97, 0xdc, 0x41, 0x03, 0xe8, 0x5d, 0x34, 0xf3,
	0x28, 0xc2, 0x86, 0xf2, 0x5f, 0xd9, 0x29, 0xba, 0x7a, 0x20, 0x21, 0xdc, 0x05, 0xf9, 0x36, 0x89,
	0x10, 0xe6, 0x15, 0x46, 0x86, 0xa7, 0x0a, 0xcd, 0x72, 0x23, 0x56, 0x9b, 0xb3, 0x6e, 0x7e, 0x92,
	0x44, 0xb0, 0x06, 0x34, 0x3b, 0x08, 0x8e, 0xfa, 0x67, 0x46, 0x36, 0x16, 0xd1, 0x50, 0x8c, 0x04,
	0xef, 0x90, 0x07, 0xc1, 0xab, 0x92, 0xc7, 0x31, 0x82, 0x10, 0xa8, 0xc7, 0x87, 0x7e, 0x64, 0xe4,
	0x38, 0x5b, 0x72, 0xd5, 0x1b, 0x1e, 0xc3, 0x2a, 0xc8, 0xb5, 0x2f, 0x67, 0x01, 0x31, 0xb4, 0x98,
	0xcc, 0x4d, 0x04, 0xb0, 0xc6, 0x6b, 0x9e, 0x19, 0xdc, 0x5b, 0xf6, 0x5c, 0x68, 0xfe, 0x95, 0xc6,
	0x2e, 0xa8, 0x8b, 0xae, 0xec, 0x8e, 0xb8, 0xff, 0xab, 0x31, 0xac, 0x6b, 0x00, 0x45, 0xb1, 0x43,
	0xa7, 0xfe, 0x8f, 0x56, 0x94, 0x8e, 0x9d, 0xd9, 0x32, 0x76, 0x76, 0x71, 0x6c, 0xab, 0xba, 0x41,
	0x9f, 0x59, 0x8f, 0xa0, 0x62, 0xb7, 0xba, 0xbd, 0xfb, 0xc1, 0x42, 0xcf, 0x54, 0x59, 0x59, 0x52,
	0xe6, 0x5e, 0xb8, 0xb2, 0x8d, 0x71, 0x98, 0xb4, 0xd4, 0xb1, 0x84, 0x22, 0xd3, 0x99, 0x8e, 0x7b,
	0xe7, 0x64, 0x96, 0x34, 0xd5, 0x7d, 0x09, 0x45, 0x86, 0x6b, 0xc5, 0x19, 0xf9, 0x0a, 0x3a, 0x92,
	0xd0, 0x82, 0x6b, 0x9d, 0x59, 0xf3, 0x5d, 0x01, 0x7a, 0x02, 0xe1, 0x01, 0x28, 0xca, 0x50, 0x3e,
	0x01, 0xac, 0xc9, 0xcd, 0xad, 0x7e, 0x44, 0xe6, 0x66, 0x9e, 0x41, 0x07, 0x94, 0xd3, 0x7a, 0x31,
	0x35, 0x34, 0xd2, 0x9b, 0xcb, 0x5b, 0x36, 0xb7, 0x65, 0x18, 0x3c, 0x01, 0xff, 0xb8, 0x45, 0x36,
	0x0c, 0x47, 0x03, 0xd2, 0x25, 0x21, 0xa3, 0x3e, 0xf2, 0x46, 0x4f, 0x04, 0xcf, 0xed, 0xac, 0x2e,
	0xcf, 0xdc, 0xcc, 0xb3, 0x56, 0xe5, 0xf5, 0xb3, 0xae, 0xbc, 0xf1, 0xf3, 0xc1, 0xcf, 0xf3, 0x57,
	0xfd, 0xcf, 0x40, 0x8b, 0xff, 0x8c, 0xfd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xcb, 0xb7,
	0xb8, 0x4b, 0x03, 0x00, 0x00,
}