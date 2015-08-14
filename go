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

# Required by Nix installer
sudo env DEBIAN_FRONTEND=noninteractive xargs apt-get --yes install <<EOF
curl
bzip2
tar
EOF

# Install pip & virtualenv from pip instead of apt since these bugs:
#
# https://github.com/docker/docker-py/issues/525#issuecomment-79428103
# https://bugs.launchpad.net/ubuntu/+source/python-pip/+bug/1306991
curl -sSL https://bootstrap.pypa.io/get-pip.py | sudo python2
hash -r
sudo pip2 install -U pip
sudo pip2 install virtualenv
hash -r
sudo pip2 install -U virtualenv

venv=~/.local/venvs/battleschool
[ ! -d "${venv}" ] || find "${venv}" -type l -delete
virtualenv "${venv}"
"${venv}/bin/pip" install -U 'git+https://github.com/dochang/battleschool.git@devel#egg=battleschool'
hash -r
export BATTLE="${venv}/bin/battle"

cache_dir=~/.battleschool/cache
[ -d "${cache_dir}" ] || mkdir -p "${cache_dir}"

"${BATTLE}" --config-file https://github.com/dochang/poweron/raw/master/poweron.yml --sudo --ask-sudo-pass --update-sources -vvv --extra-vars 'mole_state=started' "$@"

export http_proxy=http://127.0.0.1:8118
export https_proxy=http://127.0.0.1:8118
export no_proxy=localhost,127.0.0.1

curl https://nixos.org/nix/install | sh

# Clean conflicting files
rm -rf ~/.bash_logout ~/.bashrc ~/.profile ~/.nix-channels

export FRESH_LOCAL_SOURCE=dochang/dotfiles

curl -sSL get.freshshell.com | bash

"${BATTLE}" --ask-sudo-pass -vvv "$@"
