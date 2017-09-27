package geography

/**
 * 区域类
 */
type Region struct {
	C []*Coordinate
}

/**
 * 上极限坐标
 */
func (r *Region) Top() *Coordinate {
	c := r.C[0]
	for _, v := range r.C {
		if v.Longitude >= c.Longitude {
			c = v
		}
	}
	return c
}

/**
 * 下极限坐标
 */
func (r *Region) Bottom() *Coordinate {
	c := r.C[0]
	for _, v := range r.C {
		if v.Longitude <= c.Longitude {
			c = v
		}
	}
	return c
}

/**
 * 左极限坐标
 */
func (r *Region) Left() *Coordinate {
	c := r.C[0]
	for _, v := range r.C {
		if v.Latitude <= c.Latitude {
			c = v
		}
	}
	return c
}

/**
 * 右极限坐标
 */
func (r *Region) Right() *Coordinate {
	c := r.C[0]
	for _, v := range r.C {
		if v.Latitude >= c.Latitude {
			c = v
		}
	}
	return c
}
