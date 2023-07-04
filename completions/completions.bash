_senv() {
    local cur prev opts

    COMPREPLY=()

    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD - 1]}"

    opts="check to ls init --version --help"

    case "${prev}" in
    check)
        return 0
        ;;
    to)
        # COMPREPLY=($(compgen -f ${cur}))
        return 0
        ;;
    ls)
        return 0
        ;;
    init)
        return 0
        ;;
    --version)
        return 0
        ;;
    --help)
        return 0
        ;;
    *) ;;
    esac

    COMPREPLY=($(compgen -W "${opts}" -- ${cur}))

    return 0
}

complete -F _senv senv
