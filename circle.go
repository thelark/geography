package geography

type Circle struct {
	C *Coordinate
	R float64
}

/**
 * 圆区域包含某坐标
 */
func (c *Circle) ContainCoordinate(coordinate *Coordinate) bool {
	containCoordinate := false
	if (&Vector{From: c.C, To: coordinate}).Length() < c.R {
		containCoordinate = true
	}
	return containCoordinate
}

/**
 * 上极限坐标
 */
func (c *Circle) Top() *Coordinate {
	return &Coordinate{Latitude: c.C.Latitude - c.R, Longitude: c.C.Longitude}
}

/**
 * 下极限坐标
 */
func (c *Circle) Bottom() *Coordinate {
	return &Coordinate{Latitude: c.C.Latitude + c.R, Longitude: c.C.Longitude}
}

/**
 * 左极限坐标
 */
func (c *Circle) Left() *Coordinate {
	return &Coordinate{Latitude: c.C.Latitude, Longitude: c.C.Longitude - c.R}
}

/**
 * 右极限坐标
 */
func (c *Circle) Right() *Coordinate {
	return &Coordinate{Latitude: c.C.Latitude, Longitude: c.C.Longitude + c.R}
}
