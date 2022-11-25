# Terraform Provider for [Heroku Advanced Scheduler](https://www.advancedscheduler.io/)

Heavily experimental. Contributions welcome.

## Using the provider

See [examples/main.tf](examples/main.tf).

## Developing the provider

make install
cd examples
export HEROKU_ADVANCED_SCHEDULER_API_TOKEN="Bearer <the token>"
rm .terraform.lock.hcl && terraform init && terraform apply

