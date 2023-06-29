package compact

// "encoding/csv"

type Writer interface {
	Write(record []string) error
	Flush() error
}

/*
type compactWriter struct {
	w *csv.Writer
}
*/
