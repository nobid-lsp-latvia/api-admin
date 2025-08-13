// SPDX-License-Identifier: EUPL-1.2

package response

import (
	"time"

	"github.com/goccy/go-json"
)

// Attestation represents the response of the attestation.
type Attestation struct {
	// ID of the attestation
	ID string `json:"id"`
	// Name of the attestation
	Name string `json:"name"`
	// Type of the attestation
	Type string `json:"type"`
	// Format of the attestation
	Format string `json:"format"`
	// Status of the instance eg. `active`, `suspended`, `revoked`, `expired`
	Status string `json:"status"`
	// IssuedOn represents the date when the attestation was issued
	IssuedOn time.Time `json:"issuedOn"`
	// ExpiresOn represents the date when the attestation expires
	ExpiresOn time.Time `json:"expiresOn"`
	// PublicKey of the attestation
	PublicKey string `json:"publicKey"`
	// Attributes of the attestation
	Attributes json.RawMessage `json:"attributes"`
	// InstallStatus represents the status of the attestation installation (`sent`, `success`, `error`, `waiting_wi`)
	InstallStatus string `json:"installStatus"`
	// InstallMessage represents the message of the attestation installation
	InstallMessage string `json:"installMessage"`
}
