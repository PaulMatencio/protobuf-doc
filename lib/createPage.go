package lib


import "github.com/paulmatencio/protobuf-doc/src/document/documentpb"

func CreatePage(pn string, metadata string, pagen int, object *[]byte) *documentpb.Page {
	pe := &documentpb.Page{
		PageId: pn,
		PageNumber: int32(pagen),
		Metadata:   metadata,
		Object:     *object,
	}
	return pe
}

