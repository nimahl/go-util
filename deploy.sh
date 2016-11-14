#!/usr/bin/env bash

SCRIPTPATH=$(pwd)

if [ -z ${1+x} ]; then
  echo "Usage: ./deploy all <stage> | deploy <service> <stage>";
  exit 1
fi

if [ ${1} = "all" ]; then
  cd apex
  functions=$(find ./functions -maxdepth 1 -type d | cut -c 13- | tail -n +2)
  for d in ${functions} ; do
      echo "Building ${d}"
      apex build ${d} > ${d}.zip
  done
  cd ${SCRIPTPATH}
else
  cd apex
  echo "Building $1"
  apex build ${1} > ${1}.zip
  cd ${SCRIPTPATH}
fi


rm -r ${SCRIPTPATH}/.serverless/
sls deploy -s ${2}