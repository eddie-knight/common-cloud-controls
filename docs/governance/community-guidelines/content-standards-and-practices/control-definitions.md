# Control Definitions

A control definitions document is a detailed specification of security measures designed to mitigate risks and ensure compliance with established standards and frameworks. It must contain a list of all controls that pertain to a specific service category, along with relevant supporting information.

A set of control definitions should be created for each service category in the CCC Taxonomy, with each control in the document mapped to a feature of that service. A set of behavioral test definitions may be provided alongside the control definition in Gherkin to support validation of the controls.

## Common vs Specific Controls

In order to streamline maintenance, the CCC project maintains a list of [common controls].

In each service category's `controls.yaml` document, common controls are referenced in a list of IDs in the top-level value `common-controls`.

In the release pipeline, our [delivery tooling] will compile the common controls into the document alongside the specific controls. In the final output, the only difference in presentation of the controls will be the unique identifier.

## Control Definition Format

In order to create a cohesive standard that is readily useful to end users, controls must be indistinguishable from each other in format, style, and tone. Similarly, all controls must match the layout presented in the [control template](../templates/controls.yaml) prior to release.

A review from the [Communications WG] is recommended, but not required, in cases where additional support is needed to match the writing style and tone.

[!NOTE] The list of common controls follows a similar but unique format, which can be found in the [common controls] file.

### Common Control References

When documenting controls for a service category, begin by reviewing the existing [common controls]. In the event that a common control applies to this category, you may reference it from your document by adding its ID to the list `common-controls` at the top level of the controls document.

In the event that a common entry does not exist for this control, consider whether the control will apply to at least three other service categories. Or, look for a place where an existing _specific control_ can be genericized and moved to the _common controls_. After adding the new control definition to [common controls], add its ID in `common-controls`.

If a control is unique to this service category, add the full control definition within the `specific-controls` value in the controls document for this service category.

### Control Definition Values

The following list outlines the values necessary to create a new control definition using the control template. Refer to the feature or control definition for the `Service Category` name and abbreviation.

- **Category Title** - The title of the service category this control belongs to, formatted as `CCC <Service Category> Security Controls`.
- **Control ID** - A unique identifier for the control, following the format `CCC.<Service Category Abbreviation>.C#`.
- **Feature ID** - The corresponding feature ID that this control is associated with.
- **Control Title** - A brief title (3 to 10 words) that succinctly describes the control.
- **Objective** - A 1 to 3 sentence description of the control’s objective and what it aims to achieve.
- **Control Family** - Name of the [Control Family](#control-family) this group belongs to.
- **NIST CSF** - The specific ID from the NIST Cybersecurity Framework that corresponds to the control.
- **MITRE ATT&CK** - The unique identifier for the MITRE ATT&CK Tactics, Techniques, and Procedures (TTP) relevant to the control.
- **Threats** - A list of IDs for CCC threats that this control is designed to mitigate.
- **Control Mappings** - Identifiers for other frameworks that map to this control (CCM, ISO 27001, NIST 800-53, etc).
- **Test Requirements** - Detailed descriptions of the testing requirements needed to validate the control’s implementation.

### Control Family

Control family refers to a group of related security controls that are organized together based on their similar functions or objectives. Each control family addresses a particular aspect of information security.

The following list should be updated in the event that a new control family is added to the CCC Controls Catalog:

- Encryption
- Data
- Identity and Access Management
- Logging & Monitoring
- Software Supply Chain


[Communications WG]: ../../working-groups/communications/charter.md