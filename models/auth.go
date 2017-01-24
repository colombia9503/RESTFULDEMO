package models

import "github.com/colombia9503/RESTful-SqlServer/common"

type Credencial struct {
	Nombre  string `json:"Nombre"`
	Usuario string `json:"Usuario"`
	Token   string `json:"Token"`
}

type auth struct{}

var Auth = new(auth)

func (auth) Login(user Usuario) (Credencial, error) {
	var cred Credencial
	row := common.DB.QueryRow("Select first_name, user_id from pc_user_def where user_id = '" + user.Usuario + "';")
	if err := row.Scan(&cred.Nombre, &cred.Usuario); err != nil {
		return cred, err
	}

	//Validacion de la clave de usuario
	return cred, nil
}
