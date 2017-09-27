package geography

import (
	"math"
	"fmt"
)

/**
 * 坐标类
 */
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

/**
 * 两坐标点之间的距离
 */
func (c *Coordinate) Distance(cT *Coordinate) float64 {
	return float64(math.Sqrt(math.Pow(c.Longitude-cT.Longitude, 2) + math.Pow(c.Latitude-cT.Latitude, 2)))
}

/**
 * 按纬度方式输出 2°3′13″
 */
func (c *Coordinate) Lat() string {
	lat := c.Latitude
	d := math.Floor(lat)                   // 度
	f := math.Floor((lat - d) * 60)        // 分
	m := math.Floor(((lat-d)*60 - f) * 60) // 秒
	return fmt.Sprintf("%.0f°%.0f′%.0f″", d, f, m)
}

/**
 * 按经度方式输出 2°3′13″
 */
func (c *Coordinate) Lng() string {
	lat := c.Longitude
	d := math.Floor(lat)                   // 度
	f := math.Floor((lat - d) * 60)        // 分
	m := math.Floor(((lat-d)*60 - f) * 60) // 秒
	return fmt.Sprintf("%.0f°%.0f′%.0f″", d, f, m)
}

/**
 * 判断点坐标是否在向量的左侧
 */
func (c *Coordinate) IsLeftInVector(v *Vector) bool {
	/*
	 * A/v.From B/v.To C/c
	 * 用 A、B、C 三点构建三角形
	 * S(A,B,C) = ((A.X - C.X) * (B.Y - C.Y) - (A.Y - C.Y) * (B.X - C.X)) / 2
	 * 如果 S(A,B,C) > 0 则表明 ABC 三点顺时针 => C点 在 AB线段 左侧
	 * 如果 S(A,B,C) < 0 则表明 ABC 三点逆时针 => C点 在 AB线段 右侧
	 */
	return ((v.From.Latitude-c.Latitude)*(v.To.Longitude-c.Longitude)-(v.From.Longitude-c.Longitude)*(v.To.Latitude-c.Latitude))/2 > 0
}

/**
 * 判断点坐标是否在向量的右侧
 */
func (c *Coordinate) IsRightInVector(v *Vector) bool {
	/*
	 * A/v.From B/v.To C/c
	 * 用 A、B、C 三点构建三角形
	 * S(A,B,C) = ((A.X - C.X) * (B.Y - C.Y) - (A.Y - C.Y) * (B.X - C.X)) / 2
	 * 如果 S(A,B,C) > 0 则表明 ABC 三点顺时针 => C点 在 AB线段 左侧
	 * 如果 S(A,B,C) < 0 则表明 ABC 三点逆时针 => C点 在 AB线段 右侧
	 */
	return !c.IsLeftInVector(v)
}

/**
 * 判断点坐标是否在向量上
 */
func (c *Coordinate) IsInVector(v *Vector) bool {
	/*
	 * A/v.From B/v.To C/c
	 * 用 A、B、C 三点构建三角形
	 * S(A,B,C) = ((A.X - C.X) * (B.Y - C.Y) - (A.Y - C.Y) * (B.X - C.X)) / 2
	 * 如果 S(A,B,C) > 0 则表明 ABC 三点顺时针 => C点 在 AB线段 左侧
	 * 如果 S(A,B,C) < 0 则表明 ABC 三点逆时针 => C点 在 AB线段 右侧
	 * 如果 S(A,B,C) = 0 则表明 ABC 一条线 => C点 在 AB线段 上
	 */
	return ((v.From.Latitude-c.Latitude)*(v.To.Longitude-c.Longitude)-(v.From.Longitude-c.Longitude)*(v.To.Latitude-c.Latitude))/2 == 0
}

/**
 * 判断点坐标是否在区域内
 */
func (c *Coordinate) IsInRegion(r *Region) bool {
	isInRegion := true
	for k, coordinate := range r.C {
		i := k
		if k == len(r.C)-1 {
			i = 0
		}
		if c.IsLeftInVector(&Vector{coordinate, r.C[i]}) {
			isInRegion = false
			break
		}
	}
	return isInRegion
}
