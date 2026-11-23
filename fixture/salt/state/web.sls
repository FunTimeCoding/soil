{% from 'macros/helper.jinja' import greeting %}
{% set role = grains['role'] %}
include:
  - service.{{ role }}
  - absent
  - required

web_configuration:
  file.recurse:
    - name: /etc/web
    - source: salt://web/conf
