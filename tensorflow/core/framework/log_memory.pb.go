// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/log_memory.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MemoryLogStep struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	// Handle describing the feeds and fetches of the step.
	Handle string `protobuf:"bytes,2,opt,name=handle" json:"handle,omitempty"`
}

func (m *MemoryLogStep) Reset()                    { *m = MemoryLogStep{} }
func (m *MemoryLogStep) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogStep) ProtoMessage()               {}
func (*MemoryLogStep) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *MemoryLogStep) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogStep) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

type MemoryLogTensorAllocation struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	// Name of the kernel making the allocation as set in GraphDef,
	// e.g., "affine2/weights/Assign".
	KernelName string `protobuf:"bytes,2,opt,name=kernel_name,json=kernelName" json:"kernel_name,omitempty"`
	// Allocated tensor details.
	Tensor *TensorDescription `protobuf:"bytes,3,opt,name=tensor" json:"tensor,omitempty"`
}

func (m *MemoryLogTensorAllocation) Reset()                    { *m = MemoryLogTensorAllocation{} }
func (m *MemoryLogTensorAllocation) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogTensorAllocation) ProtoMessage()               {}
func (*MemoryLogTensorAllocation) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *MemoryLogTensorAllocation) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogTensorAllocation) GetKernelName() string {
	if m != nil {
		return m.KernelName
	}
	return ""
}

func (m *MemoryLogTensorAllocation) GetTensor() *TensorDescription {
	if m != nil {
		return m.Tensor
	}
	return nil
}

type MemoryLogTensorDeallocation struct {
	// Id of the tensor buffer being deallocated, used to match to a
	// corresponding allocation.
	AllocationId int64 `protobuf:"varint,1,opt,name=allocation_id,json=allocationId" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,2,opt,name=allocator_name,json=allocatorName" json:"allocator_name,omitempty"`
}

func (m *MemoryLogTensorDeallocation) Reset()                    { *m = MemoryLogTensorDeallocation{} }
func (m *MemoryLogTensorDeallocation) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogTensorDeallocation) ProtoMessage()               {}
func (*MemoryLogTensorDeallocation) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *MemoryLogTensorDeallocation) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *MemoryLogTensorDeallocation) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

type MemoryLogTensorOutput struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	// Name of the kernel producing an output as set in GraphDef, e.g.,
	// "affine2/weights/Assign".
	KernelName string `protobuf:"bytes,2,opt,name=kernel_name,json=kernelName" json:"kernel_name,omitempty"`
	// Index of the output being set.
	Index int32 `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
	// Output tensor details.
	Tensor *TensorDescription `protobuf:"bytes,4,opt,name=tensor" json:"tensor,omitempty"`
}

func (m *MemoryLogTensorOutput) Reset()                    { *m = MemoryLogTensorOutput{} }
func (m *MemoryLogTensorOutput) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogTensorOutput) ProtoMessage()               {}
func (*MemoryLogTensorOutput) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *MemoryLogTensorOutput) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogTensorOutput) GetKernelName() string {
	if m != nil {
		return m.KernelName
	}
	return ""
}

func (m *MemoryLogTensorOutput) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *MemoryLogTensorOutput) GetTensor() *TensorDescription {
	if m != nil {
		return m.Tensor
	}
	return nil
}

type MemoryLogRawAllocation struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	// Name of the operation making the allocation.
	Operation string `protobuf:"bytes,2,opt,name=operation" json:"operation,omitempty"`
	// Number of bytes in the allocation.
	NumBytes int64 `protobuf:"varint,3,opt,name=num_bytes,json=numBytes" json:"num_bytes,omitempty"`
	// Address of the allocation.
	Ptr uint64 `protobuf:"varint,4,opt,name=ptr" json:"ptr,omitempty"`
	// Id of the tensor buffer being allocated, used to match to a
	// corresponding deallocation.
	AllocationId int64 `protobuf:"varint,5,opt,name=allocation_id,json=allocationId" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,6,opt,name=allocator_name,json=allocatorName" json:"allocator_name,omitempty"`
}

