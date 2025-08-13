// SPDX-License-Identifier: EUPL-1.2

package request

import "git.zzdats.lv/edim/api-admin/routes/object"

// Person represents person identifier.
type Person struct {
	object.PersonID
	PersonCode
}
