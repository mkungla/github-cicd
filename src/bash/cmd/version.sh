#!/bin/bash

ghcicd_version_main() {
  ghcicd_help ghcicd_version_help
  ghcicd_exit 1
}

ghcicd_version_help() {
  ghcicd_log_bold "VERSION COMMANDS"
  ghcicd_log_line
  ghcicd_help_cmd "   version  " "<command> [arg...]"
  ghcicd_log_line
  ghcicd_log_line
}
