package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"importa-nfe/internal/core/ports"
	xmlService "importa-nfe/internal/services"
	"importa-nfe/pkg"
	"net/http"
	"os"
	"strconv"
)

type ImportationController struct {
	produtosRepository     ports.ProdutosRepository
	destinatarioRepository ports.DestinatarioRepository
	empresaRepository      ports.EmpresaRepository
	transactionManager     ports.TransactionManager
}

func NewImportationHandler(produtosRepository ports.ProdutosRepository, destinatarioRepository ports.DestinatarioRepository, empresaRepository ports.EmpresaRepository, transactionManager ports.TransactionManager) ImportationController {
	return ImportationController{
		produtosRepository:     produtosRepository,
		destinatarioRepository: destinatarioRepository,
		empresaRepository:      empresaRepository,
		transactionManager:     transactionManager,
	}
}

func buscaProdutos(nfe string) ([]map[string]interface{}, error) {
	xmlFile, err := os.Open("./docs/" + nfe)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	xml, err := pkg.ReadXML(xmlFile)
	if err != nil {
		return nil, err
	}

	parser, err := pkg.NewNfeParser(xml)
	if err != nil {
		return nil, err
	}

	prod, err := xmlService.Produtos(parser)
	if err != nil {
		return nil, err
	}

	var produtos []map[string]interface{}
	if err = json.Unmarshal(prod, &produtos); err != nil {
		return nil, err
	}

	return produtos, nil
}

func buscaEmitente(nfe string) (map[string]interface{}, error) {
	xmlFile, err := os.Open("./docs/" + nfe)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	data, err := pkg.ReadXML(xmlFile)
	if err != nil {
		return nil, err
	}

	parser, err := pkg.NewNfeParser(data)
	if err != nil {
		return nil, err
	}

	emit, err := xmlService.Emitente(parser)
	if err != nil {
		return nil, err
	}

	var emitente map[string]interface{}
	if err := json.Unmarshal(emit, &emitente); err != nil {
		return nil, err
	}

	return emitente, nil
}

func buscaDestinatario(nfe string) (map[string]interface{}, error) {
	xmlFile, err := os.Open("./docs/" + nfe)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	data, err := pkg.ReadXML(xmlFile)
	if err != nil {
		return nil, err
	}

	parser, err := pkg.NewNfeParser(data)
	if err != nil {
		return nil, err
	}

	dest, err := xmlService.Destinatario(parser)
	if err != nil {
		return nil, err
	}

	var destinatario map[string]interface{}
	if err := json.Unmarshal(dest, &destinatario); err != nil {
		return nil, err
	}

	return destinatario, nil
}

func (h ImportationController) GetProdutos(c *gin.Context) {
	ctx := c.Request.Context()
	empresaIDStr := c.DefaultQuery("empresa_id", "")
	empresaID, err := strconv.Atoi(empresaIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empresa_id precisa ser um número válido"})
		return
	}

	produtos, err := h.produtosRepository.FindProdutosByEmpresaID(ctx, empresaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, produtos)
}

func (h ImportationController) GetEmitente(c *gin.Context) {
	nfe := c.DefaultQuery("nfe", "")

	emitente, err := buscaEmitente(nfe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, emitente)
}

func (h ImportationController) GetDestinatario(c *gin.Context) {
	nfe := c.DefaultQuery("nfe", "")

	destinatario, err := buscaDestinatario(nfe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, destinatario)
}

func (h ImportationController) InserirNFE(c *gin.Context) {
	ctx := c.Request.Context()

	nfe := c.DefaultQuery("nfe", "")
	cnpj := c.DefaultQuery("cnpj", "")

	empresa, err := h.empresaRepository.FindByCNPJ(ctx, cnpj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	produtos, err := buscaProdutos(nfe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	prod, err := json.Marshal(produtos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	destinatario, err := buscaDestinatario(nfe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dest, err := json.Marshal(destinatario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx, err := h.transactionManager.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stringProd := string(prod)
	err = h.produtosRepository.InserirProdutos(ctx, stringProd, empresa.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_ = h.transactionManager.Rollback(tx)
		return
	}

	err = h.destinatarioRepository.InserirDest(ctx, empresa.ID, string(dest))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_ = h.transactionManager.Rollback(tx)
		return
	}

	err = h.transactionManager.Commit(tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"mensagem": "NFe inserida com sucesso"})
}
