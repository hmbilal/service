#!/usr/bin/env bash

go_file_changes=$(git diff --name-only origin/master ${CI_COMMIT_SHA} -- $(git diff --name-only origin/master...${CI_COMMIT_SHA}) | xargs -0 | grep -v "docker"| grep ".go")
[[ -z "$go_file_changes" ]] && echo "No go file changes" && exit 0

go_dir_changes=$(dirname $go_file_changes | sort | uniq | awk -F: '{if(system("[ ! -d " $0 " ]") != 0) {print $0}}')
echo $go_dir_changes
golangci-lint run $go_dir_changes