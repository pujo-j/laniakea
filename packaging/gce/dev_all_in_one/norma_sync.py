import os
import subprocess
import time

import requests

os.environ["no_proxy"] = "169.254.169.254,metadata,metadata.google.internal"
last_etag = '0'

while True:
    r = requests.get(
        url="http://metadata.google.internal/computeMetadata/v1/instance/attributes/norma",
        params={'last_etag': last_etag, 'wait_for_change': True},
        headers={
            'Metadata-Flavor': 'Google'
        })
    if r.status_code == 503:
        time.sleep(15)
        continue
    try:
        r.raise_for_status()
        last_etag = r.headers['etag']
        subprocess.run("helm upgrade norma /opt/images/norma-0.2.2.tgz -f /opt/norma.yaml", shell=True)
    except:
        time.sleep(60)
