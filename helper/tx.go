package helper

import (
	"context"
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

func AfterInsert(ctx context.Context, q *sql.Tx, total int, id int, id_penggunaan int) {
	sql2 := "INSERT INTO public.tagihan (pelanggan, penggunaan, jumlah_meter) VALUES ($1, $2, $3)"
	_, err := q.QueryContext(ctx, sql2, id, id_penggunaan, total)
	q.Commit()
	Err(err)
}

func UploadImageSql(ctx context.Context, q *sql.Tx, Name_image string, id int) {
	var (
		total        int
		tarif        int
		pelanggan_id int
	)

	sql := "UPDATE public.tagihan SET image = $1, status = 2 WHERE id_tagihan = $2 RETURNING pelanggan"
	_ = q.QueryRowContext(ctx, sql, Name_image, id).Scan(&pelanggan_id)

	sql2 := "SELECT t.tarifperkwh, ta.jumlah_meter FROM public.tagihan AS ta JOIN public.pelanggan AS pg ON ta.pelanggan = pg.id_pelanggan JOIN public.tarif AS t ON pg.tarif = t.id_tarif WHERE ta.id_tagihan = $1"
	_ = q.QueryRowContext(ctx, sql2, id).Scan(&tarif, &total)

	result := total * tarif

	sql_pembayaran := "INSERT INTO public.pembayaran (tagihan, pelanggan, total_bayar) VALUES ($1, $2, $3)"
	q.ExecContext(ctx, sql_pembayaran, id, pelanggan_id, result)
}
