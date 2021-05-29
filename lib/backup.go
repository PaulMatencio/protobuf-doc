package lib

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func Backup(indir string, outdir string ,pn string, np int) {

	doc_meta := "This is the document metadata"
	doc := CreateDocument(doc_meta)
	doc.NumberOfPages= int32(np)
	size:= 0

	for i := 1; i <= np; i++ {
		page := fmt.Sprintf("%04d", i)
		fn := filepath.Join(indir, pn+"_"+page)
		page_meta := "This is the page metadata of page " + page
		if buf, err := ioutil.ReadFile(fn); err == nil {
			p, _ := strconv.Atoi(page)
			pg := CreatePage(int32(p), page_meta, buf)
			AddPageToDucument(pg, doc)
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
