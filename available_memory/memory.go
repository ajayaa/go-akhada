package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "github.com/cloudfoundry/gosigar"

func giveScanner(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	return scanner
}

func main() {
	path := "/sys/fs/cgroup/memory/memory.limit_in_bytes"

	scanner := giveScanner(path)
	scanner.Scan()
	mem := scanner.Text()
	memInt, err := strconv.ParseUint(mem, 10, 128)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T\n", memInt)
	//memIntLong = bufio.
	fmt.Println(memInt)

	memo := sigar.Mem{}
	memo.Get()
	fmt.Println(memo.Total)
	fmt.Printf("%T\n", memo.Total)
}
