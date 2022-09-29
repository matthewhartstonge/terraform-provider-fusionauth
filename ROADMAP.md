# Roadmap

FusionAuth provides many resources and APIs which all need to be documented and built in Terraform's expected data
schema.

## Resources

Based on FusionAuth 1.36.8 the following API Endpoints are available. This does not mean that each endpoint provides a
manageable resource. As such, each endpoint will need to be cross-referenced with the API documentation to check if it
provides a candidate to come under the management of terraform.

| Legend             | Description                                                                                      |
|--------------------|--------------------------------------------------------------------------------------------------|
| :heavy_check_mark: | Yes. Can be developed.                                                                           |
| :x:                | No: Either not technically possible, or provides no benefit to being under Terraform management. |
| :microscope:       | Yet to be investigated and documented.                                                           |
| :question_mark:    | To be decided.                                                                                   |

| Endpoint                                | Supported Methods | Resource Candidate | Data Source Candidate | API Documentation Link                                                            |
|-----------------------------------------|-------------------|--------------------|-----------------------|-----------------------------------------------------------------------------------|
| `/api/application`                      | `C R U D`         | :heavy_check_mark: | :heavy_check_mark:    | https://fusionauth.io/docs/v1/tech/apis/applications                              |
| `/api/application/oauth-configuration`  | `- R - -`         | :x:                | :question_mark:       | https://fusionauth.io/docs/v1/tech/apis/applications#retrieve-oauth-configuration |
| `/api/application/role`                 | `C - U D`         | :heavy_check_mark: | :x:                   | https://fusionauth.io/docs/v1/tech/apis/applications#create-an-application-role   |
| `/api/cleanspeak/notify`                | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/connector`                        | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/consent`                          | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/email/send`                       | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/email/template`                   | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/email/template/preview`           | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/entity`                           | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/entity/grant`                     | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/entity/grant/search`              | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/entity/search`                    | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/entity/type`                      | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/entity/type/permission`           | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/entity/type/search`               | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/form`                             | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/form/field`                       | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/group`                            | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/group/member`                     | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/group/member/search`              | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/identity-provider`                | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/identity-provider/link`           | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/integration`                      | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/ip-acl`                           | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/ip-acl/search`                    | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/jwt/issue`                        | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/jwt/refresh`                      | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/jwt/validate`                     | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/jwt/vend`                         | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/key`                              | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/key/generate`                     | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/key/import`                       | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/lambda`                           | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/logger`                           | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/login`                            | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/message/template`                 | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/message/template/preview`         | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/messenger`                        | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/passwordless/start`               | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/prometheus/metrics`               | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/reactor`                          | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/reactor/metrics`                  | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/report/daily-active-user`         | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/report/login`                     | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/report/monthly-active-user`       | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/report/registration`              | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/report/totals`                    | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/scim/resource/v2/EnterpriseUsers` | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/scim/resource/v2/Groups`          | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/scim/resource/v2/Users`           | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/status`                           | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system-configuration`             | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system/audit-log`                 | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system/audit-log/export`          | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system/audit-log/search`          | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system/event-log`                 | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system/event-log/search`          | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system/log/export`                | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system/login-record/export`       | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system/login-record/search`       | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system/reindex`                   | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/system/version`                   | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/tenant`                           | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/theme`                            | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/two-factor/secret`                | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/two-factor/send`                  | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/two-factor/start`                 | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user`                             | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user-action`                      | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user-action-reason`               | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/action`                      | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/bulk`                        | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/change-password`             | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/comment`                     | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/consent`                     | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/family`                      | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/family/pending`              | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/family/request`              | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/forgot-password`             | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/import`                      | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/recent-login`                | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/refresh-token/import`        | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/registration`                | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/search`                      | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/two-factor`                  | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/two-factor/recovery-code`    | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/verify-email`                | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/user/verify-registration`         | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
| `/api/webhook`                          | `C R U D`         | :microscope:       | :microscope:          |                                                                                   |
