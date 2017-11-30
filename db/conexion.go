package db

import (
	"database/sql"
	"store/config"

	_ "github.com/go-sql-driver/mysql"
)

// Conexion es el tipo de dato que va a realizar la conexion a una db
type Conexion struct {
	Dbuser     string
	Dbpassword string
	Dbname     string
}

var (
	// db es la variable en donde almacenaremos nuestra conexion a la db
	db *sql.DB

	// err es la variable que va a revisar que no haya ningún error
	err error
)

// AbrirConexion es la funcion que abre una conexion a la db
func (este *Conexion) AbrirConexion() {
	db, err = sql.Open("mysql", config.Config["dbuser"]+":"+config.Config["dbpassword"]+"@/"+config.Config["dbname"])
}

// ObtenerConexion es la funcion que devuelve una conexion a la db
func (este *Conexion) ObtenerConexion() *sql.DB {
	return db
}

// CerrarConexion es la funcion que abre una conexion a la db
func (este *Conexion) CerrarConexion() {
	db.Close()
}
