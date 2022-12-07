package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	xlsx, err := excelize.OpenFile("./test.xlsx")

	if err != nil {
		panic(err)
	}

	// sheet name
	activeSheet := xlsx.GetActiveSheetIndex()
	activeSheetName := xlsx.GetSheetName(activeSheet)

	// get value (시트명, 셀위치)
	log.Println(xlsx.GetCellValue("activeSheetName", "A2"))

	// get rows (시트명으로)
	rows := xlsx.GetRows("activeSheetName")

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
