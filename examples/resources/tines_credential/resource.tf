resource "tines_credential" "example_text_credential" {
  name    = "Example API Key"
  team_id = 1
  value   = "example-secret-value"
}

# Secret rotation example: wire the value to an upstream secret (e.g. a rotated
# token from your CI/CD pipeline or another Terraform resource). When the value
# changes, Terraform updates the Tines Credential in place during apply.
resource "tines_credential" "rotated_token" {
  name    = "sirtmanager_gitlab-sirt_token"
  mode    = "TEXT"
  team_id = 2
  value   = var.rotated_gitlab_pat
}
