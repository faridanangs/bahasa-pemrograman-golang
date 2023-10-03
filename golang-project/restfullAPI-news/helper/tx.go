package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		err := tx.Rollback()
		IfLogingErr(err, "error terjadi di helper.tx roolback")
		panic(err)
	} else {
		err := tx.Commit()
		IfLogingErr(err, "error terjadi di helper.tx commit")
	}
}
