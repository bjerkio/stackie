# Stackie

<div style="text-align: center">
    <img src="https://raw.githubusercontent.com/bjerkio/stackie/main/.github/logo.svg" alt="Stackie" height="200px" />
    <br /><br />
</div>

[![lifecycle](https://img.shields.io/badge/lifecycle-experimental-orange.svg)](https://www.tidyverse.org/lifecycle/#experimental)

**Stackie enables developers to configure their local environment/toolchain with
ease. Made for Pulumi CLI, Google Cloud Platform (`gcloud`), and Amazon Web
Services (`aws-cli`).**

Stackie changes your `pulumi`, `gcloud` and `aws-cli` actively when you change
from project to project. Think of it as a combination of `aws-cli` and
[`nvm`](https://nvm.sh)/[`volta`](https://volta.sh/).

### Quickstart

Imagine your in your one of your personal projects. You're using
[Pulumi](https://pulumi.com/).

```shell
▶ pulumi whoami
cobraz
```

You're going to start working on that Yosemite project, but instead of logging
out of Pulumi and logging into their Pulumi account with their credentials.
Stackie has that all stored in a configuration, so all you have to do is this:

```
▶ cd ~/projects/yosemite
▶ pulumi whoami
yosemite-system-user
```

## Stackie Config

There are three possible configuration files.

```yaml
pulumi:
  cloudUrl: gs://hello-world
  stackName: bjerk/prod
```

### Project-specific configuration (`.stackie.yml`)

This configuration is probably shared with the team and pushed to the root
directory of the repository (or within the `infra` folder maybe).

### Project and Developer-specific configuration (`.stackie.local.yml`)

The personal Stackie configuration is used to store project-specific access
tokens or setups that you need. This file should probably never be pushed to
your repository.

### Developer-specific configuration (`~/.config/stackie/config.yml`)

This is where your secrets are stored away. It could look like this:

```yaml
profiles:
  - name: bobbafett
    pulumi:
      accessToken: my-token-here
  - name: fintech-company
    google:
      configuration: fintech-company
      export-application-default: true
```
