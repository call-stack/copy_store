package store

import (
	"github.com/call-stack/copy_store.git/rpc/copystore"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Validator struct{}

func (v *Validator) ValidatePasteReq(req *copystore.PasteReq) error {

	err := validation.Errors{
		"content": validation.Validate(req.Content, validation.Required, validation.Length(5, 1000)),
	}.Filter()

	return err
}
