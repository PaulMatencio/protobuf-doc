package main

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"protobuf-doc/src/document/documentpb"
	"github.com/paulmatencio/protobuf-doc/lib"
	"strconv"
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
	Backup(indir,outdir,pn,np)
	//  restore
	indir = filepath.Join(home, dest)
	outdir=  filepath.Join(home, "Tiff1")
	restore(indir,outdir,pn)
}

func backup(indir string, outdir string ,pn string, np int) {

	doc_meta := "This is the document metadata"
	doc := createDocument(doc_meta)
	doc.NumberOfPages= int32(np)
	size:= 0

	for i := 1; i <= np; i++ {
		page := fmt.Sprintf("%04d", i)
		fn := filepath.Join(indir, pn+"_"+page)
		page_meta := "This is the page metadata of page " + page
		if buf, err := ioutil.ReadFile(fn); err == nil {
			p, _ := strconv.Atoi(page)
			pg := createPage(int32(p), page_meta, buf)
			addPagetoDucument(pg, doc)
			size += len(buf)
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println(size)
	if bytes,err := proto.Marshal(doc); err == nil {
		ofn := filepath.Join(outdir, pn)
		if f, err := os.OpenFile(ofn, os.O_WRONLY|os.O_CREATE, 0600); err == nil {
			if err := Write(f,bytes);err==nil {
				fmt.Printf("%d bytes have be written to %s\n",len(bytes),ofn)
			}
		} else {
			fmt.Println(err)
		}
	}else {
		fmt.Println(err)
	}
}


func restore (indir string, outdir string, pn string) {

	var (
		ifn = filepath.Join(indir,pn)
		err error
		f *os.File
		bytes []byte
	)

	fmt.Println("Reading file ",ifn)
	if f, err = os.OpenFile(ifn, os.O_RDONLY, 0600); err == nil {
		defer f.Close()
		var doc = &documentpb.Document{}
		if bytes,err = Read(f); err == nil {
			if err= proto.Unmarshal(bytes,doc); err ==nil {
				pages := doc.GetPage()
				if len(pages) != int(doc.NumberOfPages) {
					fmt.Println("Backup of document is inconsistent %s",pn)
					os.Exit(100)
				}

				for _,page := range pages {
					  //object := page.GetObject()
					 pfn := pn +"_" +fmt.Sprintf("%04d", page.GetPageNumber())
					 if fi, err := os.OpenFile(filepath.Join(outdir,pfn), os.O_WRONLY|os.O_CREATE, 0600); err == nil {
					 	defer fi.Close()
					 	bytes:= page.GetObject()
					 	if _,err:= fi.Write(bytes); err !=nil {
					 		fmt.Printf("Error %v writing page %d",err,pfn)
						}
					 }
				}
				fmt.Println(doc.NumberOfPages)
				fmt.Println(doc.Metadata)
				fmt.Println(doc.LastUpdated)
			}
		} else {
			fmt.Printf("Error %v while reading file  %s",ifn,err)
		}
	} else {
		fmt.Printf("Error %v while opening file %s",ifn,err)
	}

}

func createPage(pagen int32, metadata string, object []byte) *documentpb.Page {
	pe := &documentpb.Page{
		PageNumber: pagen,
		Metadata:   metadata,
		Object:     object,
	}
	return pe
}

func createDocument(metadata string) *documentpb.Document {
	doc := &documentpb.Document{
		Metadata: metadata,
	}
	return doc
}

func addPagetoDucument(page *documentpb.Page, document *documentpb.Document) {
	document.Page = append(document.Page, page)
}

//  write msg tofile
func Write(w io.Writer, msg []byte) error {

	var (
		buf = make([]byte, 4)
		err error
	)

	binary.LittleEndian.PutUint32(buf, uint32(len(msg)))

	fmt.Println("Write: Size of the file",uint32(len(msg)))

	if _, err = w.Write(buf); err != nil {
		return err
	}

	if _, err = w.Write(msg); err != nil {
		return err
	} else {
		return nil
	}
}

//  read msg from a file
func Read(r io.Reader) ([]byte, error) {

	var (
		buf = make([]byte, 4)
		err error
	)

	if _, err = io.ReadFull(r, buf); err != nil {
		return nil, err
	}

	size := binary.LittleEndian.Uint32(buf)
	fmt.Println("Reading: size of the file",size)
	msg := make([]byte, size)

	if _, err = io.ReadFull(r, msg); err != nil {
		return nil, err
	}

	return msg, err
}
