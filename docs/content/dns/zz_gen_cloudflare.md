---
title: "Cloudflare"
date: 2019-03-03T16:39:46+01:00
draft: false
slug: cloudflare
dnsprovider:
  since:    "v0.3.0"
  code:     "cloudflare"
  url:      "https://www.cloudflare.com/dns/"
---

<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->
<!-- providers/dns/cloudflare/cloudflare.toml -->
<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->


Configuration for [Cloudflare](https://www.cloudflare.com/dns/).


<!--more-->

- Code: `cloudflare`
- Since: v0.3.0


Here is an example bash command using the Cloudflare provider:

```bash
CLOUDFLARE_EMAIL=you@example.com \
CLOUDFLARE_API_KEY=b9841238feb177a84330febba8a83208921177bffe733 \
lego --email you@example.com --dns cloudflare -d '*.example.com' -d example.com run

# or

CLOUDFLARE_DNS_API_TOKEN=1234567890abcdefghijklmnopqrstuvwxyz \
lego --email you@example.com --dns cloudflare -d '*.example.com' -d example.com run
```




## Credentials

| Environment Variable Name | Description |
|-----------------------|-------------|
| `CF_API_EMAIL` | Account email |
| `CF_API_KEY` | API key |
| `CF_DNS_API_TOKEN` | API token with DNS:Edit permission (since v3.1.0) |
| `CF_ZONE_API_TOKEN` | API token with Zone:Read permission (since v3.1.0) |
| `CLOUDFLARE_API_KEY` | Alias to CF_API_KEY |
| `CLOUDFLARE_DNS_API_TOKEN` | Alias to CF_DNS_API_TOKEN |
| `CLOUDFLARE_EMAIL` | Alias to CF_API_EMAIL |
| `CLOUDFLARE_ZONE_API_TOKEN` | Alias to CF_ZONE_API_TOKEN |

The environment variable names can be suffixed by `_FILE` to reference a file instead of a value.
More information [here]({{% ref "dns#configuration-and-credentials" %}}).


## Additional Configuration

| Environment Variable Name | Description |
|--------------------------------|-------------|
| `CLOUDFLARE_BASE_URL` | API base URL (Default: https://api.cloudflare.com/client/v4) |
| `CLOUDFLARE_HTTP_TIMEOUT` | API request timeout in seconds (Default: ) |
| `CLOUDFLARE_POLLING_INTERVAL` | Time between DNS propagation check in seconds (Default: 2) |
| `CLOUDFLARE_PROPAGATION_TIMEOUT` | Maximum waiting time for DNS propagation in seconds (Default: 120) |
| `CLOUDFLARE_TTL` | The TTL of the TXT record used for the DNS challenge in seconds (Default: 120) |

The environment variable names can be suffixed by `_FILE` to reference a file instead of a value.
More information [here]({{% ref "dns#configuration-and-credentials" %}}).

## Description

You may use `CF_API_EMAIL` and `CF_API_KEY` to authenticate, or `CF_DNS_API_TOKEN`, or `CF_DNS_API_TOKEN` and `CF_ZONE_API_TOKEN`.

### API keys

If using API keys (`CF_API_EMAIL` and `CF_API_KEY`), the Global API Key needs to be used, not the Origin CA Key.

Please be aware, that this in principle allows Lego to read and change *everything* related to this account.

### API tokens

With API tokens (`CF_DNS_API_TOKEN`, and optionally `CF_ZONE_API_TOKEN`),
very specific access can be granted to your resources at Cloudflare.
See this [Cloudflare announcement](https://blog.cloudflare.com/api-tokens-general-availability/) for details.

The main resources Lego cares for are the DNS entries for your Zones.
It also needs to resolve a domain name to an internal Zone ID in order to manipulate DNS entries.

Hence, you should create an API token with the following permissions:

* Zone / Zone / Read
* Zone / DNS / Edit

You also need to scope the access to all your domains for this to work.
Then pass the API token as `CF_DNS_API_TOKEN` to Lego.

**Alternatively,** if you prefer a more strict set of privileges,
you can split the access tokens:

* Create one with *Zone / Zone / Read* permissions and scope it to all your zones or just the individual zone you need to edit.
  This is needed to resolve domain names to Zone IDs and can be shared among multiple Lego installations.
  Pass this API token as `CF_ZONE_API_TOKEN` to Lego.
* Create another API token with *Zone / DNS / Edit* permissions and set the scope to the domains you want to manage with a single Lego installation.
  Pass this token as `CF_DNS_API_TOKEN` to Lego.
* Repeat the previous step for each host you want to run Lego on.
* It is possible to use the same api token for both variables if it is given `Zone:Read` and `DNS:Edit` permission for the zone.

This "paranoid" setup is mainly interesting for users who manage many zones/domains with a single Cloudflare account.
It follows the principle of least privilege and limits the possible damage, should one of the hosts become compromised.



## More information

- [API documentation](https://api.cloudflare.com/)
- [Go client](https://github.com/cloudflare/cloudflare-go)

<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->
<!-- providers/dns/cloudflare/cloudflare.toml -->
<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->
