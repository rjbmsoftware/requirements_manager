package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	client, err := sql.Open("sqlite3", "../../requirements.db")
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	f, err := os.Open("test_data_files/requirements.csv")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	insertReqSql := "INSERT INTO requirements(id, title, path, description) VALUES (?, ?, ?, ?);"
	for _, row := range rows {
		anyRow := make([]any, len(row))
		for i, v := range row {
			anyRow[i] = v
		}

		_, err := client.Exec(insertReqSql, anyRow...)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
