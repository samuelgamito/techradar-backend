package misc

import "testing"

func TestGetFriendlyNameObject(t *testing.T) {

	got := GetFriendlyName("Test 123")

	want := "test-123"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
