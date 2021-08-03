# HighMiddCovTipper

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/aURORA-JC/HighMiddConvTipper?logo=go&style=flat-square)![Golang: >=13.3 (shields.io)](https://img.shields.io/badge/Golang->%3D13.3-brightgreen?style=flat-square&logo=go)

HighMiddCovTipper is a probe script which can get the information of high & middle risk areas cased by COVID-19 in Mainland China from the National Health Commission of PRC's  API and push the formatted content to you automatically. 

This script just for ServerChan's Push Service. More: [ServerChan](https://sct.ftqq.com/) -  [@easychen](https://github.com/easychen)

## Install

1. Clone the repository from Github

```bash
git clone https://github.com/aURORA-JC/HighMiddCovTipper.git
```

2. Cd to project root path & build

```bash
go build
```

3. HighMiddCovTipper will be build in the project root path

## Usage
+ Start once

```bash
# Windows (use cmd.exe or powershell)
.\HighMiddCovTipper.exe <YOUR_SERVERCHAN_TOKEN>

# Linux
./HighMiddCovTipper <YOUR_SERVERCHAN_TOKEN>
```
+  Period task

```bash
# On Linux use crontab
# Edit crontab schedule with vim or nano
crontab -e

# Every hour run
*/60 * * * * <YOUR_SCRIPT_PATH> <YOUR_SERVERCHAN_TOKEN>
# For example, */60 * * * * /home/test/HighMiddCovTipper SCT49339TdVuPhktIqXCEUrRNIlJWxRur

# Or you can create a shell script and run that shell script instead
```

## Contributing

PRs accepted.

## License

MIT © aURORA-JC
