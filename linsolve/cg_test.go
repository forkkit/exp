// Copyright ©2017 The Gonum authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"math"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/floats"
)

func TestCG(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
testLoop:
	for _, tc := range []testCase{
		randomSPD(1, rnd),
		randomSPD(2, rnd),
		randomSPD(3, rnd),
		randomSPD(4, rnd),
		randomSPD(5, rnd),
		randomSPD(10, rnd),
		randomSPD(20, rnd),
		randomSPD(50, rnd),
		randomSPD(100, rnd),
		randomSPD(200, rnd),
		randomSPD(500, rnd),
		market("spd_100_nos4", 1e-12),
		market("spd_138_bcsstm22", 1e-10),
		market("spd_237_nos1", 1e-10),
		market("spd_468_nos5", 1e-11),
		market("spd_485_bcsstm20", 1e-12),
		market("spd_900_gr_30_30", 1e-12),
	} {
		n := tc.n
		// Compute the right-hand side b so that the vector [1,1,...,1]
		// is the solution.
		want := make([]float64, n)
		for i := range want {
			want[i] = 2 + 0.1*float64(i)
		}
		b := make([]float64, n)
		tc.mulvec(b, want, false)

		ctx := Context{
			X:        make([]float64, n),
			Residual: make([]float64, n),
			Src:      make([]float64, n),
			Dst:      make([]float64, n),
		}
		// Initial guess is a random vector.
		for i := range ctx.X {
			ctx.X[i] = rnd.NormFloat64()
		}
		// Compute the initial residual.
		tc.mulvec(ctx.Residual, ctx.X, false)
		floats.AddScaledTo(ctx.Residual, b, -1, ctx.Residual)

		var cg CG
		var itercount int
		cg.Init(n)
	cgLoop:
		for {
			op, err := cg.Iterate(&ctx)
			if err != nil {
				t.Errorf("Case %v (n=%v): unexpected error %v", tc.name, n, err)
				continue testLoop
			}
			switch op {
			case MulVec:
				tc.mulvec(ctx.Dst, ctx.Src, false)
			case PreconSolve:
				copy(ctx.Dst, ctx.Src)
			case CheckResidual:
				rnorm := floats.Norm(ctx.Residual, math.Inf(1))
				if rnorm < 1e-13 {
					ctx.Converged = true
				}
			case EndIteration:
				itercount++
				if ctx.Converged {
					break cgLoop
				}
				if itercount == tc.iters {
					t.Logf("Case %v (n=%v): %v exceeded iteration limit", tc.name, n, itercount)
					break cgLoop
				}
			}
		}

		tol := tc.tol
		if tol == 0 {
			tol = 1e-12
		}
		dist := floats.Distance(ctx.X, want, math.Inf(1))
		if dist > tol {
			t.Errorf("Case %v (n=%v): unexpected solution, |want-got|=%v", tc.name, n, dist)
		}
	}
}