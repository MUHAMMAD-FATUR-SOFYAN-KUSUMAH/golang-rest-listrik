package helper

import (
	"context"
	"database/sql"
	"fmt"
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
	fmt.Println("sdkfs")
	fmt.Println(id, id_penggunaan)
	_, err := q.QueryContext(ctx, sql2, id, id_penggunaan, total)
	q.Commit()
	fmt.Println(err)
}

func UploadImageSql(ctx context.Context, q *sql.Tx, Name_image string, id int) {

	sql := "UPDATE public.tagihan SET image = $1, status = 2 WHERE id_tagihan = $2"

	_, err := q.ExecContext(ctx, sql, Name_image, id)

	Err(err)
}
