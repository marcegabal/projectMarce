package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "marce5891"
	dbname   = "bd_pruebas"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	fmt.Println("conecto")

	insertDb := `Insert into "personas" ("id", "nombre") values($1, $2)`
	_, e := db.Exec(insertDb, 7, "Go")
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Hay errores")
		panic(err)
	}
}

/*import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/marcegabal/projectMarce/db"
)


func readPropertiesFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	properties := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			properties[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return properties, nil
}

func main() {
	filePath := "./src/config.properties"
	properties, err := readPropertiesFile(filePath)
	if err != nil {
		fmt.Println("Error reading properties file:", err)
		return
	}

	// Print the read properties
	fmt.Println("Properties:")
	/*for key, value := range properties {
	  fmt.Printf("%s = %s\n", key, value)
	}

	fmt.Println("===============================")

	fmt.Println(properties["PASSWORD_BD"])

	db.MainDb()
}*/
