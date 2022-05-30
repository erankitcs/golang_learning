package simplemaths

import "fmt"

/// Method Declarations
type SemanticVersion struct {
	// private variables . Small letter.
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

//General way.
func AnotherString(sv SemanticVersion) string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

// This is a method for semantic version type.
func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

// Value based Reciever. You have to always create a copy of object then call again for state update.
//func (sv SemanticVersion) IncrementMajor() SemanticVersion {
//	sv.major += 1
//	return sv
//}

// Pointer based method reciever.

func (sv *SemanticVersion) IncrementMajor() {
	sv.major += 1
}

func (sv *SemanticVersion) IncrementMinor() {
	sv.minor += 1
}

func (sv *SemanticVersion) IncrementPatch() {
	sv.patch += 1

}
