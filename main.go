package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/LindsayBradford/go-dbf/godbf"
	"github.com/simonvetter/modbus"
)

const dbFile = "/workspaces/go-tshp/test.dbf"
const dbEncode = "UTF-8"

func main() {
	fmt.Println("Welcome home!")

	if _, err := os.Stat(dbFile); errors.Is(err, os.ErrNotExist) {
		newTable := godbf.New(dbEncode);

		newTable.AddTextField("ADR", 4)
		newTable.AddTextField("WBBFORE", 4)
		newTable.AddTextField("WBAFTER", 4)
		newTable.AddTextField("DATE", 10)
		newTable.AddTextField("ACT", 4)

		godbf.SaveToFile(newTable, dbFile);
	}

	// workTable, err := godbf.NewFromFile(dbFile, dbEncode);
	// if err != nil {
	//  	log.Fatal(err)
	// }
	// fmt.Printf("Nuber of records: %d \n", workTable.NumberOfRecords())

	// workTable.AddNumberField("ADR", 4, 0)
	// workTable.AddNumberField("WBBFORE", 4, 0)
	// workTable.AddNumberField("WBAFTER", 4, 0)
	// workTable.AddDateField("DATE")
	// workTable.AddNumberField("ACT", 4, 0)


	// idx, err := workTable.AddNewRecord()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("New record: %d \n", idx)
	// workTable.SetFieldValueByName(idx, "ADR", "0")
	// workTable.SetFieldValueByName(idx, "ACT", "1")

	// godbf.SaveToFile(workTable, dbFile);

	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      "rtu://COM7",
        Speed:    9600,                   // default
        DataBits: 8,                       // default, optional
        Parity:   modbus.PARITY_NONE,      // default, optional
        StopBits: 1,                       // default if no parity, optional
        Timeout:  300 * time.Millisecond,
    })
	if err != nil {
		log.Fatal(err)
	}
	err = client.Open()
    if err != nil {
		log.Fatal(err)
	}
}
