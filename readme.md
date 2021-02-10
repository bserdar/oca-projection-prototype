# OCA Projection Overlay Prototype: FHIR Immunization Bundle to Verifiable Credentials

> OCA is an architecture that presents a schema as a multi-dimensional object consisting of a stable schema base and interoperable overlays. 

OCA information: https://wiki.trustoverip.org/display/HOME/Semantic+Domain+Group

OCA primarily deals with data capture. It is essentially flat, and does not directly support nested structures like JSON/XML, and thus, FHIR. There is an RFC to extend OCA to support this:

https://github.com/bserdar/oca-spec/tree/format/RFCs/003-bserdar-schema-format

The projection overlay is specified here:

https://github.com/bserdar/oca-spec/tree/projection/RFCs/004-bserdar-projection

This is a prototype of the Projection overlay that shows it is possible to use it to generate a verifiable credential from a FHIR bundle. It uses the CDC vaccie reporting specification as the output. At this point this overlay does not generate all necessary fields, but it is possible to extend it to support full VC generation.
