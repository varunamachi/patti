#!/bin/bash

scriptDir="$(cd "$(dirname "$0")" || exit ; pwd -P)"
root=$(readlink -f $scriptDir/..)

i=0
while ! nc -z localhost 5432; do   
  sleep 1 
  if [[ i == 1 ]]; then 
    echo "Waiting for postgres server at port 5432."
  fi
#   if [[ i > 7 ]]; then 
#     echo "Waited for postgres for 2 seconds, now terminating"
#     sleep 2
#     exit -1
#   fi
  echo "$i"
  ((i+=1))
done

"$root/_scripts/run.sh" "pt" "serve"