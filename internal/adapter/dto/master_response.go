package dto

type MasterResponse struct {
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	Email      string `db:"email"`
	Position   string `db:"title"`
	Bio        string `db:"bio"`
}
