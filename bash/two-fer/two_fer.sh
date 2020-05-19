#!/usr/bin/env bash
# share prints two fer
share() {
  name="you"
  if [[ $# -ne 0 ]]; then
      name=$1
  fi
  echo "One for $name, one for me."
}

# call main with all of the positional arguments
share "$@"

