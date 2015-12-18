package converter

import "io"

// Basic interface for all CX format converter

type Converter interface {

	Convert(r io.Reader, w io.Writer)
}
