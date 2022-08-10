package archive

import (
	"github.com/bodgit/sevenzip"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// FilesIn7ZipArchive 获取压缩包内的文件列表
func FilesIn7ZipArchive(sevenZipFile string, password string) (string, error) {

	// 打开 7z 文件
	rc, err := sevenzip.OpenReaderWithPassword(sevenZipFile, password)
	if err != nil {
		return "", err
	}

	// 确保关闭文件
	defer func() {
		_ = rc.Close()
	}()

	// 文件列表
	var retFileNameList []string
	for _, file := range rc.File {
		retFileNameList = append(retFileNameList, file.Name)
	}

	return strings.Join(retFileNameList, "\n"), nil
}

// Extract7ZipArchive 解压缩
func Extract7ZipArchive(sevenZipFile string, password string, rootPath string) error {

	const DefaultPermission = 0755

	// 打开 7z 文件
	rc, err := sevenzip.OpenReaderWithPassword(sevenZipFile, password)
	if err != nil {
		return err
	}

	// 确保关闭文件
	defer func() {
		_ = rc.Close()
	}()

	// 遍历文件
	for _, file := range rc.File {
		// 处理目录
		if file.FileInfo().IsDir() {
			// 目录名与指定的根目录进行拼接
			tmpPath := path.Join(rootPath, file.Name)
			// 创建目录
			err = os.MkdirAll(tmpPath, DefaultPermission)
			if err != nil {
				return err
			}
		}
	}
	var lastError error
	for _, file := range rc.File {
		// 处理文件
		if !file.FileInfo().IsDir() {
			// 文件名与指定的根目录进行拼接
			tmpFileName := path.Join(rootPath, file.Name)
			fileRc, err := file.Open()
			if err != nil {
				return err
			}
			data, err := ioutil.ReadAll(fileRc)
			if err != nil {
				return err
			}
			_ = fileRc.Close()
			err = ioutil.WriteFile(tmpFileName, data, DefaultPermission)
			if err != nil {
				lastError = err
			}
		}
	}
	return lastError
}
