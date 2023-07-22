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

// TenantResourceModel describes the resource data model.
type TenantResourceModel struct {
	Id   uuidtypes.UUID `tfsdk:"id"`
	Name types.String   `tfsdk:"name"`

	//Captcha struct {
	//	Enabled       types.Bool   `tfsdk:"enabled"`
	//	CaptchaMethod types.String `tfsdk:"captcha_method"`
	//	SecretKey     types.String `tfsdk:"secret_key"`
	//	SiteKey       types.String `tfsdk:"site_key"`
	//	Threshold     types.Float64  `tfsdk:"threshold"`
	//} `tfsdk:"captcha"`
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
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	t.client = client
}

func (t *TenantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *TenantResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := fusionauth.TenantRequest{
		SourceTenantId: data.Id.ValueString(),
		Tenant: fusionauth.Tenant{
			//Configured:                     false,
			//ConnectorPolicies:              nil,
			//Data:                           nil,
			//HttpSessionMaxInactiveInterval: 0,
			Id: data.Id.ValueString(),
			//InsertInstant:                  0,
			//Issuer:                         "",
			//LastUpdateInstant:              0,
			//LogoutURL:                      "",
			Name: data.Name.ValueString(),
			//State:                          "",
			//ThemeId:                        "",

			//AccessControlConfiguration: fusionauth.TenantAccessControlConfiguration{
			//	UiIPAccessControlListId: "",
			//},
			//CaptchaConfiguration: fusionauth.TenantCaptchaConfiguration{
			//	Enableable: fusionauth.Enableable{
			//		Enabled: false,
			//	},
			//	CaptchaMethod: "",
			//	SecretKey:     "",
			//	SiteKey:       "",
			//	Threshold:     0,
			//},
			//EmailConfiguration: fusionauth.EmailConfiguration{
			//	AdditionalHeaders:                    nil,
			//	Debug:                                false,
			//	DefaultFromEmail:                     "",
			//	DefaultFromName:                      "",
			//	EmailUpdateEmailTemplateId:           "",
			//	EmailVerifiedEmailTemplateId:         "",
			//	ForgotPasswordEmailTemplateId:        "",
			//	Host:                                 "",
			//	ImplicitEmailVerificationAllowed:     false,
			//	LoginIdInUseOnCreateEmailTemplateId:  "",
			//	LoginIdInUseOnUpdateEmailTemplateId:  "",
			//	LoginNewDeviceEmailTemplateId:        "",
			//	LoginSuspiciousEmailTemplateId:       "",
			//	Password:                             "",
			//	PasswordlessEmailTemplateId:          "",
			//	PasswordResetSuccessEmailTemplateId:  "",
			//	PasswordUpdateEmailTemplateId:        "",
			//	Port:                                 0,
			//	Properties:                           "",
			//	Security:                             "",
			//	SetPasswordEmailTemplateId:           "",
			//	TwoFactorMethodAddEmailTemplateId:    "",
			//	TwoFactorMethodRemoveEmailTemplateId: "",
			//	Unverified: fusionauth.EmailUnverifiedOptions{
			//		AllowEmailChangeWhenGated: false,
			//		Behavior:                  "",
			//	},
			//	Username:                    "",
			//	VerificationEmailTemplateId: "",
			//	VerificationStrategy:        "",
			//	VerifyEmail:                 false,
			//	VerifyEmailWhenChanged:      false,
			//},
			//EventConfiguration: fusionauth.EventConfiguration{
			//	Events: nil,
			//},
			//ExternalIdentifierConfiguration: fusionauth.ExternalIdentifierConfiguration{
			//	AuthorizationGrantIdTimeToLiveInSeconds: 0,
			//	ChangePasswordIdGenerator: fusionauth.SecureGeneratorConfiguration{
			//		Length: 0,
			//		Type:   "",
			//	},
			//	ChangePasswordIdTimeToLiveInSeconds: 0,
			//	DeviceCodeTimeToLiveInSeconds:       0,
			//	DeviceUserCodeIdGenerator: fusionauth.SecureGeneratorConfiguration{
			//		Length: 0,
			//		Type:   "",
			//	},
			//	EmailVerificationIdGenerator: fusionauth.SecureGeneratorConfiguration{
			//		Length: 0,
			//		Type:   "",
			//	},
			//	EmailVerificationIdTimeToLiveInSeconds: 0,
			//	EmailVerificationOneTimeCodeGenerator: fusionauth.SecureGeneratorConfiguration{
			//		Length: 0,
			//		Type:   "",
			//	},
			//	ExternalAuthenticationIdTimeToLiveInSeconds: 0,
			//	OneTimePasswordTimeToLiveInSeconds:          0,
			//	PasswordlessLoginGenerator: fusionauth.SecureGeneratorConfiguration{
			//		Length: 0,
			//		Type:   "",
			//	},
			//	PasswordlessLoginTimeToLiveInSeconds:  0,
			//	PendingAccountLinkTimeToLiveInSeconds: 0,
			//	RegistrationVerificationIdGenerator: fusionauth.SecureGeneratorConfiguration{
			//		Length: 0,
			//		Type:   "",
			//	},
			//	RegistrationVerificationIdTimeToLiveInSeconds: 0,
			//	RegistrationVerificationOneTimeCodeGenerator: fusionauth.SecureGeneratorConfiguration{
			//		Length: 0,
			//		Type:   "",
			//	},
			//	Samlv2AuthNRequestIdTimeToLiveInSeconds: 0,
			//	SetupPasswordIdGenerator: fusionauth.SecureGeneratorConfiguration{
			//		Length: 0,
			//		Type:   "",
			//	},
			//	SetupPasswordIdTimeToLiveInSeconds: 0,
			//	TrustTokenTimeToLiveInSeconds:      0,
			//	TwoFactorIdTimeToLiveInSeconds:     0,
			//	TwoFactorOneTimeCodeIdGenerator: fusionauth.SecureGeneratorConfiguration{
			//		Length: 0,
			//		Type:   "",
			//	},
			//	TwoFactorOneTimeCodeIdTimeToLiveInSeconds:          0,
			//	TwoFactorTrustIdTimeToLiveInSeconds:                0,
			//	WebAuthnAuthenticationChallengeTimeToLiveInSeconds: 0,
			//	WebAuthnRegistrationChallengeTimeToLiveInSeconds:   0,
			//},
			//FailedAuthenticationConfiguration: fusionauth.FailedAuthenticationConfiguration{
			//	ActionCancelPolicy: fusionauth.FailedAuthenticationActionCancelPolicy{
			//		OnPasswordReset: false,
			//	},
			//	ActionDuration:      0,
			//	ActionDurationUnit:  "",
			//	EmailUser:           false,
			//	ResetCountInSeconds: 0,
			//	TooManyAttempts:     0,
			//	UserActionId:        "",
			//},
			//FamilyConfiguration: fusionauth.FamilyConfiguration{
			//	Enableable: fusionauth.Enableable{
			//		Enabled: false,
			//	},
			//	AllowChildRegistrations:           false,
			//	ConfirmChildEmailTemplateId:       "",
			//	DeleteOrphanedAccounts:            false,
			//	DeleteOrphanedAccountsDays:        0,
			//	FamilyRequestEmailTemplateId:      "",
			//	MaximumChildAge:                   0,
			//	MinimumOwnerAge:                   0,
			//	ParentEmailRequired:               false,
			//	ParentRegistrationEmailTemplateId: "",
			//},
			//FormConfiguration: fusionauth.TenantFormConfiguration{
			//	AdminUserFormId: "",
			//},
			//JwtConfiguration: fusionauth.JWTConfiguration{
			//	Enableable: fusionauth.Enableable{
			//		Enabled: false,
			//	},
			//	AccessTokenKeyId:             "",
			//	IdTokenKeyId:                 "",
			//	RefreshTokenExpirationPolicy: "",
			//	RefreshTokenRevocationPolicy: fusionauth.RefreshTokenRevocationPolicy{
			//		OnLoginPrevented:    false,
			//		OnMultiFactorEnable: false,
			//		OnPasswordChanged:   false,
			//	},
			//	RefreshTokenTimeToLiveInMinutes: 0,
			//	RefreshTokenUsagePolicy:         "",
			//	TimeToLiveInSeconds:             0,
			//},
			//LambdaConfiguration: fusionauth.TenantLambdaConfiguration{
			//	ScimEnterpriseUserRequestConverterId:  "",
			//	ScimEnterpriseUserResponseConverterId: "",
			//	ScimGroupRequestConverterId:           "",
			//	ScimGroupResponseConverterId:          "",
			//	ScimUserRequestConverterId:            "",
			//	ScimUserResponseConverterId:           "",
			//},
			//LoginConfiguration: fusionauth.TenantLoginConfiguration{
			//	RequireAuthentication: false,
			//},
			//MaximumPasswordAge: fusionauth.MaximumPasswordAge{
			//	Enableable: fusionauth.Enableable{
			//		Enabled: false,
			//	},
			//	Days: 0,
			//},
			//MinimumPasswordAge: fusionauth.MinimumPasswordAge{
			//	Enableable: fusionauth.Enableable{
			//		Enabled: false,
			//	},
			//	Seconds: 0,
			//},
			//MultiFactorConfiguration: fusionauth.TenantMultiFactorConfiguration{
			//	Authenticator: fusionauth.MultiFactorAuthenticatorMethod{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		Algorithm:  "",
			//		CodeLength: 0,
			//		TimeStep:   0,
			//	},
			//	Email: fusionauth.MultiFactorEmailMethod{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		TemplateId: "",
			//	},
			//	LoginPolicy: "",
			//	Sms: fusionauth.MultiFactorSMSMethod{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		MessengerId: "",
			//		TemplateId:  "",
			//	},
			//},
			//OauthConfiguration: fusionauth.TenantOAuth2Configuration{
			//	ClientCredentialsAccessTokenPopulateLambdaId: "",
			//},
			//PasswordEncryptionConfiguration: fusionauth.PasswordEncryptionConfiguration{
			//	EncryptionScheme:              "",
			//	EncryptionSchemeFactor:        0,
			//	ModifyEncryptionSchemeOnLogin: false,
			//},
			//PasswordValidationRules: fusionauth.PasswordValidationRules{
			//	BreachDetection: fusionauth.PasswordBreachDetection{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		MatchMode:                 "",
			//		NotifyUserEmailTemplateId: "",
			//		OnLogin:                   "",
			//	},
			//	MaxLength: 0,
			//	MinLength: 0,
			//	RememberPreviousPasswords: fusionauth.RememberPreviousPasswords{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		Count: 0,
			//	},
			//	RequireMixedCase: false,
			//	RequireNonAlpha:  false,
			//	RequireNumber:    false,
			//	ValidateOnLogin:  false,
			//},
			//RateLimitConfiguration: fusionauth.TenantRateLimitConfiguration{
			//	FailedLogin: fusionauth.RateLimitedRequestConfiguration{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		Limit:               0,
			//		TimePeriodInSeconds: 0,
			//	},
			//	ForgotPassword: fusionauth.RateLimitedRequestConfiguration{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		Limit:               0,
			//		TimePeriodInSeconds: 0,
			//	},
			//	SendEmailVerification: fusionauth.RateLimitedRequestConfiguration{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		Limit:               0,
			//		TimePeriodInSeconds: 0,
			//	},
			//	SendPasswordless: fusionauth.RateLimitedRequestConfiguration{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		Limit:               0,
			//		TimePeriodInSeconds: 0,
			//	},
			//	SendRegistrationVerification: fusionauth.RateLimitedRequestConfiguration{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		Limit:               0,
			//		TimePeriodInSeconds: 0,
			//	},
			//	SendTwoFactor: fusionauth.RateLimitedRequestConfiguration{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		Limit:               0,
			//		TimePeriodInSeconds: 0,
			//	},
			//},
			//RegistrationConfiguration: fusionauth.TenantRegistrationConfiguration{
			//	BlockedDomains: nil,
			//},
			//ScimServerConfiguration: fusionauth.TenantSCIMServerConfiguration{
			//	Enableable: fusionauth.Enableable{
			//		Enabled: false,
			//	},
			//	ClientEntityTypeId: "",
			//	Schemas:            nil,
			//	ServerEntityTypeId: "",
			//},
			//SsoConfiguration: fusionauth.TenantSSOConfiguration{
			//	DeviceTrustTimeToLiveInSeconds: 0,
			//},
			//UserDeletePolicy: fusionauth.TenantUserDeletePolicy{
			//	Unverified: fusionauth.TimeBasedDeletePolicy{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		NumberOfDaysToRetain: 0,
			//	},
			//},
			//UsernameConfiguration: fusionauth.TenantUsernameConfiguration{
			//	Unique: fusionauth.UniqueUsernameConfiguration{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		NumberOfDigits: 0,
			//		Separator:      "",
			//		Strategy:       "",
			//	},
			//},
			//WebAuthnConfiguration: fusionauth.TenantWebAuthnConfiguration{
			//	Enableable: fusionauth.Enableable{
			//		Enabled: false,
			//	},
			//	BootstrapWorkflow: fusionauth.TenantWebAuthnWorkflowConfiguration{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		AuthenticatorAttachmentPreference: "",
			//		UserVerificationRequirement:       "",
			//	},
			//	Debug: false,
			//	ReauthenticationWorkflow: fusionauth.TenantWebAuthnWorkflowConfiguration{
			//		Enableable: fusionauth.Enableable{
			//			Enabled: false,
			//		},
			//		AuthenticatorAttachmentPreference: "",
			//		UserVerificationRequirement:       "",
			//	},
			//	RelyingPartyId:   "",
			//	RelyingPartyName: "",
			//},
		},
		//WebhookIds: nil,
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

func buildTenant(data *TenantResourceModel, t fusionauth.Tenant) {
	data.Id = uuidtypes.NewUUIDValue(t.Id)
	data.Name = types.StringValue(t.Name)
}

func (t *TenantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *TenantResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res, faErrs, err := t.client.API.RetrieveTenant(data.Id.ValueString())
	if reportedReadErrors(resp.Diagnostics, faErrs, err, "tenant") {
		tflog.Trace(ctx, "error attempting to read a tenant resource")
		return
	}
	tflog.Trace(ctx, "successfully read tenant resource ID: "+data.Id.ValueString())

	buildTenant(data, res.Tenant)
	tflog.Trace(ctx, "successfully converted read tenant response to domain model")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "successfully saved read tenant into state")
}

func (t *TenantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *TenantResourceModel

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
	var data *TenantResourceModel

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
