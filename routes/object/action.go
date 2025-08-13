// SPDX-License-Identifier: EUPL-1.2

package object

// Action represents instance action.
type Action struct {
	// ActionName represents operation e.g. suspend, unsuspend, revoke.
	ActionName string `json:"action" example:"suspend"`
}
