#!/bin/bash

# Set the backup directory and filename
BACKUP_DIR="./dbManagement/backup"
mkdir -p $BACKUP_DIR
TIMESTAMP=$(date +"%F_%T")
BACKUP_FILE="$BACKUP_DIR/mysql_backup_$TIMESTAMP.sql"

# Create the backup directory if it doesn't exist

# Dump the MySQL database
docker exec mysql-db sh -c 'exec mysqldump -uroot -prootpassword ketidatabase' > $BACKUP_FILE

# Check if the dump was successful
if [ $? -eq 0 ]; then
  echo "Backup successful: $BACKUP_FILE"
else
  echo "Backup failed!"
  exit 1
fi