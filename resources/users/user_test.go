package users

import "testing"

func TestExtractedInitialsWithCompleteInput(t *testing.T) {

	user := User{FirstName:"usama", LastName: "iqbal", Company: "xyz"}

	// check if extracted initials are "ui"
	initials := user.Initials()

	if initials != "ui"{
		 t.Fatalf("expected ui => got : %s", initials)
	}
}

func TestExtractedInitialsWithFirstNameMissing(t *testing.T) {

	user := User{FirstName:"", LastName: "iqbal", Company: "xyz"}

	// check if extracted initials are "ui"
	initials := user.Initials()

	if initials != "i"{
		t.Fatalf("expected i => got : %s", initials)
	}
}

func TestExtractedInitialsWithLastNameMissing(t *testing.T) {

	user := User{FirstName:"usama", LastName: "", Company: "xyz"}

	// check if extracted initials are "ui"
	initials := user.Initials()

	if initials != "u"{
		t.Fatalf("expected i => got : %s", initials)
	}
}

func TestExtractedInitialsWithMissingName(t *testing.T) {

	user := User{FirstName:"", LastName: "", Company: "xyz"}

	// check if extracted initials are "ui"
	initials := user.Initials()

	if initials != ""{
		t.Fatalf("expected i => got : %s", initials)
	}
}

func TestFullNameWhenBothNamesAreProvided(t *testing.T) {
	user := User{FirstName:"usama", LastName: "iqbal", Company: "xyz"}

	// check if extracted full name is "usamaiqbal"
	res := user.FullName()

	if res != "usamaiqbal" {
		t.Fatalf("expected i => got : %s", res)
	}
}

func TestFullNameWhenFirstNameIsMissing(t *testing.T) {
	user := User{FirstName:"", LastName: "iqbal", Company: "xyz"}

	// check if extracted full name is "usamaiqbal"
	res := user.FullName()

	if res != "iqbal" {
		t.Fatalf("expected i => got : %s", res)
	}
}

func TestFullNameWhenLastNameIsMissing(t *testing.T) {
	user := User{FirstName:"usama", LastName: "", Company: "xyz"}

	// check if extracted full name is "usama"
	res := user.FullName()

	if res != "usama" {
		t.Fatalf("expected i => got : %s", res)
	}
}

func TestFullNameWhenBothNamesAreMissing(t *testing.T) {
	user := User{FirstName:"", LastName: "", Company: "xyz"}

	// check if extracted full name is ""
	res := user.FullName()

	if res != "" {
		t.Fatalf("expected i => got : %s", res)
	}
}

func TestComb1WhenBothNamesAreProvided(t *testing.T) {
	user := User{FirstName:"usama", LastName: "iqbal", Company: "xyz"}

	// check if extracted full name is "uiqbal"
	res := user.Combination1()

	if res != "uiqbal" {
		t.Fatalf("expected i => got : %s", res)
	}
}

func TestComb1WhenFirstNameIsMissing(t *testing.T) {
	user := User{FirstName:"", LastName: "iqbal", Company: "xyz"}

	// check if extracted full name is "iqbal"
	res := user.Combination1()

	if res != "iqbal" {
		t.Fatalf("expected i => got : %s", res)
	}
}

func TestComb2WhenLastNameIsMissing(t *testing.T) {
	user := User{FirstName:"usama", LastName: "", Company: "xyz"}

	// check if extracted full name is "u"
	res := user.Combination1()

	if res != "u" {
		t.Fatalf("expected i => got : %s", res)
	}
}



