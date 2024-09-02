#!/bin/bash

# MySQL container details
MYSQL_CONTAINER="mysql-db"  # Ensure this matches your MySQL container name

# MySQL connection details
MYSQL_USER="root"
MYSQL_PASSWORD="rootpassword"
MYSQL_DATABASE="ketidatabase"

# Function to execute MySQL commands inside the Docker container
execute_mysql() {
    docker exec -i "$MYSQL_CONTAINER" mysql -h127.0.0.1 -P3306 -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D"$MYSQL_DATABASE" -e "$1"
}

# Query to check for entries without a response
query="SELECT file_name, request_sent_timestamp, IF(respond_received IS NULL OR respond_received = FALSE, 'False', 'True') AS respond_received FROM validity WHERE respond_received IS NULL OR respond_received = FALSE OR respond_received_timestamp IS NULL;"

# Execute the query and display the results
echo "Checking for entries with missing respond_received or respond_received_timestamp..."
results=$(execute_mysql "$query")

# Check if results are empty
if [[ -z "$results" ]]; then
    echo "All entries have been responded to."
else
    echo "Entries missing respond_received or respond_received_timestamp:"
    echo "file_name    request_sent_timestamp    respond_received"
    echo "$results" | awk 'NR > 1 { printf "%-15s %-25s %s\n", $1, $2 " " $3, $4 }'
fi
