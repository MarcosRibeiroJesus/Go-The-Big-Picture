package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log that should be analyzed.")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}

/*
$ go run . -help

Usage of /tmp/go-build757751589/b001/exe/Go-The-Big-Picture:
  -level string
        Log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL (default "ERROR")
  -path string
        The path to the log that should be analyzed. (default "myapp.log")

exit status 2
*/

/*
$ go run . -level INFO

2006-08-16 12:03:10,237 - INFO - Attempting to connect to data sources

2006-08-16 12:03:11,003 - INFO - Retrying to connect to "mydatasource"

2006-08-16 12:03:12,127 - INFO - Connection to "mydatasource" succeeded after 2 attempts

2006-08-16 12:03:13,278 - INFO - MyApp started listening on port 8080

2006-08-16 12:10:10,357 - INFO - User with ID "user27" logged in


$ go run . -level WARNING

2006-08-16 12:03:10,762 - WARNING - Failed to connect to datasource "mydatasource"
*/
