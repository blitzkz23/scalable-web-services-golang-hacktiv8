package main

import (
	"database/sql"
	"fmt"

	// Hal ini dilakukan karena sebetulnya kita tidakakan menggunakan syntax apapun dari driver Postgresql. Yangkita perlu lakukan adalah hanya meregristrasikan nya dengancara mengimportnya saja agar package databases/sqlmengetahui tentang driver jenis apa yang kita pakai.
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "naufalaldy23"
	dbname   = "hactivgosql"
)

var (
	// variabel db berisi pointer dari sql.DB
	db  *sql.DB
	err error
)

// Struct untuk menampung data employee yang akan dimasukkan ke database
type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// sql open berasal dari package database/sql, fungsi ini tidak membangun koneksi ke database, melainkan hanya memvalidasi argumen2 yang diberikan.
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// fungsi Ping untuk membangun koneksi ke database sekaligus memeriksa jika string panjang berupa info yang kitaberikan pada function Open merupakan info yang 100% valid.

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	// CreateEmployee()
	UpdateEmployee()
	GetEmployee()
	DeleteEmployee()
	GetEmployee()
}

// * Function membuat data untuk diinsert
func CreateEmployee() {
	var employee = Employee{}

	// Tanda $1, $2 dst.. merupakanrepresentasi dari nilai-nilai yang akan kita masukkan nantinya.Lalu terdapat penulisan Returning * yang memiliki arti bahwakita ingin mendapat nilai kembalian dari seluruh field yang berasal dari data yang baru dibuat ke client.
	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *`

	// ! Kita menggunakan method QueryRow.Method in digunakan untuk mengeksekusi sebuah query sqltergantung dari statement sql yang kita berikan.  Karena statement sql yang kita berikan bertujuan unutk create data,maka method QueryRow ini akan berfungsi untuk membuatdata baru dengan nilai yang kita berikan pada parameter keduadari method QueryRow.  Kemudian method QueryRow kita chaning denganmethod Scan agar kita dapat mendapatkan nilai-nilaibalikan dari statement yang telah kita buat. Karenastatement kita menggunakan Returning *
	err = db.QueryRow(sqlStatement, "Airell Jordsan", "airell@gmail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data: %+v\n", employee)
}

// * Function get data
func GetEmployee() {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees`

	// Menggunakan query ke db berdasarkan sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	// Dikarenakan method ini dapat mengembalikan satu atau lebih rows dari suatu table pada database maka haru ditutup
	defer rows.Close()

	// Method rows.Next akan menghasilkannilai true selama data nya masih ada, namun jikasudah tidak ada maka dia akan me-return falsedan proses looping akan terhenti
	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employee datas:", results)
}

// * Function update data
func UpdateEmployee() {
	sqlStatement := `
		UPDATE employees
		SET full_name = $2, email = $3, division = $4, age = $5
		WHERE id = $1;
	`

	// Dengan menggunakan exec tidak akan mengembalikan data
	res, err := db.Exec(sqlStatement, 4, "Naufal Aldy Pradana", "naufal@gmail.com", "Superman", 21)

	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated data amount:", count)
}

// * Function delete data
func DeleteEmployee() {
	sqlStatement := `
	DELETE FROM employees
	WHERE id = $1;
	`

	// Argumen $id
	res, err := db.Exec(sqlStatement, 4)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted data amount:", count)
}
