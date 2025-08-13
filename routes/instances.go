// SPDX-License-Identifier: EUPL-1.2

package routes

import (
	"git.zzdats.lv/edim/api-admin/routes/object"
	"git.zzdats.lv/edim/api-admin/routes/request"
	"git.zzdats.lv/edim/api-admin/routes/response"

	"azugo.io/azugo"
	"github.com/nobid-lsp-latvia/go-audit"
)

// @operationId instanceList
// @title Get all instances for person
// @description Method gets list of instances
// @param Person body request.Person true "Person identifier"
// @success 200 InstanceList []response.Instance "Get instance list
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 500 string string "Internal server error"
// @resource Instance
// @route /1.0/instances [post].
func (r *router) instanceList(ctx *azugo.Context) {
	req := &request.Person{}
	if err := ctx.Body.JSON(req); err != nil {
		ctx.Error(err)

		return
	}

	err := r.auditPerson(ctx, nil, nil, &req.PersonCode.PersonCode, nil, audit.ActionSearch)
	if err != nil {
		ctx.Error(err)

		return
	}

	resp := make([]response.Instance, 0)
	if err := r.Store().Exec(ctx, "wallet.get_instance_list", req, &resp); err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(resp)
}

// @operationId instanceSuspend
// @title Suspend instance
// @description Method changes instance status to suspended.
// @param id path string true "Instance identifier"
// @success 200 Instance response.Instance "Get updated instance"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 500 string string "Internal server error"
// @resource Instance
// @route /1.0/instance/{id}/suspend [post].
func (r *router) instanceSuspend(ctx *azugo.Context) {
	id := ctx.Params.String("id")
	action := "suspend"

	var (
		resp *response.Instance
		err  error
	)

	if resp, err = updateInstance(r, ctx, id, action); err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(resp)
}

// @operationId instanceUnsuspend
// @title Unsuspend instance
// @description Method changes instance status to unsuspended.
// @param id path string true "Instance identifier"
// @success 200 Instance response.Instance "Get updated instance"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 500 string string "Internal server error"
// @resource Instance
// @route /1.0/instance/{id}/unsuspend [post].
func (r *router) instanceUnsuspend(ctx *azugo.Context) {
	id := ctx.Params.String("id")
	action := "unsuspend"

	var (
		resp *response.Instance
		err  error
	)

	if resp, err = updateInstance(r, ctx, id, action); err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(resp)
}

// @operationId instanceRevoke
// @title Revokes instance
// @description Method changes instance status to unsuspended.
// @param id path string true "Instance identifier"
// @success 200 Instance response.Instance "Get updated instance"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 500 string string "Internal server error"
// @resource Instance
// @route /1.0/instance/{id}/revoke [post].
func (r *router) instanceRevoke(ctx *azugo.Context) {
	id := ctx.Params.String("id")
	action := "revoke"

	var (
		resp *response.Instance
		err  error
	)

	if resp, err = updateInstance(r, ctx, id, action); err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(resp)
}

func updateInstance(r *router, ctx *azugo.Context, id string, action string) (*response.Instance, error) {
	req := &object.InstanceAction{
		InstanceID: object.InstanceID{ID: id},
		Action:     object.Action{ActionName: action},
	}

	resp := &response.Instance{}
	if err := r.Store().Exec(ctx, "wallet.update_instance_status", req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
