package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	URlJson "lesson11/GoSearch/pkg/netsrv"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Error connecting to the server:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type 'exit' to quit.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a search query: ")
		scanner.Scan()
		query := scanner.Text()

		if query == "exit" {
			break
		}

		_, err = conn.Write([]byte(query + "\n"))
		if err != nil {
			log.Println("Error sending query:", err)
			continue
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("Error receiving response:", err)
			continue
		}

		var results []URlJson.URLWithTitle
		err = json.Unmarshal([]byte(response), &results)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		fmt.Println("Search results:")
		for _, doc := range results {
			fmt.Println("URL:", doc.URL)
			fmt.Println("Title:", doc.Title)
			fmt.Println()
		}
	}
}
