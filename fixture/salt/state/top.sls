base:
  '*':
    - common

prod:
  'web*':  # web servers
    - match: glob
    # retired:
    - web
{% if grains['id'].startswith('db') %}
  'db*':
    - conditional
{% endif %}

staging:
  'st*':
    - stagingonly
