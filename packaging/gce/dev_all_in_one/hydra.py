import os
import subprocess
import time
import requests

# Get config from metadata
r = requests.get(
    url="http://metadata.google.internal/computeMetadata/v1/instance/attributes/hydra",
    headers={
        'Metadata-Flavor': 'Google'
    })
r.raise_for_status()
last_etag = r.headers['etag']

first_load = not os.path.exists('/opt/hydra.yaml')

with open('/opt/hydra.yaml', 'w') as fd:
    fd.write(r.text)

if first_load:
    time.sleep(120)
    subprocess.run("helm install hydra /opt/images/hydra-0.2.2.tgz -f /opt/hydra.yaml", shell=True)
    time.sleep(60)
else:
    subprocess.run("helm upgrade hydra /opt/images/hydra-0.2.2.tgz -f /opt/hydra.yaml", shell=True)

while True:
    r = requests.get(
        url="http://metadata.google.internal/computeMetadata/v1/instance/attributes/hydra",
        params={'last_etag': last_etag, 'wait_for_change': True},
        headers={
            'Metadata-Flavor': 'Google'
        })
    if r.status_code != 503:
        time.sleep(15)
        continue
    try:
        r.raise_for_status()
        last_etag = r.headers['etag']
        with open('/opt/hydra.yaml', 'w') as fd:
            fd.write(r.text)
        subprocess.run("helm upgrade hydra /opt/images/hydra-0.2.2.tgz -f /opt/hydra.yaml", shell=True)
    except:
        time.sleep(60)
