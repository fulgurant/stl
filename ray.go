package stl

import (
	"math"
)

var epsilon = math.Nextafter(1, 2) - 1

type Ray struct {
	origin Vec3
	vector Vec3
}

func (ray Ray) IntersectsTriangle(triangle Triangle) (Vec3, bool) {
	edge1 := triangle.Vertices[1].Diff(triangle.Vertices[0])
	edge2 := triangle.Vertices[2].Diff(triangle.Vertices[0])
	rayCrossE2 := ray.vector.Cross(edge2)
	det := edge1.Dot(rayCrossE2)

	if det > -epsilon && det < epsilon {
		println("ray is parallel to triangle")
		// ray is parallel to triangle
		return Vec3Zero, false
	}

	invDet := 1.0 / det
	s := ray.origin.Diff(triangle.Vertices[0])
	u := invDet * s.Dot(rayCrossE2)

	if u < 0.0 || u > 1.0 {
		println("2")
		return Vec3Zero, false
	}

	sCrossE1 := s.Cross(edge1)
	v := invDet * ray.vector.Dot(sCrossE1)

	if v < 0 || u+v > 1 {
		println("3")
		return Vec3Zero, false
	}

	t := invDet * edge2.Dot(sCrossE1)

	if t > epsilon {
		println("4")
		outIntersectionPoint := ray.origin.Add(ray.vector.MultScalar(t))
		return outIntersectionPoint, true
	}

	println("5")
	return Vec3Zero, false
}
