#!/usr/bin/env bash
# share prints two fer
share() {
  name="you"
  if (( $# != 0 )); then
      name=$1
  fi
  echo "One for $name, one for me."
}

# call share with all of the positional arguments
share "$@"

