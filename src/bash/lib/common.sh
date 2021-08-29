#!/bin/bash

ghcicd_exit() {
  # shellcheck disable=SC2086
  exit $1
}

ghcicd_exec() {
  if ((GHCICD_OUTPUT_X)); then
    ghcicd_log_bold "[cmd-x ]:" "$1"
  fi

  ghcicd_task_start "exec cmd" timer

  eval "$2=\$($1)"

  local status=$?
  if [ $status -eq 0 ]; then
    ghcicd_task_done "exec cmd" "${timer:?}"
  else
    ghcicd_task_failed "exec cmd" "${timer:?}"
  fi

  return $status
}

ghcicd_is_dir() {
  if [[ -d $1 ]] && [[ -n $1 ]]; then
    return 0
  else
    return 1
  fi
}

ghcicd_file_exists() {
  if [[ -f $1 ]] && [[ -n $1 ]]; then
    return 0
  else
    return 1
  fi
}

# e.g  "CI" "true"
ghcicd_env_eq() {
  local varname="${1:=""}"
  local varval="${2:=""}"
  if [[ ! -v "${varname}" ]]; then
    ghcicd_log_info "no" $varname
    return 1
  else
    if [[ "${!varname}" == "$varval" ]]; then
      return 0
    else
      return 1
    fi
  fi
}
