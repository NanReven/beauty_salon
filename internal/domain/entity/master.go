package entity

type Position struct {
	Id    int
	Title string
}

type Master struct {
	Id         int
	UserId     int
	PositionId int
	Bio        string
	Slug       string
}
