// SPDX-License-Identifier: EUPL-1.2

package routes

import (
	app "git.zzdats.lv/edim/api-admin"
	"git.zzdats.lv/edim/api-admin/openapi"

	"github.com/nobid-lsp-latvia/go-idauth"
	oa "github.com/nobid-lsp-latvia/go-openapi"
)

type router struct {
	*app.App
	openapi *oa.OpenAPI
}

func Init(a *app.App) error {
	r := &router{
		App: a,
	}
	r.openapi = oa.NewDefaultOpenAPIHandler(openapi.OpenAPIDefinition, a.App)

	a.Get("/healthz", r.healthz)

	v1 := r.Group("/1.0")
	v1.Use(idauth.Authentication(a.App, r.Config().IDAuth))

	v1.Post("/person", idauth.UserHasScopeAtLeastLevel("admin/wallet", idauth.ScopeLevelRead, r.personID))
	v1.Delete("/person/{id}", idauth.UserHasScopeAtLeastLevel("admin/wallet", idauth.ScopeLevelDelete, r.personDelete))

	v1.Post("/instances", idauth.UserHasScopeAtLeastLevel("admin/wallet", idauth.ScopeLevelRead, r.instanceList))

	inst := v1.Group("/instance/{id}")
	{
		inst.Post("/suspend", idauth.UserHasScopeAtLeastLevel("admin/wallet", idauth.ScopeLevelWrite, r.instanceSuspend))
		inst.Post("/unsuspend", idauth.UserHasScopeAtLeastLevel("admin/wallet", idauth.ScopeLevelWrite, r.instanceUnsuspend))
		inst.Post("/revoke", idauth.UserHasScopeAtLeastLevel("admin/wallet", idauth.ScopeLevelDelete, r.instanceRevoke))
		inst.Get("/attestations", idauth.UserHasScopeAtLeastLevel("admin/wallet", idauth.ScopeLevelRead, r.attestationList))

		attest := inst.Group("/attestation/{aid}")
		{
			attest.Post("/suspend", idauth.UserHasScopeAtLeastLevel("admin/wallet", idauth.ScopeLevelWrite, r.attestationSuspend))
			attest.Post("/unsuspend", idauth.UserHasScopeAtLeastLevel("admin/wallet", idauth.ScopeLevelWrite, r.attestationUnsuspend))
			attest.Post("/revoke", idauth.UserHasScopeAtLeastLevel("admin/wallet", idauth.ScopeLevelDelete, r.attestationRevoke))
		}
	}

	return nil
}
