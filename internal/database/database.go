package database

import (
    "database/sql"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

// CreateDB create and return SQLite DB
func CreateDB(db *sql.DB, dbPath string) (*sql.DB, error) {
    file, err := os.Create(dbPath)
    if err != nil {
        return nil, err
    }
    file.Close()
    return sql.Open("sqlite3", dbPath)
}

func RunQuery(db *sql.DB, query string) (int64, error) {
    tx, _ := db.Begin()
    statement, err := tx.Prepare(query)
    if err != nil {
        return 0, err
    }
    defer statement.Close()
    res, err := statement.Exec()
    if err != nil {
        tx.Rollback()
        return 0, err
    }
    defer tx.Commit()
    return res.RowsAffected()
}

func GetData(db *sql.DB, query string) ([]string, [][]string, error) {
    row, err := db.Query(query)
    if err != nil {
        return []string{}, [][]string{}, err
    }
    defer row.Close()

    columns, _ := row.Columns()

    output := make([][]string, 0)
    rawResult := make([][]byte, len(columns))
    dest := make([]interface{}, len(columns))

    for i := range rawResult {
        dest[i] = &rawResult[i]
    }

    for row.Next() {
        row.Scan(dest...)
        res := make([]string, 0)
        for _, raw := range rawResult {
            if raw != nil {
                res = append(res, string(raw))
            }
        }
        output = append(output, res)
    }

    return columns, output, nil
}
