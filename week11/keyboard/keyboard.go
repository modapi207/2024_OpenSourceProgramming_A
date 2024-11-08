package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetInt() (int, error) {
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}
	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}
	return n, nil
}
