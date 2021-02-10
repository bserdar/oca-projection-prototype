# OCA Projection Overlay Prototype: FHIR Immunization Bundle to Verifiable Credentials

> OCA is an architecture that presents a schema as a multi-dimensional object consisting of a stable schema base and interoperable overlays. 

OCA information: https://wiki.trustoverip.org/display/HOME/Semantic+Domain+Group

OCA primarily deals with data capture. It is essentially flat, and does not directly support nested structures like JSON/XML, and thus, FHIR. There is an RFC to extend OCA to support this:

https://github.com/bserdar/oca-spec/tree/format/RFCs/003-bserdar-schema-format

The projection overlay is specified here:

https://github.com/bserdar/oca-spec/tree/projection/RFCs/004-bserdar-projection

This is a prototype of the Projection overlay that shows it is possible to use it to generate a verifiable credential from a FHIR bundle. It uses the CDC vaccine reporting specification as the output. At this point this overlay does not generate all necessary fields, but it is possible to extend it to support full output fields and VC generation.

The point of the projection overlay is to provide a prescriptive
mechanism for a governing entity to publish the schema for the
required data format as well as projection overlays for different
common data sources to minimize manual implementations. The workflow
for generating VC using projection overlay would be something like the
following:

  * Retrieve data from the source (e.g. FHIR bundle containing immunization records)
  * Validate using existing tools (FHIR validator)
  * Process data using a projection overlay, and potentially with
    other overlays to populate all data fields
  * Validate the resulting document
  * Sign the document