func (m *MemoryLogRawAllocation) Reset()                    { *m = MemoryLogRawAllocation{} }
func (m *MemoryLogRawAllocation) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogRawAllocation) ProtoMessage()               {}
func (*MemoryLogRawAllocation) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *MemoryLogRawAllocation) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *MemoryLogRawAllocation) GetNumBytes() int64 {
	if m != nil {
		return m.NumBytes
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

type MemoryLogRawDeallocation struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	// Name of the operation making the deallocation.
	Operation string `protobuf:"bytes,2,opt,name=operation" json:"operation,omitempty"`
	// Id of the tensor buffer being deallocated, used to match to a
	// corresponding allocation.
	AllocationId int64 `protobuf:"varint,3,opt,name=allocation_id,json=allocationId" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,4,opt,name=allocator_name,json=allocatorName" json:"allocator_name,omitempty"`
	// True if the deallocation is queued and will be performed later,
	// e.g. for GPU lazy freeing of buffers.
	Deferred bool `protobuf:"varint,5,opt,name=deferred" json:"deferred,omitempty"`
}

func (m *MemoryLogRawDeallocation) Reset()                    { *m = MemoryLogRawDeallocation{} }
func (m *MemoryLogRawDeallocation) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogRawDeallocation) ProtoMessage()               {}
func (*MemoryLogRawDeallocation) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *MemoryLogRawDeallocation) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogRawDeallocation) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *MemoryLogRawDeallocation) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *MemoryLogRawDeallocation) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *MemoryLogRawDeallocation) GetDeferred() bool {
	if m != nil {
		return m.Deferred
	}
	return false
}

func init() {
	proto.RegisterType((*MemoryLogStep)(nil), "tensorflow.MemoryLogStep")
	proto.RegisterType((*MemoryLogTensorAllocation)(nil), "tensorflow.MemoryLogTensorAllocation")
	proto.RegisterType((*MemoryLogTensorDeallocation)(nil), "tensorflow.MemoryLogTensorDeallocation")
	proto.RegisterType((*MemoryLogTensorOutput)(nil), "tensorflow.MemoryLogTensorOutput")
	proto.RegisterType((*MemoryLogRawAllocation)(nil), "tensorflow.MemoryLogRawAllocation")
	proto.RegisterType((*MemoryLogRawDeallocation)(nil), "tensorflow.MemoryLogRawDeallocation")
}

