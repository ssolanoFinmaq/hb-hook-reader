package entities

type Distribuitor struct {
	ID                             int
	Nombre                         string
	NIT                            string
	TipoAcuerdo                    string
	FechaFirmaAcuerdoCredito       string
	FechaFirmaAcuerdoRenting       string
	NombrePersonaQueFirma          string
	TipoDocumentoPersonaQueFirma   string
	NumeroDocumentoPersonaQueFirma string
	CalidadEnLaQueFirma            string
}

func (Distribuitor) TableName() string {
	return "distributor"
}
