package gpu

import (
	"code.google.com/p/mx3/core"
	"code.google.com/p/mx3/gpu/ptx"
	"github.com/barnex/cuda5/safe"
)

func kernMulRSymm2Dyz(fftMy, fftMz safe.Complex64s, K11, K22, K12 safe.Float32s, N1, N2 int) {
	core.Assert(K11.Len() == (N1/2+1)*N2)
	gridDim, blockDim := Make2DConf(N1, N2)
	ptx.K_kernmulRSymm2Dyz(fftMy.Pointer(), fftMz.Pointer(),
		K11.Pointer(), K22.Pointer(), K12.Pointer(),
		N1, N2, gridDim, blockDim)
}

func kernMulRSymm2Dx(fftMx safe.Complex64s, K00 safe.Float32s, N1, N2 int) {
	core.Assert(K00.Len() == (N1/2+1)*N2)
	gridDim, blockDim := Make2DConf(N1, N2)
	ptx.K_kernmulRSymm2Dx(fftMx.Pointer(), K00.Pointer(), N1, N2, gridDim, blockDim)
}

func kernMulRSymm3D(fftM [3]safe.Complex64s, K00, K11, K22, K12, K02, K01 safe.Float32s, N0, N1, N2 int) {
	core.Assert(K11.Len() == N0*N1*N2)
	gridDim, blockDim := Make2DConf(N1, N2)
	ptx.K_kernmulRSymm3D(fftM[0].Pointer(), fftM[1].Pointer(), fftM[2].Pointer(),
		K00.Pointer(), K11.Pointer(), K22.Pointer(), K12.Pointer(), K02.Pointer(), K01.Pointer(),
		N0, N1, N2, gridDim, blockDim)
}

// General kernel multiplication with general complex kernel.
// (stored in interleaved format).
// It might be more clear if the kernel were stored as safe.Complex64s.
func kernMulC(fftM [3]safe.Complex64s, K [3][3]safe.Float32s) {
	core.Assert(2*fftM[0].Len() == K[0][0].Len())
	N := fftM[0].Len()
	gridDim, blockDim := Make1DConf(N)
	ptx.K_kernmulC(fftM[0].Pointer(), fftM[1].Pointer(), fftM[2].Pointer(),
		K[0][0].Pointer(), K[1][1].Pointer(), K[2][2].Pointer(),
		K[1][2].Pointer(), K[0][2].Pointer(), K[0][1].Pointer(),
		K[2][1].Pointer(), K[2][0].Pointer(), K[1][0].Pointer(),
		N, gridDim, blockDim)
}