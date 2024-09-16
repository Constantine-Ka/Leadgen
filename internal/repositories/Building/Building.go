package Building

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strings"
)

type DB struct {
	*sqlx.DB
}

func CreateTable(db *sqlx.DB) (DB, error) {
	var dbOut DB
	content, err := os.ReadFile("./migration/building.up.sql")
	if err != nil {
		fmt.Println("Файл migration/building.up.sql не найден")
		return dbOut, err
	}
	_, err = db.Exec(string(content))
	if err != nil {
		fmt.Println(string(content))
		log.Println(err)

		return dbOut, err
	}
	dbOut.DB = db
	return dbOut, nil
}

func (db DB) InsertOne(ctx context.Context, data Building) (int64, error) {
	var execId int64
	query := fmt.Sprintf("INSERT INTO buildings (name, city, year, level) VALUES ('%s','%s', %d , %d) RETURNING id", data.Title, data.City, data.Year, data.Level)
	err := db.GetContext(ctx, &execId, query)
	if err != nil {
		return 0, err
	}

	return execId, nil
}

func (db DB) GetBuildings(ctx context.Context, filterMap map[string]string) ([]Building, error) {
	var output []Building
	var filterList []string
	query := "SELECT id, name, city, year, level FROM buildings"
	for k, v := range filterMap {
		filterList = append(filterList, fmt.Sprintf("%s = '%s'", k, v))
	}
	if len(filterList) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(filterList, " AND "))
	}
	err := db.SelectContext(ctx, &output, query)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		fmt.Println(query)
		fmt.Println(err)
		return output, err
	}
	return output, nil
}
