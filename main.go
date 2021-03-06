package main

import (
	"raytracer/geo"
)

func main() {
	renderer := geo.InitRenderer(1920, float64(16.0/9.0))

	origin := geo.Vec3{
		X: 0.0,
		Y: 0.0,
		Z: 0.0,
	}

	sphere1 := geo.Sphere{
		Center: geo.Vec3{
			X: 0.0,
			Y: 0.0,
			Z: -3.0,
		},
		Radius: 1.1,
	}

	sphere2 := geo.Sphere{
		Center: geo.Vec3{
			X: 0.0,
			Y: 2.0,
			Z: -3.0,
		},
		Radius: 0.9,
	}
	sphere3 := geo.Sphere{
		Center: geo.Vec3{
			X: 1.0,
			Y: 2.0,
			Z: -3.0,
		},
		Radius: 0.2,
	}

	lightDirection := geo.Vec3{
		X: 0.0,
		Y: 0.0,
		Z: -1,
	}

	renderer.Render(
		origin,
		[]geo.Sphere{
			sphere1,
			sphere2,
			sphere3,
		},
		lightDirection.Normalize(),
	)
}
