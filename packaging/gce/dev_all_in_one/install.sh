# Add airgap images for laniakea
gsutil -m cp gs://$SOURCE_BUCKET/lnk-norma-0.2.2/* /opt/images/
gsutil -m cp gs://$SOURCE_BUCKET/lnk-hydra-0.2.2/* /opt/images/

# And script
mv /tmp/*.py /root/

# And wait page
mv /tmp/502.html /usr/share/nginx/html/custom_502.html
(cat <<EOF
[Unit]
After=k3s.service

[Service]
Type=notify
TimeoutStartSec=800
ExecStart=/usr/bin/python3 /root/norma.py
Environment="KUBECONFIG=/etc/rancher/k3s/k3s.yaml"

[Install]
WantedBy=default.target
EOF
) >/etc/systemd/system/norma-start.service

(cat <<EOF
[Unit]
After=norma-start.service

[Service]
ExecStart=/usr/bin/python3 /root/norma_sync.py
Environment="KUBECONFIG=/etc/rancher/k3s/k3s.yaml"

[Install]
WantedBy=default.target
EOF
) >/etc/systemd/system/norma-sync.service

(cat <<EOF
[Unit]
After=norma-start.service

[Service]
ExecStart=/usr/bin/python3 /root/hydra.py
Environment="KUBECONFIG=/etc/rancher/k3s/k3s.yaml"

[Install]
WantedBy=default.target
EOF
) >/etc/systemd/system/hydra.service


systemctl enable k3s-mount.service
systemctl enable k3s.service
systemctl enable norma-start.service
systemctl enable norma-sync.service
systemctl enable hydra.service

mv /tmp/nginx.conf /etc/nginx/nginx.conf && chmod 644 /etc/nginx/nginx.conf

systemctl enable nginx.service
