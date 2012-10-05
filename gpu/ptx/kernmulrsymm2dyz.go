package ptx

const KERNMULRSYMM2DYZ = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Wed Aug  1 02:51:19 2012 (1343782279)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00001412_00000000-9_kernmulrsymm2dyz.cpp3.i"
	.file	2 "/home/arne/src/nimble-cube/gpu/ptx/kernmulrsymm2dyz.cu"

.visible .entry kernmulRSymm2Dyz(
	.param .u64 kernmulRSymm2Dyz_param_0,
	.param .u64 kernmulRSymm2Dyz_param_1,
	.param .u64 kernmulRSymm2Dyz_param_2,
	.param .u64 kernmulRSymm2Dyz_param_3,
	.param .u64 kernmulRSymm2Dyz_param_4,
	.param .u32 kernmulRSymm2Dyz_param_5,
	.param .u32 kernmulRSymm2Dyz_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<34>;
	.reg .f32 	%f<20>;
	.reg .s64 	%rd<24>;


	ld.param.u64 	%rd6, [kernmulRSymm2Dyz_param_0];
	ld.param.u64 	%rd7, [kernmulRSymm2Dyz_param_1];
	ld.param.u64 	%rd8, [kernmulRSymm2Dyz_param_2];
	ld.param.u64 	%rd9, [kernmulRSymm2Dyz_param_3];
	ld.param.u64 	%rd10, [kernmulRSymm2Dyz_param_4];
	ld.param.u32 	%r6, [kernmulRSymm2Dyz_param_5];
	ld.param.u32 	%r7, [kernmulRSymm2Dyz_param_6];
	cvta.to.global.u64 	%rd1, %rd7;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd9;
	cvta.to.global.u64 	%rd5, %rd8;
	.loc 2 29 1
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r1, %r8, %r9, %r10;
	.loc 2 30 1
	mov.u32 	%r11, %ntid.x;
	mov.u32 	%r12, %ctaid.x;
	mov.u32 	%r13, %tid.x;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	.loc 2 32 1
	setp.ge.s32 	%p1, %r2, %r7;
	setp.ge.s32 	%p2, %r1, %r6;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_5;

	.loc 2 36 1
	mad.lo.s32 	%r3, %r1, %r7, %r2;
	.loc 2 37 1
	sub.s32 	%r14, %r6, %r1;
	mad.lo.s32 	%r4, %r14, %r7, %r2;
	.loc 2 43 1
	shr.u32 	%r15, %r6, 31;
	add.s32 	%r16, %r6, %r15;
	shr.s32 	%r17, %r16, 1;
	add.s32 	%r18, %r17, 1;
	setp.lt.s32 	%p4, %r1, %r18;
	@%p4 bra 	BB0_3;

	.loc 2 50 1
	mul.wide.s32 	%rd11, %r4, 4;
	add.s64 	%rd12, %rd3, %rd11;
	ld.global.f32 	%f4, [%rd12];
	neg.ftz.f32 	%f19, %f4;
	mov.u32 	%r33, %r4;
	bra.uni 	BB0_4;

BB0_3:
	.loc 2 46 1
	mul.wide.s32 	%rd13, %r3, 4;
	add.s64 	%rd14, %rd3, %rd13;
	ld.global.f32 	%f19, [%rd14];
	mov.u32 	%r33, %r3;

BB0_4:
	.loc 2 53 1
	mov.u32 	%r5, %r33;
	.loc 2 45 1
	mul.wide.s32 	%rd15, %r5, 4;
	add.s64 	%rd16, %rd4, %rd15;
	.loc 2 44 1
	add.s64 	%rd17, %rd5, %rd15;
	.loc 2 53 1
	ld.global.f32 	%f5, [%rd16];
	shl.b32 	%r22, %r3, 1;
	.loc 2 55 1
	mul.wide.s32 	%rd18, %r22, 4;
	add.s64 	%rd19, %rd2, %rd18;
	.loc 2 56 1
	add.s32 	%r23, %r22, 1;
	mul.wide.s32 	%rd20, %r23, 4;
	add.s64 	%rd21, %rd2, %rd20;
	ld.global.f32 	%f6, [%rd21];
	.loc 2 57 1
	add.s64 	%rd22, %rd1, %rd18;
	.loc 2 58 1
	add.s64 	%rd23, %rd1, %rd20;
	ld.global.f32 	%f7, [%rd23];
	.loc 2 55 1
	ld.global.f32 	%f8, [%rd19];
	.loc 2 53 1
	ld.global.f32 	%f9, [%rd17];
	.loc 2 57 1
	ld.global.f32 	%f10, [%rd22];
	.loc 2 60 1
	mul.ftz.f32 	%f11, %f10, %f19;
	fma.rn.ftz.f32 	%f12, %f8, %f9, %f11;
	st.global.f32 	[%rd19], %f12;
	.loc 2 61 1
	mul.ftz.f32 	%f13, %f7, %f19;
	fma.rn.ftz.f32 	%f14, %f6, %f9, %f13;
	st.global.f32 	[%rd21], %f14;
	.loc 2 62 1
	mul.ftz.f32 	%f15, %f10, %f5;
	fma.rn.ftz.f32 	%f16, %f8, %f19, %f15;
	st.global.f32 	[%rd22], %f16;
	.loc 2 63 1
	mul.ftz.f32 	%f17, %f7, %f5;
	fma.rn.ftz.f32 	%f18, %f6, %f19, %f17;
	st.global.f32 	[%rd23], %f18;

BB0_5:
	.loc 2 64 2
	ret;
}


`