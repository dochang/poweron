#!/bin/sh

set -e

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
python-pip
python-virtualenv
build-essential
lsb-release
EOF

venv=~/.local/venvs/battleschool
[ ! -d "${venv}" ] || find "${venv}" -type l -delete
virtualenv "${venv}"
"${venv}/bin/pip" install -U 'git+https://github.com/dochang/battleschool.git@devel#egg=battleschool'
hash -r
export BATTLE="${venv}/bin/battle"

cache_dir=~/.battleschool/cache
[ -d "${cache_dir}" ] || mkdir -p "${cache_dir}"

"${BATTLE}" --config-file https://github.com/dochang/poweron/raw/master/poweron.yml --sudo --ask-sudo-pass --update-sources -vvv --extra-vars 'mole_state=running' "$@"
