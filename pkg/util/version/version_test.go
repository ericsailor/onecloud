package version

import "testing"

func TestVersion(t *testing.T) {
	if !LE("1.2.3", "1.2.4") {
		t.Errorf(`!LE("1.2.3", "1.2.4")`)
	}
	if !LE("1.2.4", "1.2.4") {
		t.Errorf(`!LE("1.2.4", "1.2.4")`)
	}
	if LE("1.2.5", "1.2.4") {
		t.Errorf(`!LE("1.2.5", "1.2.4")`)
	}
	if !LT("1.2.3", "1.3.1") {
		t.Errorf(`!LT("1.2.3", "1.3.1")`)
	}
	if LT("1.2.4", "1.2.4") {
		t.Errorf(`LT("1.2.4", "1.2.4")`)
	}
	if LT("1.2.4.1", "1.2.4") {
		t.Errorf(`LT("1.2.4.1", "1.2.4")`)
	}

	if GT("1.2.3", "2.3.4.1") {
		t.Errorf(`GT("1.2.3", "2.3.4.1")`)
	}
	if !GT("2.12.1", "2.9.2") {
		t.Errorf(`GT("2.12.1", "2.9.2")`)
	}
	if GT("2.12.1", "2.12.1") {
		t.Errorf(`GT("2.12.1", "2.12.1")`)
	}
	if !GE("2.12.1", "2.12.1") {
		t.Errorf(`!GE("2.12.1", "2.12.1")`)
	}
	if !GE("2.12.1", "2.9.2") {
		t.Errorf(`GE("2.12.1", "2.9.2")`)
	}
}
