#!/bin/sh
# RINGPOP NEEDS A CONCRETE BROADCAST ADDRESS. With BIND_ON_IP=0.0.0.0 the membership layer
# refuses to start -- `broadcastAddress required when listening on all interfaces (0.0.0.0/[::])`
# -- and the whole server fails to build its history service. The container's IP is assigned at
# run time, so it has to be resolved here rather than baked; upstream's own entrypoint does the
# same thing. Falls back to loopback so a host-network or odd-DNS run still starts.
# The image is a chiselled runtime with no getent, hostname or awk: the address is the one /etc/hosts
# gives the container's hostname (Docker and Kubernetes write it there), read with shell builtins.
set -e
if [ -z "${TEMPORAL_BROADCAST_ADDRESS:-}" ]; then
  _h=$(cat /etc/hostname 2>/dev/null || true)
  if [ -n "$_h" ] && [ -r /etc/hosts ]; then
    while read -r _ip _names; do
      case "$_ip" in ''|'#'*) continue ;; esac
      for _n in $_names; do
        if [ "$_n" = "$_h" ]; then TEMPORAL_BROADCAST_ADDRESS=$_ip; break 2; fi
      done
    done < /etc/hosts
  fi
  [ -n "${TEMPORAL_BROADCAST_ADDRESS:-}" ] || TEMPORAL_BROADCAST_ADDRESS=127.0.0.1
  export TEMPORAL_BROADCAST_ADDRESS
fi
echo "temporal: broadcast address ${TEMPORAL_BROADCAST_ADDRESS}"

# Temporal ships no 'default' namespace: upstream registers one in the separate
# temporalio/docker-builds auto-setup image, not in temporal-server itself. Without a
# namespace every workflow call fails NamespaceNotFound, so do it here, once, over the
# HTTP API as soon as the frontend answers. Backgrounded so the server stays PID 1's exec
# target; failures are non-fatal because a namespace already registered on a previous boot
# returns AlreadyExists and that is a success, not an error. rs-probe (the image's HTTP probe; no curl)
# sends the request body from standard input (-T /dev/stdin) with -X POST.
(
  i=0
  while [ "$i" -lt 150 ]; do
    if /usr/local/bin/rs-probe http://127.0.0.1:7243/api/v1/namespaces >/dev/null 2>&1; then
      printf '%s' '{"namespace":"default","workflowExecutionRetentionPeriod":"259200s"}' | \
        /usr/local/bin/rs-probe -X POST -T /dev/stdin http://127.0.0.1:7243/api/v1/namespaces >/dev/null 2>&1 || true
      exit 0
    fi
    i=$((i+1))
    sleep 2
  done
) &
exec temporal-server start
