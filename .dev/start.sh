#!/bin/bash

print_info() {
  local type=$1
  local message=$2

  if [ "$type" == "error" ]; then
    echo -e "\e[31m$message\e[0m"
  elif [ "$type" == "ok" ]; then
    echo -e "\e[32m$message\e[0m"
  else
    echo -e "$message"
  fi
}

if [ ! -f .env ]; then
  print_info "error" "The environment file does not exist"

  if [ ! -f .env.example ]; then
    print_info "error" "The example environment also does not exist"
    exit 1
  fi

  echo "Creating the env file..."
  cp .env.example .env

  if [ $? -eq 0 ]; then
    print_info "ok"  "File created successfully."
  else
    print_info "error" "An error occurred while creating the file."
    exit 1
  fi
fi

docker compose up --build