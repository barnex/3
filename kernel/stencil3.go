package kernel

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var stencil3_code cu.Function

type stencil3_args struct {
	arg_dst   unsafe.Pointer
	arg_src   unsafe.Pointer
	arg_w0    float32
	arg_wt    float32
	arg_wb    float32
	arg_wu    float32
	arg_wd    float32
	arg_wl    float32
	arg_wr    float32
	arg_wrap0 int
	arg_wrap1 int
	arg_wrap2 int
	arg_N0    int
	arg_N1    int
	arg_N2    int
	argptr    [15]unsafe.Pointer
}

// Wrapper for stencil3 CUDA kernel. Synchronizes before return.
func K_stencil3(dst unsafe.Pointer, src unsafe.Pointer, w0 float32, wt float32, wb float32, wu float32, wd float32, wl float32, wr float32, wrap0 int, wrap1 int, wrap2 int, N0 int, N1 int, N2 int, gridDim, blockDim cu.Dim3) {
	if stencil3_code == 0 {
		stencil3_code = cu.ModuleLoadData(stencil3_ptx).GetFunction("stencil3")
	}

	var a stencil3_args

	a.arg_dst = dst
	a.argptr[0] = unsafe.Pointer(&a.arg_dst)
	a.arg_src = src
	a.argptr[1] = unsafe.Pointer(&a.arg_src)
	a.arg_w0 = w0
	a.argptr[2] = unsafe.Pointer(&a.arg_w0)
	a.arg_wt = wt
	a.argptr[3] = unsafe.Pointer(&a.arg_wt)
	a.arg_wb = wb
	a.argptr[4] = unsafe.Pointer(&a.arg_wb)
	a.arg_wu = wu
	a.argptr[5] = unsafe.Pointer(&a.arg_wu)
	a.arg_wd = wd
	a.argptr[6] = unsafe.Pointer(&a.arg_wd)
	a.arg_wl = wl
	a.argptr[7] = unsafe.Pointer(&a.arg_wl)
	a.arg_wr = wr
	a.argptr[8] = unsafe.Pointer(&a.arg_wr)
	a.arg_wrap0 = wrap0
	a.argptr[9] = unsafe.Pointer(&a.arg_wrap0)
	a.arg_wrap1 = wrap1
	a.argptr[10] = unsafe.Pointer(&a.arg_wrap1)
	a.arg_wrap2 = wrap2
	a.argptr[11] = unsafe.Pointer(&a.arg_wrap2)
	a.arg_N0 = N0
	a.argptr[12] = unsafe.Pointer(&a.arg_N0)
	a.arg_N1 = N1
	a.argptr[13] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[14] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	str := Stream()
	cu.LaunchKernel(stencil3_code, gridDim.X, gridDim.Y, gridDim.Z, blockDim.X, blockDim.Y, blockDim.Z, 0, str, args)
	SyncAndRecycle(str)
}

