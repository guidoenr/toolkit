#!/bin/bash

# run the script as superuser
if [[ $(id -u) -ne 0 ]]; then
    echo "This script must be run with superuser privileges (sudo)."
    exit 1
fi

# clean cache and temporary files
echo "Cleaning cache and temporary files..."
apt clean
rm -rf /tmp/*
rm -rf /var/tmp/*
find /var/log -type f -exec truncate -s 0 {} \;
echo "Cache and temporary files cleaned."

# free up memory
echo "Freeing up memory..."
sync && echo 3 > /proc/sys/vm/drop_caches
echo "Memory freed."

# performance optimizations
echo "Optimizing system performance..."
echo "vm.swappiness=10" >> /etc/sysctl.conf
echo "net.core.somaxconn=1024" >> /etc/sysctl.conf
echo "vm.vfs_cache_pressure=50" >> /etc/sysctl.conf
echo "fs.inotify.max_user_watches=524288" >> /etc/sysctl.conf
sysctl -p
echo "Performance optimization complete."

# desktop environment optimization (GNOME)
echo "Optimizing desktop environment (GNOME)..."
gsettings set org.gnome.shell.app-switcher current-workspace-only true
gsettings set org.gnome.desktop.interface enable-animations false
gsettings set org.gnome.desktop.interface clock-show-seconds true
gsettings set org.gnome.desktop.peripherals.touchpad natural-scroll false
echo "Desktop environment optimized."

# kernel optimization for AMD Ryzen 7 5800X
echo "Optimizing kernel for AMD Ryzen 7 5800X..."
echo "processor.max_cstate=5" >> /etc/default/grub
update-grub
echo "Kernel optimized for AMD Ryzen 7 5800X."

echo "Optimization complete! Restart the system to apply the changes."
