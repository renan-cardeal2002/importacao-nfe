package domain

type Produto struct {
	ID             int
	EmpresaID      int
	EAN            string
	CodProduto     string
	DescrProduto   string
	uCom           string // verificar campo
	Quantidade     int
	vUnCom         float64 // verificar campo
	ValorUnitario  float64
	Custo          float64
	Preco          float64
	Margem         float64
	ValorAdicional float64
}
