#!/bin/bash

ghcicd_help_cmd() { printf "\033[1m  %-15s\033[0m %s\n" "$1" "$2"; }

################################################################################
# The command line help
################################################################################

ghcicd_help_header() {
  ghcicd_log_bold "################################################################################"
  ghcicd_log_bold "# GitHub CI/CD"
  ghcicd_log_bold "# Copyright (c) 2021 Marko Kungla"
  ghcicd_log_bold "# v$GHCICD_VERSION"
  ghcicd_log_bold "################################################################################"
  ghcicd_log_line
  ghcicd_log_line "Usage: ghcicd.sh [option...] command [arg...]" >&2
  ghcicd_log_line
  ghcicd_log_line "Why bash? Just because :-)"
  ghcicd_log_line
}

ghcicd_help_footer() {
  ghcicd_log_line
  ghcicd_log_bold "GLOBAL FLAGS"
  ghcicd_log_line
  ghcicd_log_line "  -h, --help                  show this help menu"
  ghcicd_log_line "  -v, --verbose               log verbose"
  ghcicd_log_line
}

ghcicd_help() {
  ghcicd_help_header
  # shellcheck disable=SC1091
  source "$GHCICD_PATH_SRC/help-menu.sh"
  ghcicd_help_footer
}

