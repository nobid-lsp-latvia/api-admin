// SPDX-License-Identifier: EUPL-1.2

package object

// AttestationAction represents attestation action.
type AttestationAction struct {
	// ID is the attestation identifier
	ID string `json:"attestationId" example:"1be95f00-d9ed-4778-8293-4e4269dab474"`
	// InstanceID is the instance identifier
	InstanceID string `json:"instanceId" example:"01JE3M0QX7PXSKS3DT8T3PK678"`
	Action
}
