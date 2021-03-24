#!/bin/bash

scriptDir="$(cd "$(dirname "$0")" || exit ; pwd -P)"

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

docker-compose -f "${scriptDir}/db_dc.yml" up