# josh

* Takes care of JSON response serialization and request deserialization.
* Follows JSON:API spec.
* Statically ensures that you don't make common mistakes:
  * Don't forget the status code.
  * Don't set body for statuses that don't support body.
  * Don't try sending more headers when headers are already sent.
