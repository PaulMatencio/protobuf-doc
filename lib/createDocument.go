package lib

import "github.com/paulmatencio/protobuf-doc/src/document/documentpb"

func CreateDocument(pn string, metadata string, pages int, object *[]byte )  *documentpb.Document {
	doc := &documentpb.Document{
		DocId: pn,
		Metadata: metadata,
		NumberOfPages: int32(pages),
		Object: *object,
	}
	return doc
}