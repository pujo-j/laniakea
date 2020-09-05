import os
import subprocess
import time
import sdnotify

import requests

# Get config from metadata
r = requests.get(
    url="http://metadata.google.internal/computeMetadata/v1/instance/attributes/norma",
    headers={
        'Metadata-Flavor': 'Google'
    })
r.raise_for_status()

first_load = not os.path.exists('/opt/norma.yaml')

with open('/opt/norma.yaml', 'w') as fd:
    fd.write(r.text)

if first_load:
    # Hacky hack to ensure k3s is properly started...
    print("Initial start, wait for k3s to boot")
    time.sleep(60*3)
    subprocess.run("helm install norma /opt/images/norma-0.2.2.tgz -f /opt/norma.yaml", shell=True)
    time.sleep(60)

n = sdnotify.SystemdNotifier()
n.notify("READY=1")

