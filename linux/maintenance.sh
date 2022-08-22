#!bin/bash

echo "LINUX SPEED-UP by @guidoenr"

cyan='\033[0;31m'
clear='\033[0m'
check='\xE2\x9C\x94'

printStarting(){
    printf "[task]- $1 "
}

printFinished(){
    printf "${cyan}${check}${clear} \n"
}

# clean apt
printStarting "apt-get clean"
sudo apt-get clean -y > /dev/null
printFinished 

# auto clean packages
printStarting "apt-get autoclean"
sudo apt-get autoclean -y > /dev/null
printFinished 

# auto remove packages
printStarting "apt-get autoremove"
sudo apt-get autoremove -y > /dev/null
printFinished 

# clean logs in 30 days
printStarting "cleaning journal logs"
sudo journalctl --vacuum-time=30d --quiet
printFinished 

# clean thumbnail cache
printStarting "cleaning thumbnail cache"
rm -rf ~/.cache/thumbnails/* > /dev/null
printFinished 

# remove screenshots
printStarting "removing screenshots in ~/Pictures/Screenshots"
rm -rf ~/Pictures/Screenshots/*
printFinished 

printf "All tasks finished \n"

exit


