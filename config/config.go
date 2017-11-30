package config

// Config es el mapa de datos en donde guardamos cada configuracion
var Config = make(map[string]string)

// Init es la funcion que inicializa cada parte de la configuracion
func Init() {
	Config["dbuser"] = "root"
	Config["dbpassword"] = "root"
	Config["dbname"] = "store"
	Config["puerto"] = ":8080"
}
