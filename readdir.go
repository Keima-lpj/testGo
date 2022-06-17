package main

import (
	"fmt"
	"os"
)

func main() {
	saveFile, err := os.Create("dir.txt")
	if err != nil {
		fmt.Println(err)
	}
	err = tree(".", 0, saveFile)
	if err != nil {
		fmt.Println(err)
	}
}

func tree(dstPath string, level int, file *os.File) error {
	dstF, err := os.Open(dstPath)
	if err != nil {
		return err
	}
	defer dstF.Close()
	fileInfo, err := dstF.Stat()
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() { //如果dstF是文件
		for i := 0; i < level; i++ {
			file.Write([]byte("--"))
			//fmt.Print("--")
		}
		file.Write([]byte(dstPath + "\r\n"))
		//fmt.Println(dstPath)
		return nil
	} else { //如果dstF是文件夹
		for i := 0; i < level; i++ {
			file.Write([]byte("--"))
			//fmt.Print("--")
		}
		file.Write([]byte(dstF.Name() + "\r\n"))
		//fmt.Println(dstF.Name())
		dir, err := dstF.Readdir(0) //获取文件夹下各个文件或文件夹的fileInfo
		if err != nil {
			return err
		}
		for _, fileInfo = range dir {
			err = tree(dstPath+"/"+fileInfo.Name(), level+1, file)
			if err != nil {
				return err
			}
		}
		return nil
	}

}
