package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	srcPath = `D:\downloads`
	dstPath = `D:\downloads`

	img   = "image"
	gif   = "gif"
	vid   = "video"
	exe   = "exe"
	zip   = "zip"
	other = "other"
)

var subFolderMp = map[string]bool{
	img:   true,
	gif:   true,
	vid:   true,
	exe:   true,
	zip:   true,
	other: true,
}

func main() {
	// read all files form src dir
	srcFiles, err := ioutil.ReadDir(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	// create sub-folder if not exist
	if err := createDir(dstPath, img, gif, vid, exe, zip, other); err != nil {
		log.Fatal(err)
	}

	// range over files and move to its sub-folder
	for _, file := range srcFiles {
		ext := filepath.Ext(file.Name())
		// don't move our created sub-folder
		if subFolderMp[file.Name()] {
			continue
		}
		switch ext {
		case ".msi", ".exe":
			newLocat := filepath.Join(dstPath, exe)
			if err := moveFile(srcPath, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		case ".gif":
			newLocat := filepath.Join(dstPath, gif)
			if err := moveFile(srcPath, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		case ".mp4":
			newLocat := filepath.Join(dstPath, vid)
			if err := moveFile(srcPath, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		case ".zip":
			newLocat := filepath.Join(dstPath, zip)
			if err := moveFile(srcPath, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		case ".jpg", ".png":
			newLocat := filepath.Join(dstPath, img)
			if err := moveFile(srcPath, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		default:
			newLocat := filepath.Join(dstPath, other)
			if err := moveFile(srcPath, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func createDir(dst string, subFolder ...string) error {
	if len(subFolder) != 0 {
		for _, sub := range subFolder {
			path := filepath.Join(dst, sub)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				log.Printf("create dir: %s", path)
				os.Mkdir(path, 0755)
			} else if err != nil {
				return err
			}
		}
	} else {
		if _, err := os.Stat(dst); os.IsNotExist(err) {
			os.Mkdir(dst, 0755)
		} else if err != nil {
			return err
		}
	}

	return nil
}

func moveFile(src, dst, filename string) error {
	oldLocat := filepath.Join(src, filename)
	newLocat := filepath.Join(dst, filename)

	if err := os.Rename(oldLocat, newLocat); err != nil {
		return err
	}
	return nil
}
