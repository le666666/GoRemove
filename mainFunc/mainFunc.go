package mainFunc

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type conf struct {
	Size int `yaml:"size"`
}

// MainFunc 主逻辑函数
func MainFunc(FileOrPath *[]string) {
	files := *FileOrPath
	var c conf
	size := c.readConf() // 多大的文件直接删除
	file := isRoot(files)
	if fileSize(file, int64(size)) == true {
		Remover(file)
	}else {
		Recycle(file)
	}
}

// 判断文件大小
func fileSize(FileOrPath *[]string, fileSize int64) bool {
	// 抛错，如果文件不存在会报错
	defer func() {
		if err := recover(); err != nil {
			log.Println("NotFile")
		}
	}()
	for _, value := range *FileOrPath {
		if size, _ := os.Stat(value); size.Size() >= fileSize {
		return true
	}
	}
	return false
}

// 读取配置
func (c *conf) readConf() int {
	yamlFile, err := ioutil.ReadFile("mainFunc/conf.yml")
	if err != nil {
		log.Println("Error: Configuration file conf.yml could not be found")
	}
	if err := yaml.Unmarshal(yamlFile, c); err != nil { // 读取配置
		log.Fatalf("Unmarshal: %v", err)
	}
	// fmt.Println(c.Size)
	return c.Size // 返回配置参数
}

// Remover 删除
func Remover(FileOrPath *[]string) {
	for _, value := range *FileOrPath {
		if err := os.RemoveAll(value); err != nil {
			log.Println("Error: Delete failure!")
		}
	}
}

// Recycle 回收
func Recycle(FileOrPath *[]string) {
	var tmp string
	for _, value := range *FileOrPath {
		filename := path.Base(value) // 获取文件名称
		tmp = "./tmp/" + filename // 拼接字符串
		if err := os.Rename(value, tmp); err != nil {
			log.Printf("Error: Delete failure!%v", err)
		}
	}
}

// 不允许操作根目录
func isRoot(file []string) *[]string {
	for _, root := range file {
		if root == "/" {
			log.Println("Dangerous operation: Do not delete the root directory")
		}
	}
	return &file
}

// TODO: 文件恢复