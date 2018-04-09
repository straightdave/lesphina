package lesphina

type Type int

const (
	IMPORT Type = iota
	STRUCT
	INTERFACE
	FUNCTION
	VAR
	CONST
)