const stencil3_ptx = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry stencil3(
	.param .u64 stencil3_param_0,
	.param .u64 stencil3_param_1,
	.param .f32 stencil3_param_2,
	.param .f32 stencil3_param_3,
	.param .f32 stencil3_param_4,
	.param .f32 stencil3_param_5,
	.param .f32 stencil3_param_6,
	.param .f32 stencil3_param_7,
	.param .f32 stencil3_param_8,
	.param .u32 stencil3_param_9,
	.param .u32 stencil3_param_10,
	.param .u32 stencil3_param_11,
	.param .u32 stencil3_param_12,
	.param .u32 stencil3_param_13,
	.param .u32 stencil3_param_14
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<71>;
	.reg .f32 	%f<33>;
	.reg .s64 	%rd<21>;


	ld.param.u64 	%rd3, [stencil3_param_0];
	ld.param.u64 	%rd4, [stencil3_param_1];
	ld.param.f32 	%f9, [stencil3_param_2];
	ld.param.f32 	%f10, [stencil3_param_3];
	ld.param.f32 	%f11, [stencil3_param_4];
	ld.param.f32 	%f12, [stencil3_param_5];
	ld.param.f32 	%f13, [stencil3_param_6];
	ld.param.f32 	%f14, [stencil3_param_7];
	ld.param.f32 	%f15, [stencil3_param_8];
	ld.param.u32 	%r19, [stencil3_param_12];
	ld.param.u32 	%r20, [stencil3_param_13];
	ld.param.u32 	%r21, [stencil3_param_14];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 21 1
	mov.u32 	%r1, %ntid.x;
	mov.u32 	%r2, %ctaid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 2 22 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 24 1
	setp.lt.s32 	%p1, %r8, %r21;
	setp.lt.s32 	%p2, %r4, %r20;
	and.pred  	%p3, %p2, %p1;
	.loc 2 28 1
	setp.gt.s32 	%p4, %r19, 0;
	.loc 2 24 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_11;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 28 1
	mul.lo.s32 	%r68, %r21, %r4;
	mad.lo.s32 	%r69, %r21, %r4, %r8;
	mul.lo.s32 	%r11, %r21, %r20;
	mov.u32 	%r70, 0;

BB0_2:
	mul.lo.s32 	%r15, %r70, %r20;
	mov.f32 	%f32, 0f00000000;
	.loc 2 31 1
	setp.eq.f32 	%p6, %f9, 0f00000000;
	@%p6 bra 	BB0_4;

	.loc 2 31 1
	mul.wide.s32 	%rd5, %r69, 4;
	add.s64 	%rd6, %rd2, %rd5;
	ld.global.f32 	%f17, [%rd6];
	fma.rn.f32 	%f32, %f17, %f9, 0f00000000;

BB0_4:
	.loc 2 32 1
	setp.eq.f32 	%p7, %f10, 0f00000000;
	@%p7 bra 	BB0_6;

	.loc 2 32 1
	add.s32 	%r26, %r70, 1;
	mov.u32 	%r27, 0;
	.loc 3 238 5
	max.s32 	%r28, %r26, %r27;
	.loc 2 32 1
	add.s32 	%r29, %r19, -1;
	.loc 3 210 5
	min.s32 	%r30, %r28, %r29;
	.loc 2 32 1
	mad.lo.s32 	%r31, %r30, %r20, %r4;
	mad.lo.s32 	%r32, %r31, %r21, %r8;
	mul.wide.s32 	%rd7, %r32, 4;
	add.s64 	%rd8, %rd2, %rd7;
	ld.global.f32 	%f18, [%rd8];
	add.s32 	%r34, %r70, -1;
	.loc 3 238 5
	max.s32 	%r35, %r34, %r27;
	.loc 3 210 5
	min.s32 	%r36, %r35, %r29;
	.loc 2 32 1
	mad.lo.s32 	%r37, %r36, %r20, %r4;
	mad.lo.s32 	%r38, %r37, %r21, %r8;
	mul.wide.s32 	%rd9, %r38, 4;
	add.s64 	%rd10, %rd2, %rd9;
	ld.global.f32 	%f19, [%rd10];
	mul.f32 	%f20, %f19, %f11;
	fma.rn.f32 	%f21, %f18, %f10, %f20;
	add.f32 	%f32, %f32, %f21;

BB0_6:
	.loc 2 33 1
	setp.eq.f32 	%p8, %f12, 0f00000000;
	@%p8 bra 	BB0_8;

	.loc 2 33 1
	add.s32 	%r40, %r4, 1;
	mov.u32 	%r41, 0;
	.loc 3 238 5
	max.s32 	%r42, %r40, %r41;
	.loc 2 33 1
	add.s32 	%r43, %r20, -1;
	.loc 3 210 5
	min.s32 	%r44, %r42, %r43;
	.loc 2 33 1
	add.s32 	%r45, %r44, %r15;
	mad.lo.s32 	%r46, %r45, %r21, %r8;
	mul.wide.s32 	%rd11, %r46, 4;
	add.s64 	%rd12, %rd2, %rd11;
	ld.global.f32 	%f22, [%rd12];
	add.s32 	%r48, %r4, -1;
	.loc 3 238 5
	max.s32 	%r49, %r48, %r41;
	.loc 3 210 5
	min.s32 	%r50, %r49, %r43;
	.loc 2 33 1
	add.s32 	%r51, %r50, %r15;
	mad.lo.s32 	%r52, %r51, %r21, %r8;
	mul.wide.s32 	%rd13, %r52, 4;
	add.s64 	%rd14, %rd2, %rd13;
	ld.global.f32 	%f23, [%rd14];
	mul.f32 	%f24, %f23, %f13;
	fma.rn.f32 	%f25, %f22, %f12, %f24;
	add.f32 	%f32, %f32, %f25;

BB0_8:
	.loc 2 34 1
	setp.eq.f32 	%p9, %f15, 0f00000000;
	@%p9 bra 	BB0_10;

	.loc 2 34 1
	add.s32 	%r54, %r8, 1;
	mov.u32 	%r55, 0;
	.loc 3 238 5
	max.s32 	%r56, %r54, %r55;
	.loc 2 34 1
	add.s32 	%r57, %r21, -1;
	.loc 3 210 5
	min.s32 	%r58, %r56, %r57;
	.loc 2 34 1
	add.s32 	%r59, %r58, %r68;
	mul.wide.s32 	%rd15, %r59, 4;
	add.s64 	%rd16, %rd2, %rd15;
	ld.global.f32 	%f26, [%rd16];
	add.s32 	%r61, %r8, -1;
	.loc 3 238 5
	max.s32 	%r62, %r61, %r55;
	.loc 3 210 5
	min.s32 	%r63, %r62, %r57;
	.loc 2 34 1
	add.s32 	%r64, %r63, %r68;
	mul.wide.s32 	%rd17, %r64, 4;
	add.s64 	%rd18, %rd2, %rd17;
	ld.global.f32 	%f27, [%rd18];
	mul.f32 	%f28, %f27, %f14;
	fma.rn.f32 	%f29, %f26, %f15, %f28;
	add.f32 	%f32, %f32, %f29;

BB0_10:
	.loc 2 36 1
	mul.wide.s32 	%rd19, %r69, 4;
	add.s64 	%rd20, %rd1, %rd19;
	ld.global.f32 	%f30, [%rd20];
	add.f32 	%f31, %f30, %f32;
	st.global.f32 	[%rd20], %f31;
	.loc 2 28 1
	add.s32 	%r69, %r69, %r11;
	add.s32 	%r68, %r68, %r11;
	.loc 2 28 18
	add.s32 	%r70, %r70, 1;
	.loc 2 28 1
	setp.lt.s32 	%p10, %r70, %r19;
	@%p10 bra 	BB0_2;

BB0_11:
	.loc 2 38 2
	ret;
}


`