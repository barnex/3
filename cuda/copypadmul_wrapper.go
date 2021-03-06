package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var copypadmul_code cu.Function

type copypadmul_args struct {
	arg_dst     unsafe.Pointer
	arg_D0      int
	arg_D1      int
	arg_D2      int
	arg_src     unsafe.Pointer
	arg_vol     unsafe.Pointer
	arg_S0      int
	arg_S1      int
	arg_S2      int
	arg_BsatLUT unsafe.Pointer
	arg_regions unsafe.Pointer
	argptr      [11]unsafe.Pointer
}

// Wrapper for copypadmul CUDA kernel, asynchronous.
func k_copypadmul_async(dst unsafe.Pointer, D0 int, D1 int, D2 int, src unsafe.Pointer, vol unsafe.Pointer, S0 int, S1 int, S2 int, BsatLUT unsafe.Pointer, regions unsafe.Pointer, cfg *config, str cu.Stream) {
	if copypadmul_code == 0 {
		copypadmul_code = fatbinLoad(copypadmul_map, "copypadmul")
	}

	var _a_ copypadmul_args

	_a_.arg_dst = dst
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_dst)
	_a_.arg_D0 = D0
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_D0)
	_a_.arg_D1 = D1
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_D1)
	_a_.arg_D2 = D2
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_D2)
	_a_.arg_src = src
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_src)
	_a_.arg_vol = vol
	_a_.argptr[5] = unsafe.Pointer(&_a_.arg_vol)
	_a_.arg_S0 = S0
	_a_.argptr[6] = unsafe.Pointer(&_a_.arg_S0)
	_a_.arg_S1 = S1
	_a_.argptr[7] = unsafe.Pointer(&_a_.arg_S1)
	_a_.arg_S2 = S2
	_a_.argptr[8] = unsafe.Pointer(&_a_.arg_S2)
	_a_.arg_BsatLUT = BsatLUT
	_a_.argptr[9] = unsafe.Pointer(&_a_.arg_BsatLUT)
	_a_.arg_regions = regions
	_a_.argptr[10] = unsafe.Pointer(&_a_.arg_regions)

	args := _a_.argptr[:]
	cu.LaunchKernel(copypadmul_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for copypadmul CUDA kernel, synchronized.
func k_copypadmul(dst unsafe.Pointer, D0 int, D1 int, D2 int, src unsafe.Pointer, vol unsafe.Pointer, S0 int, S1 int, S2 int, BsatLUT unsafe.Pointer, regions unsafe.Pointer, cfg *config) {
	str := stream()
	k_copypadmul_async(dst, D0, D1, D2, src, vol, S0, S1, S2, BsatLUT, regions, cfg, str)
	syncAndRecycle(str)
}

var copypadmul_map = map[int]string{0: "",
	20: copypadmul_ptx_20,
	30: copypadmul_ptx_30,
	35: copypadmul_ptx_35}

const (
	copypadmul_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry copypadmul(
	.param .u64 copypadmul_param_0,
	.param .u32 copypadmul_param_1,
	.param .u32 copypadmul_param_2,
	.param .u32 copypadmul_param_3,
	.param .u64 copypadmul_param_4,
	.param .u64 copypadmul_param_5,
	.param .u32 copypadmul_param_6,
	.param .u32 copypadmul_param_7,
	.param .u32 copypadmul_param_8,
	.param .u64 copypadmul_param_9,
	.param .u64 copypadmul_param_10
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<27>;
	.reg .f32 	%f<9>;
	.reg .s64 	%rd<22>;


	ld.param.u64 	%rd8, [copypadmul_param_0];
	ld.param.u32 	%r4, [copypadmul_param_2];
	ld.param.u32 	%r5, [copypadmul_param_3];
	ld.param.u64 	%rd9, [copypadmul_param_4];
	ld.param.u64 	%rd7, [copypadmul_param_5];
	ld.param.u32 	%r8, [copypadmul_param_6];
	ld.param.u32 	%r6, [copypadmul_param_7];
	ld.param.u32 	%r7, [copypadmul_param_8];
	ld.param.u64 	%rd10, [copypadmul_param_9];
	ld.param.u64 	%rd11, [copypadmul_param_10];
	cvta.to.global.u64 	%rd1, %rd8;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd7;
	cvta.to.global.u64 	%rd4, %rd10;
	cvta.to.global.u64 	%rd5, %rd11;
	.loc 2 9 1
	mov.u32 	%r9, %ntid.z;
	mov.u32 	%r10, %ctaid.z;
	mov.u32 	%r11, %tid.z;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	.loc 2 10 1
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	.loc 2 11 1
	mov.u32 	%r15, %ntid.x;
	mov.u32 	%r16, %ctaid.x;
	mov.u32 	%r17, %tid.x;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	.loc 2 13 1
	setp.lt.s32 	%p1, %r1, %r8;
	setp.lt.s32 	%p2, %r2, %r6;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32 	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_4;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 14 1
	mad.lo.s32 	%r18, %r1, %r6, %r2;
	mad.lo.s32 	%r19, %r18, %r7, %r3;
	.loc 2 15 1
	cvt.s64.s32 	%rd6, %r19;
	add.s64 	%rd12, %rd5, %rd6;
	ld.global.s8 	%rd13, [%rd12];
	shl.b64 	%rd14, %rd13, 2;
	add.s64 	%rd15, %rd4, %rd14;
	ld.global.f32 	%f1, [%rd15];
	.loc 2 16 1
	setp.eq.s64 	%p6, %rd7, 0;
	mov.f32 	%f8, 0f3F800000;
	.loc 2 16 1
	@%p6 bra 	BB0_3;

	shl.b64 	%rd16, %rd6, 2;
	add.s64 	%rd17, %rd3, %rd16;
	ld.global.f32 	%f8, [%rd17];

BB0_3:
	.loc 2 17 1
	shl.b64 	%rd18, %rd6, 2;
	add.s64 	%rd19, %rd2, %rd18;
	ld.global.f32 	%f5, [%rd19];
	mul.f32 	%f6, %f1, %f8;
	mul.f32 	%f7, %f6, %f5;
	mad.lo.s32 	%r24, %r1, %r4, %r2;
	mad.lo.s32 	%r25, %r24, %r5, %r3;
	mul.wide.s32 	%rd20, %r25, 4;
	add.s64 	%rd21, %rd1, %rd20;
	st.global.f32 	[%rd21], %f7;

BB0_4:
	.loc 2 19 2
	ret;
}


`
	copypadmul_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry copypadmul(
	.param .u64 copypadmul_param_0,
	.param .u32 copypadmul_param_1,
	.param .u32 copypadmul_param_2,
	.param .u32 copypadmul_param_3,
	.param .u64 copypadmul_param_4,
	.param .u64 copypadmul_param_5,
	.param .u32 copypadmul_param_6,
	.param .u32 copypadmul_param_7,
	.param .u32 copypadmul_param_8,
	.param .u64 copypadmul_param_9,
	.param .u64 copypadmul_param_10
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<27>;
	.reg .f32 	%f<9>;
	.reg .s64 	%rd<22>;


	ld.param.u64 	%rd8, [copypadmul_param_0];
	ld.param.u32 	%r4, [copypadmul_param_2];
	ld.param.u32 	%r5, [copypadmul_param_3];
	ld.param.u64 	%rd9, [copypadmul_param_4];
	ld.param.u64 	%rd7, [copypadmul_param_5];
	ld.param.u32 	%r8, [copypadmul_param_6];
	ld.param.u32 	%r6, [copypadmul_param_7];
	ld.param.u32 	%r7, [copypadmul_param_8];
	ld.param.u64 	%rd10, [copypadmul_param_9];
	ld.param.u64 	%rd11, [copypadmul_param_10];
	cvta.to.global.u64 	%rd1, %rd8;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd7;
	cvta.to.global.u64 	%rd4, %rd10;
	cvta.to.global.u64 	%rd5, %rd11;
	.loc 2 9 1
	mov.u32 	%r9, %ntid.z;
	mov.u32 	%r10, %ctaid.z;
	mov.u32 	%r11, %tid.z;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	.loc 2 10 1
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	.loc 2 11 1
	mov.u32 	%r15, %ntid.x;
	mov.u32 	%r16, %ctaid.x;
	mov.u32 	%r17, %tid.x;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	.loc 2 13 1
	setp.lt.s32 	%p1, %r1, %r8;
	setp.lt.s32 	%p2, %r2, %r6;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32 	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_4;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 14 1
	mad.lo.s32 	%r18, %r1, %r6, %r2;
	mad.lo.s32 	%r19, %r18, %r7, %r3;
	.loc 2 15 1
	cvt.s64.s32 	%rd6, %r19;
	add.s64 	%rd12, %rd5, %rd6;
	ld.global.s8 	%rd13, [%rd12];
	shl.b64 	%rd14, %rd13, 2;
	add.s64 	%rd15, %rd4, %rd14;
	ld.global.f32 	%f1, [%rd15];
	.loc 2 16 1
	setp.eq.s64 	%p6, %rd7, 0;
	mov.f32 	%f8, 0f3F800000;
	.loc 2 16 1
	@%p6 bra 	BB0_3;

	shl.b64 	%rd16, %rd6, 2;
	add.s64 	%rd17, %rd3, %rd16;
	ld.global.f32 	%f8, [%rd17];

BB0_3:
	.loc 2 17 1
	shl.b64 	%rd18, %rd6, 2;
	add.s64 	%rd19, %rd2, %rd18;
	ld.global.f32 	%f5, [%rd19];
	mul.f32 	%f6, %f1, %f8;
	mul.f32 	%f7, %f6, %f5;
	mad.lo.s32 	%r24, %r1, %r4, %r2;
	mad.lo.s32 	%r25, %r24, %r5, %r3;
	mul.wide.s32 	%rd20, %r25, 4;
	add.s64 	%rd21, %rd1, %rd20;
	st.global.f32 	[%rd21], %f7;

BB0_4:
	.loc 2 19 2
	ret;
}


`
	copypadmul_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry copypadmul(
	.param .u64 copypadmul_param_0,
	.param .u32 copypadmul_param_1,
	.param .u32 copypadmul_param_2,
	.param .u32 copypadmul_param_3,
	.param .u64 copypadmul_param_4,
	.param .u64 copypadmul_param_5,
	.param .u32 copypadmul_param_6,
	.param .u32 copypadmul_param_7,
	.param .u32 copypadmul_param_8,
	.param .u64 copypadmul_param_9,
	.param .u64 copypadmul_param_10
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<24>;
	.reg .f32 	%f<9>;
	.reg .s64 	%rd<22>;


	ld.param.u64 	%rd8, [copypadmul_param_0];
	ld.param.u32 	%r4, [copypadmul_param_2];
	ld.param.u32 	%r5, [copypadmul_param_3];
	ld.param.u64 	%rd9, [copypadmul_param_4];
	ld.param.u64 	%rd7, [copypadmul_param_5];
	ld.param.u32 	%r8, [copypadmul_param_6];
	ld.param.u32 	%r6, [copypadmul_param_7];
	ld.param.u32 	%r7, [copypadmul_param_8];
	ld.param.u64 	%rd10, [copypadmul_param_9];
	ld.param.u64 	%rd11, [copypadmul_param_10];
	cvta.to.global.u64 	%rd1, %rd8;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd7;
	cvta.to.global.u64 	%rd4, %rd10;
	cvta.to.global.u64 	%rd5, %rd11;
	.loc 3 9 1
	mov.u32 	%r9, %ntid.z;
	mov.u32 	%r10, %ctaid.z;
	mov.u32 	%r11, %tid.z;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	.loc 3 10 1
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	.loc 3 11 1
	mov.u32 	%r15, %ntid.x;
	mov.u32 	%r16, %ctaid.x;
	mov.u32 	%r17, %tid.x;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	.loc 3 13 1
	setp.lt.s32 	%p1, %r1, %r8;
	setp.lt.s32 	%p2, %r2, %r6;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32 	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_4;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 14 1
	mad.lo.s32 	%r18, %r1, %r6, %r2;
	mad.lo.s32 	%r19, %r18, %r7, %r3;
	.loc 3 15 1
	cvt.s64.s32 	%rd6, %r19;
	add.s64 	%rd12, %rd5, %rd6;
	ld.global.s8 	%rd13, [%rd12];
	shl.b64 	%rd14, %rd13, 2;
	add.s64 	%rd15, %rd4, %rd14;
	ld.global.nc.f32 	%f1, [%rd15];
	.loc 3 16 1
	setp.eq.s64 	%p6, %rd7, 0;
	mov.f32 	%f8, 0f3F800000;
	.loc 3 16 1
	@%p6 bra 	BB2_3;

	shl.b64 	%rd16, %rd6, 2;
	add.s64 	%rd17, %rd3, %rd16;
	ld.global.nc.f32 	%f8, [%rd17];

BB2_3:
	.loc 3 17 1
	shl.b64 	%rd18, %rd6, 2;
	add.s64 	%rd19, %rd2, %rd18;
	ld.global.nc.f32 	%f5, [%rd19];
	mul.f32 	%f6, %f1, %f8;
	mul.f32 	%f7, %f6, %f5;
	mad.lo.s32 	%r21, %r1, %r4, %r2;
	mad.lo.s32 	%r22, %r21, %r5, %r3;
	mul.wide.s32 	%rd20, %r22, 4;
	add.s64 	%rd21, %rd1, %rd20;
	st.global.f32 	[%rd21], %f7;

BB2_4:
	.loc 3 19 2
	ret;
}


`
)
