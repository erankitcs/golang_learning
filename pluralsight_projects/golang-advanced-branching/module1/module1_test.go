package module1

import (
	"testing"
)

// Task 1: Define vehicle interface
func TestVehicleInterfaceIsDefined(t *testing.T) {

	didFindAnInterface, didFindTheInterface := checkInterface("vehicle")

	if !didFindAnInterface || !didFindTheInterface {
		t.Error("Did not define an interface named `vehicle`")
	}

}

// Task 2: Define 3 structs
func TestMainStructsAreDefined(t *testing.T) {

	// Check for car struct
	didFindAStruct, didFindTheStruct := checkStruct("car")
	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `car`")
	}

	// Check for truck struct
	didFindAStruct, didFindTheStruct = checkStruct("truck")

	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `truck`")
	}

	// Check for bike struct
	didFindAStruct, didFindTheStruct = checkStruct("bike")

	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `bike`")
	}
}

// Task 3: Define car and truck fields
func TestCarTruckFields(t *testing.T) {

	// car fields
	if !checkStructProperties("car", "model", "string") {
		t.Error("Did not define `model` field in `car` with the proper type")
	}
	if !checkStructProperties("car", "make", "string") {
		t.Error("Did not define `make` field in `car` with the proper type")
	}
	if !checkStructProperties("car", "typeVehicle", "string") {
		t.Error("Did not define `typeVehicle` field in `car` with the proper type")
	}

	// truck fields
	if !checkStructProperties("truck", "model", "string") {
		t.Error("Did not define `model` field in `truck` with the proper type")
	}
	if !checkStructProperties("truck", "make", "string") {
		t.Error("Did not define `make` field in `truck` with the proper type")
	}
	if !checkStructProperties("truck", "typeVehicle", "string") {
		t.Error("Did not define `typeVehicle` field in `truck` with the proper type")
	}
}

// Task 4: Define bike fields
func TestBikeFields(t *testing.T) {
	// bike fields
	if !checkStructProperties("bike", "model", "string") {
		t.Error("Did not define `model` field in `bike` with the proper type")
	}
	if !checkStructProperties("bike", "make", "string") {
		t.Error("Did not define `make` field in `bike` with the proper type")
	}
}

// Task 5: Define Values struct
func TestValuesStructIsDefined(t *testing.T) {

	// Check for car struct
	didFindAStruct, didFindTheStruct := checkStruct("Values")
	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `Values`")
	}
}

// Task 6: Define Model struct
func TestModelStructIsDefined(t *testing.T) {

	// Check for car struct
	didFindAStruct, didFindTheStruct := checkStruct("Model")
	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `Model`")
	}
}

// Task 7: Add Values fields
func TestValuesFields(t *testing.T) {
	// Values fields
	if !checkStructProperties("Values", "Models", "[]Model") {
		t.Error("Did not define `Models` field in `Values` with the proper type")
	}

}

// Task 8: Add Model fields
func TestModelFields(t *testing.T) {
	// Model fields
	if !checkStructProperties("Model", "Name", "string") {
		t.Error("Did not define `Name` field in `Model` with the proper type")
	}

	if !checkStructProperties("Model", "Feedback", "[]string") {
		t.Error("Did not define `Feedback` field in `Model` with the proper type")
	}
}

// Task 9: Define feedback struct and values
func TestFeedbackResultStructIsDefinedAndHasValues(t *testing.T) {

	// Check for feedback struct
	didFindAStruct, didFindTheStruct := checkStruct("feedbackResult")
	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `feedbackResult`")
	}

	// feedbackResult fields
	if !checkStructProperties("feedbackResult", "feedbackTotal", "int") {
		t.Error("Did not define `feedbackTotal` field in `feedbackResult` with the proper type")
	}
	if !checkStructProperties("feedbackResult", "feedbackPositive", "int") {
		t.Error("Did not define `feedbackPositive` field in `feedbackResult` with the proper type")
	}
	if !checkStructProperties("feedbackResult", "feedbackNegative", "int") {
		t.Error("Did not define `feedbackNegative` field in `feedbackResult` with the proper type")
	}
	if !checkStructProperties("feedbackResult", "feedbackNeutral", "int") {
		t.Error("Did not define `feedbackNeutral` field in `feedbackResult` with the proper type")
	}

}

// Task 10: Define variables
func TestVariables(t *testing.T) {

	// vehicleResult map
	if !checkMap("vehicleResult", "string", "feedbackResult") {
		t.Error("Did not declare `vehicleResult` with the proper type.")
	}

	// inventory slice
	if !checkSlice("inventory", "vehicle") {
		t.Error("Did not declare `inventory` slice of type 'vehicle' ")
	}

}

// Task 11: Checking var initialization under func init
func TestInitializeVars(t *testing.T) {

	if !checkVarWithinFunc("init", "vehicleResult") {
		t.Error("Uncomment the `vehicleResult` assignment statement within the `init` function ")
	}

	if !checkVarWithinFunc("init", "inventory") {
		t.Error("Uncomment the `inventory` assignment statement within the `init` function ")
	}
}
