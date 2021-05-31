package lib

import "github.com/paulmatencio/protobuf-doc/src/document/documentpb"

func CreateDocument(pn string, metadata string, pages int, object *[]byte )  *documentpb.Document {
	// if big document, pages < 0 since it is initialized at the beginning
	doc := &documentpb.Document{
		DocId: pn,
		Metadata: metadata,
		// NumberOfPages: int32(np),
		PageNumber: int32(pages),
		Object: *object,
	}
	return doc
}