package gk

import (
	"math"
	"math/rand"
)

type number interface {
	int | int32 | int64 | uint | uint32 | uint64 | float32 | float64
}

func MaxU8[T number](a, b T) uint8 {
	if a > b {
		return uint8(a)
	}

	return uint8(b)
}

func MinU8[T number](a, b T) uint8 {
	if a < b {
		return uint8(a)
	}

	return uint8(b)
}

func Pow2[T number](n T) T {
	return n * n
}

func Sqrt[T number](n T) float64 {
	return math.Sqrt(float64(n))
}

func Dist[T number](x1, y1, x2, y2 T) float64 {
	return Sqrt(Pow2(x2-x1) + Pow2(y2-y1))
}

func DistPos(pos1, pos2 Pos) float64 {
	return Dist(pos1.X, pos1.Y, pos2.X, pos2.Y)
}

func Rand[T number](min, max T) float64 {
	return float64(min) + rand.Float64()*(float64(max-min))
}

func RandAngle() float64 {
	return rand.Float64() * 2 * math.Pi
}

func RandN(n int) int32 {
	return int32(rand.Intn(n))
}

func RandN2(lb, up int32) int32 {
	return int32(rand.Intn(int(up-lb+1))) + lb
}

func Scale(v int32, factor float32) int32 {
	return int32(float32(v) * factor)
}

func Clamp[T number](v, min, max T) T {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}
