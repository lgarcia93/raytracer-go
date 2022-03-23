package geo

import "math"

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vec3) Subtract(other Vec3) Vec3 {
	return Vec3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vec3) DivideBy(value float64) Vec3 {
	return Vec3{
		X: v.X / value,
		Y: v.Y / value,
		Z: v.Z / value,
	}
}

func (v Vec3) MultiplyBy(value float64) Vec3 {
	return Vec3{
		X: v.X * value,
		Y: v.Y * value,
		Z: v.Z * value,
	}
}

func (v Vec3) Dot(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec3) Normalize() Vec3 {
	//length := v.Length()

	return v.DivideBy(v.Length())
}
