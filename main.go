package main

import (
	"fmt"
	"os"
	rm "remove_Go/mainFunc"
)

func main() {
	file := os.Args
	command(&file)
}

// 判断命令行参数
func command(par *[]string)  {
	pars := *par
	switch pars[1] {
	// -r 选项直接删除
	case "-r":
		file := pars[2:]
		rm.Remover(&file)
	// -m 选项直接回收
	case "-m":
		file := pars[2:]
		rm.Recycle(&file)
	case "-h":
		fmt.Print(`
Usage:  grm [OPTIONS] FileOrPath

Options:
	-r		Delete files directly
	-m		Move files to the recycle bin
Config:
	Path: Configuration file path /etc/grm/conf.yml;
	Size: Files larger than size will be deleted and will not be reclaimed;

`)
	// TODO: 清空回收站
	//case "-u":
	//	path, _ := ioutil.ReadFile("./tmp")
	//	rm.Remover(&path)
	// 使用接口实现

	// TODO: 撤销删除
	default:
		file := pars[1:]
		rm.MainFunc(&file)
	}
}