package system

import (
	uos "github.com/tkblack/uos-go"
	"github.com/tkblack/uos-go/ecc"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `wxbio.system` contract.
func NewRegProducer(producer uos.AccountName, producerKey ecc.PublicKey, url string, location uint16) *uos.Action {
	return &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("regproducer"),
		Authorization: []uos.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(RegProducer{
			Producer:    producer,
			ProducerKey: producerKey,
			URL:         url,
			Location:    location,
		}),
	}
}

// RegProducer represents the `wxbio.system::regproducer` action
type RegProducer struct {
	Producer    uos.AccountName `json:"producer"`
	ProducerKey ecc.PublicKey   `json:"producer_key"`
	URL         string          `json:"url"`
	Location    uint16          `json:"location"` // what,s the meaning of that anyway ?
}
