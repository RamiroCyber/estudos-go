package produto

import "github.com/estudodb/db"

type Produto struct {
	ID        int64
	Nome      string
	preco     float32
	descricao string
}

func Insert(c *Produto) (id int64, err error) {
	conn, err := db.OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO produtos (nome, preco, descricao) VALUES  // ($1) RETURNING id`

	err = conn.QueryRow(sql, c.Nome).Scan(&id)

	return id, err
}
