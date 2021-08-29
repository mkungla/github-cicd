#!/bin/bash

ghcicd_version_main() {
  echo "$GHCICD_VERSION"
  ghcicd_exit 0
}

ghcicd_version_help() {
  ghcicd_log_bold "VERSION COMMANDS"
  ghcicd_log_line
  ghcicd_help_cmd "   version  " "<command> [arg...]"
  ghcicd_log_line
  ghcicd_log_line
}
