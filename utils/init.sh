#!/bin/bash
echo "Connecting to the internet"
wpa_supplicant -B -i wlp3s0 -c /etc/wpa_supplicant/wpa_supplicant.conf
dhcpcd wlp3s0
echo "Mounting storage drive"
mount /dev/sdc2 /mnt/storage
