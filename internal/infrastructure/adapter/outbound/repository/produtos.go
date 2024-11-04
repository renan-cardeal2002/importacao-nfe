package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"importa-nfe/internal/core/domain"
	"importa-nfe/internal/core/ports"
)

type produtosRepository struct {
	db *sql.DB
}

func NewProdutosRepository(db *sql.DB) ports.ProdutosRepository {
	return produtosRepository{
		db: db,
	}
}

func (r produtosRepository) FindProdutosByEmpresaID(ctx context.Context, EmpresaID int) ([]domain.Produto, error) {
	query := `
    	SELECT id_produto, 
			   id_empresa, 
			   c_prod, 
               c_ean, 
               x_prod, 
               u_com, 
               q_com, 
               v_un_com, 
               v_prod, 
               v_custo, 
               v_preco, 
               v_margem, 
               v_adicional
          FROM produtos
         WHERE id_empresa = ?`

	var produtos []domain.Produto
	rows, err := r.db.QueryContext(ctx, query, EmpresaID)
	if err != nil {
		return []domain.Produto{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var produto domain.Produto
		err = rows.Scan(
			&produto.ID,
			&produto.EmpresaID,
			&produto.CodProduto,
			&produto.EAN,
			&produto.DescrProduto,
			&produto.UnidadeMedida,
			&produto.Quantidade,
			&produto.PrecoUnitario,
			&produto.PrecoBruto,
			&produto.Custo,
			&produto.Preco,
			&produto.Margem,
			&produto.ValorAdicional,
		)
		if err != nil {
			return []domain.Produto{}, err
		}
		produtos = append(produtos, produto)
	}

	return produtos, nil
}

func (r produtosRepository) InserirProdutos(ctx context.Context, produtosJSON string, EmpresaID int) error {
	var produtos []map[string]interface{}
	if err := json.Unmarshal([]byte(produtosJSON), &produtos); err != nil {
		return err
	}

	insertStatement := `
		INSERT INTO produtos 
		    (id_empresa, c_prod, c_ean, x_prod, u_com, q_com, v_un_com, v_prod, v_custo, v_preco, v_margem, v_adicional) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	for _, produto := range produtos {
		prod, ok := produto["Prod"].(map[string]interface{})
		if !ok {
			return nil
		}

		cEAN := prod["CEAN"].(string)

		query := "SELECT count(1) as count FROM produtos WHERE c_ean = ?"
		var countEan int

		err := r.db.QueryRowContext(ctx, query, cEAN).Scan(&countEan)
		if err != nil {
			return err
		}

		if countEan > 0 {
			return errors.New("EAN já existente")
		}

		cProd := prod["CProd"].(string)
		xProd := prod["XProd"].(string)
		uCom := prod["UCom"].(string)
		qCom := prod["QCom"].(string)
		vUnCom := prod["VUnCom"].(string)

		vProd := prod["VProd"].(float64)
		vFrete := prod["VFrete"].(float64)
		vSeg := prod["VSeg"].(float64)
		vDesc := prod["VDesc"].(float64)
		vOutro := prod["VOutro"].(float64)

		vCusto := vProd + vFrete + vSeg + vOutro + vDesc

		vMargem := 0.0 // VMargem não está presente no xml

		vPreco := vCusto + vMargem
		vAdicional := 0.0 // VMargem não está presente no xml

		_, err = r.db.ExecContext(ctx, insertStatement, EmpresaID, cProd, cEAN, xProd, uCom, qCom, vUnCom, vProd, vCusto, vPreco, vMargem, vAdicional)
		if err != nil {
			return err
		}
	}

	return nil
}
