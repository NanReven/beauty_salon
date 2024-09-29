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

type MasterResponse struct {
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	Email      string `db:"email"`
	Position   string `db:"title"`
	Bio        string `db:"bio"`
}
