package lib

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	// 	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

)

func Backup(indir string, outdir string, pn string, np int) {

	doc_meta := "This is the document metadata"
	object := []byte{}
	doc := CreateDocument(pn,doc_meta,0, np,&object)
	// doc.NumberOfPages = int32(np)
	size := 0

	for i := 1; i <= np; i++ {
		page := fmt.Sprintf("%04d", i)
		fn := filepath.Join(indir, pn+"_"+page)
		page_meta := "This is the page metadata of page " + page
		if buf, err := ioutil.ReadFile(fn); err == nil {
			p, _ := strconv.Atoi(page)
			pg := CreatePage(pn,page_meta, p,  &buf)
			AddPageToDucument(pg, doc)
			size += len(buf)
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println(size)
	if bytes, err := proto.Marshal(doc); err == nil {
		ofn := filepath.Join(outdir, pn)
		if f, err := os.OpenFile(ofn, os.O_WRONLY|os.O_CREATE, 0600); err == nil {
			if err := Write(f, bytes); err == nil {
				fmt.Printf("%d bytes have be written to %s\n", len(bytes), ofn)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}


