#!/usr/bin/env bash
# ----------------------------------------------------------
# PURPOSE

# This is the test manager for ecm. It will run the testing
# sequence for ecm using docker.

# ----------------------------------------------------------
# REQUIREMENTS

# eris installed locally

# ----------------------------------------------------------
# USAGE

# test.sh

# ----------------------------------------------------------
# Set defaults

# Where are the Things?

base=github.com/eris-ltd/eris-pm
if [ "$CIRCLE_BRANCH" ]
then
  repo=${GOPATH%%:*}/src/github.com/${CIRCLE_PROJECT_USERNAME}/${CIRCLE_PROJECT_REPONAME}
  circle=true
else
  repo=$GOPATH/src/$base
  circle=false
fi
branch=${CIRCLE_BRANCH:=master}
branch=${branch/-/_}

# Other variables
was_running=0
test_exit=0

# ---------------------------------------------------------------------------
# Needed functionality

ensure_running(){
  if [[ "$(eris services ls | grep $1 | awk '{print $2}')" == "No" ]]
  then
    echo "Starting service: $1"
    eris services start $1 1>/dev/null
    early_exit
    sleep 3 # boot time
  else
    echo "$1 already started. Not starting."
    was_running=1
  fi
}

early_exit(){
  if [ $? -eq 0 ]
  then
    return 0
  fi

  echo "There was an error duing setup; keys were not properly imported. Exiting."
  if [ "$was_running" -eq 0 ]
  then
    if [ "$circle" = true ]
    then
      eris services stop keys
    else
      eris services stop -r keys
    fi
  fi
  exit 1
}

get_uuid() {
  if [[ "$(uname -s)" == "Linux" ]]
  then
    uuid=$(cat /proc/sys/kernel/random/uuid | tr -dc 'a-zA-Z0-9' | fold -w 12 | head -n 1)
  elif [[ "$(uname -s)" == "Darwin" ]]
  then
    uuid=$(uuidgen | tr -dc 'a-zA-Z0-9' | fold -w 12 | head -n 1)
  else
    uuid="62d1486f0fe5"
  fi
  return uuid
}

test_build() {
  echo ""
  echo "Building eris-cm in a docker container."
  set -e
  tests/build_tool.sh 1>/dev/null
  set +e
  if [ $? -ne 0 ]
  then
    echo "Could not build eris-cm. Debug via by directly running [`pwd`/tests/build_tool.sh]"
    exit 1
  fi
  echo "Build complete."
  echo ""
}

test_setup(){
  echo "Getting Setup"
  if [ "$circle" = true ]
  then
    export ERIS_PULL_APPROVE="true"
    eris init --yes --pull-images --source="rawgit" --testing 1>/dev/null
  fi

  ensure_running keys
  echo "Setup complete"
}

perform_tests(){
# base make a chain
# make a chain using chaintypes
# make a chain using flags
# make a chain using csv
# add a new account type
# add a new chain type
# export/inspect tars
# export/inspect zips
# accountFlags > chainTypes > csv
}

test_teardown(){
  if [ "$circle" = false ]
  then
    echo ""
    if [ "$was_running" -eq 0 ]
    then
      eris services stop -rx keys
    fi
  fi
  echo ""
  if [ "$test_exit" -eq 0 ]
  then
    echo "Tests complete! Tests are Green. :)"
  else
    echo "Tests complete. Tests are Red. :("
  fi
  cd $start
  exit $test_exit
}

# ---------------------------------------------------------------------------
# Get the things build and dependencies turned on

echo "Hello! I'm the marmot that tests the eris-cm tooling."
start=`pwd`
cd $repo
test_build
test_setup

# ---------------------------------------------------------------------------
# Go!

echo "Running Tests..."
perform_tests

# ---------------------------------------------------------------------------
# Cleaning up

test_teardown
