package GoRetro

import "math"

type Vector struct {
	X float64
	Y float64
}

func vectorAdd(v1, v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func vectorDistance(v1, v2 Vector) float64 {
	return math.Sqrt(math.Pow(v2.X-v1.X, 2) +
		math.Pow(v2.Y-v1.Y, 2))
}