func init() { proto.RegisterFile("tensorflow/core/framework/log_memory.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xdf, 0xae, 0xd2, 0x40,
	0x10, 0xc6, 0xb3, 0x16, 0x2a, 0xcc, 0x11, 0x35, 0x1b, 0x3d, 0xd6, 0x73, 0x34, 0x92, 0x1a, 0x13,
	0xe2, 0x05, 0x18, 0x8c, 0xf7, 0x4a, 0xb8, 0x21, 0x41, 0x25, 0xab, 0xf7, 0x4d, 0xa1, 0x43, 0x6d,
	0x68, 0x77, 0x9b, 0xed, 0x36, 0xc8, 0x3b, 0xf8, 0x0c, 0xbe, 0x87, 0xaf, 0xe0, 0x13, 0x79, 0x69,
	0xda, 0xad, 0xdd, 0xca, 0x9f, 0x04, 0x3d, 0x77, 0xfd, 0xa6, 0xb3, 0x33, 0xbf, 0x6f, 0x26, 0x03,
	0x2f, 0x15, 0xf2, 0x4c, 0xc8, 0x75, 0x2c, 0xb6, 0xa3, 0x95, 0x90, 0x38, 0x5a, 0x4b, 0x3f, 0xc1,
	0xad, 0x90, 0x9b, 0x51, 0x2c, 0x42, 0x2f, 0xc1, 0x44, 0xc8, 0xdd, 0x30, 0x95, 0x42, 0x09, 0x0a,
	0x26, 0xf7, 0x6a, 0x7c, 0xfa, 0x9d, 0xfe, 0xe3, 0x05, 0x98, 0xad, 0x64, 0x94, 0xaa, 0x48, 0x70,
	0xfd, 0xde, 0x7d, 0x0b, 0xbd, 0xf7, 0x65, 0xbd, 0xb9, 0x08, 0x3f, 0x29, 0x4c, 0xe9, 0x23, 0xb8,
	0x9d, 0x29, 0x4c, 0xbd, 0x28, 0x70, 0x48, 0x9f, 0x0c, 0x2c, 0x66, 0x17, 0x72, 0x16, 0xd0, 0x4b,
	0xb0, 0xbf, 0xf8, 0x3c, 0x88, 0xd1, 0xb9, 0xd5, 0x27, 0x83, 0x2e, 0xab, 0x94, 0xfb, 0x8d, 0xc0,
	0xe3, 0xba, 0xc4, 0xe7, 0xb2, 0xcf, 0xbb, 0x38, 0x16, 0x2b, 0xbf, 0xe8, 0x72, 0xba, 0xdc, 0x33,
	0xb8, 0xd8, 0xa0, 0xe4, 0x18, 0x7b, 0xdc, 0x4f, 0xfe, 0xd4, 0x04, 0x1d, 0xfa, 0xe0, 0x27, 0x48,
	0xdf, 0x80, 0xad, 0xa9, 0x1d, 0xab, 0x4f, 0x06, 0x17, 0xe3, 0xa7, 0x43, 0x63, 0x6f, 0xa8, 0xfb,
	0x4c, 0x8d, 0x1d, 0x56, 0x25, 0xbb, 0x11, 0x5c, 0xef, 0xd1, 0x4c, 0xd1, 0x37, 0x3c, 0xcf, 0xa1,
	0x67, 0x94, 0xa1, 0xba, 0x63, 0x82, 0xb3, 0x80, 0xbe, 0x80, 0xbb, 0x95, 0x16, 0xb2, 0x89, 0xd7,
	0xab, 0xa3, 0x05, 0xa1, 0xfb, 0x9d, 0xc0, 0xc3, 0xbd, 0x5e, 0x1f, 0x73, 0x95, 0xe6, 0xea, 0x06,
	0xae, 0x1f, 0x40, 0x3b, 0xe2, 0x01, 0x7e, 0x2d, 0x4d, 0xb7, 0x99, 0x16, 0x8d, 0x59, 0xb4, 0xfe,
	0x65, 0x16, 0x3f, 0x09, 0x5c, 0xd6, 0x80, 0xcc, 0xdf, 0x9e, 0xb3, 0x97, 0x27, 0xd0, 0x15, 0x29,
	0xca, 0x32, 0xab, 0xe2, 0x33, 0x01, 0x7a, 0x0d, 0x5d, 0x9e, 0x27, 0xde, 0x72, 0xa7, 0x30, 0x2b,
	0x11, 0x2d, 0xd6, 0xe1, 0x79, 0x32, 0x29, 0x34, 0xbd, 0x0f, 0x56, 0xaa, 0x34, 0x62, 0x8b, 0x15,
	0x9f, 0x87, 0xd3, 0x6e, 0x9f, 0x35, 0x6d, 0xfb, 0xd8, 0xb4, 0x7f, 0x10, 0x70, 0x9a, 0x66, 0xfe,
	0x5a, 0xeb, 0x7f, 0xda, 0x39, 0xe0, 0xb3, 0xce, 0xe2, 0x6b, 0x1d, 0xe1, 0xa3, 0x57, 0xd0, 0x09,
	0x70, 0x8d, 0x52, 0xa2, 0xb6, 0xd9, 0x61, 0xb5, 0x9e, 0xbc, 0x02, 0x47, 0xc8, 0xb0, 0xb9, 0xb4,
	0xfa, 0x34, 0x27, 0xf7, 0xe6, 0x22, 0xd4, 0xbe, 0x16, 0xc5, 0x45, 0x66, 0x0b, 0xf2, 0x8b, 0x90,
	0xa5, 0x5d, 0x9e, 0xe7, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xdc, 0x69, 0x9c, 0x0c,
	0x04, 0x00, 0x00,
}
