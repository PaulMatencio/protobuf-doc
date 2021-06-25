package lib

import (
	"encoding/binary"
	"io"
)

//  write msg tofile
func Write(w io.Writer, msg []byte) error {

	var (
		buf = make([]byte, 4)
		err error
	)
	binary.LittleEndian.PutUint32(buf, uint32(len(msg)))
	//fmt.Println("Write: Size of the file",uint32(len(msg)))
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
	// fmt.Println("Reading: size of the file",size)
	msg := make([]byte, size)

	if _, err = io.ReadFull(r, msg); err != nil {
		return nil, err
	}

	return msg, err
}

