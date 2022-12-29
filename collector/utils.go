package collector

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func execShell(command string) string {
	cmd := exec.Command("/bin/bash", "-c", command)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error: Cannot obtain stdout pipe for command: ", err)
		return ""
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error: Command execute error: ", err)
		return ""
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("Error: ReadAll error: ", err)
		return ""
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error: Wait error: ", err)
		return ""
	}
	return string(bytes)
}
