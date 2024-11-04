package ports

type ProdutosRepository interface {
	InserirProdutos(produtosJSON string, CNPJ string) error
}

type DestinatarioRepository interface {
	InserirDest(destinatarioJSON string, CNPJ string) error
}

type EmpresaRepository interface {
	VerificarEmit(cnpjEmit string, cnpjLogado string) error
}
