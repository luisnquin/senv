_senv() {
    local -a sub_commands

    sub_commands=(
        'check:check if the current working directory has `senv.yaml` or `senv.yml` files'
        'to:switch to other environment without a prompt'
        'ls:list all the environments in the working directory'
        'export:print resolved variables for an environment to stdout'
        'init:creates a new configuration file in the current directory'
    )

    _describe 'senv' sub_commands
}

compdef _senv senv
