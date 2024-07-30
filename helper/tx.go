package helper

import (
	"database/sql"
	"golang_listrik/model/domain"
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

func AfterInsert(q *sql.Tx, id int) {
	temp := domain.Tagihan{Penggunaan: id}
	var awal int
	var akhir int
	sql := "SELECT pelanggan, meter_awal, meter_ahkir from public.penggunaan WHERE id_penggunaan = $1"
	q.QueryRow(sql, id).Scan(&temp.Pelanggan, &awal, &akhir)

	temp.Total = awal - akhir
	sql2 := "INSERT INTO public.tagihan (pelanggan, penggunaan, jumlah_meter) VALUES ($1, $2, $3)"
	_, err := q.Exec(sql2, temp.Pelanggan, temp.Penggunaan, temp.Total)
	Err(err)
}
