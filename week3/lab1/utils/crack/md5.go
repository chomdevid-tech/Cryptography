
package crack

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
)

func CrackMD5(target string) {
	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		password := scanner.Text()

		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))

		fmt.Println("Trying:", password)

		if hash == target {
			fmt.Println("Password Found:", password)
			return
		}
	}

	fmt.Println("Password not found!")
}
