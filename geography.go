package geography

/**
 * TODO: 地理接口
 */
type Geography interface {
	Top() *Coordinate
	Bottom() *Coordinate
	Left() *Coordinate
	Right() *Coordinate
}

type Shape interface {
	ContainCoordinate(coordinate *Coordinate) bool
	Top() *Coordinate
	Bottom() *Coordinate
	Left() *Coordinate
	Right() *Coordinate
}
