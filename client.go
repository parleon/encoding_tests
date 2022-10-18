package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"

	//"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)



type process_info struct {
	ip   string
	port string
}

func genRandShni(x int) []byte {
	bas := make([]byte, x)
	rand.Read(bas)
	return bas
}

// parse_config takes the path pointing to a config file and translates it into a map indexed by process id containing process_info structs
func parse_config(path string) (map[string]process_info) {

	// initialize empty proccess map
	processes := make(map[string]process_info)

	// open config file and initialize a scanner
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// read values into process map
	for scanner.Scan() {
		splitline := strings.Split(scanner.Text(), " ")
		processes[splitline[0]] = process_info{splitline[1], splitline[2]}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return processes
}

func initialize_outgoing(info process_info) net.Conn {
	conn, err := net.Dial("tcp", info.ip+":"+info.port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func main() {

	processes := parse_config("config")
	gob_connection := initialize_outgoing(processes["gob"])
	json_connection := initialize_outgoing(processes["json"])
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		number, _ := strconv.Atoi(text[:len(text)-1])
		randshni := genRandShni(number)
		randshni2 := randshni
		go func() {
				enc := gob.NewEncoder(gob_connection)
				fmt.Println("gob")
				fmt.Println(time.Now().Format("2006-01-02T15:04:05.999999999Z07:00"))
				enc.Encode(randshni)	
		} ()

		go func() {
				enc := json.NewEncoder(json_connection)
				fmt.Println("json")
				fmt.Println(time.Now().Format("2006-01-02T15:04:05.999999999Z07:00"))
				enc.Encode(randshni2)
		} ()

	}
	
}