#!/bin/bash

echo "Starting.."

# Create a temporary file for storing command output
tmpfile=$(mktemp)

# Define a function to check command exit status and display error message if necessary
check_status() {
  if [ $1 -ne 0 ]; then
    echo "Error: $2 failed. See $tmpfile for details."
    exit 1
  fi
}

# Clear cache
sudo apt clean -y > $tmpfile 2>&1
check_status $? "apt clean"

# Remove old kernels
sudo apt autoremove --purge -y > $tmpfile 2>&1
check_status $? "apt autoremove"

# Remove unneeded packages
sudo apt-get autoremove -y > $tmpfile 2>&1
check_status $? "apt-get autoremove"

# Remove unnecessary files from home directory
rm -rf ~/Pictures/Screenshots/* ~/.local/share/Trash/* > $tmpfile 2>&1
check_status $? "rm"

# Update system
sudo apt update && sudo apt upgrade -y > $tmpfile 2>&1
check_status $? "apt update/upgrade"

# Clean up system logs
sudo journalctl --vacuum-size=1G > $tmpfile 2>&1
check_status $? "journalctl"

# Remove cache thumbnails
rm -rf ~/.cache/thumbnails/* > $tmpfile 2>&1
check_status $? "rm"

# Clean up temporary file
rm $tmpfile

# Display completion message
echo "Optimization completed successfully!"
