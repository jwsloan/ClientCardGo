# ADR 004: Load All Secrets from Environment Variables

## Status

Accepted

## Context

Hardcoding secrets or storing them in config files risks leaks and is not cloud-native. Using environment variables is the standard for secure secret management in 12-factor and containerized apps.

## Decision

- All secrets (database passwords, API keys, JWT secrets, etc.) are loaded from environment variables.
- No secrets are ever checked into code, config, or migrations.
- README and .env.example are explicit about this policy.
- For production, use a secrets manager (GCP Secret Manager, AWS Secrets Manager, etc.) to inject env vars.

## Consequences

- Reduced risk of accidental secret leaks.
- Easier rotation and management of sensitive credentials.
- Compatible with all major cloud deployment platforms.