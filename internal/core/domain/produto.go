package domain

type Produto struct {
	ID             int
	EmpresaID      int
	EAN            string
	CodProduto     string
	DescrProduto   string
	UnidadeMedida  string
	Quantidade     float64
	PrecoUnitario  float64
	PrecoBruto     float64
	Custo          float64
	Preco          float64
	Margem         float64
	ValorAdicional float64
}
