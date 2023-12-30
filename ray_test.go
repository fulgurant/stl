package stl

import (
	"testing"
)

func TestRayIntercectsTriangle(t *testing.T) {
	testCases := []struct {
		description    string
		triangle       Triangle
		ray            Ray
		intercept      bool
		interceptPoint Vec3
	}{
		{
			description: "horizontal t, ray above, heading down inside t",
			triangle: Triangle{
				Vertices: [3]Vec3{
					Vec3{0, 0, 0},
					Vec3{1, 0, 0},
					Vec3{0, 1, 0},
				},
			},
			ray: Ray{
				Origin:    Vec3{0.1, 0.1, 1},
				Direction: Vec3{0, 0, -1},
			},
			intercept:      true,
			interceptPoint: Vec3{0.1, 0.1, 0},
		},
		{
			description: "horizontal t, ray above, heading down on diagonal",
			triangle: Triangle{
				Vertices: [3]Vec3{
					Vec3{0, 0, 0},
					Vec3{1, 0, 0},
					Vec3{0, 1, 0},
				},
			},
			ray: Ray{
				Origin:    Vec3{0.5, 0.5, 1},
				Direction: Vec3{0, 0, -1},
			},
			intercept:      true,
			interceptPoint: Vec3{0.5, 0.5, 0},
		},
		{
			description: "horizontal t, ray above, heading down beyond diagonal",
			triangle: Triangle{
				Vertices: [3]Vec3{
					Vec3{0, 0, 0},
					Vec3{1, 0, 0},
					Vec3{0, 1, 0},
				},
			},
			ray: Ray{
				Origin:    Vec3{0.51, 0.51, 1},
				Direction: Vec3{0, 0, -1},
			},
			intercept: false,
		},
		{
			description: "horizontal t, ray above, heading down less than x",
			triangle: Triangle{
				Vertices: [3]Vec3{
					Vec3{0, 0, 0},
					Vec3{1, 0, 0},
					Vec3{0, 1, 0},
				},
			},
			ray: Ray{
				Origin:    Vec3{-0.1, 0.1, 1},
				Direction: Vec3{0, 0, -1},
			},
			intercept: false,
		},
		{
			description: "horizontal t, ray above, heading down less than y",
			triangle: Triangle{
				Vertices: [3]Vec3{
					Vec3{0, 0, 0},
					Vec3{1, 0, 0},
					Vec3{0, 1, 0},
				},
			},
			ray: Ray{
				Origin:    Vec3{0.1, -0.1, 1},
				Direction: Vec3{0, 0, -1},
			},
			intercept: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			p, i := testCase.ray.IntersectsTriangle(testCase.triangle)
			if i != testCase.intercept {
				t.Fatalf("expected intercection %v, but received %v", testCase.intercept, i)
			}

			if i && !testCase.interceptPoint.AlmostEqual(p, float32(epsilon)) {
				t.Fatalf("expected intercection point %v, but received %v", testCase.interceptPoint, p)
			}
		})
	}
}
