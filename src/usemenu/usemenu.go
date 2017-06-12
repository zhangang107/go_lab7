package main

import "fmt"
import (. "menureuse")
import "os/exec"
import "os"

func quit(agrc int, agrv [CMD_MAX_ARGV_NUM]string) int {
	os.Exit(1)
	return 0
}
func add(agrc int, agrv [CMD_MAX_ARGV_NUM]string) int {
	var x int
	var y int
	fmt.Println("pls input two int numbers")
	fmt.Scanf("%d %d",&x,&y)
	ret := x+y
	fmt.Printf("The result :%d\n",ret)
	return 0
}

func sub(agrc int, agrv [CMD_MAX_ARGV_NUM]string) int {
	var x, y int
	fmt.Println("pls input two int numbers")
	fmt.Scanf("%d %d",&x,&y)
	ret := x-y
	fmt.Printf("The result :%d\n",ret)
	return 0
}

func newfile(agrc int, agrv [CMD_MAX_ARGV_NUM]string) int {
	var cmdstr, filename string
	cmdstr = "touch"
	fmt.Println("Pls input filename")
	fmt.Scanf("%s",&filename)
	cmd :=exec.Command(cmdstr,filename)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if string(out) != "" {
		fmt.Println(string(out))
	}
	fmt.Printf("%s success!",cmdstr)
	return 0
}

func del(agrc int, agrv [CMD_MAX_ARGV_NUM]string) int {
	var cmdstr, filename string
	cmdstr = "rm"
	fmt.Println("Pls input filename")
	fmt.Scanf("%s",&filename)
	cmd :=exec.Command(cmdstr,filename)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if string(out) != "" {
		fmt.Println(string(out))
	}
	fmt.Printf("%s success!",cmdstr)
	return 0
}

func pwd(agrc int, agrv [CMD_MAX_ARGV_NUM]string) int {
	var agr string
	for i :=1; i < agrc; i++ {
		switch agrv[i] {
			case "-L":
				fmt.Println("输出连接路劲")
				agr = "-L"
			case "-P":
				fmt.Println("输出物理路劲")
				agr = "-l"
			default:
				fmt.Println("不存在的选项...将以默认方式启动命令")
				agr = ""
		}
	}
	cmd :=exec.Command("pwd", agr)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if string(out) != "" {
		fmt.Println(string(out))
	}
	return 0
}

func ls(agrc int, agrv [CMD_MAX_ARGV_NUM]string) int {
	var agr string
	for i :=1; i < agrc; i++ {
		switch agrv[i] {
			case "-a":
				fmt.Println("列出文件下所有的文件")
				agr = "-a"
			case "-l":
				fmt.Println("列出文件的详细信息")
				agr = "-l"
			case "-s":
				fmt.Println("打印出文件大小")
				agr = "-s"
			default:
				fmt.Println("该选项暂时没有...将以默认方式启动命令")
				agr = ""
		}
	}
	cmd :=exec.Command("ls", agr)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if string(out) != "" {
		fmt.Println(string(out))
	}
	return 0
}

func main() {
	MenuConfig("ls", "Show Files With opt", ls)
	MenuConfig("add","Add Two Numbers", add)
	MenuConfig("sub","Minus Two Numbers", sub)
	MenuConfig("new","New File", newfile)
	MenuConfig("del","Delete File", del)
	MenuConfig("pwd","Show The path", pwd)
	MenuConfig("quit","Quit The Menu!", quit)
	ExcuteMenu()
}