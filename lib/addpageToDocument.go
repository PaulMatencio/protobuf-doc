package lib


import "github.com/paulmatencio/protobuf-doc/src/document/documentpb"

func AddPageToDucument(page *documentpb.Page, document *documentpb.Document) {
	document.Page = append(document.Page, page)
}
