package entities

type Client struct {
	ID                                int
	TipoPersona                       string
	Nombre                            string
	Apellido                          string
	NumeroDocumento                   string
	TipoDocumento                     string
	NombreRepresentanteLegal          string
	CalidadEnLaQueActua               string
	TipoDocumentoRepresentanteLegal   string
	NumeroDocumentoRepresentanteLegal string
	Direccion                         string
	Email                             string
	TelefonoFijo                      string
	Celular                           string
	Ciudad                            string
}

func (Client) TableName() string {
	return "client"
}
