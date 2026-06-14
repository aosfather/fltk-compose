package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

type Notify func(text string)
type CmdSource struct {
	cmd     *exec.Cmd
	notify  Notify
	Workdir string
}

func (cs *CmdSource) ConfigNotify(notify Notify) {
	cs.notify = notify
}
func (cs *CmdSource) Close() {
	if cs.cmd != nil {
		cs.cmd.Process.Kill()
	}
}
func (cs *CmdSource) Open(cmd string, args ...string) {
	// args := []string{"qihuo", "-o=true"}
	// args = append(args, symal...)
	cs.cmd = exec.Command(cmd, args...)
	cs.cmd.Dir = cs.Workdir
	stdout, err := cs.cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error obtaining stdout:", err)
		return
	}
	if err := cs.cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}
	//监听output

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			if cs.notify != nil {
				cs.notify(scanner.Text())
			}
		}
	}()
}
