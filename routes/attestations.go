// SPDX-License-Identifier: EUPL-1.2

package routes

import (
	"git.zzdats.lv/edim/api-admin/routes/object"
	"git.zzdats.lv/edim/api-admin/routes/response"

	"azugo.io/azugo"
)

// @operationId attestationList
// @title Get all attestations for instance
// @description Method gets list of attestations
// @param id path string true "Instance identifier"
// @success 200 AttestationList []response.Attestation "Get attestation list
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 500 string string "Internal server error"
// @resource Attestation
// @route /1.0/instance/{id}/attestations [get].
func (r *router) attestationList(ctx *azugo.Context) {
	req := &object.InstanceID{
		ID: ctx.Params.String("id"),
	}

	resp := make([]response.Attestation, 0)
	if err := r.Store().Exec(ctx, "wallet.get_attestation_list", req, &resp); err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(resp)
}

// @operationId attestationSuspend
// @title Suspend attestation
// @description Method changes attestation status to suspended.
// @param id path string true "Instance identifier"
// @param aid path string true "Attestation identifier"
// @success 200 Attestation response.Attestation "Get updated attestation"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 500 string string "Internal server error"
// @resource Attestation
// @route /1.0/instance/{id}/attestation/{aid}/suspend [post].
func (r *router) attestationSuspend(ctx *azugo.Context) {
	id := ctx.Params.String("id")
	aid := ctx.Params.String("aid")
	action := "suspend"

	var (
		resp *response.Attestation
		err  error
	)

	if resp, err = updateAttestation(r, ctx, id, aid, action); err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(resp)
}

// @operationId attestationUnsuspend
// @title Unsuspend attestation
// @description Method changes attestation status to unsuspended.
// @param id path string true "Instance identifier"
// @param aid path string true "Attestation identifier"
// @success 200 Attestation response.Attestation "Get updated attestation"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 500 string string "Internal server error"
// @resource Attestation
// @route /1.0/instance/{id}/attestation/{aid}/unsuspend [post].
func (r *router) attestationUnsuspend(ctx *azugo.Context) {
	id := ctx.Params.String("id")
	aid := ctx.Params.String("aid")
	action := "unsuspend"

	var (
		resp *response.Attestation
		err  error
	)

	if resp, err = updateAttestation(r, ctx, id, aid, action); err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(resp)
}

// @operationId attestationRevoke
// @title Revokes attestation
// @description Method changes attestation status to unsuspended.
// @param id path string true "Instance identifier"
// @param aid path string true "Attestation identifier"
// @success 200 Attestation response.Attestation "Get updated attestation"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 500 string string "Internal server error"
// @resource Attestation
// @route /1.0/instance/{id}/attestation/{aid}/revoke [post].
func (r *router) attestationRevoke(ctx *azugo.Context) {
	id := ctx.Params.String("id")
	aid := ctx.Params.String("aid")
	action := "revoke"

	var (
		resp *response.Attestation
		err  error
	)

	if resp, err = updateAttestation(r, ctx, id, aid, action); err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(resp)
}

func updateAttestation(r *router, ctx *azugo.Context, id string, aid string, action string) (*response.Attestation, error) {
	req := &object.AttestationAction{
		ID:         aid,
		InstanceID: id,
		Action:     object.Action{ActionName: action},
	}

	resp := &response.Attestation{}
	if err := r.Store().Exec(ctx, "wallet.update_attestation_status", req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
