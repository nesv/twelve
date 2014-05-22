#!/usr/bin/env bash

export TWELVE_STRING1="test value one"

export TWELVE_BOOL_TRUE="1"
export TWELVE_BOOL_FALSE="False"
export TWELVE_BOOL_FAIL="neenerneenerneener"

go test $@