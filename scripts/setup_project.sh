#!/bin/bash

# Default values
DEFAULT_PLACEHOLDER="__PLACEHOLDER__"

# Command line arguments
PLACEHOLDER=${1:-$DEFAULT_PLACEHOLDER}
PROJECT_NAME=${2}

# Function to replace placeholders in files
replace_placeholders() {
  local placeholder=$1
  local project_name=$2

  # Find all files and replace placeholders
  find . -type f -exec sed -i "s/$placeholder/$project_name/g" {} +
}

# Check if correct number of arguments are provided
if [ "$#" -ne 1 ]; then
  echo "Value for PROJECT_NAME=$PROJECT_NAME is required!"
  exit 1
fi

echo "Using: PLACEHOLDER=$PLACEHOLDER, PROJECT_NAME=$PROJECT_NAME"

# Replace placeholders with project name
replace_placeholders "$PLACEHOLDER" "$PROJECT_NAME"

echo "Project setup successfully."