// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"math"
	"math/rand"
	"testing"

	"github.com/gonum/floats"
)

func TestBCG(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
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
		market("nos1", 1e-9),
		market("nos4", 1e-12),
		market("nos5", 1e-12),
		market("bcsstm20", 1e-10),
		market("bcsstm22", 1e-11),

		market("e05r0000", 1e-11),
		market("e05r0100", 1e-11),
		market("e05r0200", 1e-10),
		market("e05r0300", 1e-10),
		market("e05r0400", 1e-10),
		market("e05r0500", 1e-10),
		market("gre_216a", 1e-12),
		market("gre__115", 1e-12),
		market("gre__185", 1e-9),
		market("gre__343", 1e-12),
		market("impcol_b", 1e-9),
		market("impcol_c", 1e-10),
		market("steam1", 1e-8),
		market("steam3", 1e-8),
		market("west0067", 1e-11),
	} {
		n := tc.n
		// Compute the right-hand side b so that the vector [1,1,...,1]
		// is the solution.
		want := make([]float64, n)
		for i := range want {
			want[i] = 1
		}
		sys := System{
			MatVec: tc.a,
			B:      make([]float64, n),
		}
		sys.MatVec(sys.B, want, false)

		r, err := Iterative(sys, nil, &BCG{}, Settings{
			MaxIterations: 10 * tc.iters,
			Tolerance:     1e-14,
		})
		if err != nil {
			t.Errorf("Case %v (n=%v): unexpected error %v", tc.name, n, err)
			continue
		}
		dist := floats.Distance(r.X, want, math.Inf(1))
		if dist > tc.tol {
			t.Errorf("Case %v (n=%v): unexpected solution, |want-got|=%v", tc.name, n, dist)
		}
	}
}
