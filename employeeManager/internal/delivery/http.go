package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tokopedia/employeeManager/internal/repository"
)

type httpDelivery struct {
	store *repository.EmployeeStore
}

func NewHTTPDelivery(store *repository.EmployeeStore) *httpDelivery {
	return &httpDelivery{
		store: store,
	}
}

func (d *httpDelivery) Serve() {
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

		employees, err := d.store.ListEmployees(page, pageSize)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
