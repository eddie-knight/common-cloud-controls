<!-- markdownlint-disable -->
# {{ .Metadata.Id }} v{{ (lastReleaseDetails .ReleaseDetails).Version }} ({{ .Metadata.Title }})

<img height="250px" src="https://raw.githubusercontent.com/finos/branding/882d52260eb9b85a4097db38b09a52ea9bb68734/project-logos/active-project-logos/Common%20Cloud%20Controls%20Logo/Horizontal/2023_FinosCCC_Horizontal_BLK.svg" alt="CCC Logo"/>

{{ .Metadata.Description }}

## Release Notes

> {{ (lastReleaseDetails .ReleaseDetails).ReleaseManager.Summary }}

Release Manager - **{{ (lastReleaseDetails .ReleaseDetails).ReleaseManager.Name }}, {{ (lastReleaseDetails .ReleaseDetails).ReleaseManager.Company }}** ([{{ (lastReleaseDetails .ReleaseDetails).ReleaseManager.GithubId }}](https://github.com/{{ (lastReleaseDetails .ReleaseDetails).ReleaseManager.GithubId }}))

## Changes Since Last Release
{{ range (lastReleaseDetails .ReleaseDetails).ChangeLog }}
- {{ . }}
{{- end }}

## Capabilities

|Capability ID|Capability Title|
|----|----|
{{- range .Capabilities }}
|{{ .Id }}|{{ .Title }}|
{{- end }}

---
{{ range .Capabilities }}
### {{ .Id }} - {{ .Title }}

{{ .Description }}
{{- end }}

## Threats

|Threat ID|Threat Title|
|----|----|
{{- range .Threats }}
|{{ .Id }}|{{ .Title }}|
{{- end }}

---
{{ range .Threats }}
### {{ .Id }} - {{ .Title }}

{{ .Description }}

{{ if .Capabilities -}}
**Impacted Capabilities:**

| Source | Capability |
| --- | --- |
{{- range .Capabilities }}
  {{- $referenceId := .ReferenceId }}
  {{- range .Identifiers }}
| {{ $referenceId }} | {{ . }} |
  {{- end }}
{{- end }}
{{- end }}

**Related Mappings:**

| Source | Mapping |
| --- | --- |
{{- range .ExternalMappings }}
  {{- $referenceId := .ReferenceId }}
  {{- range .Identifiers }}
| {{ $referenceId }} | {{ . }} |
  {{- end }}
{{- end }}
{{ end }}

## Controls

|Control ID|Control Title|
|----|----|
{{- range .ControlFamilies }}
{{- range .Controls }}
|{{ .Id }}|{{ .Title }}|
{{- end }}
{{- end }}

---

{{- range .ControlFamilies }}
{{ $family := .Title }}
{{- range .Controls }}

### {{ .Id }} - {{ .Title }}

{{ .Objective }}

**Control Family:** {{ $family }}

{{ if .ThreatMappings -}}
#### Mitigated Threats

| Threat Catalog | Mapped Threats |
| --- | --- |
{{- range .ThreatMappings }}
  {{- $referenceId := .ReferenceId }}
  {{- range .Identifiers }}
| {{ $referenceId }} | {{ . }} |
  {{- end }}
{{- end }}
{{- end }}

{{ if .GuidelineMappings -}}
#### Associated Guidelines

| Guideline | Mapped Controls |
| --- | --- |
{{- range .GuidelineMappings }}
  {{- $referenceId := .ReferenceId }}
  {{- range .Identifiers }}
| {{ $referenceId }} | {{ . }} |
  {{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}

## Contributing Organizations

We would like to acknowledge the following organizations for their valuable contributions to this project:

|  | <img src="https://www.finos.org/hubfs/FINOS/finos-logo/FINOS_Icon_Wordmark_Name_horz_White.svg" height="80" alt="FINOS Logo"> |  |
|:--:|:--:|:--:|
| <img src="https://www.finos.org/hs-fs/hubfs/2-Jan-18-2025-03-02-33-3610-AM.png?width=1039&height=494&name=2-Jan-18-2025-03-02-33-3610-AM.png" height="80" alt="Citigroup Logo"> | <img src="https://www.finos.org/hs-fs/hubfs/69-1.png?width=2078&height=988&name=69-1.png" height="80" alt="Scott Logic Logo"> | <img src="https://www.finos.org/hs-fs/hubfs/37.png?width=2078&height=988&name=37.png" height="80" alt="Sonatype Logo"> |
| <img src="https://www.finos.org/hubfs/FINOS/finos-logo/FINOS_Icon_Wordmark_Name_horz_White.svg" height="80" alt="Logo 7"> | <img src="https://www.finos.org/hubfs/FINOS/finos-logo/FINOS_Icon_Wordmark_Name_horz_White.svg" height="80" alt="Logo 8"> | <img src="https://www.finos.org/hubfs/FINOS/finos-logo/FINOS_Icon_Wordmark_Name_horz_White.svg" height="80" alt="Logo 9"> |

<!-- Add or remove rows as needed -->

<!-- markdownlint-enable -->
