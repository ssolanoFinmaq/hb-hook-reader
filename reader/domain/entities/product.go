package entities

type Product struct {
	ID                            string
	GenerarDocumentosLegalizacion string
}

func (Product) TableName() string {
	return "products"
}
