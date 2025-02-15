---
title: "Selectel"
date: 2019-03-03T16:39:46+01:00
draft: false
slug: selectel
dnsprovider:
  since:    "v1.2.0"
  code:     "selectel"
  url:      "https://kb.selectel.com/"
---

<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->
<!-- providers/dns/selectel/selectel.toml -->
<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->


Configuration for [Selectel](https://kb.selectel.com/).


<!--more-->

- Code: `selectel`
- Since: v1.2.0


Here is an example bash command using the Selectel provider:

```bash
SELECTEL_API_TOKEN=xxxxx \
lego --email you@example.com --dns selectel -d '*.example.com' -d example.com run
```




## Credentials

| Environment Variable Name | Description |
|-----------------------|-------------|
| `SELECTEL_API_TOKEN` | API token |

The environment variable names can be suffixed by `_FILE` to reference a file instead of a value.
More information [here]({{% ref "dns#configuration-and-credentials" %}}).


## Additional Configuration

| Environment Variable Name | Description |
|--------------------------------|-------------|
| `SELECTEL_BASE_URL` | API endpoint URL |
| `SELECTEL_HTTP_TIMEOUT` | API request timeout in seconds (Default: 30) |
| `SELECTEL_POLLING_INTERVAL` | Time between DNS propagation check in seconds (Default: 2) |
| `SELECTEL_PROPAGATION_TIMEOUT` | Maximum waiting time for DNS propagation in seconds (Default: 120) |
| `SELECTEL_TTL` | The TTL of the TXT record used for the DNS challenge in seconds (Default: 60) |

The environment variable names can be suffixed by `_FILE` to reference a file instead of a value.
More information [here]({{% ref "dns#configuration-and-credentials" %}}).




## More information

- [API documentation](https://kb.selectel.com/23136054.html)

<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->
<!-- providers/dns/selectel/selectel.toml -->
<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->
