package main

type SemanticError struct {
	Message string `json:"message"`
}

func (e SemanticError) Error() string {
	return e.Message
}

var SemanticErrors = struct {
	VersionIsNotValid      SemanticError
	MajorVersionIsNotValid SemanticError
	MinorVersionIsNotValid SemanticError
	PatchVersionIsNotValid SemanticError
}{
	VersionIsNotValid:      SemanticError{Message: "Semantic version is not valid format"},
	MajorVersionIsNotValid: SemanticError{Message: "Major version is not int format"},
	MinorVersionIsNotValid: SemanticError{Message: "Minor version is not int format"},
	PatchVersionIsNotValid: SemanticError{Message: "Patch version is not int format"},
}
