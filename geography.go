package geography

/**
 * 地理接口
 */
type Geography interface {
	Top() *Coordinate
	Bottom() *Coordinate
	Left() *Coordinate
	Right() *Coordinate
}
