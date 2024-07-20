#!/bin/bash

set -e

DIR=$(dirname "$0")
CONFIG=${DIR}/../../config/settings.json
CONFIG_DIST=${DIR}/../../config.dist/settings.json.j2

mkdir -p "${CONFIG%/*}"

if [ ! -e ${CONFIG} ]; then
  cat "${CONFIG_DIST}" | sed 's/{{[[:space:]]*\([^[:space:]]*\)[[:space:]]*}}/${\1}/g' | envsubst > "${CONFIG}"
fi
