package api

import (
	"net/http"
	"time"

	"github.com/mbdeguzman/paypal_logs/models"
	"github.com/uadmin/uadmin"
)

func PaypalAPIHandler(w http.ResponseWriter, r *http.Request) {
	status := "ERROR"
	paypal := models.PaypalLogs{}
	paypal.ReferenceNumber = r.FormValue("ReferenceNumber")
	paypal.Restaurant = r.FormValue("restaurant")
	paypal.URL = r.FormValue("URL")
	paypal.Response = r.FormValue("Response")
	paypal.RequestDate = time.Now()
	uadmin.Save(&paypal)
	status = "SUCCESS"
	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"status": status,
	})
}
