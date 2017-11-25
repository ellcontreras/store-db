package db

import (
	"database/sql"
	"store/config"

	_ "github.com/go-sql-driver/mysql"
)

type Conexion struct {
	dbuser     string
	dbpassword string
	dbname     string
}

var (
	db  *sql.DB
	err error
)

func (this *Conexion) AbrirConexion() {
	db, err = sql.Open("mysql", config.Config["dbuser"]+":"+config.Config["dbpassword"]+"@/"+config.Config["dbname"])
}

func (this *Conexion) ObtenerConexion() *sql.DB {
	return db
}

func (this *Conexion) CerrarConexion() {
	db.Close()
}
