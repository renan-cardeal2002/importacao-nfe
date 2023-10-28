package xmlController

import (
	"encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
	"importa-nfe/services"
	"os"
)

func GetProdutos(c *gin.Context) {
	xmlFile, err := os.Open("./docs/41230910541434000152550010000012411749316397-nfe.xml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer xmlFile.Close()

	data, err := xmlService.LerXml(xmlFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	parser, err := xmlService.NewNfeParser(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	prod, err := parser.Produtos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var produtos []map[string]interface{}
	if err := json.Unmarshal(prod, &produtos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.IndentedJSON(http.StatusOK, produtos)
}

func GetEmitente(c *gin.Context) {
	xmlFile, err := os.Open("./docs/41230910541434000152550010000012411749316397-nfe.xml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer xmlFile.Close()

	data, err := xmlService.LerXml(xmlFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	parser, err := xmlService.NewNfeParser(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	emit, err := parser.Emitente()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var emitente map[string]interface{}
	if err := json.Unmarshal(emit, &emitente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.IndentedJSON(http.StatusOK, emitente)
}

func GetDestinatario(c *gin.Context) {
	xmlFile, err := os.Open("./docs/41230910541434000152550010000012411749316397-nfe.xml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer xmlFile.Close()

	data, err := xmlService.LerXml(xmlFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	parser, err := xmlService.NewNfeParser(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dest, err := parser.Destinatario()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var destinatario map[string]interface{}
	if err := json.Unmarshal(dest, &destinatario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.IndentedJSON(http.StatusOK, destinatario)
}

func InserirNFE(c *gin.Context) {
	xmlService.VerificarEmit("10541434000152", "10541434000152")
	// xmlService.InserirProdutos()
	// xmlService.InserirDest()

    data := gin.H{"mensagem": "NFe inserida com sucesso"}
    c.IndentedJSON(http.StatusOK, data)
}