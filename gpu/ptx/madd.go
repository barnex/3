package ptx

const MADD = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Wed Aug  1 02:51:19 2012 (1343782279)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00006815_00000000-9_madd.cpp3.i"
	.file	2 "/home/arne/src/nimble-cube/gpu/ptx/madd.cu"

.visible .entry madd(
	.param .u64 madd_param_0,
	.param .u64 madd_param_1,
	.param .f32 madd_param_2,
	.param .u64 madd_param_3,
	.param .f32 madd_param_4,
	.param .u32 madd_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<12>;
	.reg .f32 	%f<7>;
	.reg .s64 	%rd<11>;


	ld.param.u64 	%rd4, [madd_param_0];
	ld.param.u64 	%rd5, [madd_param_1];
	ld.param.f32 	%f1, [madd_param_2];
	ld.param.u64 	%rd6, [madd_param_3];
	ld.param.f32 	%f2, [madd_param_4];
	ld.param.u32 	%r2, [madd_param_5];
	cvta.to.global.u64 	%rd1, %rd4;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 2 3 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 4 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	.loc 2 5 1
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd3, %rd7;
	ld.global.f32 	%f3, [%rd8];
	add.s64 	%rd9, %rd2, %rd7;
	ld.global.f32 	%f4, [%rd9];
	mul.ftz.f32 	%f5, %f4, %f2;
	fma.rn.ftz.f32 	%f6, %f3, %f1, %f5;
	add.s64 	%rd10, %rd1, %rd7;
	st.global.f32 	[%rd10], %f6;

BB0_2:
	.loc 2 7 2
	ret;
}


`