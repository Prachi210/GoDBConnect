package helper

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	config "github.com/Prachi210/GoDBConnect/Consumer/Config"
	model "github.com/Prachi210/GoDBConnect/Consumer/model"
	_ "github.com/lib/pq"
)

// DB set up
func setupDB(config config.ConfigModel) *sql.DB {
	dbinfo := fmt.Sprintf("host= %s port= %d dbname= %s user= %s password= %s sslmode= %s",
		config.Host, config.Port, config.Dbname, config.User, config.Password, config.Sslmode)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		log.Println("Err", err.Error())
		panic(err)
	}
	return db
}

func WriteProductsIntoDB(products []model.ConsumerProduct, config config.ConfigModel) int {
	db := setupDB(config)
	defer db.Close()
	log.Println("\n Setting Up Database")
	inserts := make([]string, 0, len(products))
	for _, prod := range products {
		inserts = append(inserts, "(", strconv.Itoa(prod.ID), ",'", prod.Model, "','", prod.ItemType, "','", prod.Category, "',", prod.Price, ",current_timestamp),")

	}
	insertString := fmt.Sprint(inserts)
	insertString = strings.Replace(insertString, "[", "", 1)
	insertString = strings.Replace(insertString, ",]", "", 1)
	insertString = strings.TrimSpace(insertString)
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	stmt := fmt.Sprintf("INSERT INTO productInfoDuplicate(id, model, itemType, category, price, date) VALUES %s;", insertString)
	log.Println(" The Query is:", stmt)
	res, err := db.Exec(stmt)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return 0
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return 0
	}
	log.Printf("%d products created simulatneously", rows)

	if rows == int64(len(products)) {
		tx.Commit()
	} else {
		tx.Rollback()
		rows = 0
	}
	return int(rows)
}
