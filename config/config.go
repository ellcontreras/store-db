package config

var Config map[string]string = make(map[string]string)

func Init() {
	Config["dbuser"] = "root"
	Config["dbpassword"] = "root"
	Config["dbname"] = "ejemplo"
	Config["puerto"] = ":8080"
}
