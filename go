#!/bin/sh

set -e

[ -n "$USER" ] || {
	echo "ERROR: \$USER is not set" >&2
	exit 1
}

[ "$(id -u)" -ne 0 ] || {
	echo "ERROR: DO NOT run this script as *root* !" >&2
	exit 2
}

sudo env DEBIAN_FRONTEND=noninteractive xargs apt-get --yes install <<EOF
htop
iotop
iftop
nethogs
mtr-tiny
curl
ca-certificates
git
python
python-dev
build-essential
lsb-release
EOF

# Install pip & virtualenv from pip instead of apt since these bugs:
#
# https://github.com/docker/docker-py/issues/525#issuecomment-79428103
# https://bugs.launchpad.net/ubuntu/+source/python-pip/+bug/1306991
command -v pip2 >/dev/null 2>&1 || {
	curl -sSL https://bootstrap.pypa.io/get-pip.py | sudo python2
}
hash -r
sudo pip2 install -U pip
sudo pip2 install virtualenv
hash -r
sudo pip2 install -U virtualenv

: ${FRESH_LOCAL_SOURCE:=https://github.com/dochang/dotfiles.git}
export FRESH_LOCAL_SOURCE

[ -d ~/.dotfiles ] || {
	curl -sSL get.freshshell.com | bash
}

cd ~/.dotfiles

./bootstrap upgrade-to-sid.yml --ask-sudo-pass -vvv "$@"
