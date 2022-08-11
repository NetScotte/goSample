package myosrw

// 文件读写示例
// 可以使用bufio对文件对象进行封装
// 可以使用ioutil的Readfile()和WriteFile()读写文件
import (
	"fmt"
	"io"
	"log"
	"os"
)

// 读取文件内容
func Sample_readfile(filename string) []byte {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	counts := 0
	for {
		temp := make([]byte, 100)
		// 当temp为slice时，其内容为空
		// 如果文件中剩余的数据少于byte的长度，那么只会读取剩下的内容，即c < 100, 发生io.EOF错误时，没有读取任何内容
		c, err := file.Read(temp)
		if err == io.EOF {
			break
		}
		counts += c
		// 合并slice
		data = append(data, temp[:c]...)
	}
	fmt.Printf("read counts %d, content: \n%s\n", counts, data)
	return data
}

// 向文件中写入数据
func Sample_writefile(filecontent []byte, filename string) (count int, err error) {
	// os.O_TRUNC 清空  os.O_APPEND 追加
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	if err != nil {
		log.Fatal("open error: ", err)
		return 0, err
	}
	// Write(), WriteAt() WriteString()
	count, err = file.Write(filecontent)
	defer file.Close()
	if err != nil {
		log.Fatal("write error:", err)
		return 0, err
	}
	fmt.Printf("write %d contents to file\n", count)
	file.Close()
	return count, nil
}

// 文件系统查看
func Sample_filesystem() {
	// 切换到指定目录下
	fmt.Println("(os.Chdir) change dir")
	err := os.Chdir("../notitle/")
	if err != nil {
		log.Fatal(err)
	}
	// 获取当前目录
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("(os.Getwd) current dir is: %s(os.Getwd)\n", pwd)

	// 显示当前目录的所有文件
	file, err := os.Open("../notitle")
	if err != nil {
		log.Fatal(err)
	}
	// 获取当前目录的所有文件名，使用该方法后file.Readdir返回空
	// files, err := file.Readdirnames(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("(file.Readdirnames) files list: %v\n", files)

	// 获取当前目录第一个文件的信息
	fileinfo, err := file.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}
	// 当fileinfo没有值时，为empty slice, 而不是nil slice, 不能使用fileinfo == nil 进行判断
	if len(fileinfo) != 0 {
		fmt.Printf("(file.Readdir) first_file_info: %v: \n%+v\n", fileinfo[0].Name(), fileinfo[0])
	}
}
