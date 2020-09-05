set -x

export DEBIAN_FRONTEND=noninteractive
export K3S_VERSION="v1.18.8%2Bk3s1"
apt clean
apt update
apt -y dist-upgrade
apt -y install wget python3-pip nginx pbzip2

# Install a few python packages for tooling
pip3 install requests pyyaml sdnotify

# Install gcloud
pip3 install --no-cache-dir -U crcmod

echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list &&
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add - &&
  apt update &&
  apt install -y google-cloud-sdk

# Install k3s
wget -O /usr/local/bin/k3s "https://github.com/rancher/k3s/releases/download/$K3S_VERSION/k3s"
chmod a+x /usr/local/bin/k3s

# Install helm

## Helm
wget https://get.helm.sh/helm-v3.3.0-linux-amd64.tar.gz &&
  tar -zxvf helm-v3.3.0-linux-amd64.tar.gz &&
  mv linux-amd64/helm /usr/local/bin/helm && chmod a+x /usr/local/bin/helm

# make mount point
mkdir -p /data

# Install k3s service

cp /tmp/mount.py /root/mount.py

(
  cat <<EOF
[Unit]
After=network.target

[Service]
Type=notify
TimeoutStartSec=300
ExecStart=/usr/bin/python3 /root/mount.py
LimitNOFILE=1048576
LimitNPROC=infinity
LimitCORE=infinity

[Install]
WantedBy=default.target
EOF
) >/etc/systemd/system/k3s-mount.service

(
  cat <<EOF
[Unit]
After=k3s-mount.service

[Service]
Type=notify
ExecStart=/usr/local/bin/k3s server  -d /data/k3s/data --default-local-storage-path /data/k3s/storage --write-kubeconfig-mode 0644 --disable traefik --resolv-conf /run/systemd/resolve/resolv.conf
KillMode=process
Delegate=yes
LimitNOFILE=1048576
LimitNPROC=infinity
LimitCORE=infinity
TasksMax=infinity
TimeoutStartSec=0
Restart=always
RestartSec=5s

[Install]
WantedBy=default.target
EOF
) >/etc/systemd/system/k3s.service

# Do not enable services, inheriting images should do it
#systemctl enable k3s-mount.service
#systemctl enable k3s.service

# Add airgap images for k3s
mkdir -p /opt/images

wget -O /opt/images/k3s-airgap-images-amd64.tar https://github.com/rancher/k3s/releases/download/$K3S_VERSION/k3s-airgap-images-amd64.tar

# Precreate symbolic links

mkdir -p /data/rancher/lib && mkdir -p /data/rancher/etc
ln -s /data/rancher/lib /var/lib/rancher
ln -s /data/rancher/etc /etc/rancher
