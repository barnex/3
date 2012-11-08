package conv

import (
	"code.google.com/p/nimble-cube/gpu"
	"code.google.com/p/nimble-cube/mag"
	"code.google.com/p/nimble-cube/nimble"
	"fmt"
	"github.com/barnex/fmath"
	"math/rand"
)

// Internal: main function for conv test.
func TestSymm2(N0, N1, N2 int) {
	C := 1e-9
	mesh := nimble.NewMesh(N0, N1, N2, C, 2*C, 3*C)
	nimble.Log(mesh)
	N := mesh.NCell()

	gpu.LockCudaThread()
	hin := nimble.MakeChan3("hin", "", mesh, nimble.UnifiedMemory)
	hout := nimble.MakeChan3("hout", "", mesh, nimble.UnifiedMemory)

	acc := 1
	kern := mag.BruteKernel(nimble.ZeroPad(mesh), acc)

	// TODO: const chan
	for i := 0; i < 3; i++ {
		arr := hin.Comp(i).WriteNext(N).Host()
		initConvTestInput(nimble.Reshape(arr, hin.Size()))
		hin.Comp(i).WriteDone()
	}
	F := 10
	go func() {
		for i := 1; i < F; i++ {
			hin.WriteNext(N)
			hin.WriteDone()
		}
	}()

	//go NewSymmetricHtoD(mesh, kern, hin.NewReader(), hout).Run()
	go NewSymm2D(mesh, kern, hin, hout).Run()

	houtR := hout.NewReader()
	var outarr [3]nimble.Slice
	for i := 0; i < F; i++ {
		outarr = houtR.ReadNext(N)
		houtR.ReadDone()
	}
	outarrh := [3][]float32{outarr[0].Host(), outarr[1].Host(), outarr[2].Host()}

	arr := nimble.MakeVectors(mesh.Size())
	ref := nimble.MakeVectors(mesh.Size())
	Brute(arr, ref, kern)
	checkErr(outarrh, nimble.Contiguous3(ref))
}

// generate sparse input data
func initConvTestInput(input [][][]float32) {
	size := nimble.SizeOf(input)
	N0, N1, N2 := size[0], size[1], size[2]
	is := [...]int{N0 - 1} //	is := [...]int{0, N0 / 5, N0 / 2, N0 - 1}
	js := [...]int{N1 - 1} //	js := [...]int{0, N1 / 7, N1 / 2, N1 - 1}
	ks := [...]int{N2 - 1} //	ks := [...]int{0, N2 / 11, N2 / 2, N2 - 1}
	for _, i := range is {
		for _, j := range js {
			for _, k := range ks {
				input[i][j][k] = rnd()
			}
		}
	}
}

// Interface of any convolution.
//type Conv interface {
//	Input() [3][][][]float32     // Input data
//	Output() [3][][][]float32    // Output data
//	Kernel() [3][3][][][]float32 // Convolution kernel
//	Exec()                       // Executes the convolution
//}
//
//type Constructor func(size [3]int, kernel [3][3][][][]float32) Conv
//
//// Test if the convolution gives the same result as the brute-force implementation.
//func Test(c Conv) {
//	if !*nimble.Flag_verify {
//		nimble.Log("skipping convolution self-test")
//		return
//	}
//	input := c.Input()
//	output := c.Output()
//	size := nimble.SizeOf(input[0])
//
//	// overwrite input/output data, if any.
//	in := nimble.Contiguous3(input)
//	out := nimble.Contiguous3(output)
//	for i := range in {
//		for j := range in[i] {
//			in[i][j] = 0
//			out[i][j] = 666
//		}
//	}
//
//	initConvTestInput(input)
//
//	// Reference solution
//	bruteOut := nimble.MakeVectors(size)
//	Brute(input, bruteOut, c.Kernel())
//	ref := nimble.Contiguous3(bruteOut)
//
//	// solution under test
//	c.Exec()
//	c.Exec()
//	c.Exec() // it may fail the 2nd time, eg. 
//
//	checkErr(ref, out)
//
//	// cleanly set input/output to zero
//	for i := range in {
//		for j := range in[i] {
//			in[i][j] = 0
//			out[i][j] = 0
//		}
//	}
//}

func checkErr(ref, out [3][]float32) {
	// check if error is OK
	var maxerr float32
	for c := range ref {
		for i := range ref[c] {
			if fmath.Abs(out[c][i]-ref[c][i]) > maxerr {
				maxerr = fmath.Abs(out[c][i] - ref[c][i])
			}
		}
	}
	const tolerance = 1e-5
	if maxerr > tolerance {
		//	core.Fprint(os.Stderr, "expected:\n")
		//	core.Fprintf(os.Stderr, "% 6e", bruteOut)
		//	core.Fprint(os.Stderr, "got:\n")
		//	core.Fprintf(os.Stderr, "% 6e", c.Output())
		nimble.Fatal(fmt.Errorf("convolution self-test failed with error %v", maxerr))
	}
	nimble.Log("convolution test error:", maxerr, "OK")
}

// random number between -1 and 1.
func rnd() float32 {
	return 1 - 2*rand.Float32()
}
