// SPDX-License-Identifier: EUPL-1.2

package routes

import (
	"git.zzdats.lv/edim/api-admin/routes/object"
	"git.zzdats.lv/edim/api-admin/routes/request"

	"azugo.io/azugo"
	"github.com/nobid-lsp-latvia/go-audit"
	"github.com/valyala/fasthttp"
)

// @operationId personGet
// @title Returns person id
// @description Method returns person id.
// @param PersonCode body request.PersonCode true "Person code"
// @success 200 PersonID object.PersonID "Get person ID"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 500 string string "Internal server error"
// @resource Person
// @route /1.0/person [post].
func (r *router) personID(ctx *azugo.Context) {
	req := &request.PersonCode{}
	if err := ctx.Body.JSON(req); err != nil {
		ctx.Error(err)

		return
	}

	err := r.auditPerson(ctx, nil, nil, &req.PersonCode, nil, audit.ActionSearch)
	if err != nil {
		ctx.Error(err)

		return
	}

	resp := &object.PersonID{}

	err = r.Store().Exec(ctx, "person.get_person_id_by_code", req, &resp)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(resp)
}

// @operationId personDelete
// @title Deletes person data
// @description Method deletes all person data.
// @param id path string true "Person ID"
// @success 204 {empty} "No content"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 500 string string "Internal server error"
// @resource Person
// @route /1.0/person/{id} [delete].
func (r *router) personDelete(ctx *azugo.Context) {
	id := ctx.Params.String("id")

	personReq := &object.PersonID{
		ID: id,
	}

	err := r.Store().Exec(ctx, "person.delete_person", personReq, nil)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.StatusCode(fasthttp.StatusNoContent)
}

func (r *router) auditPerson(ctx *azugo.Context, givenName *string, familyName *string, personCode *string, identifierType *string, action audit.Action) error {
	request := audit.AuditRequest{
		ClientID: "api-admin",
		Action:   string(action),
		Person: &audit.Person{
			GivenName:      givenName,
			FamilyName:     familyName,
			Identifier:     personCode,
			IdentifierType: identifierType,
		},
	}

	err := r.App.Audit().PersonRequest(ctx, request)
	if err != nil {
		return err
	}

	return nil
}
