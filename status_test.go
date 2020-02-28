package assignment2

import "testing"

func TestGetStatusCode(t *testing.T) {
	_, err := GetStatusCode("http://google.com")
	//if get status does not return an error it works
	if err != nil {
		t.Errorf("Statuscode dont work")
	}

}
