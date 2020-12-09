# GO-OVH-DYNDNS

A little go tool that updates the DNS record for your dynamic dns on one of your TLD's. \
For this tool to work you need to [setup a dynamic dns](https://docs.ovh.com/gb/en/domains/hosting_dynhost/#objective) on one of your domains via the OVH Web control panel.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Support](#support)
- [Licence](#licence)

## Installation

### Use as a standalone tool
Clone this repository, change directory and copy configuration template:
```bash
git clone https://github.com/codefuzzler/go-dyndns-ovh.git && \
cd go-dyndns-ovh && \
```

### Use as a Docker container
```bash
docker run -it -e OVH_DNS_RECORD=sub.domain.tld \
               -e OVH_DNS_USER=domain.tld-id \
               -e OVH_DNS_PASS=changeme \
                codefuzzler/ovh-dyndns:latest
```

## Usage
To configure the tool to update the DNS Record for the subdomain "example.domain.tld" either:
* set the following environment variables:
* Copy the .env.tmpl file to .env and edit it

```bash
export OVH_DNS_RECORD=sub.domain.tld
export OVH_DNS_USER=domain.tld-id
export OVH_DNS_PASS=changeme
export OVH_DNS_LOOP=0 # milliseconds to wait before executing again, set to 0 for single execution
```


## Support

Please [open an issue](https://github.com/codefuzzler/go-dyndns-ovh/issues/new) for support.

## Licence
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <https://unlicense.org>
