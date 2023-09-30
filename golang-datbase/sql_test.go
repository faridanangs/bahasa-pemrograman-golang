package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExectSql(t *testing.T) {
	db := GetConnection()

	ctx := context.Background()
	for i := 100; i < 150; i++ {
		id := strconv.Itoa(i)
		script := "insert into customer(id, name) values('data ke " + id + "', 'anangi')"

		// jika kita menggunakan ExecContext ini tidak mengembalikan result
		_, err := db.ExecContext(ctx, script)
		if err != nil {
			panic(err)
		}

	}
	fmt.Println("data berhasil di masukan ke dalam database")

}

func TestQuerySql(t *testing.T) {
	db := GetConnection()

	ctx := context.Background()
	script := "select id,name from customer"

	// jika kita menggunakan QueryContext ini mengembalikan result
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	// jika kita menggunakan rows kita harus menclosenya supaya tidak terjadi panic/tras
	defer rows.Close()

	// kita gunakan next untuk mengambil data sesudahnya
	for rows.Next() {
		var id, name string

		// kita gunakan scan untuk mengambil data dan memasukannya ke dalam id dan kita gunakan error untuk
		// menangkap jika terjadi error pada scan
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id", id)
		fmt.Println("Name", name)

	}

}

func TestSqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	cntx := context.Background()
	script := "select id, name, email, balance, rating, create_at, birth_date, married from customer"
	rows, err := db.QueryContext(cntx, script)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		// jika kita memiliki nilai null di dalam database kemudian kita panggil di golang
		// maka akan terjadi error pada itersi tersebut cara mengatasinya kita bisa gunakan
		// sql.Null+tipedatanya apa contoh var email sql.NullString
		var id, name string
		var email sql.NullString
		var balance int32
		var rate float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rate, &createdAt, &birthDate, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("=================")
		fmt.Println("id:", id)
		fmt.Println("name:", name)
		fmt.Println("email:", email)
		fmt.Println("balance:", balance)
		fmt.Println("rate:", rate)
		if birthDate.Valid {
			fmt.Println("birth date", birthDate.Time)
		}
		fmt.Println("created at:", createdAt)
		fmt.Println("married:", married)
		fmt.Println("=================")

	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	// jika kita menggunakan kalimat seperti ini "admin' ; #" maka kita memanipulasi database atau sql injection
	// sehinggka setelah kita memasukan username tanpa password/password yang salah maka kita bisa masuk ke dalam
	// database karna kalimat setelahnya di anggap komen karna kita
	// mwnggunakan # dan karna kita menggunakan ' di akhir admin maka dia akan di anggap satu kalimat
	username := "admin' ; #"
	password := "kukik"

	cntx := context.Background()
	script := "select username from user where username = '" + username + "' and password = '" + password + "' limit 1"

	rows, err := db.QueryContext(cntx, script)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("Login sukses", username)

	} else {
		fmt.Println("Login gagal")
	}
}

func TestSqlInjectionSelf(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "farid"
	password := "farid"

	cntx := context.Background()

	// untuk mengatasi atau menangani sql injection kita cukup gunakan tanda ? di dalam script
	// contoh "select username from user where username = ? and password = ? limit 1" di sini kita masukan  2 ?
	// dan cara manggilnya kita masukan data di dalam db.QueryContext(cntx, script, username, password) sesuai jumlah ? yg ada di dalam script
	script := "select username from user where username = ? and password = ? limit 1"

	rows, err := db.QueryContext(cntx, script, username, password)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("Login sukses", username)

	} else {
		fmt.Println("Login gagal")
	}
}

func TestExectSqlSelf(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	script := "insert into user(username, password) values(?, ?)"

	// jika kita menggunakan ExecContext ini tidak mengembalikan result
	_, err := db.ExecContext(ctx, script, "farid", "farid")
	if err != nil {
		panic(err)
	}
	fmt.Println("data berhasil di masukan ke dalam database")

}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	cntx := context.Background()
	script := "insert into comments(email, comment) values(?,?)"

	email := "faridanangs@gmail.com"
	comment := "helllo ini test koment"
	result, err := db.ExecContext(cntx, script, email, comment)
	if err != nil {
		panic(err)
	}

	// untuk mengambil id yang ada di dalam database yang terakhir kita busa menggunakan
	// result.LastInsertId() yang mengembalikan 2 nilai data dan err
	incId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Data dengan index ke", incId)
}

func TestPrepareStmn(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	cntx := context.Background()
	script := "insert into comments(email, comment) values(?,?)"

	// di sini kita tidak perlu memasukan nilai ke dalam PrepareContext
	// PrepareContext ini kita gunakan untuk memanggil atau memasukan data di database secara terus
	// menerus dan hanya memanggilnya sekali saja
	statement, err := db.PrepareContext(cntx, script)
	defer statement.Close()

	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "farid " + strconv.Itoa(i) + "@gmail.com"
		comment := "Komen ke " + strconv.Itoa(i)

		// nah baru di sini kita maskan nilai yang ingin kita masukan ke dalam database melalui script
		// disini kita memasukan email dan comment dan yang paling penting adalah memasukan
		// contextnya
		result, err := statement.ExecContext(cntx, email, comment)
		if err != nil {
			panic(err)
		}

		incrId, _ := result.LastInsertId()

		fmt.Println("Index data ke", incrId)

	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	// kita gunakan begin untuk melakukan transaction ke dalam database
	// jika kita tdk menggunakan begin maka proses yang di lakuakn ke dalam
	// data base secara auto commit jika kita menggunakan begin kita bisa tdk mengirim data
	// ke dalam database walaupun prosesnya benar dan lancar
	Tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	cntx := context.Background()
	script := "insert into comments(email, comment) values(?,?)"

	statement, err := Tx.PrepareContext(cntx, script)
	defer statement.Close()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "farid " + strconv.Itoa(i) + "@gmail.com"
		comment := "Komen ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(cntx, email, comment)
		if err != nil {
			panic(err)
		}

		IncId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Data dengan index ke", IncId)
	}

	// kita bisa menggunakan commit untuk memasukan data ke dalam database
	// Tx.Commit()

	// kita bisa menggunakan roolback untuk tidak memasukan data ke dalam database
	// walaupun proses yang di lakukan sebelumnya berhasil
	Tx.Rollback()

}
