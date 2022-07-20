package module2

import (
	"testing"
)

// Task 1: Find function readJSONFile
func TestReadJSONFileFuncIsDefined(t *testing.T) {
	if !checkFunc("readJSONFile") {
		t.Error("Uncomment the function named `readJSONFile`")
	}

	if !checkImports("\"encoding/json\"") {
		t.Error("Import the package `encoding/json`")
	}
	if !checkImports("\"io/ioutil\"") {
		t.Error("Import the package `io/ioutil`")
	}
	if !checkImports("\"log\"") {
		t.Error("Import the package `log`")
	}
	if !checkImports("\"os\"") {
		t.Error("Import the package `os`")
	}
}

// Task 2: Create function generateRating()
func TestGenerateRatingFuncIsDefined(t *testing.T) {
	if !checkFunc("generateRating") {
		t.Error("Create a function `generateRating` with no parameters and no return values")
	}
}

// Task 3: Assign variable to the JSON content result.
func TestAssignVarFeedbackReturn(t *testing.T) {
	if !checkVarWithinFunc("generateRating", "f") {
		t.Error("Call the function `readJSONfile()` and assign it to a variable `f` ")
	}

}

// Task 4: Check `for _, v := range f.Models`
func TestMainForRangeModels(t *testing.T) {
	if !checkMainForWithinFunc("generateRating", "_", "v", "Models") {
		t.Error("Could not find the for statement with no key and `v` as value and `f.Models` as range")
	}
}

// Task 5: Declare variables within for stmt
func TestVarDeclInGenerateRating(t *testing.T) {

	if !checkVarDeclWithinFor("vehResult", "feedbackResult") {
		t.Error("Variable `vehresult` of type `feedbackResult` is not declared")
	}

	if !checkVarDeclWithinFor("vehRating", "rating") {
		t.Error("Variable `vehRating` of type `rating` is not declared")
	}
}

// Task 6: Check `for _, msg := range v.Feedback`
func TestForRangeFeedback(t *testing.T) {

	if !checkForStmt(mainForStmt, "_", "msg", "Feedback") {
		t.Error("Could not find the for statement with no key and `msg` as value and `v.Feedback` as range")
	}
}

// Task 7: Check `if text := strings.Split(msg, " ") ; len(text) >= 5`
func TestIfStmt(t *testing.T) {
	if !checkIfStmt(forBlock, "text", "strings", "len(text)>=5") {
		t.Error("If statment is either not defined or missing the initialization and the condition is not correct.")
	}

	if !checkImports("\"strings\"") {
		t.Error("Import the package `strings`")
	}

}

// Task 8: Set variable values
func TestSetVarValues(t *testing.T) {
	if !checkSetValues(ifBlock, "vehRating=5.0") {
		t.Error("vehRating not set")
	}
	if !checkSetValues(ifBlock, "vehResult.feedbackTotal++") {
		t.Error("feedbackTotal not added")
	}

}

// Task 9: Check for `for _, word := range text`
func TestForWithinIf(t *testing.T) {
	if !checkForWithinIf(ifBlock, "_", "word", "text") {
		t.Error("Could not find the for statement with no key and `word` as value and `text` as range")
	}
}

// Task 10: Check for `switch s := strings.Trim(strings.ToLower(word), " ,.,!,?,\t,\n,\r"); s`
func TestSwitchCalRating(t *testing.T) {
	if !checkSwitchCalRating(forWord, "s", "strings", "s") {
		t.Error("Could not find the switch statement with variable `s` assigned  and tagged.")
	}

}

// Task 11: Check for `switch`
func TestSwitchAddFeedback(t *testing.T) {
	if !checkSwitchAddFeedback(ifBlock) {
		t.Error("Could not find the switch statement with no initialization and no tag")
	}

}

// Task 12: Check for `vehicleResult[v.Name] = vehResult`
func TestVarAppendRating(t *testing.T) {
	testFail := false
	if mainForStmt == nil {
		testFail = true
	} else if !checkAppendRating(mainForStmt.Body, "vehicleResult[v.Name]=vehResult") {
		testFail = true
	}

	if testFail {
		t.Error("Assign `vehResult` to `vehicleResult[v.Name]` right before we close the first/main for statement")
	}
}

// Task 13: Call generateRating() within main()
func TestGenerateRatingFuncIsCalled(t *testing.T) {
	if !checkFuncGenerateRating("generateRating") {
		t.Error("Call the function `generateRating()` within the `main` function")
	}
}
