// SPDX-License-Identifier: EUPL-1.2

package response

// Instance represents the response of the instance.
type Instance struct {
	// ID of the instance
	ID string `json:"id"`
	// Status of the instance eg. `active`, `suspended`, `revoked`, `expired`
	Status string `json:"status"`
	// FirebaseID of the instance
	FirebaseID string `json:"firebaseId"`
}
