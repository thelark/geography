package geography

/**
 * TODO: 地理 - 区域类
 */
type Region struct {
	C []*Coordinate
}

/**
 * 获取所有坐标
 */
func (r *Region) Coordinates() []*Coordinate {
	return r.C
}

/**
 * 获取所有向量 - 顺时针
 */
func (r *Region) Vectors() []*Vector {
	vectors := make([]*Vector, 0)
	for index, coordinate := range r.Coordinates() {
		i := index + 1
		if index == len(r.C)-1 {
			i = 0
		}
		vectors = append(vectors, &Vector{coordinate, r.C[i]})
	}
	return vectors
}

/**
 * 计算周长
 */
func (r *Region) Perimeter() float64 {
	perimeter := 0.0
	for _, vector := range r.Vectors() {
		perimeter += vector.Length()
	}
	return perimeter
}

/**
 * 区域包含某坐标
 */
func (r *Region) ContainCoordinate(coordinate *Coordinate) bool {
	j := len(r.C) - 1
	containCoordinate := false
	for index := 0; index < len(r.C); index++ {
		if r.C[index].Latitude < coordinate.Latitude && r.C[j].Latitude >= coordinate.Latitude ||
			r.C[j].Latitude < coordinate.Latitude && r.C[index].Latitude >= coordinate.Latitude {
			if r.C[index].Longitude+(coordinate.Latitude-r.C[index].Latitude)/(r.C[j].Latitude-r.C[index].Latitude)*(r.C[j].Longitude-r.C[index].Longitude) < coordinate.Longitude {
				containCoordinate = !containCoordinate
			}
		}
		j = index
	}
	return containCoordinate
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
