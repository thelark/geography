package geography

import "math"

/**
 * 向量类
 */
type Vector struct {
	From *Coordinate
	To   *Coordinate
}

/**
 * 向量长度
 */
func (v *Vector) Length() float64 {
	return math.Sqrt(math.Pow(float64(v.To.Latitude-v.From.Latitude), 2) + math.Pow(float64(v.To.Longitude-v.From.Longitude), 2))
}
