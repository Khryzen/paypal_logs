package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type PaypalLogs struct {
	uadmin.Model
	ReferenceNumber string
	Restaurant      string
	URL             string
	Response        string
	RequestDate     time.Time
}
