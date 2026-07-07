---
name: build-threat-catalog
description: Create a CCC threat catalog (threats.yaml) for a cloud service — import applicable core threats, define service-specific threats mapped to the service's capabilities, and map each to external frameworks (MITRE ATT&CK, MITRE D3FEND, CISA KEV, CWE, OWASP), validating against schemas/threats-schema.json. Requires the service folder to already contain metadata.yaml and capabilities.yaml. Use when the user asks to "identify threats for `service`", "create a threat catalog for `service`", or "create threats.yaml" for a service.
---

# build-threat-catalog (pointer)

The canonical skill lives at [`skills/build-threat-catalog/SKILL.md`](../../../skills/build-threat-catalog/SKILL.md).
Read that file now and follow it exactly — including any per-step output formats
and confirmation gates. Do not act on this pointer file alone.
