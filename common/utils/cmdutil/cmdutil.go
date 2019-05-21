package cmdutil

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func CmdResult(cmd string, args []string) string {
	run := exec.Command(cmd, args...)
	out, err := run.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(out)
}

func CmdResults(cmd string, args []string) (result [][]string) {
	run := exec.Command(cmd, args...)
	out, err := run.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	_ = run.Start()
	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)
	re := regexp.MustCompile(`\s+`)
	for scanner.Scan() {
		m := scanner.Text()
		line := fmt.Sprintln(m)
		lineArr := re.Split(strings.TrimSpace(line), -1)
		result = append(result, lineArr)
	}
	return
}
