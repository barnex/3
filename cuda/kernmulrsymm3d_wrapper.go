package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var kernmulRSymm3D_code cu.Function

type kernmulRSymm3D_args struct {
	arg_fftMx  unsafe.Pointer
	arg_fftMy  unsafe.Pointer
	arg_fftMz  unsafe.Pointer
	arg_fftKxx unsafe.Pointer
	arg_fftKyy unsafe.Pointer
	arg_fftKzz unsafe.Pointer
	arg_fftKyz unsafe.Pointer
	arg_fftKxz unsafe.Pointer
	arg_fftKxy unsafe.Pointer
	arg_N0     int
	arg_N1     int
	arg_N2     int
	argptr     [12]unsafe.Pointer
}

// Wrapper for kernmulRSymm3D CUDA kernel, asynchronous.
func k_kernmulRSymm3D_async(fftMx unsafe.Pointer, fftMy unsafe.Pointer, fftMz unsafe.Pointer, fftKxx unsafe.Pointer, fftKyy unsafe.Pointer, fftKzz unsafe.Pointer, fftKyz unsafe.Pointer, fftKxz unsafe.Pointer, fftKxy unsafe.Pointer, N0 int, N1 int, N2 int, cfg *config, str cu.Stream) {
	if kernmulRSymm3D_code == 0 {
		kernmulRSymm3D_code = fatbinLoad(kernmulRSymm3D_map, "kernmulRSymm3D")
	}

	var a kernmulRSymm3D_args

	a.arg_fftMx = fftMx
	a.argptr[0] = unsafe.Pointer(&a.arg_fftMx)
	a.arg_fftMy = fftMy
	a.argptr[1] = unsafe.Pointer(&a.arg_fftMy)
	a.arg_fftMz = fftMz
	a.argptr[2] = unsafe.Pointer(&a.arg_fftMz)
	a.arg_fftKxx = fftKxx
	a.argptr[3] = unsafe.Pointer(&a.arg_fftKxx)
	a.arg_fftKyy = fftKyy
	a.argptr[4] = unsafe.Pointer(&a.arg_fftKyy)
	a.arg_fftKzz = fftKzz
	a.argptr[5] = unsafe.Pointer(&a.arg_fftKzz)
	a.arg_fftKyz = fftKyz
	a.argptr[6] = unsafe.Pointer(&a.arg_fftKyz)
	a.arg_fftKxz = fftKxz
	a.argptr[7] = unsafe.Pointer(&a.arg_fftKxz)
	a.arg_fftKxy = fftKxy
	a.argptr[8] = unsafe.Pointer(&a.arg_fftKxy)
	a.arg_N0 = N0
	a.argptr[9] = unsafe.Pointer(&a.arg_N0)
	a.arg_N1 = N1
	a.argptr[10] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[11] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	cu.LaunchKernel(kernmulRSymm3D_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for kernmulRSymm3D CUDA kernel, synchronized.
func k_kernmulRSymm3D(fftMx unsafe.Pointer, fftMy unsafe.Pointer, fftMz unsafe.Pointer, fftKxx unsafe.Pointer, fftKyy unsafe.Pointer, fftKzz unsafe.Pointer, fftKyz unsafe.Pointer, fftKxz unsafe.Pointer, fftKxy unsafe.Pointer, N0 int, N1 int, N2 int, cfg *config) {
	str := stream()
	k_kernmulRSymm3D_async(fftMx, fftMy, fftMz, fftKxx, fftKyy, fftKzz, fftKyz, fftKxz, fftKxy, N0, N1, N2, cfg, str)
	syncAndRecycle(str)
}

var kernmulRSymm3D_map = map[int]string{0: "",
	20: kernmulRSymm3D_ptx_20,
	30: kernmulRSymm3D_ptx_30,
	35: kernmulRSymm3D_ptx_35}

const (
	kernmulRSymm3D_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry kernmulRSymm3D(
	.param .u64 kernmulRSymm3D_param_0,
	.param .u64 kernmulRSymm3D_param_1,
	.param .u64 kernmulRSymm3D_param_2,
	.param .u64 kernmulRSymm3D_param_3,
	.param .u64 kernmulRSymm3D_param_4,
	.param .u64 kernmulRSymm3D_param_5,
	.param .u64 kernmulRSymm3D_param_6,
	.param .u64 kernmulRSymm3D_param_7,
	.param .u64 kernmulRSymm3D_param_8,
	.param .u32 kernmulRSymm3D_param_9,
	.param .u32 kernmulRSymm3D_param_10,
	.param .u32 kernmulRSymm3D_param_11
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<71>;
	.reg .f32 	%f<39>;
	.reg .s64 	%rd<36>;


	ld.param.u64 	%rd4, [kernmulRSymm3D_param_0];
	ld.param.u64 	%rd5, [kernmulRSymm3D_param_1];
	ld.param.u64 	%rd6, [kernmulRSymm3D_param_2];
	ld.param.u64 	%rd7, [kernmulRSymm3D_param_3];
	ld.param.u64 	%rd8, [kernmulRSymm3D_param_4];
	ld.param.u64 	%rd9, [kernmulRSymm3D_param_5];
	ld.param.u64 	%rd10, [kernmulRSymm3D_param_6];
	ld.param.u64 	%rd11, [kernmulRSymm3D_param_7];
	ld.param.u64 	%rd12, [kernmulRSymm3D_param_8];
	ld.param.u32 	%r14, [kernmulRSymm3D_param_9];
	ld.param.u32 	%r15, [kernmulRSymm3D_param_10];
	ld.param.u32 	%r16, [kernmulRSymm3D_param_11];
	.loc 2 35 1
	mov.u32 	%r17, %ctaid.y;
	mov.u32 	%r18, %ntid.y;
	mov.u32 	%r19, %tid.y;
	mad.lo.s32 	%r20, %r18, %r17, %r19;
	.loc 2 36 1
	mov.u32 	%r21, %ntid.x;
	mov.u32 	%r22, %ctaid.x;
	mov.u32 	%r23, %tid.x;
	mad.lo.s32 	%r24, %r21, %r22, %r23;
	.loc 2 38 1
	setp.lt.s32 	%p1, %r24, %r16;
	setp.lt.s32 	%p2, %r20, %r15;
	and.pred  	%p3, %p2, %p1;
	.loc 2 44 1
	setp.gt.s32 	%p4, %r14, 0;
	.loc 2 38 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_6;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 55 1
	shr.u32 	%r26, %r15, 31;
	add.s32 	%r27, %r15, %r26;
	shr.s32 	%r28, %r27, 1;
	add.s32 	%r1, %r28, 1;
	.loc 2 44 1
	sub.s32 	%r37, %r15, %r20;
	mad.lo.s32 	%r70, %r16, %r37, %r24;
	mad.lo.s32 	%r69, %r16, %r20, %r24;
	shl.b32 	%r64, %r69, 1;
	mov.u32 	%r65, 0;
	cvta.to.global.u64 	%rd13, %rd4;
	cvta.to.global.u64 	%rd15, %rd5;
	cvta.to.global.u64 	%rd16, %rd6;
	cvta.to.global.u64 	%rd27, %rd11;
	cvta.to.global.u64 	%rd30, %rd9;
	cvta.to.global.u64 	%rd32, %rd8;
	cvta.to.global.u64 	%rd34, %rd7;

BB0_2:
	.loc 2 44 1
	mov.u32 	%r67, %r70;
	mov.u32 	%r66, %r69;
	mov.u32 	%r7, %r67;
	mov.u32 	%r5, %r66;
	setp.lt.s32 	%p6, %r20, %r1;
	.loc 2 48 1
	mul.wide.s32 	%rd14, %r64, 4;
	add.s64 	%rd1, %rd13, %rd14;
	ld.global.f32 	%f1, [%rd1];
	.loc 2 49 1
	ld.global.f32 	%f2, [%rd1+4];
	.loc 2 50 1
	add.s64 	%rd2, %rd15, %rd14;
	ld.global.f32 	%f3, [%rd2];
	.loc 2 51 1
	ld.global.f32 	%f4, [%rd2+4];
	.loc 2 52 1
	add.s64 	%rd3, %rd16, %rd14;
	ld.global.f32 	%f5, [%rd3];
	.loc 2 53 1
	ld.global.f32 	%f6, [%rd3+4];
	.loc 2 55 1
	@%p6 bra 	BB0_4;

	cvta.to.global.u64 	%rd17, %rd10;
	.loc 2 67 1
	mul.wide.s32 	%rd18, %r7, 4;
	add.s64 	%rd19, %rd17, %rd18;
	ld.global.f32 	%f13, [%rd19];
	neg.f32 	%f38, %f13;
	cvta.to.global.u64 	%rd20, %rd12;
	.loc 2 69 1
	add.s64 	%rd21, %rd20, %rd18;
	ld.global.f32 	%f14, [%rd21];
	neg.f32 	%f37, %f14;
	mov.u32 	%r68, %r7;
	bra.uni 	BB0_5;

BB0_4:
	cvta.to.global.u64 	%rd22, %rd10;
	.loc 2 59 1
	mul.wide.s32 	%rd23, %r5, 4;
	add.s64 	%rd24, %rd22, %rd23;
	ld.global.f32 	%f38, [%rd24];
	cvta.to.global.u64 	%rd25, %rd12;
	.loc 2 61 1
	add.s64 	%rd26, %rd25, %rd23;
	ld.global.f32 	%f37, [%rd26];
	mov.u32 	%r68, %r5;

BB0_5:
	.loc 2 72 1
	mov.u32 	%r9, %r68;
	.loc 2 60 1
	mul.wide.s32 	%rd28, %r9, 4;
	add.s64 	%rd29, %rd27, %rd28;
	.loc 2 58 1
	add.s64 	%rd31, %rd30, %rd28;
	.loc 2 57 1
	add.s64 	%rd33, %rd32, %rd28;
	.loc 2 56 1
	add.s64 	%rd35, %rd34, %rd28;
	.loc 2 72 1
	ld.global.f32 	%f15, [%rd31];
	ld.global.f32 	%f16, [%rd33];
	ld.global.f32 	%f17, [%rd35];
	mul.f32 	%f18, %f3, %f37;
	fma.rn.f32 	%f19, %f1, %f17, %f18;
	ld.global.f32 	%f20, [%rd29];
	fma.rn.f32 	%f21, %f5, %f20, %f19;
	st.global.f32 	[%rd1], %f21;
	.loc 2 73 1
	mul.f32 	%f22, %f4, %f37;
	fma.rn.f32 	%f23, %f2, %f17, %f22;
	fma.rn.f32 	%f24, %f6, %f20, %f23;
	st.global.f32 	[%rd1+4], %f24;
	.loc 2 74 1
	mul.f32 	%f25, %f3, %f16;
	fma.rn.f32 	%f26, %f1, %f37, %f25;
	fma.rn.f32 	%f27, %f5, %f38, %f26;
	st.global.f32 	[%rd2], %f27;
	.loc 2 75 1
	mul.f32 	%f28, %f4, %f16;
	fma.rn.f32 	%f29, %f2, %f37, %f28;
	fma.rn.f32 	%f30, %f6, %f38, %f29;
	st.global.f32 	[%rd2+4], %f30;
	.loc 2 76 1
	mul.f32 	%f31, %f3, %f38;
	fma.rn.f32 	%f32, %f1, %f20, %f31;
	fma.rn.f32 	%f33, %f5, %f15, %f32;
	st.global.f32 	[%rd3], %f33;
	.loc 2 77 1
	mul.f32 	%f34, %f4, %f38;
	fma.rn.f32 	%f35, %f2, %f20, %f34;
	fma.rn.f32 	%f36, %f6, %f15, %f35;
	st.global.f32 	[%rd3+4], %f36;
	.loc 2 44 1
	mul.lo.s32 	%r62, %r16, %r15;
	mad.lo.s32 	%r70, %r16, %r15, %r7;
	shl.b32 	%r63, %r62, 1;
	add.s32 	%r64, %r64, %r63;
	mad.lo.s32 	%r69, %r16, %r15, %r5;
	.loc 2 44 18
	add.s32 	%r65, %r65, 1;
	.loc 2 44 1
	setp.lt.s32 	%p7, %r65, %r14;
	@%p7 bra 	BB0_2;

BB0_6:
	.loc 2 79 2
	ret;
}


`
	kernmulRSymm3D_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry kernmulRSymm3D(
	.param .u64 kernmulRSymm3D_param_0,
	.param .u64 kernmulRSymm3D_param_1,
	.param .u64 kernmulRSymm3D_param_2,
	.param .u64 kernmulRSymm3D_param_3,
	.param .u64 kernmulRSymm3D_param_4,
	.param .u64 kernmulRSymm3D_param_5,
	.param .u64 kernmulRSymm3D_param_6,
	.param .u64 kernmulRSymm3D_param_7,
	.param .u64 kernmulRSymm3D_param_8,
	.param .u32 kernmulRSymm3D_param_9,
	.param .u32 kernmulRSymm3D_param_10,
	.param .u32 kernmulRSymm3D_param_11
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<61>;
	.reg .f32 	%f<39>;
	.reg .s64 	%rd<36>;


	ld.param.u64 	%rd4, [kernmulRSymm3D_param_0];
	ld.param.u64 	%rd5, [kernmulRSymm3D_param_1];
	ld.param.u64 	%rd6, [kernmulRSymm3D_param_2];
	ld.param.u64 	%rd7, [kernmulRSymm3D_param_3];
	ld.param.u64 	%rd8, [kernmulRSymm3D_param_4];
	ld.param.u64 	%rd9, [kernmulRSymm3D_param_5];
	ld.param.u64 	%rd10, [kernmulRSymm3D_param_6];
	ld.param.u64 	%rd11, [kernmulRSymm3D_param_7];
	ld.param.u64 	%rd12, [kernmulRSymm3D_param_8];
	ld.param.u32 	%r22, [kernmulRSymm3D_param_9];
	ld.param.u32 	%r23, [kernmulRSymm3D_param_10];
	ld.param.u32 	%r24, [kernmulRSymm3D_param_11];
	.loc 2 35 1
	mov.u32 	%r1, %ctaid.y;
	mov.u32 	%r2, %ntid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 2 36 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r25, %r5, %r6, %r7;
	.loc 2 38 1
	setp.lt.s32 	%p1, %r25, %r24;
	setp.lt.s32 	%p2, %r4, %r23;
	and.pred  	%p3, %p2, %p1;
	.loc 2 44 1
	setp.gt.s32 	%p4, %r22, 0;
	.loc 2 38 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_6;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 55 1
	shr.u32 	%r27, %r23, 31;
	add.s32 	%r28, %r23, %r27;
	shr.s32 	%r29, %r28, 1;
	add.s32 	%r8, %r29, 1;
	.loc 2 44 1
	sub.s32 	%r32, %r23, %r4;
	mad.lo.s32 	%r60, %r24, %r32, %r25;
	mul.lo.s32 	%r10, %r24, %r23;
	mad.lo.s32 	%r59, %r24, %r4, %r25;
	shl.b32 	%r54, %r59, 1;
	mov.u32 	%r55, 0;
	cvta.to.global.u64 	%rd13, %rd4;
	cvta.to.global.u64 	%rd15, %rd5;
	cvta.to.global.u64 	%rd16, %rd6;
	cvta.to.global.u64 	%rd27, %rd11;
	cvta.to.global.u64 	%rd30, %rd9;
	cvta.to.global.u64 	%rd32, %rd8;
	cvta.to.global.u64 	%rd34, %rd7;

BB0_2:
	.loc 2 44 1
	mov.u32 	%r57, %r60;
	mov.u32 	%r56, %r59;
	mov.u32 	%r15, %r57;
	mov.u32 	%r13, %r56;
	.loc 2 48 1
	mul.wide.s32 	%rd14, %r54, 4;
	add.s64 	%rd1, %rd13, %rd14;
	ld.global.f32 	%f1, [%rd1];
	.loc 2 49 1
	ld.global.f32 	%f2, [%rd1+4];
	.loc 2 50 1
	add.s64 	%rd2, %rd15, %rd14;
	ld.global.f32 	%f3, [%rd2];
	.loc 2 51 1
	ld.global.f32 	%f4, [%rd2+4];
	.loc 2 52 1
	add.s64 	%rd3, %rd16, %rd14;
	ld.global.f32 	%f5, [%rd3];
	.loc 2 53 1
	ld.global.f32 	%f6, [%rd3+4];
	setp.lt.s32 	%p6, %r4, %r8;
	.loc 2 55 1
	@%p6 bra 	BB0_4;

	cvta.to.global.u64 	%rd17, %rd10;
	.loc 2 67 1
	mul.wide.s32 	%rd18, %r15, 4;
	add.s64 	%rd19, %rd17, %rd18;
	ld.global.f32 	%f13, [%rd19];
	neg.f32 	%f38, %f13;
	cvta.to.global.u64 	%rd20, %rd12;
	.loc 2 69 1
	add.s64 	%rd21, %rd20, %rd18;
	ld.global.f32 	%f14, [%rd21];
	neg.f32 	%f37, %f14;
	mov.u32 	%r58, %r15;
	bra.uni 	BB0_5;

BB0_4:
	cvta.to.global.u64 	%rd22, %rd10;
	.loc 2 59 1
	mul.wide.s32 	%rd23, %r13, 4;
	add.s64 	%rd24, %rd22, %rd23;
	ld.global.f32 	%f38, [%rd24];
	cvta.to.global.u64 	%rd25, %rd12;
	.loc 2 61 1
	add.s64 	%rd26, %rd25, %rd23;
	ld.global.f32 	%f37, [%rd26];
	mov.u32 	%r58, %r13;

BB0_5:
	.loc 2 72 1
	mov.u32 	%r17, %r58;
	.loc 2 60 1
	mul.wide.s32 	%rd28, %r17, 4;
	add.s64 	%rd29, %rd27, %rd28;
	.loc 2 58 1
	add.s64 	%rd31, %rd30, %rd28;
	.loc 2 57 1
	add.s64 	%rd33, %rd32, %rd28;
	.loc 2 56 1
	add.s64 	%rd35, %rd34, %rd28;
	.loc 2 72 1
	ld.global.f32 	%f15, [%rd31];
	ld.global.f32 	%f16, [%rd33];
	ld.global.f32 	%f17, [%rd35];
	mul.f32 	%f18, %f3, %f37;
	fma.rn.f32 	%f19, %f1, %f17, %f18;
	ld.global.f32 	%f20, [%rd29];
	fma.rn.f32 	%f21, %f5, %f20, %f19;
	st.global.f32 	[%rd1], %f21;
	.loc 2 73 1
	mul.f32 	%f22, %f4, %f37;
	fma.rn.f32 	%f23, %f2, %f17, %f22;
	fma.rn.f32 	%f24, %f6, %f20, %f23;
	st.global.f32 	[%rd1+4], %f24;
	.loc 2 74 1
	mul.f32 	%f25, %f3, %f16;
	fma.rn.f32 	%f26, %f1, %f37, %f25;
	fma.rn.f32 	%f27, %f5, %f38, %f26;
	st.global.f32 	[%rd2], %f27;
	.loc 2 75 1
	mul.f32 	%f28, %f4, %f16;
	fma.rn.f32 	%f29, %f2, %f37, %f28;
	fma.rn.f32 	%f30, %f6, %f38, %f29;
	st.global.f32 	[%rd2+4], %f30;
	.loc 2 76 1
	mul.f32 	%f31, %f3, %f38;
	fma.rn.f32 	%f32, %f1, %f20, %f31;
	fma.rn.f32 	%f33, %f5, %f15, %f32;
	st.global.f32 	[%rd3], %f33;
	.loc 2 77 1
	mul.f32 	%f34, %f4, %f38;
	fma.rn.f32 	%f35, %f2, %f20, %f34;
	fma.rn.f32 	%f36, %f6, %f15, %f35;
	st.global.f32 	[%rd3+4], %f36;
	.loc 2 44 1
	add.s32 	%r60, %r15, %r10;
	shl.b32 	%r53, %r10, 1;
	add.s32 	%r54, %r54, %r53;
	add.s32 	%r59, %r13, %r10;
	.loc 2 44 18
	add.s32 	%r55, %r55, 1;
	.loc 2 44 1
	setp.lt.s32 	%p7, %r55, %r22;
	@%p7 bra 	BB0_2;

BB0_6:
	.loc 2 79 2
	ret;
}


`
	kernmulRSymm3D_ptx_35 = `
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

.visible .entry kernmulRSymm3D(
	.param .u64 kernmulRSymm3D_param_0,
	.param .u64 kernmulRSymm3D_param_1,
	.param .u64 kernmulRSymm3D_param_2,
	.param .u64 kernmulRSymm3D_param_3,
	.param .u64 kernmulRSymm3D_param_4,
	.param .u64 kernmulRSymm3D_param_5,
	.param .u64 kernmulRSymm3D_param_6,
	.param .u64 kernmulRSymm3D_param_7,
	.param .u64 kernmulRSymm3D_param_8,
	.param .u32 kernmulRSymm3D_param_9,
	.param .u32 kernmulRSymm3D_param_10,
	.param .u32 kernmulRSymm3D_param_11
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<53>;
	.reg .f32 	%f<39>;
	.reg .s64 	%rd<36>;


	ld.param.u64 	%rd4, [kernmulRSymm3D_param_0];
	ld.param.u64 	%rd5, [kernmulRSymm3D_param_1];
	ld.param.u64 	%rd6, [kernmulRSymm3D_param_2];
	ld.param.u64 	%rd7, [kernmulRSymm3D_param_3];
	ld.param.u64 	%rd8, [kernmulRSymm3D_param_4];
	ld.param.u64 	%rd9, [kernmulRSymm3D_param_5];
	ld.param.u64 	%rd10, [kernmulRSymm3D_param_6];
	ld.param.u64 	%rd11, [kernmulRSymm3D_param_7];
	ld.param.u64 	%rd12, [kernmulRSymm3D_param_8];
	ld.param.u32 	%r22, [kernmulRSymm3D_param_9];
	ld.param.u32 	%r23, [kernmulRSymm3D_param_10];
	ld.param.u32 	%r24, [kernmulRSymm3D_param_11];
	.loc 3 35 1
	mov.u32 	%r1, %ctaid.y;
	mov.u32 	%r2, %ntid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 3 36 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r25, %r5, %r6, %r7;
	.loc 3 38 1
	setp.lt.s32 	%p1, %r25, %r24;
	setp.lt.s32 	%p2, %r4, %r23;
	and.pred  	%p3, %p2, %p1;
	.loc 3 44 1
	setp.gt.s32 	%p4, %r22, 0;
	.loc 3 38 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_6;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 55 1
	shr.u32 	%r27, %r23, 31;
	add.s32 	%r28, %r23, %r27;
	shr.s32 	%r29, %r28, 1;
	add.s32 	%r8, %r29, 1;
	.loc 3 44 1
	sub.s32 	%r32, %r23, %r4;
	mad.lo.s32 	%r52, %r24, %r32, %r25;
	mul.lo.s32 	%r10, %r24, %r23;
	mad.lo.s32 	%r51, %r24, %r4, %r25;
	shl.b32 	%r46, %r51, 1;
	mov.u32 	%r47, 0;
	cvta.to.global.u64 	%rd13, %rd4;
	cvta.to.global.u64 	%rd15, %rd5;
	cvta.to.global.u64 	%rd16, %rd6;
	cvta.to.global.u64 	%rd27, %rd11;
	cvta.to.global.u64 	%rd30, %rd9;
	cvta.to.global.u64 	%rd32, %rd8;
	cvta.to.global.u64 	%rd34, %rd7;

BB2_2:
	.loc 3 44 1
	mov.u32 	%r49, %r52;
	mov.u32 	%r48, %r51;
	mov.u32 	%r15, %r49;
	mov.u32 	%r13, %r48;
	.loc 3 48 1
	mul.wide.s32 	%rd14, %r46, 4;
	add.s64 	%rd1, %rd13, %rd14;
	ld.global.f32 	%f1, [%rd1];
	.loc 3 49 1
	ld.global.f32 	%f2, [%rd1+4];
	.loc 3 50 1
	add.s64 	%rd2, %rd15, %rd14;
	ld.global.f32 	%f3, [%rd2];
	.loc 3 51 1
	ld.global.f32 	%f4, [%rd2+4];
	.loc 3 52 1
	add.s64 	%rd3, %rd16, %rd14;
	ld.global.f32 	%f5, [%rd3];
	.loc 3 53 1
	ld.global.f32 	%f6, [%rd3+4];
	setp.lt.s32 	%p6, %r4, %r8;
	.loc 3 55 1
	@%p6 bra 	BB2_4;

	cvta.to.global.u64 	%rd17, %rd10;
	.loc 3 67 1
	mul.wide.s32 	%rd18, %r15, 4;
	add.s64 	%rd19, %rd17, %rd18;
	ld.global.nc.f32 	%f13, [%rd19];
	neg.f32 	%f38, %f13;
	cvta.to.global.u64 	%rd20, %rd12;
	.loc 3 69 1
	add.s64 	%rd21, %rd20, %rd18;
	ld.global.nc.f32 	%f14, [%rd21];
	neg.f32 	%f37, %f14;
	mov.u32 	%r50, %r15;
	bra.uni 	BB2_5;

BB2_4:
	cvta.to.global.u64 	%rd22, %rd10;
	.loc 3 59 1
	mul.wide.s32 	%rd23, %r13, 4;
	add.s64 	%rd24, %rd22, %rd23;
	ld.global.nc.f32 	%f38, [%rd24];
	cvta.to.global.u64 	%rd25, %rd12;
	.loc 3 61 1
	add.s64 	%rd26, %rd25, %rd23;
	ld.global.nc.f32 	%f37, [%rd26];
	mov.u32 	%r50, %r13;

BB2_5:
	.loc 3 72 1
	mov.u32 	%r17, %r50;
	.loc 3 60 1
	mul.wide.s32 	%rd28, %r17, 4;
	add.s64 	%rd29, %rd27, %rd28;
	.loc 3 58 1
	add.s64 	%rd31, %rd30, %rd28;
	.loc 3 57 1
	add.s64 	%rd33, %rd32, %rd28;
	.loc 3 56 1
	add.s64 	%rd35, %rd34, %rd28;
	.loc 3 72 1
	ld.global.nc.f32 	%f15, [%rd29];
	ld.global.nc.f32 	%f16, [%rd31];
	ld.global.nc.f32 	%f17, [%rd33];
	ld.global.nc.f32 	%f18, [%rd35];
	mul.f32 	%f19, %f3, %f37;
	fma.rn.f32 	%f20, %f1, %f18, %f19;
	fma.rn.f32 	%f21, %f5, %f15, %f20;
	st.global.f32 	[%rd1], %f21;
	.loc 3 73 1
	mul.f32 	%f22, %f4, %f37;
	fma.rn.f32 	%f23, %f2, %f18, %f22;
	fma.rn.f32 	%f24, %f6, %f15, %f23;
	st.global.f32 	[%rd1+4], %f24;
	.loc 3 74 1
	mul.f32 	%f25, %f3, %f17;
	fma.rn.f32 	%f26, %f1, %f37, %f25;
	fma.rn.f32 	%f27, %f5, %f38, %f26;
	st.global.f32 	[%rd2], %f27;
	.loc 3 75 1
	mul.f32 	%f28, %f4, %f17;
	fma.rn.f32 	%f29, %f2, %f37, %f28;
	fma.rn.f32 	%f30, %f6, %f38, %f29;
	st.global.f32 	[%rd2+4], %f30;
	.loc 3 76 1
	mul.f32 	%f31, %f3, %f38;
	fma.rn.f32 	%f32, %f1, %f15, %f31;
	fma.rn.f32 	%f33, %f5, %f16, %f32;
	st.global.f32 	[%rd3], %f33;
	.loc 3 77 1
	mul.f32 	%f34, %f4, %f38;
	fma.rn.f32 	%f35, %f2, %f15, %f34;
	fma.rn.f32 	%f36, %f6, %f16, %f35;
	st.global.f32 	[%rd3+4], %f36;
	.loc 3 44 1
	add.s32 	%r52, %r15, %r10;
	shl.b32 	%r45, %r10, 1;
	add.s32 	%r46, %r46, %r45;
	add.s32 	%r51, %r13, %r10;
	.loc 3 44 18
	add.s32 	%r47, %r47, 1;
	.loc 3 44 1
	setp.lt.s32 	%p7, %r47, %r22;
	@%p7 bra 	BB2_2;

BB2_6:
	.loc 3 79 2
	ret;
}


`
)