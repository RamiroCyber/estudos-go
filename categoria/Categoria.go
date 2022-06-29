package categoria

import "github.com/estudodb/db"

type Categoria struct {
	ID   int64
	Nome string
}

func Insert(c Categoria) (id int64, err error) {
	conn, err := db.OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO categorias (nome) VALUES ($1) RETURNING id`

	err = conn.QueryRow(sql, c.Nome).Scan(&id)

	return id, err
}
