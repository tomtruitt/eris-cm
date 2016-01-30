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

base=github.com/eris-ltd/eris-cm
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
chains_dir=$HOME/.eris/chains

export ERIS_PULL_APPROVE="true"
export ERIS_MIGRATE_APPROVE="true"

# ---------------------------------------------------------------------------
# Needed functionality

ensure_running(){
  if [[ "$(eris services ls -qr | grep $1)" == "$1" ]]
  then
    echo "$1 already started. Not starting."
    was_running=1
  else
    echo "Starting service: $1"
    eris services start $1 1>/dev/null
    early_exit
    sleep 3 # boot time
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
  echo $uuid
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
    eris init --yes --pull-images=true --testing=true 1>/dev/null
  fi

  ensure_running keys
  echo "Setup complete"
}

check_test(){
  # check chain is running
  chain=( $(eris chains ls --quiet --running | grep $uuid) )
  if [ ${#chain[@]} -ne 1 ]
  then
    echo "chain does not appear to be running"
    test_exit=1
  fi

  # check results file exists
  if [ ! -e "$chains_dir/$uuid/accounts.csv" ]
  then
    echo "accounts.csv not present"
    ls -la $chains_dir/$uuid
    pwd
    ls -la $chains_dir
    test_exit=1
  fi

  # check genesis.json
  genOut=$(cat $dir_to_use/genesis.json)
  genIn=$(eris chains plop $uuid genesis)
  if [[ "$genOut" != "$genIn" ]]
  then
    test_exit=1
    echo "genesis.json's do not match"
    echo
    echo "expected"
    echo
    echo $genOut
    echo
    echo "received"
    echo
    echo $genIn
  fi

  # check priv_validator
  privOut=$(cat $dir_to_use/priv_validator.json)
  privIn=$(eris data exec $uuid "cat /home/eris/.eris/chains/$uuid/priv_validator.json")
  if [[ "$privOut" != "$privIn" ]]
  then
    test_exit=1
    echo "priv_validator.json's do not match"
    echo
    echo "expected"
    echo
    echo $privOut
    echo
    echo "received"
    echo
    echo $privIn
  fi
}

run_test(){
  echo -e "Running Test:\t$@"
  $@
  if [ $? -ne 0 ]
  then
    test_exit=1
    return 1
  fi
  dir_to_use=$chains_dir/$uuid/$direct
  eris chains new $uuid --dir $uuid/$direct
  if [ $? -ne 0 ]
  then
    test_exit=1
    return 1
  fi
  sleep 3
  eris chains stop -f $uuid
  eris chains rm -xf $uuid
  rm -rf $chains_dir/$uuid
}

perform_tests(){
  echo
  # simplest test
  uuid=$(get_uuid)
  direct=""
  run_test eris chains make $uuid --account-types=Full:1
  if [ $test_exit -eq 1 ]
  then
    return 1
  fi

  # more complex flags test
  uuid=$(get_uuid)
  direct="$uuid"_validator_000
  run_test eris chains make $uuid --account-types=Root:2,Developer:2,Participant:2,Validator:1
  if [ $test_exit -eq 1 ]
  then
    return 1
  fi

  # chain-type test
  uuid=$(get_uuid)
  direct=""
  run_test eris chains make $uuid --chain-type=simplechain
  if [ $test_exit -eq 1 ]
  then
    return 1
  fi

  # add a new account type
  uuid=$(get_uuid)
  direct=""
  cp $repo/tests/fixtures/tester.toml $chains_dir/account-types/.
  run_test eris chains make $uuid --account-types=Test:1
  if [ $test_exit -eq 1 ]
  then
    return 1
  fi
  rm $chains_dir/account-types/tester.toml

  # add a new chain type
  uuid=$(get_uuid)
  direct="$uuid"_full_000
  cp $repo/tests/fixtures/testchain.toml $chains_dir/chain-types/.
  run_test eris chains make $uuid --chain-type=testchain
  if [ $test_exit -eq 1 ]
  then
    return 1
  fi
  rm $chains_dir/chain-types/testchain.toml

  # export/inspect tars
  uuid=$(get_uuid)
  direct=""
  eris chains make $uuid --account-types=Full:2 --tar
  if [ $? -ne 0 ]
  then
    test_exit=1
    return 1
  fi
  tar -xzf $chains_dir/$uuid/"$uuid"_full_000.tar.gz -C $chains_dir/$uuid/.
  run_test echo "tar test"
  if [ $test_exit -eq 1 ]
  then
    return 1
  fi

  # export/inspect zips
  # todo

  # make a chain using csv
  uuid=$(get_uuid)
  direct=""
  eris chains make $uuid --account-types=Full:1
  if [ $? -ne 0 ]
  then
    test_exit=1
    return 1
  fi
  rm $chains_dir/$uuid/genesis.json
  prev_dir=`pwd`
  cd $chains_dir/$uuid
  eris chains make $uuid --known --accounts accounts.csv --validators validators.csv > $chains_dir/$uuid/genesis.json
  run_test echo "known test"
  cd $prev_dir
  if [ $test_exit -eq 1 ]
  then
    return 1
  fi
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
