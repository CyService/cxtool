package converter

// Basic interface for all CX format converter

type Converter interface {

	Convert(sourceFileName string)

	ConvertFromStdin()
}
