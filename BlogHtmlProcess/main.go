package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	//"path/filepath"
	"strconv"
	"strings"
)

func main(){
	//ex, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//currentPath := filepath.Dir(ex)+"/"

	// currentPath := "/Users/wontun/Desktop/ARCHIVE/HTML/blog/post/md/"
	currentPath := "./"
	fileSet := make(map[string]int64)
	meta := make(map[string]int64)

	files, _ := ioutil.ReadDir(currentPath)
	for _, f := range files {
		if strings.LastIndexAny(f.Name(), ".md") == len(f.Name())-1 {
			fileSet[f.Name()] = f.ModTime().Unix()
		}
	}

	metaFile, err := os.Open(currentPath+"meta")
	defer func() {
		err = metaFile.Close()
		if err != nil{
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(metaFile)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic("MetaFile format error!")
			}
		}
		line = strings.Trim(line, "\n")
		if strings.Compare(line, "") == 0{
			continue
		}
		filename, filedate := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		meta[filename], _ = strconv.ParseInt(filedate, 10, 64)
	}

	// 比较meta和fileSet的差异，修改meta
	for key, val := range fileSet {
		_, ok := meta[key]
		// 产生新的文件，则添加到meta中
		if !ok {
			fmt.Printf("New File Found %v, Recoding to Meta\n", key)
			meta[key] = val
		}

		// 如果日期小于fileSet：
		if meta[key] < val {
			fmt.Printf("Old File Found %v, Modifying the Meta\n", key)
			meta[key] = val
		}
	}

	// 移除Meta里不存在fileSet里面的文件
	for key, _ := range meta {
		_, ok := fileSet[key]
		if !ok {			// 不存在
			fmt.Printf("File %v does not exist but in meta entry, delete")
			delete(meta, key)
			// 删除对应的html文件
			del_filename := currentPath + strings.Split(key, ".")[0]+".html"
			_ = os.Remove(del_filename)
			fmt.Printf("File %v deleted", del_filename)
		}
	}

	// 将meta写入
	file, err := os.OpenFile(currentPath+"meta", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for key, val := range meta {
		title, date, tags, desc := getMetaFromMD(currentPath+key)
		filename := strings.TrimRight(key, ".md")
		filename = filename + ".html"

		_, err = write.WriteString(fmt.Sprintf("%s,%d,%s,%s,%s,%s,%s\n", key, val, filename,title,date,tags,desc))
		if err != nil{
			println(err)
			break
		}
	}
	// Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Meta Updated Successfully")
	return
}

func getMetaFromMD(filepath string)(string, string, string, string) {
	var title string	= ""
	var date string 	= ""
	var tags string 	= ""
	var desc string 	= ""

	file, err := os.Open(filepath)
	defer func() {
		err = file.Close()
		if err != nil{
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(file)
	buf.ReadString('\n')				// 第一个 ---
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic("MarkDown File Read error!")
			}
		}
		line = strings.Trim(line, "\n")
		if strings.Compare(line, "") == 0{
			continue
		}
		if strings.Contains(line,"---") {			// 出口
			break
		}
		if len(strings.Split(line,":")) != 2 {
			continue
		}

		key, value := strings.Split(line, ":")[0], strings.Split(line, ":")[1]
		key = strings.Trim(key," ")
		value = strings.Trim(value, " ")

		if strings.Compare(key, "title") == 0 {
			title = value
		}
		if strings.Compare(key, "date") == 0 {
			date = value
		}
		if strings.Compare(key, "tags") == 0 {
			value = strings.ReplaceAll(value, "[", "")
			value = strings.ReplaceAll(value, "]", "")
			valArr := strings.Split(value, ",")
			value = strings.Join(valArr, " ")
			tags = value
		}
		if strings.Compare(key, "description") == 0 {
			desc = value
		}
	}
	return title, date, tags, desc
}


