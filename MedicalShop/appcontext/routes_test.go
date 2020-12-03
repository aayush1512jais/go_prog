package appcontext

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aayush1512jais/go_prog/MedicalShop/controller"
)

func TestSetupRoutes(t *testing.T) {

	var medicineController *controller.Controller
	r := SetupRoutes(medicineController)

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/MedicalShop/status")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(body)
	expected := "Server Running!!"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}
