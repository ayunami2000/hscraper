package seen

import (
	"bufio"
	"os"
	"sync"
)

var file *os.File
var seen = make(map[string]struct{})
var mutex = sync.RWMutex{}

// TODO: Make this a module?

func init() {
	var err error

	file, err = os.OpenFile("./seen.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		seen[scanner.Text()] = struct{}{}
	}
}

func Seen(str string) bool {
	mutex.RLock()
	_, exists := seen[str]
	mutex.RUnlock()
	return exists
}

func Add(str string) {
	mutex.Lock()
	seen[str] = struct{}{}
	_, err := file.WriteString(str + "\n")
	mutex.Unlock()

	if err != nil {
		panic(err)
	}
}
