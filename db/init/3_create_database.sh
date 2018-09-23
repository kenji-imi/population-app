#!/bin/sh

echo "CREATE DATABASE IF NOT EXISTS \`dump_schema\` ;" | "${mysql[@]}"
echo "GRANT ALL ON \`dump_schema\`.* TO '"$MYSQL_USER"'@'%' ;" | "${mysql[@]}"
echo 'FLUSH PRIVILEGES ;' | "${mysql[@]}"
