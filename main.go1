package main

import (
	"os"
	"path/filepath"
	 doc "github.com/paulmatencio/protobuf-doc/lib"

)

func main() {
	// backup
	pn := "ST33_E10305447221"
	np := 232;
	src:= "Tiff"
	dest:= "backup"
	home, _ := os.UserHomeDir()
	indir := filepath.Join(home, src)
	outdir := filepath.Join(home, dest)
	// Backup
	doc.Backup(indir,outdir,pn,np)
	// Restore
	indir = filepath.Join(home, dest)
	outdir=  filepath.Join(home, "Tiff1")
	doc.Restore(indir,outdir,pn)
}




