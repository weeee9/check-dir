package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	img   = "image"
	gif   = "gif"
	vid   = "video"
	exe   = "exe"
	zip   = "zip"
	other = "other"
)

var subs = []string{img, gif, vid, exe, zip, other}

var subFolderMp = map[string]string{}

type config struct {
	Src        string
	Dst        string
	SubFolders []string
}

func main() {
	app := &cli.App{
		Name:     "Check Dir",
		Usage:    "Check and move files in src directory to dst directory",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name: "weeee9",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "src",
				Aliases:  []string{"source"},
				Usage:    "source directory",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "dst",
				Aliases:  []string{"destination"},
				Usage:    "source directory",
				Required: true,
			},
		},
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	return exec(config{
		Src: c.String("src"),
		Dst: c.String("dst"),
	})
}

func exec(cfg config) error {
	// read all files form src dir
	srcFiles, err := ioutil.ReadDir(cfg.Src)
	if err != nil {
		log.Fatal(err)
	}

	// create sub-folder if not exist
	if err := createDir(cfg.Dst, subs...); err != nil {
		log.Fatal(err)
	}

	for _, sub := range subs {
		subFolderMp[sub] = sub
	}

	// range over files and move to its sub-folder
	for _, file := range srcFiles {
		ext := filepath.Ext(file.Name())

		// don't move our created sub-folder
		if in(file.Name(), subs) {
			continue
		}

		switch ext {
		case ".msi", ".exe":
			newLocat := filepath.Join(cfg.Dst, exe)
			if err := moveFile(cfg.Src, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		case ".gif":
			newLocat := filepath.Join(cfg.Dst, gif)
			if err := moveFile(cfg.Src, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		case ".mp4":
			newLocat := filepath.Join(cfg.Dst, vid)
			if err := moveFile(cfg.Src, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		case ".zip":
			newLocat := filepath.Join(cfg.Dst, zip)
			if err := moveFile(cfg.Src, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		case ".jpg", ".png":
			newLocat := filepath.Join(cfg.Dst, img)
			if err := moveFile(cfg.Src, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		default:
			newLocat := filepath.Join(cfg.Dst, other)
			if err := moveFile(cfg.Src, newLocat, file.Name()); err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
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

func in(name string, s []string) bool {
	for _, n := range s {
		if n == name {
			return true
		}
	}
	return false
}
