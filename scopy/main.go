package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"errors"
)

type ServerInfo struct {
	Name       string
	User       string
	PathPrefix string
}

var default_servers = map[string]ServerInfo{
	"code1":   ServerInfo{Name: "code1", User: "steven", PathPrefix: "/data/dev/steven.xibao.com/"},
	"router1": ServerInfo{Name: "router1", User: "steven", PathPrefix: "/var/www/steven.xibao.com/"},
}

func run_linux_shell_cmd(cmd string) (string, error) {
	//fmt.Println("Run shell cmd: ", cmd)
	f, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func run_win_shell_cmd(cmd string) (string, error) {
	//fmt.Println("Run shell cmd: ", cmd)
	f, err := exec.Command("cmd", "/C", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func get_path_from_filename(filename string) string {
	pos := strings.LastIndex(filename, string(os.PathSeparator))
	if pos < 0 {
		pos = strings.LastIndex(filename, "/")
	}
	if pos < 1 {
		return ""
	} else {
		return filename[0:pos]
	}
}

func parseFileAndFlag(s string) (filename, flag string, err error) {
	pos := strings.LastIndex(s, "@")
	if pos < 1 {
		err = errors.New("Invalid filename")
		return
	}

	filename = s[0:pos]
	flag = s[pos + 1:]
	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: scopy.exe <xxxx/yyy>@<server>")
		return
	}
	filename, serverFlag, err := parseFileAndFlag(os.Args[1])
	server, ok := default_servers[serverFlag]
	if !ok {
		fmt.Println("invalid remote server.")
		return
	}
	cmd := fmt.Sprintf(`scp %s %s@%s:%s%s/`, filename, server.User, server.Name, server.PathPrefix, get_path_from_filename(filename))
	ret, err := run_win_shell_cmd(cmd)
	if err != nil {
		fmt.Printf("ERR: %v\n", err.Error())
		return
	}
	fmt.Println(ret)
}
