package dummy

var Nombre string

func init() {
	Nombre = "Fernando"
}
func HolaMundo() string {
	return "Hola " + Nombre
}
