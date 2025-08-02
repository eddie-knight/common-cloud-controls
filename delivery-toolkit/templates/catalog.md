<!-- markdownlint-disable -->
<img width="50%" src="https://raw.githubusercontent.com/finos/branding/882d52260eb9b85a4097db38b09a52ea9bb68734/project-logos/active-project-logos/Common%20Cloud%20Controls%20Logo/Horizontal/2023_FinosCCC_Horizontal_BLK.svg" alt="CCC Logo"/>

# {{ .Metadata.Id }} v{{ (lastReleaseDetails .ReleaseDetails).Version }} ({{ .Metadata.Title }})

{{ .Metadata.Description }}

## Notes from the Release Manager:

> _{{ (lastReleaseDetails .ReleaseDetails).ReleaseManager.Summary|safe }}_
>
> _- {{ (lastReleaseDetails .ReleaseDetails).ReleaseManager.Name }}, {{ (lastReleaseDetails .ReleaseDetails).ReleaseManager.Company }} ([{{ (lastReleaseDetails .ReleaseDetails).ReleaseManager.GithubId }}](https://github.com/{{ (lastReleaseDetails .ReleaseDetails).ReleaseManager.GithubId }}))_

### Changes Since Last Release:

{{ range (lastReleaseDetails .ReleaseDetails).ChangeLog }}
- {{ . }}
{{- end }}

## Capabilities

The following capabilities are required to be present on a resource for it to be considered a {{ .Metadata.Title }} service. Threats outlined in this catalog are assesssed based on the presence of these capabilities.

|Capability ID|Capability Title|Description|
|----|----|----|
{{- range .Capabilities }}
|{{ .Id }}|{{ .Title }}|{{.Description|safe}}|
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

<div style="display: flex; width: 100%;">
  <div style="flex: 1; padding-right: 10px;">
  {{ if .Capabilities -}}
  <i>This threat is associated with the following capabilities:</i>
  <ul>
    {{ range .Capabilities }}
      {{ range .Identifiers }}
  <li>{{ . }}</li>
      {{- end }}
    {{- end }}
  </ul>
  {{- end }}
  </div>
  <div style="flex: 1; padding-left: 10px;">
    {{ if .ExternalMappings -}}
    <i>The authors considered the following external elements noteworthy for this threat:</i>
    <table cellpadding="5" style="width:100%; border-collapse: collapse; border-style: hidden;">
      <thead>
        <tr>
          <th style="border: 1px solid #ddd; padding: 8px;">Source</th>
          <th style="border: 1px solid #ddd; padding: 8px;">Mapping</th>
        </tr>
      </thead>
      <tbody>
        {{- range .ExternalMappings }}
          {{- $referenceId := .ReferenceId }}
          {{- range .Identifiers }}
        <tr>
          <td style="border: 1px solid #ddd; padding: 8px;">{{ $referenceId }}</td>
          <td style="border: 1px solid #ddd; padding: 8px;">{{ . }}</td>
        </tr>
          {{- end }}
        {{- end }}
      </tbody>
    </table>
    {{- end }}
  </div>
</div>
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
### {{ .Title }}
{{- range .Controls }}

#### {{ .Id }}

**{{ .Title }}**

{{ .Objective }}

<div style="display: flex; width: 100%;">
  <div style="flex: 1; padding-right: 10px;">
    {{ if .ThreatMappings -}}
    <table cellpadding="5" style="width:100%; border-collapse: collapse; border-style: hidden;">
      <thead>
        <tr>
          <th style="border: 1px solid #ddd; padding: 8px;">Threat Catalog</th>
          <th style="border: 1px solid #ddd; padding: 8px;">Related Threat</th>
        </tr>
      </thead>
      <tbody>
        {{- range .ThreatMappings }}
          {{- $referenceId := .ReferenceId }}
          {{- range .Identifiers }}
        <tr>
          <td style="border: 1px solid #ddd; padding: 8px;">{{ $referenceId }}</td>
          <td style="border: 1px solid #ddd; padding: 8px;">{{ . }}</td>
        </tr>
          {{- end }}
        {{- end }}
      </tbody>
    </table>
    {{- end }}
  </div>
  <div style="flex: 1; padding-left: 10px;">
    {{ if .GuidelineMappings -}}
    <table cellpadding="5" style="width:100%; border-collapse: collapse; border-style: hidden;">
      <thead>
        <tr>
          <th style="border: 1px solid #ddd; padding: 8px;">Guideline</th>
          <th style="border: 1px solid #ddd; padding: 8px;">Related Guidance</th>
        </tr>
      </thead>
      <tbody>
        {{- range .GuidelineMappings }}
          {{- $referenceId := .ReferenceId }}
          {{- range .Identifiers }}
        <tr>
          <td style="border: 1px solid #ddd; padding: 8px;">{{ $referenceId }}</td>
          <td style="border: 1px solid #ddd; padding: 8px;">{{ . }}</td>
        </tr>
          {{- end }}
        {{- end }}
      </tbody>
    </table>
    {{- end }}
  </div>
</div>
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
