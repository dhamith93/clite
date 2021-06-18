package display

import (
    "os"
    
    "github.com/olekukonko/tablewriter"
)

func PrintTable(columns []string, data [][]string) {
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader(columns)
    table.SetAutoFormatHeaders(false)
    for _, v := range data {
        table.Append(v)
    }
    table.Render()
}
