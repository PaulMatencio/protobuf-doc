package lib

import "protobuf-doc/src/document/documentpb"

func CreatePage(pagen int32, metadata string, object []byte) *documentpb.Page {
	pe := &documentpb.Page{
		PageNumber: pagen,
		Metadata:   metadata,
		Object:     object,
	}
	return pe
}

