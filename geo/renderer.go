package geo

import (
	"image/color"
	"math"
)

type renderer struct {
	imageWidth     int
	aspectRatio    float64
	imageHeight    int
	viewportWidth  float64
	viewportHeight float64
	focalLenght    float64
}

func InitRenderer(outputWidth int, aspectRatio float64) *renderer {
	return &renderer{
		imageWidth:     outputWidth,
		imageHeight:    int(float64(outputWidth) / aspectRatio),
		aspectRatio:    aspectRatio,
		viewportWidth:  2.0 * aspectRatio,
		viewportHeight: 2.0,
		focalLenght:    1.0,
	}
}

func (r *renderer) Render(Origin Vec3, spheres []Sphere, LightDirection Vec3) {

	img := CreateImage(r.imageWidth, r.aspectRatio)

	horizontalVec := Vec3{X: r.viewportWidth, Y: 0.0, Z: 0.0}
	verticalVec := Vec3{X: 0.0, Y: r.viewportHeight, Z: 0.0}
	focalLengthVec := Vec3{X: 0.0, Y: 0.0, Z: r.focalLenght}

	lowerLeftCorner := Origin.Subtract(horizontalVec.DivideBy(2.0)).Subtract(verticalVec.DivideBy(2.0)).Subtract(focalLengthVec)

	for i := 0; i < r.imageHeight; i++ {
		for j := 0; j < r.imageWidth; j++ {
			img.SetPixelColor(j, i, color.RGBA{0, 0, 0, 255})
		}
	}

	for i := 0; i < r.imageHeight; i++ {
		for j := 0; j < r.imageWidth; j++ {
			for _, sphere := range spheres {

				u := float64(j) / float64(r.imageWidth-1)
				v := float64(i) / float64(r.imageHeight-1)

				direction := lowerLeftCorner.Add(horizontalVec.MultiplyBy(u)).Add(verticalVec.MultiplyBy(v)).Subtract(Origin)

				ray := Ray{
					Origin:    Origin,
					Direction: direction,
				}

				k := ray.Origin.Subtract(sphere.Center)

				a := ray.Direction.Dot(ray.Direction)

				b := 2 * k.Dot(ray.Direction)

				c := k.Dot(k) - sphere.Radius*sphere.Radius

				discriminant := b*b - 4.0*a*c

				t1 := (-b + math.Sqrt(discriminant)) / (2.0 * a)

				if discriminant >= 0 {
					iSectPoint := ray.Origin.Add(ray.Direction.MultiplyBy(t1))

					surfaceNormal := iSectPoint.Subtract(sphere.Center).Normalize()

					shading := LightDirection.Dot(surfaceNormal)

					// Clamp the shading result
					if shading < 0 {
						shading = 0
					}

					if shading > 1 {
						shading = 1
					}

					img.SetPixelColor(j, i, color.RGBA{uint8(255 * shading), uint8(255 * shading), uint8(255 * shading), 255})
				}
			}
		}

	}

	img.Save("render.png")
}
