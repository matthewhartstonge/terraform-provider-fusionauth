/*
 * Copyright (c) Matthew Hartstonge <matt@mykro.co.nz>
 * SPDX-License-Identifier: MPL-2.0
 */

package provider

import (
	// Standard Library Imports
	"context"
	"fmt"

	// External Imports
	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/matthewhartstonge/terraform-plugin-framework-type-uuid/uuidtypes"

	// Internal Imports
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
		MarkdownDescription: "FusionAuth Tenant",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				CustomType:          uuidtypes.UUIDType{},
				Computed:            true,
				Description:         "The unique identifier for this Tenant.",
				MarkdownDescription: "The unique identifier for this Tenant.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The unique name of the Tenant.",
				MarkdownDescription: "The unique name of the Tenant.",
			},
			//"captcha": schema.ListNestedAttribute{
			//	NestedObject:        schema.NestedAttributeObject{},
			//	Optional:            true,
			//	Description:         "",
			//	MarkdownDescription: "",
			//	DeprecationMessage:  "",
			//	Validators:          nil,
			//	PlanModifiers:       nil,
			//	Default:             nil,
			//},
			//"email": schema.ListNestedAttribute{
			//	NestedObject:        schema.NestedAttributeObject{},
			//	Optional:            true,
			//	Description:         "",
			//	MarkdownDescription: "",
			//	DeprecationMessage:  "",
			//	Validators:          nil,
			//	PlanModifiers:       nil,
			//	Default:             nil,
			//},
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
	var data *tenantResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tenant := fusionauth.Tenant{
		Id:   data.ID.ValueString(),
		Name: data.Name.ValueString(),
	}

	payload := fusionauth.TenantRequest{
		SourceTenantId: data.ID.ValueString(),
		Tenant:         tenant,
	}

	res, faErrs, err := t.client.API.CreateTenant(t.client.TenantID, payload)
	if reportedCreateErrors(resp.Diagnostics, faErrs, err, "tenant") {
		tflog.Trace(ctx, "error attempting to create a tenant resource")
		return
	}
	tflog.Trace(ctx, "successfully created a tenant resource")

	buildTenant(data, res.Tenant)
	tflog.Trace(ctx, "successfully converted created tenant response to domain model")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "successfully saved created tenant into state")
}

func buildTenant(data *tenantResourceModel, t fusionauth.Tenant) {
	data.ID = uuidtypes.NewUUIDValue(t.Id)
	data.Name = types.StringValue(t.Name)
}

func (t *TenantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *tenantResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res, faErrs, err := t.client.API.RetrieveTenant(data.ID.ValueString())
	if reportedReadErrors(resp.Diagnostics, faErrs, err, "tenant") {
		tflog.Trace(ctx, "error attempting to read a tenant resource")
		return
	}
	tflog.Trace(ctx, "successfully read tenant resource ID: "+data.ID.ValueString())

	buildTenant(data, res.Tenant)
	tflog.Trace(ctx, "successfully converted read tenant response to domain model")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "successfully saved read tenant into state")
}

func (t *TenantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *tenantResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (t *TenantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *tenantResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}

func (t *TenantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
