package lesphina

// Kind represents different entry types in lesphina.
type Kind int

// Kinds of a entry
const (
	KindImport Kind = iota
	KindElement
	KindConst
	KindVar
	KindInterface
	KindStruct
	KindFunction
	KindInterfaceMethod
)
