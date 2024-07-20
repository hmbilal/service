#!/bin/bash

while getopts ":t:" option; do
  case $option in
  t)
    tag_name="$OPTARG"
    ;;
  *)
    echo "Usage: $0 [-t tag_name]"
    exit 1
    ;;
  esac
done

if [ -z "$tag_name" ]; then
  tag_name="1.0.0"
fi

echo "Tag name: $tag_name"

docker build -t registry.digitalarsenal.net/elpaso/backend/activities/app:latest docker/images/app
docker build -t registry.digitalarsenal.net/elpaso/backend/activities/app:$tag_name docker/images/app

docker push registry.digitalarsenal.net/elpaso/backend/activities/app:latest
docker push registry.digitalarsenal.net/elpaso/backend/activities/app:$tag_name
