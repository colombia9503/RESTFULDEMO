package models

import "github.com/colombia9503/RESTful-SqlServer/common"

type Costo struct {
	Codigo string `json:"codigo" db:"codigo"`
	Costo  int    `json:"costo" db:"costo"`
}

type Verificado struct {
	Existe bool `json:"existe"`
}

var Costos = new(costos)

type costos struct{}

func (costos) SelectAll() ([]Costo, error) {
	cts := []Costo{}
	rows, err := common.DB.Query("select codigo, costo from costoitem;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Costo
		if err := rows.Scan(&c.Codigo, &c.Costo); err != nil {
			return nil, err
		}
		cts = append(cts, c)
	}
	return cts, nil
}

func (costos) Exists(cod string) Verificado {
	var c Costo
	row := common.DB.QueryRow("select codigo, costo from costoitem where codigo=?;", cod)
	if err := row.Scan(&c.Codigo, &c.Costo); err != nil {
		return Verificado{Existe: false}
	}
	return Verificado{Existe: true}
}

func (costos) SelectOne(cod string) (Costo, error) {
	var c Costo
	row := common.DB.QueryRow("select codigo, costo from costoitem where codigo=?;", cod)
	return c, row.Scan(&c.Codigo, &c.Costo)
}

func (costos) Insert(cost Costo) (*Verificado, error) {
	var verf *Verificado
	stmt, err := common.DB.Prepare("insert into costoitem (codigo, costo) values (?,?);")
	if err != nil {
		verf = &Verificado{
			Existe: false,
		}
		return verf, err
	}
	_, err = stmt.Exec(cost.Codigo, cost.Costo)
	if err != nil {
		verf = &Verificado{
			Existe: false,
		}
		return verf, err
	}
	defer stmt.Close()
	verf = &Verificado{
		Existe: true,
	}
	return verf, nil
}
