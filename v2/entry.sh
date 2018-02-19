#!/bin/sh

set -e

datadir=${APP_DATADIR:="/var/lib/data"}
host=${APP_HOST:="127.0.0.1"}
port=${APP_PORT:="26257"}
username=${APP_USERNAME:=""}
database=${APP_DATABASE:=""}

cat <<EOF > /etc/config.json
{
    "datadir": "${datadir}",
    "host": "$host",
    "port": "$port",
    "username": "$username",
    "database": "$database"
}
EOF

mkdir -p ${APP_DATADIR}

exec "/app"