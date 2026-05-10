package models

type MoveEvent struct {
	EntityID  uint64
	Pos       Vec2
	TargetPos Vec2
	Rotation  float32
	Speed     float32
}
