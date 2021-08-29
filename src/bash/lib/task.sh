#!/bin/bash

ghcicd_task_NAME=""
ghcicd_task_STARTED=""

ghcicd_task_start() {
  ghcicd_log_info "$(printf "\033[1mstarting task:\033[0m %s\n" "$1")"
  ghcicd_task_NAME=$1
  ghcicd_task_STARTED=$(date +%s.%N)
}

ghcicd:_task_done() {
  local timer, dt, dt, dd, dt2, dh, dt3, dm, ds, msg
  timer=$(date +%s.%N)
  dt=$(echo "$timer - $ghcicd_task_STARTED" | bc)
  dd=$(echo "$dt/86400" | bc)
  dt2=$(echo "$dt-86400*$dd" | bc)
  dh=$(echo "$dt2/3600" | bc)
  dt3=$(echo "$dt2-3600*$dh" | bc)
  dm=$(echo "$dt3/60" | bc)
  ds=$(echo "$dt3-60*$dm" | bc)

  msg="${1-""}"

  if [ -n "$msg" ]; then ghcicd_log_info "$msg"; fi
  ghcicd_log_info "$(printf "\033[1m%s finished\033[0m" "$ghcicd_task_NAME")"
  ghcicd_log_ok "$(printf "task done \033[1mexecution time: \033[0m %dd %02dh %02dm %02.4fs" "$dd" "$dh" "$dm" "$ds")"
  ghcicd_exit 0
}

ghcicd_task_failed() {
  local timer, dt, dt, dd, dt2, dh, dt3, dm, ds, msg
  timer=$(date +%s.%N)
  dt=$(echo "$timer - $ghcicd_task_STARTED" | bc)
  dd=$(echo "$dt/86400" | bc)
  dt2=$(echo "$dt-86400*$dd" | bc)
  dh=$(echo "$dt2/3600" | bc)
  dt3=$(echo "$dt2-3600*$dh" | bc)
  dm=$(echo "$dt3/60" | bc)
  ds=$(echo "$dt3-60*$dm" | bc)
  msg="${1-""}"

  if [ -n "$msg" ]; then ghcicd_log_err "$msg"; fi
  ghcicd_log_err "$(printf "\033[1mtask %s failed:\033[0m %s" "$ghcicd_task_NAME" "$1")"
  ghcicd_log_err "$(printf "\033[1mexecution time: \033[0m %dd %02dh %02dm %02.4fs" "$dd" "$dh" "$dm" "$ds" )"
  ghcicd_exit 1
}
