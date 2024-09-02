#!/bin/bash
MYSQL_USER="root"
MYSQL_PASSWORD="rootpassword"
MYSQL_DATABASE="ketidatabase"

# Define the backup directory inside the container
BACKUP_DIR="/backup"

# Print contents of the backup directory for debugging
echo "Listing files in $BACKUP_DIR:"
ls -l $BACKUP_DIR

# Locate the most recent backup file
LATEST_BACKUP=$(ls -t $BACKUP_DIR/mysql_backup_*.sql 2>/dev/null | head -n 1)

if [ -z "$LATEST_BACKUP" ]; then
  echo "No backup file found. Continuing without restoration."
else
  echo "Restoring from backup: $LATEST_BACKUP"

  # Restore the latest backup
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" "$MYSQL_DATABASE" < "$LATEST_BACKUP"

  if [ $? -eq 0 ]; then
    echo "Restoration successful."
  else
    echo "Restoration failed."
    exit 1
  fi
fi

# Continue with the original entrypoint command
exec "$@"
