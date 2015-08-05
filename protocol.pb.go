// Code generated by protoc-gen-go.
// source: protocol.proto
// DO NOT EDIT!

/*
Package kadiyadb is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	Metadata
	EpochMetrics
	Metrics
*/
package kadiyadb

import proto "github.com/golang/protobuf/proto"
import index "github.com/kadirahq/kadiyadb/index"
import block "github.com/kadirahq/kadiyadb/block"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Metadata struct {
	// path to store files
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// point resolution in nano seconds
	Resolution int64 `protobuf:"varint,2,opt,name=resolution" json:"resolution,omitempty"`
	// retention time in nano seconds
	Retention int64 `protobuf:"varint,3,opt,name=retention" json:"retention,omitempty"`
	// epoch duration in nano seconds
	Duration int64 `protobuf:"varint,4,opt,name=duration" json:"duration,omitempty"`
	// payload size in bytes
	PayloadSize uint32 `protobuf:"varint,5,opt,name=payloadSize" json:"payloadSize,omitempty"`
	// records per segment
	SegmentSize uint32 `protobuf:"varint,6,opt,name=segmentSize" json:"segmentSize,omitempty"`
	// maximum read-only epochs
	MaxROEpochs uint32 `protobuf:"varint,7,opt,name=maxROEpochs" json:"maxROEpochs,omitempty"`
	// maximum read-write epochs
	MaxRWEpochs uint32 `protobuf:"varint,8,opt,name=maxRWEpochs" json:"maxRWEpochs,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}

type EpochMetrics struct {
	// performance metrics of index
	Index *index.Metrics `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	// performance metrics of block
	Block *block.Metrics `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
}

func (m *EpochMetrics) Reset()         { *m = EpochMetrics{} }
func (m *EpochMetrics) String() string { return proto.CompactTextString(m) }
func (*EpochMetrics) ProtoMessage()    {}

func (m *EpochMetrics) GetIndex() *index.Metrics {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *EpochMetrics) GetBlock() *block.Metrics {
	if m != nil {
		return m.Block
	}
	return nil
}

type Metrics struct {
	REpochs map[int64]*EpochMetrics `protobuf:"bytes,1,rep,name=rEpochs" json:"rEpochs,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	WEpochs map[int64]*EpochMetrics `protobuf:"bytes,2,rep,name=wEpochs" json:"wEpochs,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}

func (m *Metrics) GetREpochs() map[int64]*EpochMetrics {
	if m != nil {
		return m.REpochs
	}
	return nil
}

func (m *Metrics) GetWEpochs() map[int64]*EpochMetrics {
	if m != nil {
		return m.WEpochs
	}
	return nil
}
