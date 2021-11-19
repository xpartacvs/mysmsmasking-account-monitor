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
docker container run --rm -it -e ... mysmsmasking-account-monitor:latest
```

> **PENTING**: Minimal ada 5 (lima) environment variable yang harus di-_assign_ yaitu environment variable yang di beri tanda `checklist` (√) pada [tabel dibawah ini](#konfigurasi)

## Konfigurasi

Konfigurasi aplikasi ini dapat dilakukan dengan menggunakan environment variables.

| **Variable**            | **Type**  | **Req** | **Default**                  | **Description**                                                                                                                 |
| :---                    | :---      | :---:   | :---                         | :---                                                                                                                            |
| `DISCORD_WEBHOOKURL`    | `string`  | √       |                              | URL webhook Discord.                                                                                                            |
| `DISCORD_BOT_NAME`      | `string`  |         | suka-suka discord            | Nama bot yang akan muncul di channel Discord.                                                                                   |
| `DISCORD_BOT_AVATARURL` | `string`  |         | suka-suka discord            | URL ke file gambar yang akan digunakan sebagai avatar bot discord.                                                              |
| `DISCORD_BOT_MESSAGE`   | `string`  |         | `Reminder akun MySMSMAsking` | Pesan yang akan ditulis bot discord perihal status akun MySMSMasking.                                                           |
| `LOGMODE`               | `string`  |         | `disabled`                   | Mode log aplikasi: `debug`, `info`, `warn`, `error`, dan `disabled`.                                                            |
| `MYSMSMASKING_USER`     | `string`  | √       |                              | Username akun MySMSMasking.                                                                                                     |
| `MYSMSMASKING_PASSWORD` | `string`  | √       |                              | Password akun MySMSMasking.                                                                                                     |
| `BALANCE_LIMIT`         | `integer` |         | `300000`                     | Jika saldo kurang dari nilai variabel ini maka alert via discord webhook akan terpicu.                                          |
| `GRACE_PERIOD`          | `integer` |         | `14`                         | Jumlah hari menjelang tanggal kedaluarsa akun. Alert akan terpicu jika tanggal sekarang >= (tanggal kedaluarsa - variabel ini). |
| `SCHEDULE`              | `string`  | √       |                              | Jadwal pemeriksaan status akun MySMSMasking (dalam format CRON).                                                                |

## Lisensi

MIT kok. Insya Allah _Open Source_ selamanya.

## Cara Kontribusi

- Silahkan PR saja.
- **WAJIB** menggunakan bahasa Indonesia jika ingin mengubah atau menambahkan info di `README.md`.
- Jika ingin ditiru, mohon pertimbangkan ini: _ATM lebih baik dari ATP_  (**ATM**=Amati Tiru Modifikasi, **ATP**=Amati Tiru Plek-plek)
- Oh iya aplikasi ini dibuat dengan bahasa pemrograman _GO_ ya.
