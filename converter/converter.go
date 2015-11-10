package converter

type Converter interface {

	Convert(sourceFileName, targetFileName string)
}


