#!/bin/bash

ghcicd_task_start() {
  local name="${1-""}"
  ghcicd_log_info "$(printf "\033[1mstarting task:\033[0m %s\n" "$name")"
  if [ $# -eq 2 ]; then
    eval "$2=\$(date +%s.%N)"
  fi
}

ghcicd_task_done() {
  local name="${1-""}"
  local timer dt dt dd dt2 dh dt3 dm ds

  if ! ((GHCICD_VERBOSE)); then
    return
  fi
  if [ $# -eq 1 ]; then
    ghcicd_log_ok "$(printf "%s done" "$name")"
  elif [ $# -eq 2 ]; then
    timer=$(date +%s.%N)
    dt=$(echo "$timer - $2" | bc)
    dd=$(echo "$dt/86400" | bc)
    dt2=$(echo "$dt-86400*$dd" | bc)
    dh=$(echo "$dt2/3600" | bc)
    dt3=$(echo "$dt2-3600*$dh" | bc)
    dm=$(echo "$dt3/60" | bc)
    ds=$(echo "$dt3-60*$dm" | bc)
    ghcicd_log_ok "$(printf "%s done \033[1mexecution time: \033[0m %dd %02dh %02dm %02.4fs" \
      "$name" "$dd" "$dh" "$dm" "$ds")"
  fi
}

ghcicd_task_failed() {
  local name="${1-""}"
  local timer dt dt dd dt2 dh dt3 dm ds

  if ! ((GHCICD_VERBOSE)); then
    return
  fi
  if [ $# -eq 1 ]; then
    ghcicd_log_ok "$(printf "%s done" "$name")"
  elif [ $# -eq 2 ]; then
    timer=$(date +%s.%N)
    dt=$(echo "$timer - $2" | bc)
    dd=$(echo "$dt/86400" | bc)
    dt2=$(echo "$dt-86400*$dd" | bc)
    dh=$(echo "$dt2/3600" | bc)
    dt3=$(echo "$dt2-3600*$dh" | bc)
    dm=$(echo "$dt3/60" | bc)
    ds=$(echo "$dt3-60*$dm" | bc)
    ghcicd_log_err "$(printf "%s failed \033[1mexecution time: \033[0m %dd %02dh %02dm %02.4fs" \
      "$name" "$dd" "$dh" "$dm" "$ds")"
  fi
}
