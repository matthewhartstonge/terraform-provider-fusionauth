/*
 * Copyright (c) Matthew Hartstonge <matt@mykro.co.nz>
 * SPDX-License-Identifier: MPL-2.0
 */

package provider

import (
	"fmt"
	"strings"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/matthewhartstonge/terraform-plugin-framework-type-uuid/uuidtypes"
)

func newTenant(data *tenantResourceModel) fusionauth.Tenant {
	return fusionauth.Tenant{
		Id:   data.ID.ValueString(),
		Name: data.Name.ValueString(),
		CaptchaConfiguration: fusionauth.TenantCaptchaConfiguration{
			Enableable: fusionauth.Enableable{
				Enabled: data.Captcha.Enabled.ValueBool(),
			},
			CaptchaMethod: fusionauth.CaptchaMethod(data.Captcha.Method.ValueString()),
			SecretKey:     data.Captcha.SecretKey.ValueString(),
			SiteKey:       data.Captcha.SiteKey.ValueString(),
			Threshold:     data.Captcha.Threshold.ValueFloat64(),
		},
	}
}

func setTenantState(data *tenantResourceModel, t fusionauth.Tenant) {
	data.ID = uuidtypes.NewUUIDValue(t.Id)
	data.Configured = types.BoolValue(t.Configured)
	data.Name = types.StringValue(t.Name)

	data.Captcha.Enabled = types.BoolValue(t.CaptchaConfiguration.Enabled)
	data.Captcha.Method = types.StringValue(t.CaptchaConfiguration.CaptchaMethod.String())
	data.Captcha.SecretKey = types.StringValue(t.CaptchaConfiguration.SecretKey)
	data.Captcha.SiteKey = types.StringValue(t.CaptchaConfiguration.SiteKey)
	data.Captcha.Threshold = types.Float64Value(t.CaptchaConfiguration.Threshold)
}

// tenantResourceModel describes the resource data model.
type tenantResourceModel struct {
	ID         uuidtypes.UUID `tfsdk:"id"`
	Configured types.Bool     `tfsdk:"configured"`
	Name       types.String   `tfsdk:"name"`

	Captcha tenantCaptchaModel `tfsdk:"captcha"`

	//AccessControl
	//ConnectorPolicies
	//Data
	//Email
	//Event
	//ExternalIdentifier
	//FailedAuthentication
	//Family
	//Form
	//HttpSessionMaxInactiveInterval
	//InsertInstant
	//Issuer
	//JWT
	//Lambda
	//LastUpdateInstant
	//Login
	//LogoutURL
	//MaximumPasswordAge
	//MinimumPasswordAge
	//MultiFactor
	//OAuth
	//PasswordEncryption
	//PasswordValidationRules
	//RateLimit
	//Registration
	//ScimServer
	//SSO
	//State
	//ThemeID
	//UserDeletePolicy
	//Username
	//WebAuthn
}

// tenantCaptchaModel maps tenant captcha configuration.
type tenantCaptchaModel struct {
	Enabled   types.Bool    `tfsdk:"enabled"`
	Method    types.String  `tfsdk:"method"`
	SecretKey types.String  `tfsdk:"secret_key"`
	SiteKey   types.String  `tfsdk:"site_key"`
	Threshold types.Float64 `tfsdk:"threshold"`
}

func tenantCaptchaModelSchema() schema.Attribute {
	captchaMethods := []string{
		fusionauth.CaptchaMethod_GoogleRecaptchaV2.String(),
		fusionauth.CaptchaMethod_GoogleRecaptchaV3.String(),
		fusionauth.CaptchaMethod_HCaptcha.String(),
		fusionauth.CaptchaMethod_HCaptchaEnterprise.String(),
	}

	return schema.SingleNestedAttribute{
		Optional:            true,
		Description:         "Provides CAPTCHA configuration for the tenant.",
		MarkdownDescription: "Provides CAPTCHA configuration for the tenant.",

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Description:         "Whether captcha configuration is enabled.",
				MarkdownDescription: "Whether captcha configuration is enabled.",
			},
			"method": schema.StringAttribute{
				Optional:            true,
				Description:         "The type of captcha method to use.",
				MarkdownDescription: fmt.Sprintf("The type of captcha method to use. Must be One Of {%s}", strings.Join(captchaMethods, ", ")),
				Validators: []validator.String{
					stringvalidator.OneOf(captchaMethods...),
				},
			},
			"secret_key": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				Description:         "The secret key for this captcha method.",
				MarkdownDescription: "The secret key for this captcha method.",
			},
			"site_key": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				Description:         "The site key for this captcha method.",
				MarkdownDescription: "The site key for this captcha method.",
			},
			"threshold": schema.Float64Attribute{
				Optional:            true,
				Description:         "The numeric threshold which separates a passing score from a failing one. This value only applies if using either the Google v3 or HCaptcha Enterprise method, otherwise this value is ignored.",
				MarkdownDescription: "The numeric threshold which separates a passing score from a failing one. This value only applies if using either the Google v3 or HCaptcha Enterprise method, otherwise this value is ignored.",
			},
		},
	}
}
