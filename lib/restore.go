package lib

import (
	"fmt"
	"os"
	"path/filepath"
    "github.com/paulmatencio/protobuf-doc/src/document/documentpb"
	"github.com/golang/protobuf/proto"
)

func Restore(indir string, outdir string, pn string) {

	var (
		ifn   = filepath.Join(indir, pn)
		err   error
		f     *os.File
		bytes []byte
	)

	fmt.Println("Reading file ", ifn)
	if f, err = os.OpenFile(ifn, os.O_RDONLY, 0600); err == nil {
		defer f.Close()
		var doc = &documentpb.Document{}
		if bytes, err = Read(f); err == nil {
			if err = proto.Unmarshal(bytes, doc); err == nil {
				pages := doc.GetPage()
				if len(pages) != int(doc.NumberOfPages) {
					fmt.Println("Backup of document is inconsistent %s", pn)
					os.Exit(100)
				}

				for _, page := range pages {
					//object := page.GetObject()
					pfn := pn + "_" + fmt.Sprintf("%04d", page.GetPageNumber())
					if fi, err := os.OpenFile(filepath.Join(outdir, pfn), os.O_WRONLY|os.O_CREATE, 0600); err == nil {
						defer fi.Close()
						bytes := page.GetObject()
						if _, err := fi.Write(bytes); err != nil {
							fmt.Printf("Error %v writing page %d", err, pfn)
						}
					}
				}
				fmt.Println(doc.NumberOfPages)
				fmt.Println(doc.Metadata)
				fmt.Println(doc.LastUpdated)
			}
		} else {
			fmt.Printf("Error %v while reading file  %s", ifn, err)
		}
	} else {
		fmt.Printf("Error %v while opening file %s", ifn, err)
	}

}
