
# Senv

![go-ci](https://github.com/luisnquin/senv/actions/workflows/go.yml/badge.svg)
<img alt="built-with-nix" src="https://builtwithnix.org/badge.svg" width="80px">

## Why?

It's annoying to have to manually update each **environment variable** or rename **.env** files in order to change to specific configurations for a project.
The problem is multiplied if you are working on more than a single project at the same time so the idea here is to have a **senv.yaml** in each one of your
projects when you think that you're dealing with multiple **.env** settings.

You can call the program in any subdirectory of your project and the program will automatically find the **senv.yaml**.

It's highly inspired by [VSCode - .ENV Switcher](https://marketplace.visualstudio.com/items?itemName=EcksDy.env-switcher) but only deals with a single
configuration file and can be called directly from the command line.

## Demo

[![demo](https://asciinema.org/a/eZrIbb4eDxX0tO7fWyFop2Zg8.svg)](https://asciinema.org/a/eZrIbb4eDxX0tO7fWyFop2Zg8)

## Install

### Via Go

```bash
# Requires go >=1.18
$ go install github.com/luisnquin/senv@latest
```

### Via Nix flakes

```nix
{
  inputs = {
    senv.url = "github:luisnquin/senv";
  };

  outputs = inputs @ {
    senv,
    ...
  }: let
    system = "x86_64-linux";

    specialArgs = {
      senv-switcher = senv.defaultPackage.${system};
    };
  in {
    # Home manager
    homeConfigurations."..." = home-manager.lib.homeManagerConfiguration {
      extraSpecialArgs = specialArgs;

      modules = [
        # add "senv-switcher" to your home.packages
        ./home.nix
      ];
    };

    # Or NixOS configuration
    nixosConfigurations."..." = lib.nixosSystem {
        specialArgs = specialArgs;

        modules = [
          # add "senv-switcher" to your environment.systemPackages
          ./configuration.nix
        ];
      };
    };
}

```

## Try it with Nix ❄️

```bash
# Creates a senv.yaml file in your current folder
$ nix run github:luisnquin/senv -- init

# Lists your declared environments
$ nix run github:luisnquin/senv -- ls

# Run a prompt that allows you to select an environment
$ nix run github:luisnquin/senv

# Help
$ nix run github:luisnquin/senv -- --help

# ...
```

## Settings file

The program requires a **senv.yaml** or **senv.yml** file and you can call the program in any subdirectory of the project.

### Example

Suppose we have a **senv.yaml** file like this:

```yaml
# senv.yaml
envFile: ./app/.env # optional absolute/relative path
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
  DATABASE_USER: admin # both variables will be added to the
  DATABASE_PORT: 5432 # selected environment if not declared
useExportPrefix: false # optional
```

When the .env file is generated from **preprod** it will look like this:

```bash
#_preprod_#

DATABASE_PASSWORD="root123"
DATABASE_HOST="localhost"
DATABASE_USER="admin"
DATABASE_PORT="5432"
```

## Completions

### Supported shells

- zsh

```bash
 # Add this line to your .zshrc file
 $ source <(senv completion zsh)
```

- bash

```bash
 # Add this line to your .bashrc file
  $ source <(senv completion bash)
```

## Integrations

### [Starship](https://starship.rs/)

```toml
# starship.toml
format = "${custom.environment_name}" # more...

[custom.environment_name]
command = "senv out"
description = "Displays the name of your current senv environment"
format = "using [($output )]($style)"
shell = ["bash", "--noprofile", "--norc"]
style = "#a8e046"
when = "senv out"
```

## LICENSE

[MIT](../LICENSE)
