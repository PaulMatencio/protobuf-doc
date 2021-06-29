package lib
import "github.com/paulmatencio/protobuf-doc/src/document/documentpb"
func CreatePdf(pn string, metadata string, object *[]byte) *documentpb.Pdf {
	pd := &documentpb.Pdf{
		PdfId:    pn, // docid
		Size:     (int64)(len(*object)),
		Metadata: metadata,
		Pdf:      *object,
	}
	return pd
}