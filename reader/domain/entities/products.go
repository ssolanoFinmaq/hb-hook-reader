package entities

type Products struct {
	ID                             string `json:"-"`
	ContactId                      int
	Client                         Client `gorm:"foreignKey:ContactId"`
	CompanyId                      int
	Distribuitor                   Distribuitor `gorm:"foreignKey:CompanyId"`
	Tipo                           string
	Sector                         string
	NumeroProducto                 string
	ConceptoFinal                  string
	Valor                          string
	Plazo                          string
	Tasa                           string
	Cuota                          string
	Garantias                      string
	ComentariosComite              string
	MiembrosComite                 string
	MontoAFinanciar                string
	ValorGMF                       string
	DescripcionMaquina             string
	CupoRotativoInicial            string
	ValorGMFCupoRotativo           string
	CupoRotativoConGMF             string
	TipoOferta                     string
	NumeroOferta                   string
	SerialEquiposAFinanciar        string
	CiudadMaquina                  string
	DepartamentoMaquina            string
	DireccionMaquina               string
	CaracteristicasTecnicasMaquina string
	ValorMaquinaLetras             string
	ValorMaquina                   string
	PorcentajeAFinanciar           string
	Riesgo2B                       string
	NombreCodeudor1                string
	NombreCodeudor2                string
	NombreCodeudor3                string
	NombreCodeudor4                string
	NombreCodeudor5                string
	TipoDocumentoCodeudor1         string
	TipoDocumentoCodeudor2         string
	TipoDocumentoCodeudor3         string
	TipoDocumentoCodeudor4         string
	TipoDocumentoCodeudor5         string
	NumeroDocumentoCodeudor1       string
	NumeroDocumentoCodeudor2       string
	NumeroDocumentoCodeudor3       string
	NumeroDocumentoCodeudor4       string
	NumeroDocumentoCodeudor5       string
	MarcaEquipo                    string
	Canon                          string
	IVACanon                       string
	SeguroTodoRiesgo               string
	CanonTotal                     string
	ValorSalvamento                string
}
