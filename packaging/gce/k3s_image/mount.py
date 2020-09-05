import subprocess

import sdnotify


def shellEx(cmd: str):
    res = subprocess.run(cmd, shell=True)
    return res.returncode


def shell(cmd: str):
    res = subprocess.run(cmd, shell=True)
    res.check_returncode()


# Mount data drive

if shellEx("mount |grep /data") == 1:
    disk_device = "/dev/nvme0n1"
    try:
        print("mounting data drive")
        shell(f"mount -o discard,defaults {disk_device} /data")
    except:
        print("New data drive, preparing folders")
        # Format it if necessary
        shell(f"mkfs.ext4 -E lazy_itable_init=0,lazy_journal_init=0,discard {disk_device}")
        shell(f"mount -o discard,defaults,noatime {disk_device} /data")
        # And create persistent folders
        shell("mkdir -p /data/k3s/data/agent/images && mkdir -p /data/k3s/storage && mkdir -p /data/rancher/etc && mkdir -p /data/rancher/lib")
        # Copy images to data dir
        print("Copying images")
        shell("cp /opt/images/k3s-airgap-images-amd64.tar /data/k3s/data/agent/images/")
        shell("cp /opt/images/*.bz2 /data/k3s/data/agent/images/ && mkdir -p /data/k3s/storage")
        # Decompress images to data dir
        shell("cd /data/k3s/data/agent/images/ && pbzip2 -d *.tar.bz2")

print("Data drive ready")
n = sdnotify.SystemdNotifier()
n.notify("READY=1")
