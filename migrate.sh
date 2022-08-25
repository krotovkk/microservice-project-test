#!/bin/sh

Help()
{
   echo "-t: for test data base operation"
   echo "-d: for migration status display"
}

export MIGRATION_DIR="./migrations"
export DB_DSN="host=localhost port=80 user=user password=password dbname=products sslmode=disable"
export TEST_DB_DSN="host=localhost port=81 user=user password=password dbname=products sslmode=disable"

operation="up"
db=$DB_DSN

while getopts ":htd" option; do
   case $option in
      h)
        Help
        exit;;
      t)
        db=$TEST_DB_DSN;;
      d)
        operation="status";;
   esac
done

goose -v -dir ${MIGRATION_DIR} postgres "${db}" "${operation}"