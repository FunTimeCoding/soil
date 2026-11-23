{%- import_text 'common/files/banner.txt' as banner %}
{% import_yaml "common/files/kernel.yml" as kernel %}

include:
  - .child

motd:
  file.managed:
    - name: /etc/motd
    - source: salt://common/files/motd

common_packages:
  pkg.installed:
    - pkgs:
      - orphan
