#!/usr/bin/env bash

# wait-for-it.sh

set -e

# Timeout in seconds (default 30)
TIMEOUT=${WAIT_FOR_IT_TIMEOUT:-30}
# Interval between checks (default 1)
INTERVAL=${WAIT_FOR_IT_INTERVAL:-1}

# Function to check if the service is up
check_service() {
  local host="$1"
  local port="$2"
  
  nc -z "$host" "$port" > /dev/null 2>&1
}

# Parsing arguments
host="$1"
port="$2"
shift 2

# Wait for the service to be up
echo "Waiting for $host:$port to be available..."

timeout="$TIMEOUT"
while [ $timeout -gt 0 ]; do
  if check_service "$host" "$port"; then
    echo "$host:$port is available."
    exec "$@"
    exit 0
  fi
  sleep "$INTERVAL"
  timeout=$((timeout - INTERVAL))
done

echo "Timeout: $host:$port did not become available after $TIMEOUT seconds."
exit 1
