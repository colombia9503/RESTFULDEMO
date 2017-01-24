package models

type Usuario struct {
	Nombre  string `json:"Nombre" db:"first_name"`
	Usuario string `json:"Usuario" db:"user_id"`
}
