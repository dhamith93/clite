package main

import (
    "bufio"
    "database/sql"
    "fmt"
    "os"
    "strings"

    "github.com/dhamith93/clite/internal/database"
    "github.com/dhamith93/clite/internal/display"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    var (
        db  *sql.DB
        err error
    )
    dbOpened := false
    for {
        fmt.Print("> ")
        cmd, _ := reader.ReadString('\n')
        cmd = strings.ToUpper(strings.TrimRight(cmd, "\n"))

        if cmd == "EXIT" {
            fmt.Println("Good bye")
            break
        }

        cmdArr := strings.Fields(cmd)

        if len(cmdArr) > 1 {

            if cmdArr[0] == "OPEN" {
                db, err = database.CreateDB(db, cmdArr[1])
                if err != nil {
                    fmt.Printf("Error opening DB: %v\n", err.Error())
                    continue
                }

                dbOpened = true
                fmt.Printf("Database %v opened\n", cmdArr[1])
                continue
            }

            if cmdArr[0] == "SELECT" && dbOpened {
                columns, data, err := database.GetData(db, cmd)
                if err != nil {
                    fmt.Printf("Error in %v : %v\n", cmd, err.Error())
                    continue
                }

                display.PrintTable(columns, data)
                continue
            }

            affectedRows, err := database.RunQuery(db, cmd)
            if err != nil {
                fmt.Printf("Error in %v : %v\n", cmd, err.Error())
                continue
            }

            fmt.Printf("%v row(s) affected...\n", affectedRows)
        }

    }
}
