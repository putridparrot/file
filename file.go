package file

import (
	"os"
	"path/filepath"
	"io"
	"log"
	"strings"
)

type IFileCopy interface {
	Copy(src string, dest string) error
}

type FileCopy struct {
	WhatIf bool
}

func NewFileCopy() *FileCopy {
	return &FileCopy {}
}

func (fc *FileCopy) copyFile(src string, dest string) error {

	_, srcFilename := filepath.Split(src)
	destFileName := dest + "\\" + srcFilename

	log.Println("Copy File: " + src + " to " + destFileName)
	if !fc.WhatIf {
		srcFile, err := os.Open(src)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err  := os.Create(destFileName)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}

		err = destFile.Sync()
		if err != nil {
			return err
		}
	}
	return nil
}

func (fc *FileCopy) copy(srcRoot string, src string, dest string) error {
	return filepath.Walk(src, func(path string, fi os.FileInfo, err error) error {
		switch mode := fi.Mode(); {
		case mode.IsDir():
			child := strings.Replace(path, srcRoot, "", 1)
			newDest := dest + child
			log.Println("Copy Directory " + path + " to " + newDest)
			if !fc.WhatIf {
				if _, err := os.Stat(newDest); os.IsNotExist(err) {
					os.Mkdir(newDest, os.ModeDir)
				}
			}
		case mode.IsRegular():
			child := strings.Replace(filepath.Dir(path), srcRoot, "", 1)
			newDest := dest + child
			return fc.copyFile(path, newDest)
		}
		return nil
	})
}

func (fc *FileCopy) Copy(src string, dest string) error {
	fi, err := os.Stat(src)
	if err != nil {
		return nil
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return fc.copy(src, src, dest)
	case mode.IsRegular():
		fc.copyFile(src, dest)
	}

	return nil
}
