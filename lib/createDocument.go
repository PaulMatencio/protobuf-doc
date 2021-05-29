package lib

import "protobuf-doc/src/document/documentpb"

func CreateDocument(metadata string) *documentpb.Document {
	doc := &documentpb.Document{
		Metadata: metadata,
	}
	return doc
}