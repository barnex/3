package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var uniaxialanisotropy_code cu.Function

type uniaxialanisotropy_args struct {
	arg_Bx unsafe.Pointer
	arg_By unsafe.Pointer
	arg_Bz unsafe.Pointer
	arg_Mx unsafe.Pointer
	arg_My unsafe.Pointer
	arg_Mz unsafe.Pointer
	arg_Ux float32
	arg_Uy float32
	arg_Uz float32
	arg_N  int
	argptr [10]unsafe.Pointer
}

// Wrapper for uniaxialanisotropy CUDA kernel, asynchronous.
func k_uniaxialanisotropy_async(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, Mx unsafe.Pointer, My unsafe.Pointer, Mz unsafe.Pointer, Ux float32, Uy float32, Uz float32, N int, cfg *config, str cu.Stream) {
	if uniaxialanisotropy_code == 0 {
		uniaxialanisotropy_code = fatbinLoad(uniaxialanisotropy_map, "uniaxialanisotropy")
	}

	var a uniaxialanisotropy_args

	a.arg_Bx = Bx
	a.argptr[0] = unsafe.Pointer(&a.arg_Bx)
	a.arg_By = By
	a.argptr[1] = unsafe.Pointer(&a.arg_By)
	a.arg_Bz = Bz
	a.argptr[2] = unsafe.Pointer(&a.arg_Bz)
	a.arg_Mx = Mx
	a.argptr[3] = unsafe.Pointer(&a.arg_Mx)
	a.arg_My = My
	a.argptr[4] = unsafe.Pointer(&a.arg_My)
	a.arg_Mz = Mz
	a.argptr[5] = unsafe.Pointer(&a.arg_Mz)
	a.arg_Ux = Ux
	a.argptr[6] = unsafe.Pointer(&a.arg_Ux)
	a.arg_Uy = Uy
	a.argptr[7] = unsafe.Pointer(&a.arg_Uy)
	a.arg_Uz = Uz
	a.argptr[8] = unsafe.Pointer(&a.arg_Uz)
	a.arg_N = N
	a.argptr[9] = unsafe.Pointer(&a.arg_N)

	args := a.argptr[:]
	cu.LaunchKernel(uniaxialanisotropy_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for uniaxialanisotropy CUDA kernel, synchronized.
func k_uniaxialanisotropy(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, Mx unsafe.Pointer, My unsafe.Pointer, Mz unsafe.Pointer, Ux float32, Uy float32, Uz float32, N int, cfg *config) {
	str := stream()
	k_uniaxialanisotropy_async(Bx, By, Bz, Mx, My, Mz, Ux, Uy, Uz, N, cfg, str)
	syncAndRecycle(str)
}

var uniaxialanisotropy_map = map[int]string{0: "",
	20: uniaxialanisotropy_ptx_20,
	30: uniaxialanisotropy_ptx_30,
	35: uniaxialanisotropy_ptx_35}

const (
	uniaxialanisotropy_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry uniaxialanisotropy(
	.param .u64 uniaxialanisotropy_param_0,
	.param .u64 uniaxialanisotropy_param_1,
	.param .u64 uniaxialanisotropy_param_2,
	.param .u64 uniaxialanisotropy_param_3,
	.param .u64 uniaxialanisotropy_param_4,
	.param .u64 uniaxialanisotropy_param_5,
	.param .f32 uniaxialanisotropy_param_6,
	.param .f32 uniaxialanisotropy_param_7,
	.param .f32 uniaxialanisotropy_param_8,
	.param .u32 uniaxialanisotropy_param_9
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<39>;
	.reg .s64 	%rd<22>;


	ld.param.u64 	%rd8, [uniaxialanisotropy_param_0];
	ld.param.u64 	%rd9, [uniaxialanisotropy_param_1];
	ld.param.u64 	%rd10, [uniaxialanisotropy_param_2];
	ld.param.u64 	%rd11, [uniaxialanisotropy_param_3];
	ld.param.u64 	%rd12, [uniaxialanisotropy_param_4];
	ld.param.u64 	%rd13, [uniaxialanisotropy_param_5];
	ld.param.f32 	%f13, [uniaxialanisotropy_param_6];
	ld.param.f32 	%f14, [uniaxialanisotropy_param_7];
	ld.param.f32 	%f15, [uniaxialanisotropy_param_8];
	ld.param.u32 	%r2, [uniaxialanisotropy_param_9];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd8;
	cvta.to.global.u64 	%rd4, %rd13;
	cvta.to.global.u64 	%rd5, %rd12;
	cvta.to.global.u64 	%rd6, %rd11;
	.loc 2 9 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 10 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_6;

	.loc 2 12 1
	cvt.s64.s32 	%rd7, %r1;
	mul.wide.s32 	%rd14, %r1, 4;
	add.s64 	%rd15, %rd6, %rd14;
	add.s64 	%rd16, %rd5, %rd14;
	add.s64 	%rd17, %rd4, %rd14;
	ld.global.f32 	%f1, [%rd15];
	ld.global.f32 	%f2, [%rd16];
	.loc 2 13 1
	mul.f32 	%f17, %f2, %f2;
	fma.rn.f32 	%f18, %f1, %f1, %f17;
	.loc 2 12 1
	ld.global.f32 	%f3, [%rd17];
	.loc 2 13 1
	fma.rn.f32 	%f19, %f3, %f3, %f18;
	.loc 3 991 5
	sqrt.rn.f32 	%f4, %f19;
	mov.f32 	%f16, 0f00000000;
	.loc 2 14 1
	setp.eq.f32 	%p2, %f4, 0f00000000;
	mov.f32 	%f38, %f16;
	@%p2 bra 	BB0_3;

	rcp.rn.f32 	%f5, %f4;
	mov.f32 	%f38, %f5;

BB0_3:
	.loc 2 14 1
	mov.f32 	%f6, %f38;
	mul.f32 	%f7, %f6, %f1;
	mul.f32 	%f8, %f6, %f2;
	mul.f32 	%f9, %f6, %f3;
	.loc 2 16 1
	mul.f32 	%f21, %f14, %f14;
	fma.rn.f32 	%f22, %f13, %f13, %f21;
	fma.rn.f32 	%f23, %f15, %f15, %f22;
	.loc 3 991 5
	sqrt.rn.f32 	%f10, %f23;
	.loc 2 16 1
	setp.eq.f32 	%p3, %f10, 0f00000000;
	mov.f32 	%f37, %f16;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f37, %f10;

BB0_5:
	.loc 2 18 1
	add.f32 	%f24, %f10, %f10;
	.loc 4 2399 3
	div.rn.f32 	%f25, %f24, %f4;
	.loc 2 16 1
	mul.f32 	%f26, %f37, %f13;
	mul.f32 	%f27, %f37, %f14;
	.loc 2 18 1
	mul.f32 	%f28, %f8, %f27;
	fma.rn.f32 	%f29, %f7, %f26, %f28;
	.loc 2 16 1
	mul.f32 	%f30, %f37, %f15;
	.loc 2 18 1
	fma.rn.f32 	%f31, %f9, %f30, %f29;
	mul.f32 	%f32, %f25, %f31;
	mul.f32 	%f33, %f32, %f26;
	mul.f32 	%f34, %f32, %f27;
	mul.f32 	%f35, %f32, %f30;
	.loc 2 20 1
	shl.b64 	%rd18, %rd7, 2;
	add.s64 	%rd19, %rd3, %rd18;
	st.global.f32 	[%rd19], %f33;
	.loc 2 21 1
	add.s64 	%rd20, %rd2, %rd18;
	st.global.f32 	[%rd20], %f34;
	.loc 2 22 1
	add.s64 	%rd21, %rd1, %rd18;
	st.global.f32 	[%rd21], %f35;

BB0_6:
	.loc 2 24 2
	ret;
}


`
	uniaxialanisotropy_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry uniaxialanisotropy(
	.param .u64 uniaxialanisotropy_param_0,
	.param .u64 uniaxialanisotropy_param_1,
	.param .u64 uniaxialanisotropy_param_2,
	.param .u64 uniaxialanisotropy_param_3,
	.param .u64 uniaxialanisotropy_param_4,
	.param .u64 uniaxialanisotropy_param_5,
	.param .f32 uniaxialanisotropy_param_6,
	.param .f32 uniaxialanisotropy_param_7,
	.param .f32 uniaxialanisotropy_param_8,
	.param .u32 uniaxialanisotropy_param_9
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<39>;
	.reg .s64 	%rd<22>;


	ld.param.u64 	%rd8, [uniaxialanisotropy_param_0];
	ld.param.u64 	%rd9, [uniaxialanisotropy_param_1];
	ld.param.u64 	%rd10, [uniaxialanisotropy_param_2];
	ld.param.u64 	%rd11, [uniaxialanisotropy_param_3];
	ld.param.u64 	%rd12, [uniaxialanisotropy_param_4];
	ld.param.u64 	%rd13, [uniaxialanisotropy_param_5];
	ld.param.f32 	%f13, [uniaxialanisotropy_param_6];
	ld.param.f32 	%f14, [uniaxialanisotropy_param_7];
	ld.param.f32 	%f15, [uniaxialanisotropy_param_8];
	ld.param.u32 	%r2, [uniaxialanisotropy_param_9];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd8;
	cvta.to.global.u64 	%rd4, %rd13;
	cvta.to.global.u64 	%rd5, %rd12;
	cvta.to.global.u64 	%rd6, %rd11;
	.loc 2 9 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 10 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_6;

	.loc 2 12 1
	cvt.s64.s32 	%rd7, %r1;
	mul.wide.s32 	%rd14, %r1, 4;
	add.s64 	%rd15, %rd6, %rd14;
	add.s64 	%rd16, %rd5, %rd14;
	add.s64 	%rd17, %rd4, %rd14;
	ld.global.f32 	%f1, [%rd15];
	ld.global.f32 	%f2, [%rd16];
	.loc 2 13 1
	mul.f32 	%f17, %f2, %f2;
	fma.rn.f32 	%f18, %f1, %f1, %f17;
	.loc 2 12 1
	ld.global.f32 	%f3, [%rd17];
	.loc 2 13 1
	fma.rn.f32 	%f19, %f3, %f3, %f18;
	.loc 3 991 5
	sqrt.rn.f32 	%f4, %f19;
	mov.f32 	%f16, 0f00000000;
	.loc 2 14 1
	setp.eq.f32 	%p2, %f4, 0f00000000;
	mov.f32 	%f38, %f16;
	@%p2 bra 	BB0_3;

	rcp.rn.f32 	%f5, %f4;
	mov.f32 	%f38, %f5;

BB0_3:
	.loc 2 14 1
	mov.f32 	%f6, %f38;
	mul.f32 	%f7, %f6, %f1;
	mul.f32 	%f8, %f6, %f2;
	mul.f32 	%f9, %f6, %f3;
	.loc 2 16 1
	mul.f32 	%f21, %f14, %f14;
	fma.rn.f32 	%f22, %f13, %f13, %f21;
	fma.rn.f32 	%f23, %f15, %f15, %f22;
	.loc 3 991 5
	sqrt.rn.f32 	%f10, %f23;
	.loc 2 16 1
	setp.eq.f32 	%p3, %f10, 0f00000000;
	mov.f32 	%f37, %f16;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f37, %f10;

BB0_5:
	.loc 2 18 1
	add.f32 	%f24, %f10, %f10;
	.loc 4 2399 3
	div.rn.f32 	%f25, %f24, %f4;
	.loc 2 16 1
	mul.f32 	%f26, %f37, %f13;
	mul.f32 	%f27, %f37, %f14;
	.loc 2 18 1
	mul.f32 	%f28, %f8, %f27;
	fma.rn.f32 	%f29, %f7, %f26, %f28;
	.loc 2 16 1
	mul.f32 	%f30, %f37, %f15;
	.loc 2 18 1
	fma.rn.f32 	%f31, %f9, %f30, %f29;
	mul.f32 	%f32, %f25, %f31;
	mul.f32 	%f33, %f32, %f26;
	mul.f32 	%f34, %f32, %f27;
	mul.f32 	%f35, %f32, %f30;
	.loc 2 20 1
	shl.b64 	%rd18, %rd7, 2;
	add.s64 	%rd19, %rd3, %rd18;
	st.global.f32 	[%rd19], %f33;
	.loc 2 21 1
	add.s64 	%rd20, %rd2, %rd18;
	st.global.f32 	[%rd20], %f34;
	.loc 2 22 1
	add.s64 	%rd21, %rd1, %rd18;
	st.global.f32 	[%rd21], %f35;

BB0_6:
	.loc 2 24 2
	ret;
}


`
	uniaxialanisotropy_ptx_35 = `
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

.visible .entry uniaxialanisotropy(
	.param .u64 uniaxialanisotropy_param_0,
	.param .u64 uniaxialanisotropy_param_1,
	.param .u64 uniaxialanisotropy_param_2,
	.param .u64 uniaxialanisotropy_param_3,
	.param .u64 uniaxialanisotropy_param_4,
	.param .u64 uniaxialanisotropy_param_5,
	.param .f32 uniaxialanisotropy_param_6,
	.param .f32 uniaxialanisotropy_param_7,
	.param .f32 uniaxialanisotropy_param_8,
	.param .u32 uniaxialanisotropy_param_9
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<12>;
	.reg .f32 	%f<39>;
	.reg .s64 	%rd<22>;


	ld.param.u64 	%rd8, [uniaxialanisotropy_param_0];
	ld.param.u64 	%rd9, [uniaxialanisotropy_param_1];
	ld.param.u64 	%rd10, [uniaxialanisotropy_param_2];
	ld.param.u64 	%rd11, [uniaxialanisotropy_param_3];
	ld.param.u64 	%rd12, [uniaxialanisotropy_param_4];
	ld.param.u64 	%rd13, [uniaxialanisotropy_param_5];
	ld.param.f32 	%f13, [uniaxialanisotropy_param_6];
	ld.param.f32 	%f14, [uniaxialanisotropy_param_7];
	ld.param.f32 	%f15, [uniaxialanisotropy_param_8];
	ld.param.u32 	%r2, [uniaxialanisotropy_param_9];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd8;
	cvta.to.global.u64 	%rd4, %rd13;
	cvta.to.global.u64 	%rd5, %rd12;
	cvta.to.global.u64 	%rd6, %rd11;
	.loc 3 9 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 3 10 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB2_6;

	.loc 3 12 1
	cvt.s64.s32 	%rd7, %r1;
	mul.wide.s32 	%rd14, %r1, 4;
	add.s64 	%rd15, %rd6, %rd14;
	ld.global.nc.f32 	%f1, [%rd15];
	add.s64 	%rd16, %rd5, %rd14;
	ld.global.nc.f32 	%f2, [%rd16];
	add.s64 	%rd17, %rd4, %rd14;
	ld.global.nc.f32 	%f3, [%rd17];
	.loc 3 13 1
	mul.f32 	%f17, %f2, %f2;
	fma.rn.f32 	%f18, %f1, %f1, %f17;
	fma.rn.f32 	%f19, %f3, %f3, %f18;
	.loc 4 991 5
	sqrt.rn.f32 	%f4, %f19;
	mov.f32 	%f16, 0f00000000;
	.loc 3 14 1
	setp.eq.f32 	%p2, %f4, 0f00000000;
	mov.f32 	%f38, %f16;
	@%p2 bra 	BB2_3;

	rcp.rn.f32 	%f5, %f4;
	mov.f32 	%f38, %f5;

BB2_3:
	.loc 3 14 1
	mov.f32 	%f6, %f38;
	mul.f32 	%f7, %f6, %f1;
	mul.f32 	%f8, %f6, %f2;
	mul.f32 	%f9, %f6, %f3;
	.loc 3 16 1
	mul.f32 	%f21, %f14, %f14;
	fma.rn.f32 	%f22, %f13, %f13, %f21;
	fma.rn.f32 	%f23, %f15, %f15, %f22;
	.loc 4 991 5
	sqrt.rn.f32 	%f10, %f23;
	.loc 3 16 1
	setp.eq.f32 	%p3, %f10, 0f00000000;
	mov.f32 	%f37, %f16;
	@%p3 bra 	BB2_5;

	rcp.rn.f32 	%f37, %f10;

BB2_5:
	.loc 3 18 1
	add.f32 	%f24, %f10, %f10;
	.loc 5 2399 3
	div.rn.f32 	%f25, %f24, %f4;
	.loc 3 16 1
	mul.f32 	%f26, %f37, %f13;
	mul.f32 	%f27, %f37, %f14;
	.loc 3 18 1
	mul.f32 	%f28, %f8, %f27;
	fma.rn.f32 	%f29, %f7, %f26, %f28;
	.loc 3 16 1
	mul.f32 	%f30, %f37, %f15;
	.loc 3 18 1
	fma.rn.f32 	%f31, %f9, %f30, %f29;
	mul.f32 	%f32, %f25, %f31;
	mul.f32 	%f33, %f32, %f26;
	mul.f32 	%f34, %f32, %f27;
	mul.f32 	%f35, %f32, %f30;
	.loc 3 20 1
	shl.b64 	%rd18, %rd7, 2;
	add.s64 	%rd19, %rd3, %rd18;
	st.global.f32 	[%rd19], %f33;
	.loc 3 21 1
	add.s64 	%rd20, %rd2, %rd18;
	st.global.f32 	[%rd20], %f34;
	.loc 3 22 1
	add.s64 	%rd21, %rd1, %rd18;
	st.global.f32 	[%rd21], %f35;

BB2_6:
	.loc 3 24 2
	ret;
}


`
)