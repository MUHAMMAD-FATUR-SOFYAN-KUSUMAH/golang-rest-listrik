package helper

import (
	"database/sql"
)

func Tx(q *sql.Tx) {
	err := recover()
	if err != nil {
		errRolback := q.Rollback()
		Err(errRolback)
		panic(err)
	} else {
		errCommit := q.Commit()
		Err(errCommit)
	}
}

func AfterInsert(q *sql.Tx, total int, id int, id_penggunaan int) {
	sql2 := "INSERT INTO public.tagihan (pelanggan, penggunaan, jumlah_meter) VALUES ($1, $2, $3)"
	_, err := q.Exec(sql2, id, id_penggunaan, total)
	Err(err)
}

func AfterUpdate(q *sql.Tx, id int, total int) {
	sql := "UPDATE public.tagihan SET jumlah_meter = $1 WHERE id_penggunaan = $2"
	_, err := q.Exec(sql, total, id)
	Err(err)
}
