
# Senv

## Install

```bash
# Requires go >=1.20
go install github.com/luisnquin/senv@latest
```

## Configuration

At this moment, it requires an `senv.yaml` or `senv.yml` file in your working directory, you can call the program in any subdirectory, it automatically searches for the configuration file recursively searching in the parent folders.  

### File example

```yaml
envs:
- name: dev
    variables:
      DATABASE_USER: admin
      DATABASE_PASSWORD: pwd123
      DATABASE_HOST: localhost
- name: dev-2
    variables:
      DATABASE_USER: admin
      DATABASE_PASSWORD: pwd321
      DATABASE_HOST: test-host
- name: prod
    variables:
      DATABASE_USER: admin
      DATABASE_PASSWORD: test-password
      DATABASE_HOST: test-host
- name: preprod
    variables:
      DATABASE_PASSWORD: root123
      DATABASE_HOST: localhost
defaults:
  DATABASE_USER: admin
  DATABASE_PORT: 5432
```
