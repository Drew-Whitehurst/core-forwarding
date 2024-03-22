#!/bin/bash

go build apx-proxy.go
go build ouster-proxy.go
go build camera-proxy.go

mkdir /root/dev-tools

cp apx-proxy ouster-proxy camera-proxy /root/dev-tools

cp *.service /etc/systemd/system

systemctl enable core_forwarding_ouster.service
systemctl enable core_forwarding_apx.service
systemctl enable core_forwarding_camera.service

systemctl start core_forwarding_ouster.service
systemctl start core_forwarding_apx.service
systemctl start core_forwarding_camera.service
