IP2WHOIS Go CLI
===============
This Go command line tool enables user to query the WHOIS domain information for a specific domain name.

This program requires an API key to function. You may sign up for a free API key at https://www.ip2location.io/pricing.


Installation
============

#### `go install` Installation

```bash
go install github.com/ip2whois/ip2whois-cli/ip2whois@latest
```


#### Git Installation

```bash
git clone https://github.com/ip2whois/ip2whois-cli ip2whois-cli
cd ip2whois-cli
go install ./ip2whois/
$GOPATH/bin/ip2whois
```


#### Debian/Ubuntu (amd64)

```bash
curl -LO https://github.com/ip2whois/ip2whois-cli/releases/download/v1.0.0/ip2whois-1.0.0.deb
sudo dpkg -i ip2whois-1.0.0.deb
```


### Windows Powershell

Launch Powershell as administrator then run the below:

```bash
iwr -useb https://github.com/ip2whois/ip2whois-cli/releases/download/v1.0.0/windows.ps1 | iex
```


### Download pre-built binaries

Supported OS/architectures below:

```
darwin_amd64
darwin_arm64
dragonfly_amd64
freebsd_386
freebsd_amd64
freebsd_arm
freebsd_arm64
linux_386
linux_amd64
linux_arm
linux_arm64
netbsd_386
netbsd_amd64
netbsd_arm
netbsd_arm64
openbsd_386
openbsd_amd64
openbsd_arm
openbsd_arm64
solaris_amd64
windows_386
windows_amd64
windows_arm
```

After choosing a platform `PLAT` from above, run:

```bash
# for Windows, use ".zip" instead of ".tar.gz"
curl -LO https://github.com/ip2whois/ip2whois-cli/releases/download/v1.0.0/ip2whois_1.0.0_${PLAT}.tar.gz
# OR
wget https://github.com/ip2whois/ip2whois-cli/releases/download/v1.0.0/ip2whois_1.0.0_${PLAT}.tar.gz

tar -xvf ip2whois_1.0.0_${PLAT}.tar.gz
mv ip2whois_1.0.0_${PLAT} /usr/local/bin/ip2whois
```


Usage Examples
==============

### Display help
```bash
ip2whois -h
```

### Configure API key
```bash
ip2whois config <API KEY>
```

### Query WHOIS for specific domain (JSON)
```bash
ip2whois locaproxy.com
```

### Query WHOIS for specific domain (pretty print)
```bash
ip2whois -o pretty locaproxy.com
```

### Query WHOIS for specific domain and show only specific result fields
```bash
ip2whois -f domain,domain_id,status,create_date,registrar.name locaproxy.com
```

### Convert normal domain name to Punycode
```bash
ip2whois normal2puny <DOMAIN>
```

### Convert Punycode to normal domain name
```bash
ip2whois puny2normal <DOMAIN>
```


Example API Response
====================
```json
{
  "domain": "locaproxy.com",
  "domain_id": "1710914405_DOMAIN_COM-VRSN",
  "status": "clientTransferProhibited https://icann.org/epp#clientTransferProhibited",
  "create_date": "2012-04-03T02:34:32Z",
  "update_date": "2021-12-03T02:54:57Z",
  "expire_date": "2024-04-03T02:34:32Z",
  "domain_age": 4280,
  "whois_server": "whois.godaddy.com",
  "registrar": {
    "iana_id": "146",
    "name": "GoDaddy.com, LLC",
    "url": "https://www.godaddy.com"
  },
  "registrant": {
    "name": "Registration Private",
    "organization": "Domains By Proxy, LLC",
    "street_address": "DomainsByProxy.com",
    "city": "Tempe",
    "region": "Arizona",
    "zip_code": "85284",
    "country": "US",
    "phone": "+1.4806242599",
    "fax": "",
    "email": "Select Contact Domain Holder link at https://www.godaddy.com/whois/results.aspx?domain=LOCAPROXY.COM"
  },
  "admin": {
    "name": "Registration Private",
    "organization": "Domains By Proxy, LLC",
    "street_address": "DomainsByProxy.com",
    "city": "Tempe",
    "region": "Arizona",
    "zip_code": "85284",
    "country": "US",
    "phone": "+1.4806242599",
    "fax": "",
    "email": "Select Contact Domain Holder link at https://www.godaddy.com/whois/results.aspx?domain=LOCAPROXY.COM"
  },
  "tech": {
    "name": "Registration Private",
    "organization": "Domains By Proxy, LLC",
    "street_address": "DomainsByProxy.com",
    "city": "Tempe",
    "region": "Arizona",
    "zip_code": "85284",
    "country": "US",
    "phone": "+1.4806242599",
    "fax": "",
    "email": "Select Contact Domain Holder link at https://www.godaddy.com/whois/results.aspx?domain=LOCAPROXY.COM"
  },
  "billing": {
    "name": "",
    "organization": "",
    "street_address": "",
    "city": "",
    "region": "",
    "zip_code": "",
    "country": "",
    "phone": "",
    "fax": "",
    "email": ""
  },
  "nameservers": [
    "vera.ns.cloudflare.com",
    "walt.ns.cloudflare.com"
  ]
}
```


LICENCE
=====================
See the LICENSE file.
