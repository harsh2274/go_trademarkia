package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main2() {

	database, err := sql.Open("postgres", "postgres://root:855fc1@localhost/mydb?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer database.Close()

	_, err = database.Exec("CREATE TABLE Infor IF NOT EXISTS  (id SERIAL PRIMARY KEY, xml_data XML)")
	if err != nil {
		fmt.Println(err)
	}

	// to insert in multiple data
	stmt, err := database.Prepare("INSERT INTO Infor (xml_data) VALUES($1)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	xmlData := "<prosecution-entry><identifier>1</identifier><code>158</code><type-code>X</type-code><date>20040725</date><history-text>APPEAL TO BOARD</history-text></prosecution-entry>"

	_, err = stmt.Exec(xmlData)
	if err != nil {
		panic(err)
	}

	xmlData := os.Open("C:/Users/Asus/Desktop/ML_Project/tt230101.xml")
	_, err = stmt.Exec(xmlData)
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := database.Prepare("SELECT xml_data FROM Infor WHERE id=$1")
	if err != nil {
		// handle error
	}
	defer stmt.Close()

	fmt.Println("XML data inserted successfully.")
}
