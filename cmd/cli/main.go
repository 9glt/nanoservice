package main

import (
	"dbsqlite/pkg/records"
	"flag"
	"log"
	"os"

	// Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details

	"gorm.io/gorm"
)

var (
	db *gorm.DB

	flagAction   = flag.String("action", "", "Command to run")
	flagCategory = flag.String("category", "", "Category to run")
	flagDomain   = flag.String("domain", "root.toor", "Domain to run")
)

func main() {
	flag.Parse()

	if *flagAction == "" {
		log.Fatal("No action specified")
	}

	if *flagCategory == "" {
		log.Fatal("No category specified")
	}

	flagz := records.Flags{
		Action:   *flagAction,
		Category: *flagCategory,
		Domain:   *flagDomain,
		Args:     flag.Args(),
		Rdr:      os.Stdin,
	}

	records.DBCSqlite("db.db")

	switch *flagAction {
	case "create":
		record, _ := records.Create(&flagz)
		log.Printf("Created: %v", record)
	case "read":
		record, _ := records.Read(&flagz)
		log.Printf("Read: %v", record)
	case "list":
		recs, _ := records.List(&flagz)
		log.Printf("List: %v", recs)
	case "delete":
		record, _ := records.Delete(&flagz)
		log.Printf("Deleted: %v", record)
	case "update":
		record, _ := records.Update(&flagz)
		log.Printf("Updated: %v", record)

	}
}
