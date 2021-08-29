#!/bin/bash
################################################################################
# LOG
################################################################################
set -a
ghcicd_log_info() {
  if ((GHCICD_VERBOSE)); then
    printf "\033[94m%s\033[0m %s\n" "[ghcicd]:" "$*";
  fi
}
ghcicd_log_err() { printf "\033[91m%s\033[0m %s\n" "[ghcicd]:" "$*" >&2; }
ghcicd_log_warn() { printf "\033[33m%s\033[0m %s\n" "[ghcicd]:" "$*"; }
ghcicd_log_ok() { printf "\033[32m%s\033[0m %s\n" "[ghcicd]:" "$*"; }
ghcicd_log_line() { printf "%s\n" "$*"; }
ghcicd_log_debug() {
  if ((GHCICD_DEBUG)); then
    printf "\033[35m%s\033[0m %s\n" "[ghcicd]:" "$*";
  fi
}
ghcicd_log_mute() { printf "\033[39m%s\033[0m\n" "$*"; }
ghcicd_log_bold() { printf "\033[1m%s\033[0m\n" "$*"; }
set +a
