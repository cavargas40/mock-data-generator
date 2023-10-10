package csv

import (
	"encoding/csv"
	"os"
)

// CSVWriter is a wrapper around the standard CSV writer.
type CSVWriter struct {
	writer *csv.Writer
}

// NewCSVWriter creates a new CSVWriter instance.
func NewCSVWriter(file *os.File) *CSVWriter {
	return &CSVWriter{
		writer: csv.NewWriter(file),
	}
}

// Write writes a single row of data to the CSV file.
func (cw *CSVWriter) Write(data []string) error {
	return cw.writer.Write(data)
}

// Close closes the CSVWriter and flushes any buffered data to the underlying writer.
func (cw *CSVWriter) Close() {
	cw.writer.Flush()
}

// Flush flushes any buffered data to the underlying writer.
func (cw *CSVWriter) Flush() {
	cw.writer.Flush()
}
