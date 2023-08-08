/*
 * Copyright (c) Matthew Hartstonge <matt@mykro.co.nz>
 * SPDX-License-Identifier: MPL-2.0
 */

package provider

import (
	"context"
	"fmt"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/matthewhartstonge/terraform-plugin-framework-type-uuid/uuidtypes"

	faClient "github.com/matthewhartstonge/terraform-provider-fusionauth/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &TenantResource{}
	_ resource.ResourceWithImportState = &TenantResource{}
)

func NewTenantResource() resource.Resource {
	return &TenantResource{}
}

// TenantResource defines the resource implementation.
type TenantResource struct {
	client *faClient.Client
}

func (t *TenantResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant"
}

func (t *TenantResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description:         "FusionAuth Tenant",
		MarkdownDescription: "A FusionAuth Tenant is a named object that represents a discrete namespace for Users, Applications and Groups. A user is unique by email address or username within a tenant.\n\nTenants may be useful to support a multi-tenant application where you wish to use a single instance of FusionAuth but require the ability to have duplicate users across the tenants in your own application. In this scenario a user may exist multiple times with the same email address and different passwords across tenants.\n\nTenants may also be useful in a test or staging environment to allow multiple users to call APIs and create and modify users without possibility of collision.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				CustomType:          uuidtypes.UUIDType{},
				Description:         "The unique identifier for this Tenant.",
				MarkdownDescription: "The unique identifier for this Tenant.",
				Computed:            true,
			},
			"configured": schema.BoolAttribute{
				Computed:            true,
				Description:         "Indicates the tenant has been configured. It is always true, except for default tenant when the setup wizard has not been completed, in which case it is false.",
				MarkdownDescription: "Indicates the tenant has been configured. It is always `true`, except for default tenant when the setup wizard has not been completed, in which case it is `false`.",
			},
			"name": schema.StringAttribute{
				Description:         "The unique name of the Tenant.",
				MarkdownDescription: "The unique name of the Tenant.",
				Required:            true,
			},
			"captcha": tenantCaptchaModelSchema(),
		},
	}
}

func (t *TenantResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*faClient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *faClient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	t.client = client
}

func (t *TenantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan *tenantResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := fusionauth.TenantRequest{
		Tenant: newTenant(plan),
	}

	res, faErrs, err := t.client.API.CreateTenant(t.client.TenantID, payload)
	if reportedCreateErrors(resp.Diagnostics, faErrs, err, "tenant") {
		tflog.Trace(ctx, "error attempting to create a tenant resource")
		return
	}
	tflog.Trace(ctx, "successfully created a tenant resource")

	setTenantState(plan, res.Tenant)
	tflog.Trace(ctx, "successfully converted created tenant response to domain model")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	tflog.Trace(ctx, "successfully saved created tenant into state")
}

func (t *TenantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *tenantResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res, faErrs, err := t.client.API.RetrieveTenant(state.ID.ValueString())
	if reportedReadErrors(resp.Diagnostics, faErrs, err, "tenant") {
		tflog.Trace(ctx, "error attempting to read a tenant resource")
		return
	}
	tflog.Trace(ctx, "successfully read tenant resource ID: "+state.ID.ValueString())

	setTenantState(state, res.Tenant)
	tflog.Trace(ctx, "successfully converted read tenant response to domain model")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	tflog.Trace(ctx, "successfully saved read tenant into state")
}

func (t *TenantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan *tenantResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := fusionauth.TenantRequest{
		Tenant: newTenant(plan),
	}

	res, faErrs, err := t.client.API.UpdateTenant(plan.ID.ValueString(), payload)
	if reportedUpdateErrors(resp.Diagnostics, faErrs, err, "tenant") {
		tflog.Trace(ctx, "error attempting to read a tenant resource")
		return
	}
	tflog.Trace(ctx, "successfully updated tenant resource ID: "+plan.ID.ValueString())

	setTenantState(plan, res.Tenant)
	tflog.Trace(ctx, "successfully converted updated tenant response to domain model")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	tflog.Trace(ctx, "successfully saved updated tenant into state")
}

func (t *TenantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *tenantResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}
	_, faErrs, err := t.client.API.DeleteTenant(state.ID.ValueString())
	if reportedDeleteErrors(resp.Diagnostics, faErrs, err, "tenant") {
		tflog.Trace(ctx, "error attempting to delete a tenant resource")
		return
	}

	tflog.Trace(ctx, "successfully deleted tenant resource ID: "+state.ID.ValueString())
}

func (t *TenantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
