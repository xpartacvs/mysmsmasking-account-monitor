# MySMSMasking Account Monitor

Monitor status saldo dan tanggal kedaluarsa akun [MySMSMasking](https://mysmsmasking.com/)

## Cara Pakai

Yang paling enak ya pakai docker. Berikut ini step-stepnya:

```bash
# Clone repository dulu dan masuk ke directorynya
git clone git@github.com:xpartacvs/mysmsmasking-account-monitor.git
cd mysmsmasking-account-monitor

# Build image docker-nya
docker image build -t mysmsmasking-account-monitor:latest .

# Run container-nya
docker container run \
    -it \
    -e DISCORD_WEBHOOKURL=... \
    -e RAJASMS_API_URL=... \
    -e RAJASMS_API_KEY=... \
    mysmsmasking-account-monitor:latest
```
