package cuda

import (
	"github.com/barnex/cuda5/cu"
	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
	"unsafe"
)

// 3D byte slice, used for region lookup.
type Bytes struct {
	Ptr unsafe.Pointer
	Len int
}

// Construct new 3D byte slice for given mesh.
func NewBytes(m *data.Mesh) *Bytes {
	Len := int64(m.NCell())
	ptr := cu.MemAlloc(Len)
	cu.MemsetD8(cu.DevicePtr(ptr), 0, Len)
	return &Bytes{unsafe.Pointer(uintptr(ptr)), int(Len)}
}

// Upload src (host) to dst (gpu)
func (dst *Bytes) Upload(src []byte) {
	util.Argument(int(dst.Len) == len(src))
	cu.MemcpyHtoD(cu.DevicePtr(uintptr(dst.Ptr)), unsafe.Pointer(&src[0]), int64(dst.Len))
}

// Frees the GPU memory and disables the slice.
func (b *Bytes) Free() {
	cu.MemFree(cu.DevicePtr(uintptr(b.Ptr)))
	b.Ptr = nil
	b.Len = 0
}
