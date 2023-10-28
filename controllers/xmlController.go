package xmlController

import (
	"encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
	"importa-nfe/services"
	"os"
)

func buscaProdutos(nfe string) ([]map[string]interface{}, error) {
	xmlFile, err := os.Open("./docs/" + nfe)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	data, err := xmlService.LerXml(xmlFile)
	if err != nil {
		return nil, err
	}

	parser, err := xmlService.NewNfeParser(data)
	if err != nil {
		return nil, err
	}

	prod, err := parser.Produtos()
	if err != nil {
		return nil, err
	}

	var produtos []map[string]interface{}
	if err := json.Unmarshal(prod, &produtos); err != nil {
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

	data, err := xmlService.LerXml(xmlFile)
	if err != nil {
		return nil, err
	}

	parser, err := xmlService.NewNfeParser(data)
	if err != nil {
		return nil, err
	}

	emit, err := parser.Emitente()
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

	data, err := xmlService.LerXml(xmlFile)
	if err != nil {
		return nil, err
	}

	parser, err := xmlService.NewNfeParser(data)
	if err != nil {
		return nil, err
	}

	dest, err := parser.Destinatario()
	if err != nil {
		return nil, err
	}

	var destinatario map[string]interface{}
	if err := json.Unmarshal(dest, &destinatario); err != nil {
		return nil, err
	}

	return destinatario, nil
}

func GetProdutos(c *gin.Context) {
	nfe := c.DefaultQuery("nfe", "")

	produtos, err := buscaProdutos(nfe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.IndentedJSON(http.StatusOK, produtos)
}

func GetEmitente(c *gin.Context) {
	nfe := c.DefaultQuery("nfe", "")

	emitente, err := buscaEmitente(nfe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.IndentedJSON(http.StatusOK, emitente)
}

func GetDestinatario(c *gin.Context) {
	nfe := c.DefaultQuery("nfe", "")

	destinatario, err := buscaDestinatario(nfe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.IndentedJSON(http.StatusOK, destinatario)
}

func InserirNFE(c *gin.Context) {
	nfe := c.DefaultQuery("nfe", "")
	cnpj := c.DefaultQuery("cnpj", "")

    produtos, err := buscaProdutos(nfe)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Converte os produtos em JSON
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

    // Converte o destinatario em JSON
    dest, err := json.Marshal(destinatario)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Insere os produtos no banco de dados
    err = xmlService.InserirProdutos(string(prod), cnpj)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Realiza outras operações, como inserção de destinatário (se necessário)
    err = xmlService.InserirDest(string(dest), cnpj)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    data := gin.H{"mensagem": "NFe inserida com sucesso"}
    c.IndentedJSON(http.StatusOK, data)
}