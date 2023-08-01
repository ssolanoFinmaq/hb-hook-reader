package dtos

import (
	"fmt"
	"time"

	"github.com/FinMaq/hook-reader/reader/domain/entities"
)

type Output struct {
	Fecha                             string
	TipoPersona                       string
	TipoProducto                      string
	Sector                            string
	Nombre                            string
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
	NombreDistribuidor                string
	NITDistribuidor                   string
	TipoAcuerdo                       string
	FechaFirmaAcuerdoCredito          string
	FechaFirmaAcuerdoRenting          string
	NombrePersonaQueFirma             string
	TipoDocumentoPersonaQueFirma      string
	NumeroDocumentoPersonaQueFirma    string
	CalidadEnLaQueFirma               string
	NumeroProducto                    string
	ConceptoFinal                     string
	Valor                             string
	Plazo                             string
	Tasa                              string
	Cuota                             string
	Garantias                         string
	ComentariosComite                 string
	MiembrosComite                    string
	MontoAFinanciar                   string
	ValorGMF                          string
	DescripcionMaquina                string
	CupoRotativoInicial               string
	ValorGMFCupoRotativo              string
	CupoRotativoConGMF                string
	TipoOferta                        string
	NumeroOferta                      string
	SerialEquiposAFinanciar           string
	CiudadMaquina                     string
	DepartamentoMaquina               string
	DireccionMaquina                  string
	CaracteristicasTecnicasMaquina    string
	ValorMaquinaLetras                string
	ValorMaquina                      string
	PorcentajeAFinanciar              string
	Riesgo2B                          string
	NombreCodeudor1                   string
	NombreCodeudor2                   string
	NombreCodeudor3                   string
	NombreCodeudor4                   string
	NombreCodeudor5                   string
	TipoDocumentoCodeudor1            string
	TipoDocumentoCodeudor2            string
	TipoDocumentoCodeudor3            string
	TipoDocumentoCodeudor4            string
	TipoDocumentoCodeudor5            string
	NumeroDocumentoCodeudor1          string
	NumeroDocumentoCodeudor2          string
	NumeroDocumentoCodeudor3          string
	NumeroDocumentoCodeudor4          string
	NumeroDocumentoCodeudor5          string
	MarcaEquipo                       string
	Canon                             string
	IVACanon                          string
	SeguroTodoRiesgo                  string
	CanonTotal                        string
	ValorSalvamento                   string
}

func ProductsToOutputs(products []entities.Products) []Output {
	outputs := []Output{}
	for _, product := range products {
		output := Output{
			Fecha:                             time.Now().Format("DD/MM/YYYY"),
			TipoPersona:                       product.Client.TipoPersona,
			TipoProducto:                      product.Tipo,
			Sector:                            product.Sector,
			Nombre:                            fmt.Sprint(product.Client.Nombre, " ", product.Client.Apellido),
			NumeroDocumento:                   product.Client.NumeroDocumento,
			TipoDocumento:                     product.Client.TipoDocumento,
			NombreRepresentanteLegal:          product.Client.NombreRepresentanteLegal,
			CalidadEnLaQueActua:               product.Client.CalidadEnLaQueActua,
			TipoDocumentoRepresentanteLegal:   product.Client.TipoDocumentoRepresentanteLegal,
			NumeroDocumentoRepresentanteLegal: product.Client.NumeroDocumentoRepresentanteLegal,
			Direccion:                         product.Client.Direccion,
			Email:                             product.Client.Email,
			TelefonoFijo:                      product.Client.TelefonoFijo,
			Celular:                           product.Client.Celular,
			Ciudad:                            product.Client.Ciudad,
			NombreDistribuidor:                product.Distribuitor.Nombre,
			NITDistribuidor:                   product.Distribuitor.NIT,
			TipoAcuerdo:                       product.Distribuitor.TipoAcuerdo,
			FechaFirmaAcuerdoCredito:          product.Distribuitor.FechaFirmaAcuerdoCredito,
			FechaFirmaAcuerdoRenting:          product.Distribuitor.FechaFirmaAcuerdoRenting,
			NombrePersonaQueFirma:             product.Distribuitor.NombrePersonaQueFirma,
			TipoDocumentoPersonaQueFirma:      product.Distribuitor.TipoDocumentoPersonaQueFirma,
			NumeroDocumentoPersonaQueFirma:    product.Distribuitor.NumeroDocumentoPersonaQueFirma,
			CalidadEnLaQueFirma:               product.Distribuitor.CalidadEnLaQueFirma,
			NumeroProducto:                    product.NumeroProducto,
			ConceptoFinal:                     product.ConceptoFinal,
			Valor:                             product.Valor,
			Plazo:                             product.Plazo,
			Tasa:                              product.Tasa,
			Cuota:                             product.Cuota,
			Garantias:                         product.Garantias,
			ComentariosComite:                 product.ComentariosComite,
			MiembrosComite:                    product.MiembrosComite,
			MontoAFinanciar:                   product.MontoAFinanciar,
			ValorGMF:                          product.ValorGMF,
			DescripcionMaquina:                product.DescripcionMaquina,
			CupoRotativoInicial:               product.CupoRotativoInicial,
			ValorGMFCupoRotativo:              product.ValorGMFCupoRotativo,
			CupoRotativoConGMF:                product.CupoRotativoConGMF,
			TipoOferta:                        product.TipoOferta,
			NumeroOferta:                      product.NumeroOferta,
			SerialEquiposAFinanciar:           product.SerialEquiposAFinanciar,
			CiudadMaquina:                     product.CiudadMaquina,
			DepartamentoMaquina:               product.DepartamentoMaquina,
			DireccionMaquina:                  product.DireccionMaquina,
			CaracteristicasTecnicasMaquina:    product.CaracteristicasTecnicasMaquina,
			ValorMaquinaLetras:                product.ValorMaquinaLetras,
			ValorMaquina:                      product.ValorMaquina,
			PorcentajeAFinanciar:              product.PorcentajeAFinanciar,
			Riesgo2B:                          product.Riesgo2B,
			NombreCodeudor1:                   product.NombreCodeudor1,
			NombreCodeudor2:                   product.NombreCodeudor2,
			NombreCodeudor3:                   product.NombreCodeudor3,
			NombreCodeudor4:                   product.NombreCodeudor4,
			NombreCodeudor5:                   product.NombreCodeudor5,
			TipoDocumentoCodeudor1:            product.TipoDocumentoCodeudor1,
			TipoDocumentoCodeudor2:            product.TipoDocumentoCodeudor2,
			TipoDocumentoCodeudor3:            product.TipoDocumentoCodeudor3,
			TipoDocumentoCodeudor4:            product.TipoDocumentoCodeudor4,
			TipoDocumentoCodeudor5:            product.TipoDocumentoCodeudor5,
			NumeroDocumentoCodeudor1:          product.NumeroDocumentoCodeudor1,
			NumeroDocumentoCodeudor2:          product.NumeroDocumentoCodeudor2,
			NumeroDocumentoCodeudor3:          product.NumeroDocumentoCodeudor3,
			NumeroDocumentoCodeudor4:          product.NumeroDocumentoCodeudor4,
			NumeroDocumentoCodeudor5:          product.NumeroDocumentoCodeudor5,
			MarcaEquipo:                       product.MarcaEquipo,
			Canon:                             product.Canon,
			IVACanon:                          product.IVACanon,
			SeguroTodoRiesgo:                  product.SeguroTodoRiesgo,
			CanonTotal:                        product.CanonTotal,
			ValorSalvamento:                   product.ValorSalvamento,
		}
		outputs = append(outputs, output)
	}
	return outputs
}
