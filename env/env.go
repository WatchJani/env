package env

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Loader(file *os.File) (map[string]string, error) {
	res := make(map[string]string)

	reader := bufio.NewScanner(file)

	err := reader.Err()
	if err != nil {
		log.Println(err)
	}

	for reader.Scan() {
		env := strings.Split(reader.Text(), "=")
		res[env[0]] = env[1][1 : len(env[1])-1]
	}

	return res, nil
}

func DefaultName(name []string) []string {
	if len(name) == 0 {
		return []string{".env"}
	}

	return name
}

func Load(filename ...string) (envMap map[string]string, err error) {
	filename = DefaultName(filename)
	return readFile(filename[0])
}

func readFile(filename string) (envMap map[string]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	return Loader(file)
}
