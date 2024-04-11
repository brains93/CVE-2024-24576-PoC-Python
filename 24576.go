package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Enter payload here:")
	reader := bufio.NewReader(os.Stdin)
	inputPayload, _ := reader.ReadString('\n')
	inputPayload = strings.TrimSpace(inputPayload)

	cmd := exec.Command("cmd", "/C", "C:\\Users\\Hacker\\source\\repos\\CVE-2024-24567-PoC-Python\\test.bat", inputPayload)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Output:\n", string(stdoutStderr))
}
