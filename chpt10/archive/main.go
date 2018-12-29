package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	src, desc := os.Args[1], os.Args[2]
	fmt.Println(src, desc)
	err := unZip(src, desc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unzip!")
}

func unZip(src, destination string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		fPath := path.Join(destination, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fPath, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(fPath), f.Mode())

			outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return err
			}
		}
	}
	return nil
}
