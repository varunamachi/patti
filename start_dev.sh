#!/bin/bash

scriptDir="$(cd "$(dirname "$0")" || exit ; pwd -P)"

# Right now both scriptDir and root will be same, but if we have to move the 
# script to somewhere else the argument to readlink can be a path relative to
# script dir and the root will represent the project root
root=$(readlink -f $scriptDir) 


function checkExec() {
    ver="-v"
    if [[ ! -z $2 ]]; then 
        ver="$2"
    fi

    $1 $ver foo >/dev/null 2>&1 ||\
        { echo >&2 "Program '$1' is required for this script to work"; exit 1; }
}

checkExec "docker"
checkExec "docker-compose"
checkExec "tmux" "-V"

tmux new-session -d docker-compose -f "$root/pt/_scripts/db_dc.yml" up
# tmux split-window -v top
# tmux split-window -v docker-compose -f "$root/pt/_scripts/db_dc.yml" up
tmux split-window -v "bash $root/pt/_scripts/serve.sh" 
tmux select-layout even-vertical
tmux attach-session
