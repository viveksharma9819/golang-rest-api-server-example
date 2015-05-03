package customers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sogko/golang-rest-api-server-example/server"
	"net/http"
)

// HandleCustomersGet is the controller for GET /customers
func HandleCustomersGet(w http.ResponseWriter, req *http.Request) {
	r := server.RendererCtx(req)
	db := server.DbCtx(req)

	var message interface{}
	customers, err := GetCustomers(db)
	success := (err == nil)

	r.JSON(w, http.StatusOK, map[string]interface{}{
		"customers": customers,
		"success":   success,
		"message":   message,
	})
}

// HandleCustomersPost is the controller for POST /customers
func HandleCustomersPost(w http.ResponseWriter, req *http.Request) {
	r := server.RendererCtx(req)
	db := server.DbCtx(req)

	var customer Customer

	// decode JSON body into Customer
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&customer)
	if err != nil {
		r.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Malformed JSON body",
			"success": false,
		})
		return
	}

	// ensure that customer obj is valid
	if !customer.IsValid() {
		r.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid customer object",
			"success": false,
		})
		return
	}

	err = CreateCustomer(db, &customer)
	if err != nil {
		r.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to save customer object",
			"success": false,
		})
		return
	}

	r.JSON(w, http.StatusCreated, map[string]interface{}{
		"customer": customer,
		"success":  true,
	})
}

// HandleCustomerGet is the controller for GET /customer/{id}
func HandleCustomerGet(w http.ResponseWriter, req *http.Request) {
	r := server.RendererCtx(req)
	db := server.DbCtx(req)
	params := mux.Vars(req)
	id := params["id"]

	var message interface{}
	customer, err := GetCustomer(db, id)
	if err != nil {
		message = err.Error()
		customer = nil
	}
	success := (err == nil)

	r.JSON(w, http.StatusOK, map[string]interface{}{
		"customer": customer,
		"success":  success,
		"message":  message,
	})
}